package main

import (
	"github.com/suryansh74/banking/app"
	"github.com/suryansh74/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
