
-- name: FindTaskById :one
SELECT * 
FROM tasks
WHERE id = $1; 