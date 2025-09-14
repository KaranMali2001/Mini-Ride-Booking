-- name: GetBookingByID :one
SELECT * FROM bookings
WHERE booking_id = $1;




