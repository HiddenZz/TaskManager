package task

import "context"

type Repository interface {
	Create(ctx context.Context)
	GetById(ctx context.Context)
}
