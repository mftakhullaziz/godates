package daily_quotas

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain"
)

type DailyQuotasEntities interface {
	UpdateOrInsertDailyQuotaEntities(ctx context.Context, tx *sql.Tx, dto domain.DailyQuotasDto) error
}
