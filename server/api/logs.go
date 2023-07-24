package api

import (
	"context"
	"errors"
	"io"
	"time"

	accesslogv3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	accesslogservicev3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	"github.com/gogo/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/apoxy-dev/proximal/core/log"
	serverdb "github.com/apoxy-dev/proximal/server/db"
	sqlc "github.com/apoxy-dev/proximal/server/db/sql"

	logsv1 "github.com/apoxy-dev/proximal/api/logs/v1"
)

// LogsService manages request logs using Envoy's access log and tap services.
type LogsService struct {
	db *serverdb.DB
}

// NewLogsService creates a new LogsService.
func NewLogsService(db *serverdb.DB) *LogsService {
	return &LogsService{
		db: db,
	}
}

func logFromRow(row sqlc.AccessLog) (*logsv1.Log, error) {
	var entry accesslogv3.HTTPAccessLogEntry
	if err := protojson.Unmarshal(row.Entry, &entry); err != nil {
		return nil, err
	}

	return &logsv1.Log{
		Id:        entry.Request.RequestId,
		Timestamp: entry.CommonProperties.StartTime,
		Http:      &entry,
	}, nil
}

// GetLogs returns a stream of logs.
func (s *LogsService) GetLogs(ctx context.Context, req *logsv1.GetLogsRequest) (*logsv1.GetLogsResponse, error) {
	var start, end time.Time
	if req.Start != nil {
		start = req.Start.AsTime()
	}
	if req.End != nil {
		end = req.End.AsTime()
	} else {
		end = time.Now()
	}

	if req.PageToken != "" {
		dTs, err := decodeNextPageToken(req.PageToken)
		if err != nil {
			log.Errorf("failed to decode page token: %v", err)
			return nil, status.Error(codes.InvalidArgument, "invalid page token")
		}

		log.Debugf("decoded page token: %s", string(dTs))

		start, err = time.Parse(time.RFC3339Nano, string(dTs))
		if err != nil {
			log.Errorf("failed to parse page token: %v", err)
			return nil, status.Error(codes.InvalidArgument, "invalid page token")
		}

	}
	pageSize := req.PageSize
	if pageSize == 0 {
		pageSize = 100
	}

	logs, err := s.db.Queries().GetAccessLogsByStartAndEnd(ctx, sqlc.GetAccessLogsByStartAndEndParams{
		StartTime: start.UnixNano(),
		EndTime:   end.UnixNano(),
		Limit:     int64(pageSize),
	})
	if err != nil {
		return nil, err
	}

	logpbs := make([]*logsv1.Log, 0, len(logs))
	for _, l := range logs {
		logpb, err := logFromRow(l)
		if err != nil {
			return nil, err
		}
		logpbs = append(logpbs, logpb)
	}

	var nextPageToken string
	if len(logpbs) == int(pageSize) {
		nextPageToken = encodeNextPageToken(logpbs[len(logpbs)-1].Timestamp.AsTime().Format(time.RFC3339Nano))
	}

	return &logsv1.GetLogsResponse{
		Logs:          logpbs,
		NextPageToken: nextPageToken,
	}, nil
}

func (s *LogsService) GetFullLog(ctx context.Context, req *logsv1.GetFullLogRequest) (*logsv1.FullLog, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *LogsService) storeHTTPLogEntry(ctx context.Context, entry *accesslogv3.HTTPAccessLogEntry) error {
	if entry.Request.RequestId == "" {
		return errors.New("missing request ID")
	}
	entryJson, err := protojson.Marshal(entry)
	if err != nil {
		return err
	}
	return s.db.Queries().AddAccessLog(ctx, sqlc.AddAccessLogParams{
		RequestID: entry.Request.RequestId,
		StartTime: entry.CommonProperties.StartTime.AsTime().UnixNano(),
		EndTime:   entry.CommonProperties.StartTime.AsTime().Add(entry.CommonProperties.Duration.AsDuration()).UnixNano(),
		Entry:     entryJson,
	})
}

func (s *LogsService) StreamAccessLogs(stream accesslogservicev3.AccessLogService_StreamAccessLogsServer) error {
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}

		if req.Identifier != nil {
			log.Infof("Initiating access log stream %v: %s/%s", req.Identifier.LogName, req.Identifier.Node.Cluster, req.Identifier.Node.Id)
		}

		switch entries := req.LogEntries.(type) {
		case *accesslogservicev3.StreamAccessLogsMessage_HttpLogs:
			for _, entry := range entries.HttpLogs.LogEntry {
				log.Debugf("Received HTTP access log entry: %v", entry)

				if entry == nil {
					continue
				}
				common := entry.CommonProperties
				if common == nil {
					log.Warnf("Received HTTP access log entry without common properties: %v", entry)
					continue
				}
				req := entry.Request
				if req == nil {
					log.Warnf("Received HTTP access log entry without request: %v", entry)
					continue
				}
				resp := entry.Response
				if resp == nil {
					log.Warnf("Received HTTP access log entry without response: %v", entry)
					continue
				}

				log.Debugf("%s: %v -> %v %s %s %d id=%s",
					entry.CommonProperties.StartTime.AsTime().Format(time.RFC3339),
					common.DownstreamRemoteAddress,
					common.UpstreamRemoteAddress,
					req.RequestMethod.String(),
					req.Path,
					resp.ResponseCode.GetValue(),
					req.RequestId,
				)

				if err := s.storeHTTPLogEntry(stream.Context(), entry); err != nil {
					log.Errorf("Failed to store log entry: %v", err)
				}
			}
		default:
			log.Debugf("Received unknown access log entry: %v", entries)
		}

	}
	return nil
}

func (s *LogsService) RegisterALS(srv *grpc.Server) {
	accesslogservicev3.RegisterAccessLogServiceServer(srv, s)
}
