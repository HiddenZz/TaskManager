package app

import (
	"fmt"
	"os"
	"taskmanager.com/config"
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
}
