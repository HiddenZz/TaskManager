package task

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"taskmanager.com/helpers/types"
	domain "taskmanager.com/internal/domain/tasks"
	"taskmanager.com/internal/generated/repository"
)

type Queries interface {
	FindTaskById(context.Context, int32) (repository.Task, error)
	CreateTask(context.Context, repository.CreateTaskParams) (int32, error)
	CheckExistsTask(context.Context, repository.CheckExistsTaskParams) (bool, error)
	DeleteTask(context.Context, int32) error
}

type Repository struct {
	q Queries
}

func NewRepository(p *pgxpool.Pool) *Repository {
	return &Repository{q: repository.New(p)}
}

func (r Repository) Delete(ctx context.Context, id int) error {
	err := r.q.DeleteTask(ctx, int32(id))
	if err != nil {
		return fmt.Errorf("error during deletion")
	}
	return err
}

func (r Repository) Create(ctx context.Context, createTask func() (*domain.Task, error)) (*domain.Task, error) {
	task, err := createTask()
	if err != nil {
		return nil, fmt.Errorf("error when create Task %v", err)
	}

	taskExists, err := r.q.CheckExistsTask(ctx, repository.CheckExistsTaskParams{
		Name:       task.Name(),
		CreateDate: types.Timestamp(task.CreateDate()),
	})
	if err != nil {
		return nil, fmt.Errorf("error during check exist task %v", err)
	}
	if taskExists {
		return nil, fmt.Errorf("task exist %v", task)
	}

	data, err := r.q.CreateTask(ctx, repository.CreateTaskParams{
		Name:       task.Name(),
		Desc:       types.Text(task.Desc()),
		CreateDate: types.Timestamp(task.CreateDate()),
	})

	if err != nil {
		return nil, fmt.Errorf("error during insert new task:  %v %v", err, task)
	}
	task, err = domain.NewTask(int(data), task.Name(), task.Desc(), task.CreateDate())
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r Repository) GetById(ctx context.Context, id int) (*domain.Task, error) {
	data, err := r.q.FindTaskById(ctx, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("task not found")
		}
		return nil, fmt.Errorf("failed to select by id %d, %v", id, err)
	}

	task, err := domain.NewTask(int(data.ID), data.Name, data.Desc.String, data.CreateDate.Time)
	if err != nil {
		return nil, err
	}
	return task, nil
}
