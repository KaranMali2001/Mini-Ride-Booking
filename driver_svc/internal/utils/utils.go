package utils

import (
	"encoding/json"
	"net/http"

	"github.com/KaranMali2001/mini-ride-booking/common/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Response struct {
	Success bool                   `json:"success"`
	Message map[string]interface{} `json:"message,omitempty"`
	Err     error                  `json:"error,omitempty"`
}

func SendJson(success bool, message map[string]interface{}, err error, w http.ResponseWriter) {
	res := Response{
		Success: success,
		Message: message,
		Err:     err,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		logger.Errorln("errror while sending json message", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
func UUIDFromString(s string) (pgtype.UUID, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return pgtype.UUID{
		Bytes: u, // uuid.UUID is [16]byte
		Valid: true,
	}, nil
}
