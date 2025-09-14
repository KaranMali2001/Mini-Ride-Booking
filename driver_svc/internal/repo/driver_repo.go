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
func (r *DriverRepo) CreateJob(ctx context.Context, req models.CreateJobRequest) (generated.Job, error) {
	params := generated.CreateJobParams{}
	copier.Copy(&params, &req)
	job, err := r.q.CreateJob(ctx, params)
	if err != nil {
		return generated.Job{}, err
	}
	return job, nil
}
func (r *DriverRepo) UpdateDriver(ctx context.Context, req models.UpdateDriverRequest) error {
	params := generated.UpdateDriverParams{}
	copier.Copy(&params, &req)
	return r.q.UpdateDriver(ctx, params)
}
func (r *DriverRepo) GetJobByBookingId(ctx context.Context, bookingID pgtype.UUID) (generated.Job, error) {
	return r.q.GetJobByBookingId(ctx, bookingID)
}
func (r *DriverRepo) GetALlJobs(ctx context.Context) ([]generated.Job, error) {
	return r.q.GetAllJobs(ctx)
}
