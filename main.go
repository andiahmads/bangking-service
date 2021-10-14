package main

import (
	"github.com/andiahmads/bangking-service/app"
	"github.com/andiahmads/bangking-service/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
