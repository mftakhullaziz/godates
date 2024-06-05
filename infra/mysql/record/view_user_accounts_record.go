package record

import "time"

// ViewedUserAccountRecord represents a profile that has been viewed by a user
type ViewedUserAccountRecord struct {
	ViewID    int64     `db:"view_id"`
	AccountID int64     `db:"account_id"`
	UserID    int64     `db:"user_id"`
	Date      time.Time `db:"date"`
}

func (ViewedUserAccountRecord) TableName() string {
	return "viewed_user_accounts"
}
