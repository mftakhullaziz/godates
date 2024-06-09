package packages

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain"
)

type PackageEntity interface {
	GetAllPackagesEntity(ctx context.Context, tx *sql.Tx) ([]domain.PackageDto, error)
}
