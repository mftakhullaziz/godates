package domain

import "time"

type UserDto struct {
	UserID      int64
	AccountID   int64
	FullName    *string
	DateOfBirth time.Time
	Age         int
	Gender      string
	Address     string
	Bio         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Users struct {
	UserID      int64
	AccountID   int64
	FullName    *string
	DateOfBirth *time.Time
	Age         int
	Gender      string
	Address     string
	Bio         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AllUsers struct {
	UserID    int64
	AccountID int64
	Verified  bool
}

type AllUserViews struct {
	UserID    int64
	AccountID int64
	FullName  *string
	Username  string
	Age       int
	Gender    string
	Address   string
	Bio       string
	Verified  bool
}

type UserViewsResponse struct {
	UserID    int64    `json:"user_id"`
	AccountID int64    `json:"account_id"`
	FullName  *string  `json:"full_name"`
	Username  string   `json:"username"`
	Photos    []string `json:"photos"`
	Videos    []string `json:"videos"`
	Age       int      `json:"age"`
	Gender    string   `json:"gender"`
	Address   string   `json:"address"`
	Bio       string   `json:"bio"`
	Verified  bool     `json:"verified"`
}

type UserViewNilResponse struct {
	Message string `json:"message"`
}
