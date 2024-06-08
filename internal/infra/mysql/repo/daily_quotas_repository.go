package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type DailyQuotasRepository interface {
	UpdateOrInsertDailyQuota(ctx context.Context, tx *sql.Tx, dailyQuota record.DailyQuotaRecord) error
}
