package daily_quotas

import (
	"context"
	"database/sql"
	"godating-dealls/internal/common"
	dailyQuotas "godating-dealls/internal/core/entities/daily_quotas"
	userEntities "godating-dealls/internal/core/entities/users"
	"godating-dealls/internal/domain"
	"log"
)

type DailyQuotasUsecase struct {
	DB  *sql.DB
	Dqe dailyQuotas.DailyQuotasEntities
	Ue  userEntities.UserEntities
}

func NewDailyQuotasUsecase(db *sql.DB, dqe dailyQuotas.DailyQuotasEntities, ue userEntities.UserEntities) InputDailyQuotaBoundary {
	return &DailyQuotasUsecase{
		DB:  db,
		Dqe: dqe,
		Ue:  ue,
	}
}

func (d DailyQuotasUsecase) ExecuteAutoUpdateDailyQuotaUsecase(ctx context.Context) error {
	fn := func(tx *sql.Tx) error {
		users, err := d.Ue.FindAllUserEntities(ctx, tx)
		common.HandleErrorReturn(err)
		log.Printf("users: %v", users)

		for _, user := range users {
			dailyQuotaDto := domain.DailyQuotasDto{
				AccountID:      user.AccountID,
				UserIsVerified: user.Verified,
			}

			// Check user is premium
			if user.Verified == true {
				dailyQuotaDto.UserIsVerified = true
			}
			common.PrintJSON("usecase | daily quotas", dailyQuotaDto)

			err := d.Dqe.UpdateOrInsertDailyQuotaEntities(ctx, tx, dailyQuotaDto)
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
