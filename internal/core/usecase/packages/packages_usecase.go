package packages

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/accounts"
	"godating-dealls/internal/core/entities/daily_quotas"
	"godating-dealls/internal/core/entities/packages"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/jsonwebtoken"
	"log"
)

type PackageUsecase struct {
	DB                *sql.DB
	PackageEntity     packages.PackageEntity
	AccountEntity     accounts.AccountEntity
	DailyQuotasEntity daily_quotas.DailyQuotasEntity
}

func NewPackageUsecase(db *sql.DB,
	packageEntity packages.PackageEntity,
	accountEntity accounts.AccountEntity,
	dailyQuotasEntity daily_quotas.DailyQuotasEntity) InputPackageBoundary {
	return &PackageUsecase{
		DB:                db,
		PackageEntity:     packageEntity,
		AccountEntity:     accountEntity,
		DailyQuotasEntity: dailyQuotasEntity,
	}
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

func (p PackageUsecase) ExecutePurchasedPackages(ctx context.Context, token string, request domain.PurchasePackageRequest, boundary BoundaryPackageOutput) error {
	fn := func(tx *sql.Tx) error {
		claims, err := jsonwebtoken.VerifyJWTToken(token)
		if err != nil {
			return errors.New("invalid token")
		}

		packageDto := domain.PackageDto{
			PackageID:                request.PackageID,
			PackageName:              request.PackageName,
			Price:                    request.Price,
			PackageDurationInMonthly: request.PackageDurationInMonthly,
			UnlimitedSwipes:          request.UnlimitedSwipes,
			AccountID:                claims.AccountId,
		}
		err = p.PackageEntity.PurchasePackage(ctx, tx, packageDto)
		if err != nil {
			return errors.New("could not purchase package")
		}

		// if success purchase update total quota today to unlimited and account to verified
		err = p.AccountEntity.UpdateAccountVerified(ctx, tx, claims.AccountId)
		if err != nil {
			return errors.New("could not update account verified")
		}

		err = p.DailyQuotasEntity.UpdateTotalQuotasInPremiumAccount(ctx, tx, claims.AccountId)
		if err != nil {
			return errors.New("could not update total quotas")
		}

		boundary.PurchasePackageResponse(domain.PurchasePackageResponse{
			PackageID: request.PackageID,
			Price:     request.Price,
			Message:   "Purchased package successfully",
		}, nil)

		return nil
	}

	err := common.WithExecuteTransactionalManager(ctx, p.DB, fn)
	if err != nil {
		return errors.New("transactional failed")
	}
	return err
}
