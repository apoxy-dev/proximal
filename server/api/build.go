package api

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gogo/status"
	"github.com/google/uuid"
	tclient "go.temporal.io/sdk/client"
	"golang.org/x/exp/maps"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/apoxy-dev/proximal/core/log"
	serverdb "github.com/apoxy-dev/proximal/server/db"
	sqlc "github.com/apoxy-dev/proximal/server/db/sql"
	"github.com/apoxy-dev/proximal/server/ingest"

	middlewarev1 "github.com/apoxy-dev/proximal/api/middleware/v1"
)

func buildFromRow(b *sqlc.Build) *middlewarev1.Build {
	return &middlewarev1.Build{
		Sha:          b.Sha,
		Status:       middlewarev1.Build_BuildStatus(middlewarev1.Build_BuildStatus_value[b.Status]),
		StatusDetail: b.StatusDetail,
		StartedAt:    timestamppb.New(b.StartedAt.Time),
		UpdatedAt:    timestamppb.New(b.UpdatedAt.Time),
	}
}

func (s *MiddlewareService) ListBuilds(ctx context.Context, req *middlewarev1.ListBuildsRequest) (*middlewarev1.ListBuildsResponse, error) {
	log.Infof("list middlewares: %s", protojson.Format(req))

	ts := time.Now()
	if req.PageToken != "" {
		dTs, err := decodeNextPageToken(req.PageToken)
		if err != nil {
			log.Errorf("failed to decode page token: %v", err)
			return nil, status.Error(codes.InvalidArgument, "invalid page token")
		}

		log.Debugf("decoded page token: %s", string(dTs))

		ts, err = time.Parse(time.RFC3339Nano, string(dTs))
		if err != nil {
			log.Errorf("failed to parse page token: %v", err)
			return nil, status.Error(codes.InvalidArgument, "invalid page token")
		}
	}
	pageSize := req.PageSize
	if pageSize == 0 {
		pageSize = 100
	}

	bs, err := s.db.Queries().ListBuildsByMiddlewareSlug(ctx, sqlc.ListBuildsByMiddlewareSlugParams{
		MiddlewareSlug: req.Slug,
		Datetime:       ts,
		Limit:          int64(pageSize),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list middleware: %w", err)
	}

	builds := make([]*middlewarev1.Build, 0, len(bs))
	for _, b := range bs {
		builds = append(builds, buildFromRow(&b))
	}

	var nextPageToken string
	if len(builds) == int(pageSize) {
		nextPageToken = encodeNextPageToken(bs[len(bs)-1].StartedAt.Time.Format(time.RFC3339Nano))
	}

	return &middlewarev1.ListBuildsResponse{
		Builds:        builds,
		NextPageToken: nextPageToken,
	}, nil
}

func (s *MiddlewareService) GetBuild(ctx context.Context, req *middlewarev1.GetBuildRequest) (*middlewarev1.Build, error) {
	b, err := s.db.Queries().GetBuildByMiddlewareSlugAndSha(ctx, sqlc.GetBuildByMiddlewareSlugAndShaParams{
		MiddlewareSlug: req.Slug,
		Sha:            req.Sha,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "build not found for middleware", req.Slug)
		}
		return nil, status.Errorf(codes.Internal, "unable to get build: %v", err)
	}
	return buildFromRow(&b), nil
}

func (s *MiddlewareService) GetLiveBuild(ctx context.Context, req *middlewarev1.GetLiveBuildRequest) (*middlewarev1.Build, error) {
	b, err := s.db.Queries().GetLiveReadyBuildByMiddlewareSlug(ctx, req.Slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "build not found for middleware", req.Slug)
		}
		return nil, status.Errorf(codes.Internal, "unable to get latest build: %v", err)
	}
	return buildFromRow(&b), nil
}

func (s *MiddlewareService) startBuildWorkflow(ctx context.Context, slug string, params *middlewarev1.MiddlewareIngestParams) error {
	tOpts := tclient.StartWorkflowOptions{
		ID:                       fmt.Sprintf("%s-%s", slug, uuid.New()),
		TaskQueue:                ingest.MiddlewareIngestQueue,
		WorkflowExecutionTimeout: 60 * time.Second,
	}
	in := &ingest.MiddlewareIngestParams{
		Slug:   slug,
		Params: params,
	}
	sw, err := s.tc.ExecuteWorkflow(ctx, tOpts, ingest.StartBuildWorkflow, in)
	if err != nil {
		return fmt.Errorf("unable to start ingest workflow: %v", err)
	}
	return sw.Get(ctx, nil)
}

func (s *MiddlewareService) TriggerBuild(ctx context.Context, req *middlewarev1.TriggerBuildRequest) (*emptypb.Empty, error) {
	log.Infof("TriggerBuild: %v", req)

	tx, err := s.db.Begin()
	if err != nil {
		log.Errorf("failed to begin transaction: %v", err)
		return nil, status.Error(codes.Internal, "failed to trigger build")
	}
	defer tx.Rollback()
	qtx := s.db.Queries().WithTx(tx)

	m, err := qtx.GetMiddlewareBySlug(ctx, req.Slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "middleware %q not found", req.Slug)
		}
		log.Errorf("unable to get middleware: %v", err)
		return nil, status.Errorf(codes.Internal, "unable to get middleware: %v", err)
	}

	var ingestParams middlewarev1.MiddlewareIngestParams
	if err := protojson.Unmarshal(m.IngestParamsJson, &ingestParams); err != nil {
		log.Errorf("unable to unmarshal ingest params: %v", err)
		return nil, status.Errorf(codes.Internal, "unable to unmarshal ingest params: %v", err)
	}

	if err := s.startBuildWorkflow(ctx, req.Slug, &ingestParams); err != nil {
		log.Errorf("unable to start build workflow: %v", err)
		return nil, status.Errorf(codes.Internal, "unable to start build workflow: %v", err)
	}

	_, err = qtx.UpdateMiddlewareStatus(ctx, sqlc.UpdateMiddlewareStatusParams{
		Slug:   req.Slug,
		Status: serverdb.MiddlewareStatusPendingReady,
		StatusDetail: sql.NullString{
			String: "build triggered",
			Valid:  true,
		},
		LiveBuildSha: m.LiveBuildSha,
	})
	if err != nil {
		log.Errorf("failed to update middleware status: %v", err)
		return nil, status.Error(codes.Internal, "failed to trigger build")
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("failed to commit transaction: %v", err)
		return nil, status.Error(codes.Internal, "failed to trigger build")
	}

	return &emptypb.Empty{}, nil
}

func (s *MiddlewareService) GetBuildOutput(ctx context.Context, req *middlewarev1.GetBuildOutputRequest) (*middlewarev1.BuildOutput, error) {
	log.Infof("GetBuildOutput: %v", req)

	allowedType := map[string]bool{
		"stdout":   true,
		"stderr":   true,
		"wasm.out": true,
	}
	if !allowedType[req.OutputType] {
		return nil, status.Errorf(codes.InvalidArgument, "invalid output type: %q (must be one of %v)", req.OutputType, maps.Keys(allowedType))
	}

	b, err := s.db.Queries().GetBuildByMiddlewareSlugAndSha(ctx, sqlc.GetBuildByMiddlewareSlugAndShaParams{
		MiddlewareSlug: req.Slug,
		Sha:            req.Sha,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "build not found for middleware %q", req.Slug)
		}
		return nil, status.Errorf(codes.Internal, "unable to get build: %v", err)
	}

	// TODO(dilyevsky): Make this stuff more streaming. Devs love streaming.
	if b.OutputPath.String == "" {
		log.Errorf("output path for build %s is empty", b.Sha)
		return nil, status.Errorf(codes.NotFound, "output not found for build %s", b.Sha)
	}
	path := filepath.Join(b.OutputPath.String, req.OutputType)
	output, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Errorf("output file not found in store for build %s: %s", b.Sha, path)
			return nil, status.Errorf(codes.NotFound, "output not found for build %s", b.Sha)
		}
		log.Errorf("unable to read output: %v", err)
		return nil, status.Errorf(codes.Internal, "output not found for build %s", b.Sha)
	}

	return &middlewarev1.BuildOutput{
		Build:  buildFromRow(&b),
		Output: output,
	}, nil
}

func (s *MiddlewareService) GetLiveBuildOutput(ctx context.Context, req *middlewarev1.GetLiveBuildOutputRequest) (*middlewarev1.BuildOutput, error) {
	r, err := s.GetLiveBuild(ctx, &middlewarev1.GetLiveBuildRequest{
		Slug: req.Slug,
	})
	if err != nil {
		return nil, err
	}
	return s.GetBuildOutput(ctx, &middlewarev1.GetBuildOutputRequest{
		Slug:       req.Slug,
		Sha:        r.Sha,
		OutputType: req.OutputType,
	})
}
