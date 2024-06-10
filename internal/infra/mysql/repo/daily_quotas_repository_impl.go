package repo

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/common"
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

func (d DailyQuotasRepositoryImpl) FindDailyQuotasByUserId(ctx context.Context, tx *sql.Tx, accountId int64) (record.DailyQuotaRecord, error) {
	query := "SELECT d.quota_id, d.account_id, d.swipe_count, d.total_quota, d.date FROM daily_quotas d INNER JOIN accounts a ON d.account_id = a.account_id WHERE d.account_id = ? AND a.verified = FALSE"
	row := tx.QueryRowContext(ctx, query, accountId)

	var records record.DailyQuotaRecord
	err := row.Scan(
		&records.QuotaID,
		&records.AccountID,
		&records.SwipeCount,
		&records.TotalQuota,
		&records.Date,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record.DailyQuotaRecord{}, errors.New("error finding daily quota when scan rows")
		}
		return record.DailyQuotaRecord{}, errors.New("error finding daily quotas")
	}
	return records, nil
}

func (d DailyQuotasRepositoryImpl) UpdateIncreaseSwipeCount(ctx context.Context, tx *sql.Tx, dailyQuota record.DailyQuotaRecord) error {
	query := "UPDATE daily_quotas SET swipe_count  = swipe_count  + 1 WHERE account_id = ? AND date = CURDATE()"
	common.PrintJSON("printed query", query)
	_, err := tx.ExecContext(ctx, query, dailyQuota.AccountID)
	return err
}

func (d DailyQuotasRepositoryImpl) UpdateDecreaseTotalCount(ctx context.Context, tx *sql.Tx, dailyQuota record.DailyQuotaRecord) error {
	query := "UPDATE daily_quotas SET total_quota  = total_quota  - 1 WHERE account_id = ? AND date = CURDATE()"
	common.PrintJSON("printed query", query)
	_, err := tx.ExecContext(ctx, query, dailyQuota.AccountID)
	return err
}

func (d DailyQuotasRepositoryImpl) UpdateTotalQuotaInPremiumAccount(ctx context.Context, tx *sql.Tx, dailyQuota record.DailyQuotaRecord) error {
	query := "UPDATE daily_quotas SET total_quota  = -1 WHERE account_id = ? AND date = CURDATE()"
	common.PrintJSON("printed query", query)
	_, err := tx.ExecContext(ctx, query, dailyQuota.AccountID)
	return err
}

func (d DailyQuotasRepositoryImpl) FindTotalQuotaByAccountId(ctx context.Context, tx *sql.Tx, accountId int64) (record.DailyQuotaRecord, error) {
	query := "SELECT d.quota_id, d.account_id, d.swipe_count, d.total_quota, d.date FROM daily_quotas d INNER JOIN accounts a ON d.account_id = a.account_id WHERE d.account_id = ?"
	row := tx.QueryRowContext(ctx, query, accountId)

	var records record.DailyQuotaRecord
	err := row.Scan(
		&records.QuotaID,
		&records.AccountID,
		&records.SwipeCount,
		&records.TotalQuota,
		&records.Date,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record.DailyQuotaRecord{}, errors.New("error finding daily quota when scan rows")
		}
		return record.DailyQuotaRecord{}, errors.New("error finding daily quotas")
	}
	return records, nil
}
