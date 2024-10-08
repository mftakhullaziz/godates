package domain

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type RegisterResponse struct {
	AccountId int64  `json:"account_id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type AccountDto struct {
	Email    *string
	Username *string
	Password string
}

type Accounts struct {
	AccountId int64
	Email     string
	Username  string
	Password  string
	CreateAt  time.Time
	UpdateAt  time.Time
}

type AccountDetail struct {
	AccountId int64
	Email     string
	Username  string
	Verified  bool
}
