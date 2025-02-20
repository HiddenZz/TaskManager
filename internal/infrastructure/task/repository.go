package task

import "context"

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r Repository) Create(ctx context.Context) {

}

func (r Repository) GetById(ctx context.Context) {

}
