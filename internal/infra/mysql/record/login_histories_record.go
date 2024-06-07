package record

import "time"

type LoginHistoriesRecord struct {
	LoginHistoriesID  int64      `db:"login_histories_id"`
	UserID            int64      `db:"user_id"`
	AccountID         int64      `db:"account_id"`
	LoginAt           *time.Time `db:"login_at"`
	LogoutAt          *time.Time `db:"logout_at"`
	DurationInSeconds *float64   `db:"duration_in_seconds"`
}

func (LoginHistoriesRecord) TableName() string {
	return "login_histories"
}
