package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/models"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/service"
	"github.com/KaranMali2001/mini-ride-booking/booking_svc/internal/utils"
	"github.com/go-playground/validator/v10"
)

type BookingHandler struct {
	bookingService *service.BookingService
	validator      *validator.Validate
}

func NewBookingHandler(s *service.BookingService) *BookingHandler {
	return &BookingHandler{
		bookingService: s,
		validator:      validator.New(),
	}
}
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.validator.Struct(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	booking, err := h.bookingService.CreateBooking(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.SendJson(true, map[string]interface{}{"booking": booking}, nil, w)
}

func (h *BookingHandler) GetAllBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.bookingService.GetAllBookings(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJson(true, map[string]interface{}{"bookings": bookings}, nil, w)
}
