package task

import (
	"context"
	domain "taskmanager.com/internal/domain/tasks"
)

type Repository interface {
	Create(ctx context.Context, createTask func() (*domain.Task, error)) (*domain.Task, error)
	GetById(ctx context.Context, id int) (*domain.Task, error)
	Delete(ctx context.Context, id int) error
}
