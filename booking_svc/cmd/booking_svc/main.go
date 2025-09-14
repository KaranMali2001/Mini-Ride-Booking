package main

import (
	"log"

	"github.com/KaranMali2001/mini-ride-booking/common/logger"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("WARNING: .env file not found")
		os.Exit(1)
	}

	isProd := os.Getenv("LOG_PROD") == "true"

	_, logger, err := logger.Init(isProd)
	if err != nil {
		log.Fatal("ERROR while starting the logger", err)
		return
	}

	logger.Info("Booking Service Started")

}
