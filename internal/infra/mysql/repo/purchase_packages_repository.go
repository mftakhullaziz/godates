package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type PurchasePackagesRepository interface {
	PurchasePackagesByAccount(ctx context.Context, tx *sql.Tx, record record.AccountPremiumRecord) error
	FindAccountPremiumByAccountId(ctx context.Context, tx *sql.Tx, accountId int64) (record.AccountPremiumRecord, error)
}
