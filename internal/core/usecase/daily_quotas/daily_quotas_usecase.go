package daily_quotas

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/accounts"
	"godating-dealls/internal/core/entities/daily_quotas"
	"godating-dealls/internal/core/entities/packages"
	"godating-dealls/internal/core/entities/users"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/jsonwebtoken"
	"log"
	"strconv"
)

type DailyQuotasUsecase struct {
	DB                *sql.DB
	DailyQuotasEntity daily_quotas.DailyQuotasEntity
	UserEntity        users.UserEntity
	AccountEntity     accounts.AccountEntity
	PackageEntity     packages.PackageEntity
}

func NewDailyQuotasUsecase(
	db *sql.DB,
	dailyQuotasEntity daily_quotas.DailyQuotasEntity,
	userEntity users.UserEntity,
	accountEntity accounts.AccountEntity,
	packageEntity packages.PackageEntity) InputDailyQuotaBoundary {
	return &DailyQuotasUsecase{
		DB:                db,
		DailyQuotasEntity: dailyQuotasEntity,
		UserEntity:        userEntity,
		AccountEntity:     accountEntity,
		PackageEntity:     packageEntity,
	}
}

func (d DailyQuotasUsecase) ExecuteAutoUpdateDailyQuotaUsecase(ctx context.Context) error {
	fn := func(tx *sql.Tx) error {
		usersList, err := d.UserEntity.FindAllUserEntities(ctx, tx)
		common.HandleErrorReturn(err)
		common.PrintJSON("daily usecase | users", usersList)

		for _, user := range usersList {
			dailyQuotaDto := domain.DailyQuotasDto{
				AccountID:      user.AccountID,
				UserIsVerified: user.Verified,
			}

			// Check user is premium
			if user.Verified == true {
				dailyQuotaDto.UserIsVerified = true
			}
			common.PrintJSON("daily usecase | daily quotas", dailyQuotaDto)

			err := d.DailyQuotasEntity.UpdateOrInsertDailyQuotaEntities(ctx, tx, dailyQuotaDto)
			common.HandleErrorReturn(err)
		}
		return nil
	}

	err := common.WithExecuteTransactionalManager(ctx, d.DB, fn)
	if err != nil {
		log.Println("Transaction failed:", err)
	}
	return err
}

func (d DailyQuotasUsecase) ExecuteFindDailyQuotaUsecase(ctx context.Context, token string, boundary DailyQuotasOutputBoundary) error {
	fn := func(tx *sql.Tx) error {
		claims, err := jsonwebtoken.VerifyJWTToken(token)
		if err != nil {
			return errors.New("invalid token")
		}

		verifiedAccounts, err := d.AccountEntity.FindAccountVerifiedEntities(ctx, tx, claims.AccountId)
		if err != nil {
			return errors.New("invalid find accounts")
		}

		quota, err := d.DailyQuotasEntity.FindTotalDailyQuotasAndSwipeCount(ctx, tx, claims.AccountId)
		if err != nil {
			return errors.New("invalid find quota account")
		}

		res := domain.DailyQuotaResponse{
			TotalQuotas: strconv.FormatInt(quota.TotalQuota, 10),
			SwipeCount:  quota.SwipeCount,
		}

		if verifiedAccounts {
			expireQuota, err := d.PackageEntity.FindAccountPremiumPackage(ctx, tx, claims.AccountId)
			if err != nil {
				return errors.New("invalid find account premium package")
			}
			res.TotalQuotas = "Unlimited Until " + expireQuota.ExpiresIn.Format("2006-01-02 15:04:05")
		}

		boundary.DailyQuotaResponse(res, nil)
		return nil
	}

	err := common.WithExecuteTransactionalManager(ctx, d.DB, fn)
	if err != nil {
		return errors.New("execute transactional manager failed: " + err.Error())
	}
	return err
}
