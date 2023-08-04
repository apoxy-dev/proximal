package db

import (
	middlewarev1 "github.com/apoxy-dev/proximal/api/middleware/v1"
)

var (
	SourceTypeDirect = middlewarev1.MiddlewareIngestParams_DIRECT.String()
	SourceTypeGitHub = middlewarev1.MiddlewareIngestParams_GITHUB.String()
)

var (
	MiddlewareStatusUnknown      = middlewarev1.Middleware_UNKNOWN.String()
	MiddlewareStatusPending      = middlewarev1.Middleware_PENDING.String()
	MiddlewareStatusPendingReady = middlewarev1.Middleware_PENDING_READY.String()
	MiddlewareStatusReady        = middlewarev1.Middleware_READY.String()
	MiddlewareStatusErrored      = middlewarev1.Middleware_ERRORED.String()
)
