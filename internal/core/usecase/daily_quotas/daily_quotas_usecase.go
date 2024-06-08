package daily_quotas

import (
	"database/sql"
	"godating-dealls/internal/core/entities/daily_quotas"
)

type DailyQuotasUsecase struct {
	DB                 *sql.DB
	DailyQuotaEntities daily_quotas.DailyQuotasEntities
}

func NewDailyQuotasUsecase(db *sql.DB, dailyQuotasEntities daily_quotas.DailyQuotasEntities) InputDailyQuotaBoundary {
	return &DailyQuotasUsecase{
		DB:                 db,
		DailyQuotaEntities: dailyQuotasEntities,
	}
}

func (d DailyQuotasUsecase) ExecuteAutoUpdateDailyQuotaUsecase() error {
	//TODO implement me
	panic("implement me")
}
