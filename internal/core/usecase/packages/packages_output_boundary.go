package packages

import "godating-dealls/internal/domain"

type BoundaryPackageOutput interface {
	PackageResponse([]domain.PackageResponse, error)
	PurchasePackageResponse(domain.PurchasePackageResponse, error)
}
