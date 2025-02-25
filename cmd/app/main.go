package main

import (
	"taskmanager.com/config"
	"taskmanager.com/internal"
)

func main() {
	internal.Run(config.Create())

}
