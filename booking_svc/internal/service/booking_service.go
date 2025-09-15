package service

import (
	"context"

	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/models"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/repo"
)

type BookingService struct {
	repo *repo.BookingRepo
}

func NewBookingService(repo *repo.BookingRepo) *BookingService {
	return &BookingService{repo: repo}
}
func (s *BookingService) CreateBooking(ctx context.Context, req models.CreateBookingRequest) (generated.BookingBooking, error) {
	return s.repo.CreateBooking(ctx, req)
}
