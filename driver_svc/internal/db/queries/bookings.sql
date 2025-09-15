-- name: GetBookingByID :one
SELECT * FROM booking.bookings
WHERE booking_id = $1;




