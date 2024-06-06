package common

type DefaultResponse struct {
	StatusCode int         `json:"status_code"`
	IsSuccess  bool        `json:"is_success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
