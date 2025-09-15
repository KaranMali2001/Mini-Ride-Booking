package repo

import (
	"context"

	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/db/generated"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jinzhu/copier"
)

type BookingRepo struct {
	q *generated.Queries
}

func NewBookingRepo(q *generated.Queries) *BookingRepo {
	return &BookingRepo{q: q}
}
func (r *BookingRepo) CreateBooking(ctx context.Context, req models.CreateBookingRequest) (generated.BookingBooking, error) {
	params := generated.CreateBookingParams{}
	copier.Copy(&params, &req)
	booking, err := r.q.CreateBooking(ctx, params)
	if err != nil {
		return generated.BookingBooking{}, err
	}
	return booking, nil
}
func (r *BookingRepo) GetBookingByID(ctx context.Context, bookingID pgtype.UUID) (generated.BookingBooking, error) {
	return r.q.GetBookingByID(ctx, bookingID)
}
