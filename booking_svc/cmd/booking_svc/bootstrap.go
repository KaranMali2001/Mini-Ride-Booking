package main

import (
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/handlers"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/repo"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/service"
)

// bootstrapHandler initializes logger, DB, repo, service, and handler
func bootstrapHandler(queries *generated.Queries) (*handlers.BookingHandler, error) {

	// Repo -> Service -> Handler
	repo := repo.NewBookingRepo(queries)
	svc := service.NewBookingService(repo)
	handler := handlers.NewBookingHandler(svc)

	return handler, nil
}
