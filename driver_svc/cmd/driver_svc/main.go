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
	log.Println("Starting Driver Service...")

	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("WARNING: .env file not found, using defaults")

	}

	isProd := os.Getenv("LOG_PROD") == "true"
	_, logger, err := logger.Init(isProd)
	if err != nil {
		log.Fatal("FATAL: Error while starting the logger:", err)
		return
	}
	logger.Info("Logger initialized successfully")

	// DB
	logger.Info("Connecting to database...")
	_, queries, err := config.ConnectDb()
	if err != nil {
		log.Fatal("FATAL: Error while connecting to database:", err)
		return
	}
	logger.Info("Database connected successfully")

	// Bootstrap handler with RabbitMQ
	logger.Info("Bootstrapping handler with RabbitMQ...")
	handler, err := bootstrapHandler(queries)
	if err != nil {
		log.Fatal("FATAL: Failed to bootstrap handler:", err)
		return
	}
	logger.Info("Handler bootstrapped successfully")

	// Setup routes
	logger.Info("Setting up routes...")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/jobs/{id}/accept", handler.AcceptJob)
	r.Get("/jobs", handler.GetAllJobs)
	r.Get("/drivers", handler.GetAllDrivers)
	log.Println("Routes configured")

	port := config.LoadConfig().Port
	if port == "" {
		port = "8081"
	}

	logger.Info("Driver Service Started successfully on port:", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		logger.Error("HTTP server failed:", err)
		return
	}
}
