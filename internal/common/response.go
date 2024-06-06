package common

type DefaultResponse struct {
	StatusCode int         `json:"status_code"`
	IsSuccess  bool        `json:"is_success"`
	Message    string      `json:"message"`
	RequestAt  string      `json:"request_at"`
	Data       interface{} `json:"data"`
}
