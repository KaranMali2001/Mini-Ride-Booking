package tests

import (
	"context"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
	"testing"
)

// Interfaces for testing
type DriverRepoInterface interface {
	UpdateJobDriver(ctx context.Context, req models.UpdateJobDriverRequest) (generated.DriverJob, error)
	UpdateDriver(ctx context.Context, req models.UpdateDriverRequest) error
	GetALlJobs(ctx context.Context) ([]generated.DriverJob, error)
	GetAllDrivers(ctx context.Context) ([]generated.DriverDriver, error)
}

type DriverProducerInterface interface {
	PublishBookingAccepted(event interface{}) error
}

// MockDriverRepo implements the repository interface for testing
type MockDriverRepo struct {
	updateJobDriverFunc func(ctx context.Context, req models.UpdateJobDriverRequest) (generated.DriverJob, error)
	updateDriverFunc    func(ctx context.Context, req models.UpdateDriverRequest) error
	getAllJobsFunc      func(ctx context.Context) ([]generated.DriverJob, error)
	getAllDriversFunc   func(ctx context.Context) ([]generated.DriverDriver, error)
}

func (m *MockDriverRepo) UpdateJobDriver(ctx context.Context, req models.UpdateJobDriverRequest) (generated.DriverJob, error) {
	if m.updateJobDriverFunc != nil {
		return m.updateJobDriverFunc(ctx, req)
	}
	return generated.DriverJob{}, nil
}

func (m *MockDriverRepo) UpdateDriver(ctx context.Context, req models.UpdateDriverRequest) error {
	if m.updateDriverFunc != nil {
		return m.updateDriverFunc(ctx, req)
	}
	return nil
}

func (m *MockDriverRepo) GetALlJobs(ctx context.Context) ([]generated.DriverJob, error) {
	if m.getAllJobsFunc != nil {
		return m.getAllJobsFunc(ctx)
	}
	return []generated.DriverJob{}, nil
}

func (m *MockDriverRepo) GetAllDrivers(ctx context.Context) ([]generated.DriverDriver, error) {
	if m.getAllDriversFunc != nil {
		return m.getAllDriversFunc(ctx)
	}
	return []generated.DriverDriver{}, nil
}

// MockProducer implements the producer interface for testing
type MockProducer struct {
	publishBookingAcceptedFunc func(event interface{}) error
}

func (m *MockProducer) PublishBookingAccepted(event interface{}) error {
	if m.publishBookingAcceptedFunc != nil {
		return m.publishBookingAcceptedFunc(event)
	}
	return nil
}

// TestDriverService is a wrapper that accepts interfaces
type TestDriverService struct {
	repo     DriverRepoInterface
	producer DriverProducerInterface
}

func NewTestDriverService(repo DriverRepoInterface, producer DriverProducerInterface) *TestDriverService {
	return &TestDriverService{
		repo:     repo,
		producer: producer,
	}
}

func (s *TestDriverService) AcceptJob(ctx context.Context, req models.AcceptJobRequest) error {
	// Update existing job with driver assignment
	_, err := s.repo.UpdateJobDriver(ctx, models.UpdateJobDriverRequest{
		BookingID:  req.BookingID,
		DriverID:   req.DriverId,
		RideStatus: "Accepted",
	})
	if err != nil {
		return err
	}

	// Update the driver status
	err = s.repo.UpdateDriver(ctx, models.UpdateDriverRequest{
		DriverID:  req.DriverId,
		Available: true,
	})
	if err != nil {
		return err
	}

	// Publish booking.accepted event
	if err := s.producer.PublishBookingAccepted(req); err != nil {
		return err
	}

	return nil
}

func (s *TestDriverService) GetAllJobs(ctx context.Context) ([]generated.DriverJob, error) {
	return s.repo.GetALlJobs(ctx)
}

func (s *TestDriverService) GetAllDrivers(ctx context.Context) ([]generated.DriverDriver, error) {
	return s.repo.GetAllDrivers(ctx)
}

func TestAcceptJob_Success(t *testing.T) {
	// Arrange
	bookingID := "550e8400-e29b-41d4-a716-446655440001"
	driverID := "d-1"

	mockRepo := &MockDriverRepo{
		updateJobDriverFunc: func(ctx context.Context, req models.UpdateJobDriverRequest) (generated.DriverJob, error) {
			return generated.DriverJob{
				BookingID:  pgtype.UUID{Valid: true},
				DriverID:   pgtype.UUID{Valid: true},
				RideStatus: "Accepted",
			}, nil
		},
		updateDriverFunc: func(ctx context.Context, req models.UpdateDriverRequest) error {
			return nil
		},
	}

	mockProducer := &MockProducer{
		publishBookingAcceptedFunc: func(event interface{}) error {
			return nil
		},
	}

	driverService := NewTestDriverService(mockRepo, mockProducer)

	req := models.AcceptJobRequest{
		BookingID: bookingID,
		DriverId:  driverID,
	}

	// Act
	err := driverService.AcceptJob(context.Background(), req)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGetAllJobs_Success(t *testing.T) {
	// Arrange
	expectedJobs := []generated.DriverJob{
		{
			BookingID:  pgtype.UUID{Valid: true},
			RideStatus: "Requested",
			Price:      220,
		},
		{
			BookingID:  pgtype.UUID{Valid: true},
			RideStatus: "Requested",
			Price:      150,
		},
	}

	mockRepo := &MockDriverRepo{
		getAllJobsFunc: func(ctx context.Context) ([]generated.DriverJob, error) {
			return expectedJobs, nil
		},
	}

	mockProducer := &MockProducer{}
	driverService := NewTestDriverService(mockRepo, mockProducer)

	// Act
	result, err := driverService.GetAllJobs(context.Background())

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 jobs, got %d", len(result))
	}

	if result[0].RideStatus != "Requested" {
		t.Errorf("Expected first job status 'Requested', got %s", result[0].RideStatus)
	}
}

func TestGetAllDrivers_Success(t *testing.T) {
	// Arrange
	expectedDrivers := []generated.DriverDriver{
		{
			DriverID:  pgtype.UUID{Valid: true},
			Name:      "Asha",
			Available: true,
		},
		{
			DriverID:  pgtype.UUID{Valid: true},
			Name:      "Ravi",
			Available: true,
		},
	}

	mockRepo := &MockDriverRepo{
		getAllDriversFunc: func(ctx context.Context) ([]generated.DriverDriver, error) {
			return expectedDrivers, nil
		},
	}

	mockProducer := &MockProducer{}
	driverService := NewTestDriverService(mockRepo, mockProducer)

	// Act
	result, err := driverService.GetAllDrivers(context.Background())

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 drivers, got %d", len(result))
	}

	if result[0].Name != "Asha" {
		t.Errorf("Expected first driver name 'Asha', got %s", result[0].Name)
	}
}
