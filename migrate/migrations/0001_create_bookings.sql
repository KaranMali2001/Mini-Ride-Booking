-- +goose Up
CREATE TABLE IF NOT EXISTS bookings (
    booking_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pickuploc_lat DOUBLE PRECISION NOT NULL,
    pickuploc_lng DOUBLE PRECISION NOT NULL,
    dropoff_lat DOUBLE PRECISION NOT NULL,
    dropoff_lng DOUBLE PRECISION NOT NULL,
    price INT NOT NULL,
    ride_status TEXT NOT NULL,
    driver_id UUID NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS bookings;
