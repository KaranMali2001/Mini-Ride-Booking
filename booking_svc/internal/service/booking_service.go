package service

import (
	"context"

	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/models"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/mq"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/repo"
	"github.com/KaranMali2001/mini-ride-booking/common/logger"
)

type BookingService struct {
	repo     *repo.BookingRepo
	producer *mq.Producer
}

func NewBookingService(repo *repo.BookingRepo, producer *mq.Producer) *BookingService {
	return &BookingService{
		repo:     repo,
		producer: producer,
	}
}

func (s *BookingService) CreateBooking(ctx context.Context, req models.CreateBookingRequest) (generated.BookingBooking, error) {
	booking, err := s.repo.CreateBooking(ctx, req)
	if err != nil {
		return booking, err
	}

	// Publish booking.created event
	bookingIDStr := booking.BookingID.String()
	event := mq.BookingCreatedEvent{
		BookingID: bookingIDStr,
		PickupLoc: mq.LatLng{
			Lat: booking.PickuplocLat,
			Lng: booking.PickuplocLng,
		},
		Dropoff: mq.LatLng{
			Lat: booking.DropoffLat,
			Lng: booking.DropoffLng,
		},
		Price:      float64(booking.Price),
		RideStatus: booking.RideStatus,
	}

	if err := s.producer.PublishBookingCreated(event); err != nil {
		logger.Errorln("Failed to publish booking.created event:", err)
		return booking, err
	}

	return booking, nil
}

func (s *BookingService) GetAllBookings(ctx context.Context) ([]generated.BookingBooking, error) {
	return s.repo.GetAllBookings(ctx)
}
