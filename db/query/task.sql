
-- name: FindTaskById :one
SELECT * 
FROM tasks
WHERE id = $1;



-- name: CreateTask :one
INSERT INTO tasks ("name", "desc", "create_date") VALUES ($1, $2, $3)
    RETURNING  id;


-- name: CheckExistsTask :one
SELECT EXISTS (SELECT 1 FROM tasks WHERE "name" = $1 AND "create_date" = $2);


-- name: DeleteTask :exec
DELETE  FROM tasks WHERE id = $1;


-- name: UpdateTask :one
UPDATE tasks
SET
    "name" = $2,
    "desc" = $3
WHERE id = $1
RETURNING *
;
