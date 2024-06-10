package repo

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/infra/mysql/record"
)

type PurchasePackagesRepositoryImpl struct {
	PurchasePackagesRepository PurchasePackagesRepository
}

func NewPurchasePackagesRepositoryImpl() PurchasePackagesRepository {
	return &PurchasePackagesRepositoryImpl{}
}

func (p PurchasePackagesRepositoryImpl) PurchasePackagesByAccount(ctx context.Context, tx *sql.Tx, record record.AccountPremiumRecord) error {
	query := `
		INSERT INTO account_premiums (account_id, package_id, purchase_date, expiry_date, unlimited_swipes_active, status)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := tx.ExecContext(ctx, query,
		record.AccountID,
		record.PackageID,
		record.PurchaseDate,
		record.ExpiryDate,
		record.UnlimitedSwipesActive,
		record.Status,
	)

	if err != nil {
		return errors.New("insert into accounts premium failed")
	}
	return nil
}

func (p PurchasePackagesRepositoryImpl) FindAccountPremiumByAccountId(ctx context.Context, tx *sql.Tx, accountId int64) (record.AccountPremiumRecord, error) {
	query := `SELECT account_id, package_id, purchase_date, expiry_date, unlimited_swipes_active, status FROM account_premiums WHERE account_id = ?`

	row, err := tx.QueryContext(ctx, query, accountId)
	if err != nil {
		return record.AccountPremiumRecord{}, errors.New("query account premium failed")
	}
	defer row.Close()

	var accountPremiumRecord record.AccountPremiumRecord
	if row.Next() {
		err = row.Scan(
			&accountPremiumRecord.AccountID,
			&accountPremiumRecord.PackageID,
			&accountPremiumRecord.PurchaseDate,
			&accountPremiumRecord.ExpiryDate,
			&accountPremiumRecord.UnlimitedSwipesActive,
			&accountPremiumRecord.Status,
		)
		if err != nil {
			return record.AccountPremiumRecord{}, errors.New("query account premium failed")
		}
		return accountPremiumRecord, nil
	}

	return record.AccountPremiumRecord{}, errors.New("no account premium found")
}
