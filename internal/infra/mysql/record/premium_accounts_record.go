package record

import "time"

// AccountPremiumRecord represents a premium package purchased by a user
type AccountPremiumRecord struct {
	PurchaseID            int64     `db:"purchase_id"`
	AccountID             int64     `db:"account_id"`
	PackageID             int64     `db:"package_id"`
	PurchaseDate          time.Time `db:"purchase_date"`
	ExpiryDate            time.Time `db:"expiry_date"`
	UnlimitedSwipesActive bool      `db:"unlimited_swipes_active"`
	Status                bool      `db:"status"`
}

func (AccountPremiumRecord) TableName() string {
	return "account_premiums"
}
