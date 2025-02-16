package main

import (
	"taskmanager.com/config"
	"taskmanager.com/internal/app"
)

func main() {
	app.Run(config.Create())
}
