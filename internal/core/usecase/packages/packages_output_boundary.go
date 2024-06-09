package packages

import "godating-dealls/internal/domain"

type BoundaryPackageOutput interface {
	PackageResponse([]domain.PackageResponse, error)
}
