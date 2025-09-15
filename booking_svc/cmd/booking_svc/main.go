package main

import (
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/config"
	"github.com/KaranMali2001/mini-ride-booking/common/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting Booking Service...")

	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println(" WARNING: .env file not found, using defaults")
	} else {

	}

	isProd := os.Getenv("LOG_PROD") == "true"
	_, logger, err := logger.Init(isProd)
	if err != nil {
		log.Fatal("FATAL: Error while starting the logger:", err)
		return
	}
	logger.Info("Logger initialized successfully")

	_, queries, err := config.ConnectDb()
	if err != nil {
		log.Fatal("FATAL: Error while connecting to database:", err)
		return
	}
	logger.Info("Database connected successfully")

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

	r.Post("/bookings", handler.CreateBooking)
	r.Get("/bookings", handler.GetAllBookings)
	logger.Info("Routes configured")

	port := config.LoadConfig().Port
	if port == "" {
		port = "8080"
	}

	logger.Info("Booking Service Started successfully on port:", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		logger.Error("HTTP server failed:", err)
		return
	}
}
