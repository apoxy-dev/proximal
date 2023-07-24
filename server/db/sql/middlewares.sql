-- name: CreateMiddleware :one
INSERT INTO middlewares (
  slug,
  source_type,
  ingest_params_json,
  runtime_params_json,
  status,
  status_detail
) VALUES (
  ?,
  ?,
  ?,
  ?,
  ?,
  ?
)
RETURNING *;

-- name: GetMiddlewareBySlug :one
SELECT *
FROM middlewares
WHERE slug = ?
LIMIT 1;

-- name: UpdateMiddleware :one
UPDATE middlewares
SET
  ingest_params_json = ?,
  runtime_params_json = ?,
  updated_at = STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')
WHERE slug = ?
RETURNING *;

-- name: UpdateMiddlewareStatus :one
UPDATE middlewares
SET
  status = ?,
  status_detail = ?,
  live_build_sha = ?,
  updated_at = STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')
WHERE slug = ?
RETURNING *;

-- name: DeleteMiddleware :exec
DELETE FROM middlewares
WHERE slug = ?;

-- name: ListMiddlewares :many
SELECT *
FROM middlewares
WHERE datetime(created_at) < datetime(?)
ORDER BY created_at DESC
LIMIT ?;

-- name: ListMiddlewaresAll :many
SELECT *
FROM middlewares
ORDER BY created_at DESC;

-- name: ListMiddlewaresBySlugs :many
SELECT *
FROM middlewares
WHERE slug IN (sqlc.slice('slugs'))
ORDER BY created_at DESC;

-- name: CreateBuild :one
INSERT INTO builds (
  sha,
  middleware_slug,
  status,
  status_detail
) VALUES (
  ?,
  ?,
  ?,
  ?
)
RETURNING *;

-- name: ListBuildsByMiddlewareSlug :many
SELECT *
FROM builds
WHERE middleware_slug = ? AND deleted_at IS NULL AND datetime(started_at) < datetime(?)
ORDER BY started_at DESC
LIMIT ?;

-- name: GetBuildByMiddlewareSlugAndSha :one
SELECT *
FROM builds
WHERE middleware_slug = ? AND sha = ? AND deleted_at IS NULL
LIMIT 1;

-- name: GetLiveReadyBuildByMiddlewareSlug :one
SELECT *
FROM builds
WHERE middleware_slug IN (
  SELECT live_build_sha
  FROM middlewares
  WHERE slug = ? AND deleted_at IS NULL
) AND status = 'READY' AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateBuildStatus :one
UPDATE builds
SET
  status = ?,
  status_detail = ?,
  output_path = ?,
  updated_at = STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')
WHERE middleware_slug = ? AND sha = ?
RETURNING *;

-- name: DeleteBuild :exec
UPDATE builds
SET deleted_at = STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')
WHERE middleware_slug = ? AND sha = ?
