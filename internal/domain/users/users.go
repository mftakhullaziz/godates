package users

import "time"

type UserDto struct {
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
