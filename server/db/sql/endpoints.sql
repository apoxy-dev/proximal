-- name: CreateEndpoint :one
INSERT INTO endpoints (
  cluster,
  is_domain,
  use_tls,
  lookup_family
) VALUES (
  ?,
  ?,
  ?,
  ?
)
RETURNING *;

-- name: CountEndpoints :one
SELECT COUNT(*) FROM endpoints;

-- name: GetEndpointByCluster :one
SELECT *
FROM endpoints
WHERE cluster = ?
LIMIT 1;

-- name: ListEndpoints :many
SELECT * FROM endpoints;

-- name: UpdateEndpoint :one
UPDATE endpoints
SET
  is_domain = ?,
  use_tls = ?,
  lookup_family = ?,
  updated_at = STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')
WHERE cluster = ?
RETURNING *;

-- name: DeleteEndpoint :exec
DELETE FROM endpoints
WHERE cluster = ?;

-- name: GetDefaultUpstream :one
SELECT *
FROM endpoints
WHERE cluster IN (
  SELECT cluster
  FROM endpoint_default
  LIMIT 1
)
LIMIT 1;

-- name: InitDefaultUpstream :exec
INSERT INTO endpoint_default (cluster)
VALUES (?);

-- name: SetDefaultUpstream :exec
UPDATE endpoint_default
SET
  cluster = ?,
  updated_at = STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW');

-- name: CreateEndpointAddress :one
INSERT INTO endpoint_addresses (
  cluster,
  host,
  port
) VALUES (
  ?,
  ?,
  ?
)
RETURNING *;

-- name: GetEndpointAddressesByCluster :many
SELECT *
FROM endpoint_addresses
WHERE cluster = ?;

-- name: ListEndpointAddresses :many
SELECT *
FROM endpoint_addresses;

-- name: DeleteEndpointAddress :exec
DELETE FROM endpoint_addresses
WHERE cluster = ? AND host = ? AND port = ?;
