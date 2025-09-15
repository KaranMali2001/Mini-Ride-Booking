package repo

import (
	"context"

	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/models"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/jinzhu/copier"
)

type DriverRepo struct {
	q *generated.Queries
}

func NewDriverRepo(q *generated.Queries) *DriverRepo {
	return &DriverRepo{q: q}
}
func (r *DriverRepo) CreateJob(ctx context.Context, req models.CreateJobRequest) (generated.DriverJob, error) {
	params := generated.CreateJobParams{}
	copier.Copy(&params, &req)
	job, err := r.q.CreateJob(ctx, params)
	if err != nil {
		return generated.DriverJob{}, err
	}
	return job, nil
}

//	func (r *DriverRepo) UpdateDriver(ctx context.Context, req models.UpdateDriverRequest) error {
//		params := generated.UpdateDriverParams{}
//		copier.Copy(&params, &req)
//		return r.q.UpdateDriver(ctx, params)
//	}
func (r *DriverRepo) GetJobByBookingId(ctx context.Context, bookingID pgtype.UUID) (generated.DriverJob, error) {
	return r.q.GetJobByBookingId(ctx, bookingID)
}
func (r *DriverRepo) GetALlJobs(ctx context.Context) ([]generated.DriverJob, error) {
	return r.q.GetAllJobs(ctx)
}

func (r *DriverRepo) GetAllDrivers(ctx context.Context) ([]generated.DriverDriver, error) {
	return r.q.GetDrivers(ctx)
}

func (r *DriverRepo) UpdateJobDriver(ctx context.Context, req models.UpdateJobDriverRequest) (generated.DriverJob, error) {
	params := generated.UpdateJobDriverParams{}
	copier.Copy(&params, &req)
	return r.q.UpdateJobDriver(ctx, params)
}
func (r *DriverRepo) UpdateDriverStatus(ctx context.Context, req models.UpdateDriverRequest) error {
	params := generated.UpdateDriverStatusParams{}
	copier.Copy(&params, &req)
	return r.q.UpdateDriverStatus(ctx, params)
}
