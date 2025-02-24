package task

import (
	"context"
	domain "taskmanager.com/internal/domain/tasks"
)

type Repository interface {
	Create(ctx context.Context)
	GetById(ctx context.Context, id int32) (*domain.Task, error)
}
