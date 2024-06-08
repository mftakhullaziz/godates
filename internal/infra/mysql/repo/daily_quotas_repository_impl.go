package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
	"log"
)

type DailyQuotasRepositoryImpl struct {
	DailyQuotasRepo DailyQuotasRepository
}

func NewDailyQuotasRepositoryImpl() DailyQuotasRepository {
	return &DailyQuotasRepositoryImpl{}
}

func (d DailyQuotasRepositoryImpl) UpdateOrInsertDailyQuota(ctx context.Context, tx *sql.Tx, dailyQuota record.DailyQuotaRecord) error {
	query := `
        INSERT INTO daily_quotas (account_id, date, swipe_count, total_quota)
        VALUES (?, ?, ?, ?)
        ON CONFLICT (account_id, date)
        DO UPDATE SET
            swipe_count = EXCLUDED.swipe_count,
            total_quota = EXCLUDED.total_quota
    `
	log.Printf("query: %s", query)
	_, err := tx.ExecContext(ctx, query, dailyQuota.AccountID, dailyQuota.Date, dailyQuota.SwipeCount, dailyQuota.TotalQuota)
	return err
}
