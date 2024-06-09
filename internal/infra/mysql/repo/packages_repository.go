package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type PackagesRepository interface {
	GetAllPackages(ctx context.Context, tx *sql.Tx) ([]record.PremiumPackageRecord, error)
}
