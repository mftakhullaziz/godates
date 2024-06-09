package packages

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/packages"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/jsonwebtoken"
	"log"
)

type PackageUsecase struct {
	DB            *sql.DB
	PackageEntity packages.PackageEntity
}

func NewPackageUsecase(db *sql.DB, packageEntity packages.PackageEntity) InputPackageBoundary {
	return &PackageUsecase{DB: db, PackageEntity: packageEntity}
}

func (p PackageUsecase) ExecuteGetAllPackages(ctx context.Context, token string, boundary BoundaryPackageOutput) error {
	fn := func(tx *sql.Tx) error {
		// Verify token is not expired
		_, err := jsonwebtoken.VerifyJWTToken(token)
		if err != nil {
			return errors.New("invalid token")
		}

		res, err := p.PackageEntity.GetAllPackagesEntity(ctx, tx)
		if err != nil {
			return errors.New("could not get all packages entity")
		}

		var packageRes []domain.PackageResponse
		for _, pkg := range res {
			packageRes = append(packageRes, domain.PackageResponse{
				PackageID:                pkg.PackageID,
				PackageName:              pkg.PackageName,
				Description:              pkg.Description,
				PackageDurationInMonthly: pkg.PackageDurationInMonthly,
				Price:                    pkg.Price,
				UnlimitedSwipes:          pkg.UnlimitedSwipes,
				Status:                   pkg.Status,
			})
		}

		boundary.PackageResponse(packageRes, nil)

		return nil
	}

	err := common.WithReadOnlyTransactionManager(ctx, p.DB, fn)
	if err != nil {
		log.Println("Transaction failed:", err)
	}
	return err
}
