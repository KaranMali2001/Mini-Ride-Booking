-- name: GetBookingByID :one
SELECT * FROM booking.bookings
WHERE booking_id = $1;

-- name: CreateBooking :one
INSERT INTO booking.bookings (
     pickuploc_lat, pickuploc_lng, dropoff_lat, dropoff_lng, price, ride_status, driver_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetAllBookings :many
SELECT * FROM booking.bookings
ORDER BY created_at DESC;

-- name: UpdateBookingStatus :exec
UPDATE booking.bookings
SET ride_status = $2, driver_id = $3
WHERE booking_id = $1;