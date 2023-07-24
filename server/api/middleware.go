package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gogo/status"
	"github.com/golang/protobuf/ptypes/empty"
	tclient "go.temporal.io/sdk/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/apoxy-dev/proximal/core/log"
	serverdb "github.com/apoxy-dev/proximal/server/db"
	sqlc "github.com/apoxy-dev/proximal/server/db/sql"
	"github.com/apoxy-dev/proximal/server/watcher"

	middlewarev1 "github.com/apoxy-dev/proximal/api/middleware/v1"
)

// MiddlewareService implements the MiddlewareServiceServer gRPC service.
type MiddlewareService struct {
	db      *serverdb.DB
	tc      tclient.Client
	watcher *watcher.Watcher
}

// NewMiddlewareService returns a new MiddlewareService.
func NewMiddlewareService(db *serverdb.DB, tc tclient.Client, w *watcher.Watcher) *MiddlewareService {
	return &MiddlewareService{
		db:      db,
		tc:      tc,
		watcher: w,
	}
}

func (s *MiddlewareService) List(ctx context.Context, req *middlewarev1.ListRequest) (*middlewarev1.ListResponse, error) {
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

	ms, err := s.db.Queries().ListMiddlewares(ctx, sqlc.ListMiddlewaresParams{
		Datetime: ts,
		Limit:    int64(pageSize),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list middleware: %w", err)
	}

	mpbs := make([]*middlewarev1.Middleware, 0, len(ms))
	for _, m := range ms {
		mw, err := middlewareFromRow(&m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert middleware: %w", err)
		}
		mpbs = append(mpbs, mw)
	}

	var nextPageToken string
	if len(mpbs) == int(pageSize) {
		nextPageToken = encodeNextPageToken(mpbs[len(mpbs)-1].CreatedAt.AsTime().Format(time.RFC3339Nano))
	}

	return &middlewarev1.ListResponse{
		Middlewares:   mpbs,
		NextPageToken: nextPageToken,
	}, nil

}

func middlewareFromRow(m *sqlc.Middleware) (*middlewarev1.Middleware, error) {
	var ingestParams middlewarev1.MiddlewareIngestParams
	if err := protojson.Unmarshal(m.IngestParamsJson, &ingestParams); err != nil {
		return nil, fmt.Errorf("unable to unmarshal ingest params: %v", err)
	}
	var runtimeParams middlewarev1.MiddlewareRuntimeParams
	if err := protojson.Unmarshal(m.RuntimeParamsJson, &runtimeParams); err != nil {
		return nil, fmt.Errorf("unable to unmarshal runtime params: %v", err)
	}
	status, ok := middlewarev1.Middleware_MiddlewareStatus_value[strings.ToUpper(m.Status)]
	if !ok {
		return nil, fmt.Errorf("unknown middleware status: %v", m.Status)
	}
	return &middlewarev1.Middleware{
		Slug:          m.Slug,
		IngestParams:  &ingestParams,
		RuntimeParams: &runtimeParams,
		Status:        middlewarev1.Middleware_MiddlewareStatus(status),
		LiveBuildSha:  m.LiveBuildSha.String,
		CreatedAt:     timestamppb.New(m.CreatedAt.Time),
		UpdatedAt:     timestamppb.New(m.UpdatedAt.Time),
	}, nil
}

func (s *MiddlewareService) Get(ctx context.Context, req *middlewarev1.GetRequest) (*middlewarev1.Middleware, error) {
	m, err := s.db.Queries().GetMiddlewareBySlug(ctx, req.Slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "middleware not found")
		}
		log.Errorf("unable to get middleware: %v", err)
		return nil, status.Error(codes.Internal, "unable to get middleware")
	}

	pb, err := middlewareFromRow(&m)
	if err != nil {
		log.Errorf("unable to create middleware from row: %v", err)
		return nil, status.Error(codes.Internal, "unable to create middleware from row")
	}
	return pb, nil
}

// Create starts a new ingest workflow for the given middleware.
func (s *MiddlewareService) Create(ctx context.Context, req *middlewarev1.CreateRequest) (*middlewarev1.Middleware, error) {
	paramsJson, err := protojson.Marshal(req.Middleware.IngestParams)
	if err != nil {
		log.Errorf("unable to marshal params: %v", err)
		return nil, status.Error(codes.Internal, "unable to marshal params")
	}
	runParamsJson, err := protojson.Marshal(req.Middleware.RuntimeParams)
	if err != nil {
		log.Errorf("unable to marshal runtime params: %v", err)
		return nil, status.Error(codes.Internal, "unable to marshal runtime params")
	}

	_, err = s.db.Queries().GetMiddlewareBySlug(ctx, req.Middleware.Slug)
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "middleware already exists")
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("unable to get middleware by slug: %v", err)
		return nil, status.Error(codes.Internal, "unable to get middleware by slug")
	}

	tx, err := s.db.Begin()
	if err != nil {
		log.Errorf("failed to begin transaction: %v", err)
		return nil, status.Error(codes.Internal, "failed to start ingest")
	}
	defer tx.Rollback()
	qtx := s.db.Queries().WithTx(tx)

	m, err := qtx.CreateMiddleware(ctx, sqlc.CreateMiddlewareParams{
		Slug:              req.Middleware.Slug,
		SourceType:        serverdb.SourceTypeGitHub,
		IngestParamsJson:  paramsJson,
		RuntimeParamsJson: runParamsJson,
		Status:            serverdb.MiddlewareStatusPending,
		StatusDetail: sql.NullString{
			String: "pending",
			Valid:  true,
		},
	})
	if err != nil {
		log.Errorf("unable to create middleware: %v", err)
		return nil, status.Error(codes.Internal, "unable to create middleware")
	}

	if err := s.startBuildWorkflow(ctx, req.Middleware.Slug, req.Middleware.IngestParams); err != nil {
		log.Errorf("unable to start build workflow: %v", err)
		return nil, status.Error(codes.Internal, "unable to start build workflow")
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("unable to commit transaction: %v", err)
		return nil, status.Error(codes.Internal, "unable to commit transaction")
	}

	pb, err := middlewareFromRow(&m)
	if err != nil {
		log.Errorf("unable to create middleware from row: %v", err)
		return nil, status.Error(codes.Internal, "unable to create middleware from row")
	}
	return pb, nil
}

func (s *MiddlewareService) Update(ctx context.Context, req *middlewarev1.UpdateRequest) (*middlewarev1.Middleware, error) {
	m, err := s.db.Queries().GetMiddlewareBySlug(ctx, req.Middleware.Slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "middleware not found")
		}
		log.Errorf("unable to get middleware: %v", err)
		return nil, status.Error(codes.Internal, "unable to get middleware")
	}

	mpb, err := middlewareFromRow(&m)
	if err != nil {
		log.Errorf("unable to create middleware from row: %v", err)
		return nil, status.Error(codes.Internal, "unable to create middleware from row")
	}

	if mpb.Type != req.Middleware.Type {
		return nil, status.Error(codes.InvalidArgument, "cannot change middleware type")
	}
	if req.Middleware.IngestParams != nil {
		if req.Middleware.IngestParams.Type != mpb.IngestParams.Type {
			return nil, status.Error(codes.InvalidArgument, "cannot change ingest type")
		}
		if req.Middleware.IngestParams.Type == middlewarev1.MiddlewareIngestParams_GITHUB {
			if req.Middleware.IngestParams.GetGithubRepo() != mpb.IngestParams.GetGithubRepo() {
				return nil, status.Error(codes.InvalidArgument, "cannot change github repo")
			}
		} else if req.Middleware.IngestParams.Type == middlewarev1.MiddlewareIngestParams_DIRECT {
			if req.Middleware.IngestParams.GetWatchDir() != mpb.IngestParams.GetWatchDir() {
				return nil, status.Error(codes.InvalidArgument, "cannot change watch dir")
			}
		}
	}

	paramsJson, err := protojson.Marshal(req.Middleware.IngestParams)
	if err != nil {
		log.Errorf("unable to marshal params: %v", err)
		return nil, status.Error(codes.Internal, "unable to marshal params")
	}
	runParamsJson, err := protojson.Marshal(req.Middleware.RuntimeParams)
	if err != nil {
		log.Errorf("unable to marshal runtime params: %v", err)
		return nil, status.Error(codes.Internal, "unable to marshal runtime params")
	}

	upd, err := s.db.Queries().UpdateMiddleware(ctx, sqlc.UpdateMiddlewareParams{
		Slug:              req.Middleware.Slug,
		IngestParamsJson:  paramsJson,
		RuntimeParamsJson: runParamsJson,
	})
	if err != nil {
		log.Errorf("unable to update middleware: %v", err)
		return nil, status.Error(codes.Internal, "unable to update middleware")
	}

	upb, err := middlewareFromRow(&upd)
	if err != nil {
		log.Errorf("unable to create middleware from row: %v", err)
		return nil, status.Error(codes.Internal, "unable to create middleware from row")
	}

	return upb, nil
}

func (s *MiddlewareService) Delete(ctx context.Context, req *middlewarev1.DeleteRequest) (*emptypb.Empty, error) {
	log.Infof("deleting middleware %v", req.Slug)
	m, err := s.db.Queries().GetMiddlewareBySlug(ctx, req.Slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "middleware not found")
		}
		log.Errorf("unable to get middleware: %v", err)
		return nil, status.Error(codes.Internal, "unable to get middleware")
	}

	mpb, err := middlewareFromRow(&m)
	if err != nil {
		log.Errorf("unable to create middleware from row: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete middleware")
	}

	if mpb.IngestParams.Type == middlewarev1.MiddlewareIngestParams_DIRECT {
		if err := s.watcher.Remove(m.Slug); err != nil {
			log.Warnf("unable to remove watch dir: %v", err)
		}
	}

	if err := s.db.Queries().DeleteMiddleware(ctx, req.Slug); err != nil {
		log.Errorf("unable to delete middleware: %v", err)
		return nil, status.Error(codes.Internal, "unable to delete middleware")
	}
	return &emptypb.Empty{}, nil
}

func (s *MiddlewareService) InternalList(ctx context.Context, _ *empty.Empty) (*middlewarev1.ListResponse, error) {
	mws, err := s.db.Queries().ListMiddlewaresAll(ctx)
	if err != nil {
		log.Errorf("unable to list middleware: %v", err)
		return nil, status.Error(codes.Internal, "unable to list middleware")
	}

	var resp middlewarev1.ListResponse
	for _, m := range mws {
		pb, err := middlewareFromRow(&m)
		if err != nil {
			log.Errorf("unable to create middleware from row: %v", err)
			return nil, status.Error(codes.Internal, "unable to create middleware from row")
		}
		resp.Middlewares = append(resp.Middlewares, pb)
	}
	return &resp, nil
}
