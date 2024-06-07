package common

import (
	"encoding/json"
	"net/http"
)

type DefaultResponse struct {
	StatusCode int         `json:"status_code"`
	IsSuccess  bool        `json:"is_success"`
	Message    string      `json:"message"`
	RequestAt  string      `json:"request_at"`
	Data       interface{} `json:"data"`
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := DefaultResponse{
		StatusCode: statusCode,
		Message:    message,
		IsSuccess:  statusCode >= 200 && statusCode < 300,
		RequestAt:  FormatTime(),
		Data:       data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	HandleErrorReturn(err)
}
