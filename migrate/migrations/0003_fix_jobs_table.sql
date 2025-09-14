-- +goose Up
-- Drop the existing jobs table and recreate with proper foreign keys
DROP TABLE IF EXISTS jobs;

CREATE TABLE IF NOT EXISTS jobs (
    job_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL,
    driver_id UUID NULL,
    job_status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    FOREIGN KEY (booking_id) REFERENCES bookings(booking_id) ON DELETE CASCADE,
    FOREIGN KEY (driver_id) REFERENCES drivers(driver_id) ON DELETE SET NULL
);
ALTER TABLE jobs ADD CONSTRAINT unique_booking_id UNIQUE (booking_id);
-- +goose Down
-- Recreate the old jobs table structure
DROP TABLE IF EXISTS jobs;

CREATE TABLE IF NOT EXISTS jobs (
    booking_id UUID PRIMARY KEY,
    driver_id UUID NULL,
    ride_status TEXT NOT NULL,
    pickuploc_lat DOUBLE PRECISION NOT NULL,
    pickuploc_lng DOUBLE PRECISION NOT NULL,
    dropoff_lat DOUBLE PRECISION NOT NULL,
    dropoff_lng DOUBLE PRECISION NOT NULL,
    price INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);