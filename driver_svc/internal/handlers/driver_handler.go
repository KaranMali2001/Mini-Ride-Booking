package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/models"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/service"
	"github.com/KaranMali2001/mini-ride-booking/driver_svc/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type DriverHandler struct {
	DriverService *service.DriverService
	Validator     *validator.Validate
}

func NewDriverHandler(svc *service.DriverService) *DriverHandler {
	return &DriverHandler{DriverService: svc, Validator: validator.New()}
}
func (h *DriverHandler) AcceptJob(w http.ResponseWriter, r *http.Request) {
	var req models.AcceptJobRequest
	req.BookingID = chi.URLParam(r, "id")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.Validator.Struct(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.DriverService.AcceptJob(r.Context(), req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJson(true, nil, nil, w)
}
func (h *DriverHandler) GetAllJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := h.DriverService.GetAllJobs(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJson(true, map[string]interface{}{"jobs": jobs}, nil, w)
}

func (h *DriverHandler) GetAllDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, err := h.DriverService.GetAllDrivers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJson(true, map[string]interface{}{"drivers": drivers}, nil, w)
}
