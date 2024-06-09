package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type PurchasePackagesRepository interface {
	PurchasePackagesByAccount(ctx context.Context, tx *sql.Tx, record record.AccountPremiumRecord) error
}
