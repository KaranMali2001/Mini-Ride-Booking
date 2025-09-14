package service

import (
	"context"

	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/models"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/repo"
)

type DriverService struct {
	repo *repo.DriverRepo
}

func NewDriverService(repo *repo.DriverRepo) *DriverService {
	return &DriverService{repo: repo}
}
func (s *DriverService) AcceptJob(ctx context.Context, req models.AcceptJobRequest) error {
	//first parse both bookingid and jobId to uuid

	//try creating new job with given bookingId and driverId , it will fail if job is already created by someone because we have unique constraint on booking_id
	_, err := s.repo.CreateJob(ctx, models.CreateJobRequest{
		BookingID: req.BookingID,
		DriverID:  req.DriverId,
		JobStatus: "pending",
	})
	if err != nil {
		return err
	}
	//update the driver status
	err = s.repo.UpdateDriver(ctx, models.UpdateDriverRequest{
		DriverID:  req.DriverId,
		JobStatus: "accepted",
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *DriverService) GetAllJobs(ctx context.Context) ([]generated.Job, error) {
	return s.repo.GetALlJobs(ctx)
}
