package models

type CreateJobRequest struct {
	BookingID string `json:"booking_id" validate:"required,uuid4"`
	DriverID  string `json:"driver_id" validate:"required,uuid4"`
	JobStatus string `json:"job_status" validate:"required"`
}
type UpdateDriverRequest struct {
	DriverID  string `json:"driver_id" validate:"required,uuid4"`
	JobStatus string `json:"job_status" validate:"required"`
}
type AcceptJobRequest struct {
	BookingID string `json:"booking_id" validate:"required,uuid4"`
	DriverId  string `json:"driver_id" validate:"required,uuid4"`
}
