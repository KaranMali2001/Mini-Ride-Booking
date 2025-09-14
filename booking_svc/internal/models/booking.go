package models

type CreateBookingRequest struct {
	PickupLat  float64 `json:"pickup_lat" validate:"required"`
	PickupLng  float64 `json:"pickup_lng" validate:"required"`
	DropoffLat float64 `json:"dropoff_lat" validate:"required"`
	DropoffLng float64 `json:"dropoff_lng" validate:"required"`
	Price      int     `json:"price" validate:"required"`
	RideStatus string  `json:"ride_status" validate:"required"`
	DriverID   *string `json:"driver_id,omitempty" validate:"omitempty,uuid4"`
}
