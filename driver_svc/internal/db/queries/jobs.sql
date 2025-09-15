-- name: CreateJob :one
INSERT INTO driver.jobs (job_id, booking_id, driver_id, ride_status, pickuploc_lat, pickuploc_lng, dropoff_lat, dropoff_lng, price)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetJobByBookingId :one
SELECT * FROM driver.jobs
WHERE booking_id = $1;

-- name: GetAllJobs :many
SELECT * FROM driver.jobs;