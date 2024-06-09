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
