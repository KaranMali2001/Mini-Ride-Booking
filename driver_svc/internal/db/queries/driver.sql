
-- name: UpdateDriver :exec
UPDATE driver.jobs
SET driver_id = $1, ride_status = $2
WHERE booking_id = $3;

-- name: GetDrivers :many
SELECT * FROM driver.drivers;

-- name: GetDriverByID :one
SELECT * FROM driver.drivers
WHERE driver_id = $1;
