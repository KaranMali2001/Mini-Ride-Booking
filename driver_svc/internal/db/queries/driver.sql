
-- name: UpdateDriver :exec
UPDATE jobs
SET driver_id = $1, job_status = $2
WHERE booking_id = $3;

-- name: GetDrivers :many
SELECT * FROM drivers;

-- name: GetDriverByID :one
SELECT * FROM drivers
WHERE driver_id = $1;
