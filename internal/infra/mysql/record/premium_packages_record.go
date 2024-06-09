package record

import "time"

// PremiumPackageRecord represents a premium package available for purchase
type PremiumPackageRecord struct {
	PackageID                int64     `db:"package_id"`
	PackageName              string    `db:"package_name"`
	Description              string    `db:"description"`
	PackageDurationInMonthly int64     `db:"package_duration_in_monthly"`
	Price                    float64   `db:"price"`
	UnlimitedSwipes          bool      `db:"unlimited_swipes"`
	Status                   bool      `db:"status"`
	CreatedAt                time.Time `db:"created_at"`
	UpdatedAt                time.Time `db:"updated_at"`
}

func (PremiumPackageRecord) TableName() string {
	return "packages"
}
