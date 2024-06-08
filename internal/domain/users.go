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
	DateOfBirth time.Time
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
	UserID    int64
	AccountID int64
	FullName  *string
	Username  string
	Photos    []string
	Videos    []string
	Age       int
	Gender    string
	Address   string
	Bio       string
	Verified  bool
}
