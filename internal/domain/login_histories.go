package domain

import "time"

type LoginHistoriesDto struct {
	LoginHistoriesID   int64
	UserID             int64
	AccountID          int64
	LoginAt            *time.Time
	LogoutAt           *time.Time
	UserActiveDuration *time.Time
}
