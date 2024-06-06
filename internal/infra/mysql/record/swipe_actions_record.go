package record

import "time"

// SwipeRecord represents a swipe action in the system
type SwipeRecord struct {
	SwipeID   int64     `db:"swipe_id"`
	AccountID int64     `db:"account_id"`
	UserID    int64     `db:"user_id"`
	Action    string    `db:"action"`
	SwipeDate time.Time `db:"swipe_date"`
}

func (SwipeRecord) TableName() string {
	return "swipes"
}
