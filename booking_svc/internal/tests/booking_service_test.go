package tests

import (
	"context"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
	"testing"
)

// Interfaces for testing
type BookingRepoInterface interface {
	CreateBooking(ctx context.Context, req models.CreateBookingRequest) (generated.BookingBooking, error)
	GetAllBookings(ctx context.Context) ([]generated.BookingBooking, error)
}

type ProducerInterface interface {
	PublishBookingCreated(event interface{}) error
}

// MockRepo implements the repository interface for testing
type MockBookingRepo struct {
	createBookingFunc  func(ctx context.Context, req models.CreateBookingRequest) (generated.BookingBooking, error)
	getAllBookingsFunc func(ctx context.Context) ([]generated.BookingBooking, error)
}

func (m *MockBookingRepo) CreateBooking(ctx context.Context, req models.CreateBookingRequest) (generated.BookingBooking, error) {
	if m.createBookingFunc != nil {
		return m.createBookingFunc(ctx, req)
	}
	return generated.BookingBooking{}, nil
}

func (m *MockBookingRepo) GetAllBookings(ctx context.Context) ([]generated.BookingBooking, error) {
	if m.getAllBookingsFunc != nil {
		return m.getAllBookingsFunc(ctx)
	}
	return []generated.BookingBooking{}, nil
}

// MockProducer implements the producer interface for testing
type MockProducer struct {
	publishBookingCreatedFunc func(event interface{}) error
}

func (m *MockProducer) PublishBookingCreated(event interface{}) error {
	if m.publishBookingCreatedFunc != nil {
		return m.publishBookingCreatedFunc(event)
	}
	return nil
}

// TestBookingService is a wrapper that accepts interfaces
type TestBookingService struct {
	repo     BookingRepoInterface
	producer ProducerInterface
}

func NewTestBookingService(repo BookingRepoInterface, producer ProducerInterface) *TestBookingService {
	return &TestBookingService{
		repo:     repo,
		producer: producer,
	}
}

func (s *TestBookingService) CreateBooking(ctx context.Context, req models.CreateBookingRequest) (generated.BookingBooking, error) {
	booking, err := s.repo.CreateBooking(ctx, req)
	if err != nil {
		return booking, err
	}

	// Simulate publishing event
	if err := s.producer.PublishBookingCreated(booking); err != nil {
		return booking, err
	}

	return booking, nil
}

func (s *TestBookingService) GetAllBookings(ctx context.Context) ([]generated.BookingBooking, error) {
	return s.repo.GetAllBookings(ctx)
}

func TestCreateBooking_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockBookingRepo{
		createBookingFunc: func(ctx context.Context, req models.CreateBookingRequest) (generated.BookingBooking, error) {
			return generated.BookingBooking{
				BookingID:    pgtype.UUID{Valid: true},
				PickuplocLat: req.PickupLat,
				PickuplocLng: req.PickupLng,
				DropoffLat:   req.DropoffLat,
				DropoffLng:   req.DropoffLng,
				Price:        int32(req.Price),
				RideStatus:   "Requested",
			}, nil
		},
	}

	mockProducer := &MockProducer{
		publishBookingCreatedFunc: func(event interface{}) error {
			return nil
		},
	}

	bookingService := NewTestBookingService(mockRepo, mockProducer)

	req := models.CreateBookingRequest{
		PickupLat:  12.9,
		PickupLng:  77.6,
		DropoffLat: 12.95,
		DropoffLng: 77.64,
		Price:      220,
		RideStatus: "Requested",
	}

	// Act
	result, err := bookingService.CreateBooking(context.Background(), req)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.RideStatus != "Requested" {
		t.Errorf("Expected RideStatus 'Requested', got %s", result.RideStatus)
	}

	if result.Price != 220 {
		t.Errorf("Expected Price 220, got %d", result.Price)
	}
}

func TestGetAllBookings_Success(t *testing.T) {
	// Arrange
	expectedBookings := []generated.BookingBooking{
		{
			BookingID:  pgtype.UUID{Valid: true},
			RideStatus: "Requested",
			Price:      220,
		},
		{
			BookingID:  pgtype.UUID{Valid: true},
			RideStatus: "Accepted",
			Price:      150,
		},
	}

	mockRepo := &MockBookingRepo{
		getAllBookingsFunc: func(ctx context.Context) ([]generated.BookingBooking, error) {
			return expectedBookings, nil
		},
	}

	mockProducer := &MockProducer{}
	bookingService := NewTestBookingService(mockRepo, mockProducer)

	// Act
	result, err := bookingService.GetAllBookings(context.Background())

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 bookings, got %d", len(result))
	}

	if result[0].RideStatus != "Requested" {
		t.Errorf("Expected first booking status 'Requested', got %s", result[0].RideStatus)
	}
}
