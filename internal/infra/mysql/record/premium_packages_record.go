package record

import "time"

// PremiumPackageRecord represents a premium package available for purchase
type PremiumPackageRecord struct {
	PackageID       int64     `db:"package_id"`
	Description     string    `db:"description"`
	Price           float64   `db:"price"`
	UnlimitedSwipes bool      `db:"unlimited_swipes"`
	Verified        bool      `db:"verified"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

func (PremiumPackageRecord) TableName() string {
	return "premium_packages"
}
