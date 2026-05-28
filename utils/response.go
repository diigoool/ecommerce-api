package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success   bool   `json:"success" example:"true"`
	Message   string `json:"message,omitempty" example:"success"`
	Code      string `json:"code,omitempty"`
	Data      any    `json:"data,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := APIResponse{
		Success: true,
		Code:    "SUCCESS",
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, status int, code string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := APIResponse{
		Success: false,
		Code:    code,
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "failed to encode error", http.StatusInternalServerError)
	}
}

func JSONWithMeta(w http.ResponseWriter, status int, data any, page, limit int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := map[string]any{
		"success": true,
		"data":    data,
		"meta": map[string]int{
			"page":  page,
			"limit": limit,
		},
	}

	json.NewEncoder(w).Encode(res)
}
