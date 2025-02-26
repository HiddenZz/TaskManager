package app

import (
	"net/http"
	"taskmanager.com/internal/app/task"
)

func SetupHttpServer(repository task.Repository) *http.ServeMux {
	mux := http.NewServeMux()
	task.NewRouter(mux, repository)
	return mux
}
