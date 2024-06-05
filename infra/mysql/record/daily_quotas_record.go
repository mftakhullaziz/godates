package record

import "time"

// DailyQuotaRecord represents the daily swipe quota for a user
type DailyQuotaRecord struct {
	QuotaID    int64     `db:"quota_id"`
	AccountID  int64     `db:"account_id"`
	Date       time.Time `db:"date"`
	SwipeCount int       `db:"swipe_count"`
}

func (DailyQuotaRecord) TableName() string {
	return "daily_quotas"
}
