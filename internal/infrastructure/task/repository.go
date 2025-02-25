package task

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	domain "taskmanager.com/internal/domain/tasks"
	"taskmanager.com/internal/generated/repository"
)

type Queries interface {
	FindTaskById(context.Context, int32) (repository.Task, error)
}

type Repository struct {
	q Queries
}

func NewRepository(p *pgxpool.Pool) *Repository {
	return &Repository{q: repository.New(p)}
}

func (r Repository) Create(ctx context.Context) {

}

func (r Repository) GetById(ctx context.Context, id int32) (*domain.Task, error) {
	data, err := r.q.FindTaskById(ctx, id)
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
