-- name: AddAccessLog :exec
INSERT INTO access_logs (request_id, start_time, end_time, entry)
VALUES (?, ?, ?, ?);

-- name: GetAccessLogsByStartAndEnd :many
SELECT *
FROM access_logs
WHERE start_time >= ? AND end_time <= ?
ORDER BY start_time DESC
LIMIT ?;

-- name: GetAccessLogsByRequestID :one
SELECT *
FROM access_logs
WHERE request_id = ?;

-- name: AddTrace :exec	
INSERT INTO traces (request_id, trace)
VALUES (?, ?);

-- name: GetTraceByRequestID :one
SELECT *
FROM traces
WHERE request_id = ?;
