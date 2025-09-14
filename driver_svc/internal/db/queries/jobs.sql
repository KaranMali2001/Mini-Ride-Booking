-- name: CreateJob :one
INSERT INTO jobs (job_id, booking_id, driver_id, job_status)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetJobByBookingId :one
SELECT * FROM jobs
WHERE booking_id = $1;

-- name: GetAllJobs :many
SELECT * FROM jobs;