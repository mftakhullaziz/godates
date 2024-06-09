package packages

import "context"

type InputPackageBoundary interface {
	ExecuteGetAllPackages(ctx context.Context, token string, boundary BoundaryPackageOutput) error
}
