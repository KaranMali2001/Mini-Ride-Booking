-- +goose Up
-- Create drivers table
CREATE TABLE IF NOT EXISTS drivers (
    driver_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    available BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Optional: jobs table to store consumed booking events
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

-- +goose Down
DROP TABLE IF EXISTS jobs;
DROP TABLE IF EXISTS drivers;
