package models

type CreateJobRequest struct {
	BookingID    string  `json:"booking_id" validate:"required,uuid4"`
	DriverID     string  `json:"driver_id" validate:"required,uuid4"`
	RideStatus   string  `json:"ride_status" validate:"required"`
	PickuplocLat float64 `json:"pickuploc_lat" validate:"required"`
	PickuplocLng float64 `json:"pickuploc_lng" validate:"required"`
	DropoffLat   float64 `json:"dropoff_lat" validate:"required"`
	DropoffLng   float64 `json:"dropoff_lng" validate:"required"`
	Price        int32   `json:"price" validate:"required"`
}
type UpdateDriverRequest struct {
	DriverID  string `json:"driver_id" validate:"required,uuid4"`
	Available bool   `json:"available" validate:"required"`
}
type AcceptJobRequest struct {
	BookingID string `json:"booking_id" validate:"required,uuid4"`
	DriverId  string `json:"driver_id" validate:"required,uuid4"`
}

type UpdateJobDriverRequest struct {
	BookingID  string `json:"booking_id" validate:"required,uuid4"`
	DriverID   string `json:"driver_id" validate:"required,uuid4"`
	RideStatus string `json:"ride_status" validate:"required"`
}
