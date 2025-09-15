
-- +goose Up
CREATE SCHEMA IF NOT EXISTS driver;
CREATE TABLE IF NOT EXISTS driver.drivers (
    driver_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    available BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS driver.jobs (
    job_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id UUID NOT NULL,
    driver_id UUID NULL,
    ride_status TEXT NOT NULL CHECK (ride_status IN ('Requested','Accepted')),
    pickuploc_lat DOUBLE PRECISION NOT NULL,
    pickuploc_lng DOUBLE PRECISION NOT NULL,
    dropoff_lat DOUBLE PRECISION NOT NULL,
    dropoff_lng DOUBLE PRECISION NOT NULL,
    price INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT unique_booking_id UNIQUE (booking_id)
);

-- +goose Down
DROP TABLE IF EXISTS driver.jobs;
DROP TABLE IF EXISTS driver.drivers;
