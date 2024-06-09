package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type DailyQuotasRepository interface {
	UpdateOrInsertDailyQuota(ctx context.Context, tx *sql.Tx, dailyQuota record.DailyQuotaRecord) error
	FindDailyQuotasByUserId(ctx context.Context, tx *sql.Tx, accountId int64) (record.DailyQuotaRecord, error)
	UpdateIncreaseSwipeCount(ctx context.Context, tx *sql.Tx, dailyQuota record.DailyQuotaRecord) error
	UpdateDecreaseTotalCount(ctx context.Context, tx *sql.Tx, dailyQuota record.DailyQuotaRecord) error
}
