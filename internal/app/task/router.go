package task

import (
	"net/http"
)

type Handler struct {
	repository Repository
}

func NewHandler(repository Repository) Handler {
	return Handler{repository: repository}
}

func NewRouter(router *http.ServeMux, repository Repository) {

	handler := NewHandler(repository)
	router.HandleFunc("GET /tasks/{id}", handler.GetById)
	router.HandleFunc("POST /tasks/", handler.Create)
	router.HandleFunc("DELETE /tasks/{id}", handler.Delete)
}
