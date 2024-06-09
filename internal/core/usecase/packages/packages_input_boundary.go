package packages

import (
	"context"
	"godating-dealls/internal/domain"
)

type InputPackageBoundary interface {
	ExecuteGetAllPackages(ctx context.Context, token string, boundary BoundaryPackageOutput) error
	ExecutePurchasedPackages(ctx context.Context, token string, request domain.PurchasePackageRequest, boundary BoundaryPackageOutput) error
}
