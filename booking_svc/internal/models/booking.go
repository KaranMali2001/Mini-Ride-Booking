package models

type CreateBookingRequest struct {
	PickupLat  float64 `json:"pickup_lat" validate:"required,ne=0"`
	PickupLng  float64 `json:"pickup_lng" validate:"required,ne=0"`
	DropoffLat float64 `json:"dropoff_lat" validate:"required,ne=0"`
	DropoffLng float64 `json:"dropoff_lng" validate:"required,ne=0"`
	Price      int     `json:"price" validate:"required,gt=0"`
	RideStatus string  `json:"ride_status" validate:"required"`
}
