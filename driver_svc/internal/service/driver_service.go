package service

import (
	"context"
	"fmt"

	"github.com/KaranMali2001/mini-ride-booking/common/logger"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/models"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/mq"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/repo"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type DriverService struct {
	repo     *repo.DriverRepo
	producer *mq.Producer
}

func NewDriverService(repo *repo.DriverRepo, producer *mq.Producer) *DriverService {
	return &DriverService{
		repo:     repo,
		producer: producer,
	}
}

func (s *DriverService) AcceptJob(ctx context.Context, req models.AcceptJobRequest) error {

	bookingUUID, err := uuid.Parse(req.BookingID)
	if err != nil {
		logger.Errorln("error while parsing the booking id", err)
	}

	pgUUID := pgtype.UUID{}
	pgUUID.Bytes = bookingUUID
	pgUUID.Valid = true

	job, err := s.repo.GetJobByBookingId(ctx, pgUUID)
	if err != nil {
		return err
	}

	if job.DriverID.Valid {
		return fmt.Errorf("driver already assigned to this job")
	}
	// Update existing job with driver assignment
	_, err = s.repo.UpdateJobDriver(ctx, models.UpdateJobDriverRequest{
		BookingID:  req.BookingID,
		DriverID:   req.DriverId,
		RideStatus: "Accepted",
	})
	if err != nil {
		return err
	}

	err = s.repo.UpdateDriverStatus(ctx, models.UpdateDriverRequest{
		DriverID:  req.DriverId,
		Available: false,
	})
	if err != nil {
		return err
	}
	//update Driver status

	// Publish booking.accepted event
	event := mq.BookingAcceptedEvent{
		BookingID:  req.BookingID,
		DriverID:   req.DriverId,
		RideStatus: "Accepted",
	}

	if err := s.producer.PublishBookingAccepted(event); err != nil {
		logger.Errorln("Error while publishing booking accepted event:", err)
		return err
	}

	return nil
}

func (s *DriverService) GetAllJobs(ctx context.Context) ([]generated.DriverJob, error) {
	return s.repo.GetALlJobs(ctx)
}

func (s *DriverService) GetAllDrivers(ctx context.Context) ([]generated.DriverDriver, error) {
	return s.repo.GetAllDrivers(ctx)
}
