-- name: CreateLog :one
INSERT INTO logs (
  service, level, message
) VALUES ( $1, $2, $3 ) RETURNING *;

-- name: ListLogs :many
SELECT * FROM logs
ORDER BY created_at DESC
LIMIT $1;
