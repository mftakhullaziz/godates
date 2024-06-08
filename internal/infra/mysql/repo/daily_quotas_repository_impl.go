package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/queries"
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
	query := queries.InsertIntoDailyQuotaRecord
	log.Printf("query: %s", query)
	_, err := tx.ExecContext(ctx, query, dailyQuota.AccountID, dailyQuota.SwipeCount, dailyQuota.TotalQuota)
	return err
}
