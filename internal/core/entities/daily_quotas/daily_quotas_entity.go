package daily_quotas

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain"
)

type DailyQuotasEntity interface {
	UpdateOrInsertDailyQuotaEntities(ctx context.Context, tx *sql.Tx, dto domain.DailyQuotasDto) error
	FetchTotalDailyQuotas(ctx context.Context, tx *sql.Tx, accountId int64) (int64, error)
	UpdateIncreaseSwipeCountAndDecreaseTotalQuota(ctx context.Context, tx *sql.Tx, accountId int64) error
	UpdateIncreaseSwipeCount(ctx context.Context, tx *sql.Tx, accountId int64) error
	UpdateTotalQuotasInPremiumAccount(ctx context.Context, tx *sql.Tx, accountId int64) error
}
