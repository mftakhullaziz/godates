package auths

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	AccountId int64  `json:"account_id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type AccountDto struct {
	Email       string
	Username    string
	Password    string
	PhoneNumber string
}

type Accounts struct {
	AccountId   int64
	Email       string
	Username    string
	Password    string
	PhoneNumber string
	CreateAt    time.Time
	UpdateAt    time.Time
}
