package daily_quotas

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/common"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/mysql/record"
	"godating-dealls/internal/infra/mysql/repo"
	"log"
	"time"
)

type DailyQuotasEntitiesImpl struct {
	DailyQuotaRepository repo.DailyQuotasRepository
	Validate             *validator.Validate
}

func NewDailyQuotasEntitiesImpl(validate *validator.Validate, dailyQuotaRepository repo.DailyQuotasRepository) DailyQuotasEntities {
	return &DailyQuotasEntitiesImpl{Validate: validate, DailyQuotaRepository: dailyQuotaRepository}
}

func (d DailyQuotasEntitiesImpl) UpdateOrInsertDailyQuotaEntities(ctx context.Context, tx *sql.Tx, dto domain.DailyQuotasDto) error {
	today := time.Now().Format("2006-01-02")
	todayTime, err := time.Parse(today, "2006-01-02")
	common.HandleErrorReturn(err)

	dailyQuota := record.DailyQuotaRecord{
		AccountID:  dto.AccountID,
		Date:       todayTime,
		SwipeCount: 0,  // Default swipe count
		TotalQuota: 10, // Set the default quota
	}

	// Set total quota is unlimited
	if dto.UserIsVerified == true {
		dailyQuota.TotalQuota = -1
	}
	log.Printf("daily quota entities %v", dailyQuota)

	err = d.DailyQuotaRepository.UpdateOrInsertDailyQuota(ctx, tx, dailyQuota)
	common.HandleErrorReturn(err)
	return nil
}
