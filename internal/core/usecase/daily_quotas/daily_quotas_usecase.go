package daily_quotas

import (
	"context"
	"database/sql"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/daily_quotas"
	"godating-dealls/internal/core/entities/users"
	"godating-dealls/internal/domain"
	"log"
)

type DailyQuotasUsecase struct {
	DB                *sql.DB
	DailyQuotasEntity daily_quotas.DailyQuotasEntity
	UserEntity        users.UserEntity
}

func NewDailyQuotasUsecase(db *sql.DB, dailyQuotasEntity daily_quotas.DailyQuotasEntity, userEntity users.UserEntity) InputDailyQuotaBoundary {
	return &DailyQuotasUsecase{
		DB:                db,
		DailyQuotasEntity: dailyQuotasEntity,
		UserEntity:        userEntity,
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
