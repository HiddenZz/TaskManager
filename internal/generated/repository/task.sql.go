// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: task.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkExistsTask = `-- name: CheckExistsTask :one
SELECT EXISTS (SELECT 1 FROM tasks WHERE "name" = $1 AND "create_date" = $2)
`

type CheckExistsTaskParams struct {
	Name       string
	CreateDate pgtype.Timestamp
}

func (q *Queries) CheckExistsTask(ctx context.Context, arg CheckExistsTaskParams) (bool, error) {
	row := q.db.QueryRow(ctx, checkExistsTask, arg.Name, arg.CreateDate)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createTask = `-- name: CreateTask :one
INSERT INTO tasks ("name", "desc", "create_date") VALUES ($1, $2, $3)
    RETURNING  id
`

type CreateTaskParams struct {
	Name       string
	Desc       pgtype.Text
	CreateDate pgtype.Timestamp
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (int32, error) {
	row := q.db.QueryRow(ctx, createTask, arg.Name, arg.Desc, arg.CreateDate)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const findTaskById = `-- name: FindTaskById :one
SELECT id, name, "desc", create_date 
FROM tasks
WHERE id = $1
`

func (q *Queries) FindTaskById(ctx context.Context, id int32) (Task, error) {
	row := q.db.QueryRow(ctx, findTaskById, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Desc,
		&i.CreateDate,
	)
	return i, err
}
