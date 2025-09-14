package main

import (
	"log"
	"net/http"
	"os"

	"github.com/KaranMali2001/mini-ride-booking/common/logger"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("WARNING: .env file not found")
	}

	// Logger
	isProd := os.Getenv("LOG_PROD") == "true"
	_, logger, err := logger.Init(isProd)
	if err != nil {
		log.Fatal("ERROR while starting the logger", err)
		return
	}

	// DB
	_, queries, err := config.ConnectDb()
	if err != nil {
		log.Fatal("ERROR while connecting to database", err)
		return
	}
	handler, err := bootstrapHandler(queries)
	if err != nil {
		log.Fatal("Failed to bootstrap handler:", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/jobs/:id/accept", handler.AcceptJob) // POST
	r.Get("/jobs", handler.GetAllJobs)

	logger.Info("Driver Service Started on ", config.LoadConfig().Port)
	if err := http.ListenAndServe(":"+config.LoadConfig().Port, r); err != nil {
		logger.Error("HTTP server failed", err)
	}

}
