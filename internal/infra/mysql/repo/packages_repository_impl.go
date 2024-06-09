package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type PackagesRepositoryImpl struct {
	PackagesRepository PackagesRepository
}

func NewPackagesRepositoryImpl() PackagesRepository {
	return &PackagesRepositoryImpl{}
}

func (p PackagesRepositoryImpl) GetAllPackages(ctx context.Context, tx *sql.Tx) ([]record.PremiumPackageRecord, error) {
	query := "SELECT * FROM packages WHERE status = TRUE"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var packages []record.PremiumPackageRecord
	for rows.Next() {
		var pkg record.PremiumPackageRecord
		err := rows.Scan(
			&pkg.PackageID,
			&pkg.PackageName,
			&pkg.Description,
			&pkg.PackageDurationInMonthly,
			&pkg.Price,
			&pkg.UnlimitedSwipes,
			&pkg.Status,
			&pkg.CreatedAt,
			&pkg.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		packages = append(packages, pkg)
	}

	// Check for any error encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return packages, nil
}
