-- name: GetBookingByID :one
SELECT * FROM booking.bookings
WHERE booking_id = $1;

-- name: CreateBooking :one
INSERT INTO booking.bookings (
    booking_id, pickuploc_lat, pickuploc_lng, dropoff_lat, dropoff_lng, price, ride_status, driver_id, created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;