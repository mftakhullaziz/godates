package domain

import "time"

type PackageDto struct {
	PackageID                int64
	PackageName              string
	Description              string
	PackageDurationInMonthly int64
	Price                    float64
	UnlimitedSwipes          bool
	Status                   bool
	AccountID                int64
}

type PackageResponse struct {
	PackageID                int64   `json:"package_id"`
	PackageName              string  `json:"package_name"`
	Description              string  `json:"description"`
	PackageDurationInMonthly int64   `json:"package_duration_in_monthly"`
	Price                    float64 `json:"price"`
	UnlimitedSwipes          bool    `json:"unlimited_swipes"`
	Status                   bool    `json:"status"`
}

type PurchasePackageRequest struct {
	PackageID                int64   `json:"package_id"`
	Price                    float64 `json:"price"`
	PackageName              string  `json:"package_name"`
	PackageDurationInMonthly int64   `json:"package_duration_in_monthly"`
	UnlimitedSwipes          bool    `json:"unlimited_swipes"`
}

type PurchasePackageResponse struct {
	PackageID int64   `json:"package_id"`
	Price     float64 `json:"price"`
	Message   string  `json:"message"`
}

type AccountPurchasePackage struct {
	AccountID int64
	ExpiresIn time.Time
}
