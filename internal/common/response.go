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
	TotalData  int64       `json:"total_data"`
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, message string, data interface{}, totalData int64) {
	response := DefaultResponse{
		StatusCode: statusCode,
		Message:    message,
		IsSuccess:  statusCode >= 200 && statusCode < 300,
		RequestAt:  FormatTime(),
		Data:       data,
		TotalData:  totalData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	HandleErrorReturn(err)
}
