package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/suryansh74/banking/app"
	"github.com/suryansh74/banking/logger"
)

func main() {
	err := godotenv.Load() // loads .env from current directory
	if err != nil {
		log.Println("Warning: .env file not found or failed to load")
	}

	logger.Info("Starting the application...")
	app.Start()
}
