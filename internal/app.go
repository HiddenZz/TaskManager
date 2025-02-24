package internal

import (
	"fmt"
	"net/http"
	"os"
	"taskmanager.com/config"
	"taskmanager.com/internal/app"
	taskImpl "taskmanager.com/internal/infrastructure/task"
	"taskmanager.com/pkg/db"
)

func Run(config *config.Config) {
	database := db.DB{}
	err := database.Connect(config.BuildConnectionString(map[string]string{}))
	defer func(database *db.DB) {
		_ = database.Close()
	}(&database)
	if err != nil {
		fmt.Printf("run d.Connect: complete with error %v", err)
		os.Exit(1)
	}

	taskRepository := taskImpl.NewRepository(database.Pool)
	httpServer := app.SetupHttpServer(taskRepository)

	server := http.Server{
		Handler: httpServer,
		Addr:    "localhost:8080",
	}

	server.ListenAndServe()
}
