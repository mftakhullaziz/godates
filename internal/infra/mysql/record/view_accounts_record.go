package record

import "time"

// ViewAccountRecord represents a profile that has been viewed by a user
type ViewAccountRecord struct {
	ViewID    int64     `db:"view_id"`
	AccountID int64     `db:"account_id"`
	UserID    int64     `db:"user_id"`
	Date      time.Time `db:"date"`
}

func (ViewAccountRecord) TableName() string {
	return "view_accounts"
}
