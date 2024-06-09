package daily_quotas

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/common"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/mysql/record"
	"godating-dealls/internal/infra/mysql/repo"
	"time"
)

type DailyQuotasEntityImpl struct {
	DailyQuotaRepository repo.DailyQuotasRepository
	Validate             *validator.Validate
}

func NewDailyQuotasEntityImpl(validate *validator.Validate, dailyQuotaRepository repo.DailyQuotasRepository) DailyQuotasEntity {
	return &DailyQuotasEntityImpl{Validate: validate, DailyQuotaRepository: dailyQuotaRepository}
}

func (d DailyQuotasEntityImpl) UpdateOrInsertDailyQuotaEntities(ctx context.Context, tx *sql.Tx, dto domain.DailyQuotasDto) error {
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
	common.PrintJSON("entities | daily quota entities", dailyQuota)

	err = d.DailyQuotaRepository.UpdateOrInsertDailyQuota(ctx, tx, dailyQuota)
	common.HandleErrorReturn(err)
	return nil
}

func (d DailyQuotasEntityImpl) FetchTotalDailyQuotas(ctx context.Context, tx *sql.Tx, accountId int64) (int64, error) {
	dailyQuotaRecord, err := d.DailyQuotaRepository.FindDailyQuotasByUserId(ctx, tx, accountId)
	common.HandleErrorReturn(err)
	return dailyQuotaRecord.TotalQuota, nil
}

func (d DailyQuotasEntityImpl) UpdateIncreaseSwipeCountAndDecreaseTotalQuota(ctx context.Context, tx *sql.Tx, accountId int64) error {
	err := d.DailyQuotaRepository.UpdateDecreaseTotalCount(ctx, tx, record.DailyQuotaRecord{AccountID: accountId})
	common.HandleErrorReturn(err)
	err = d.DailyQuotaRepository.UpdateIncreaseSwipeCount(ctx, tx, record.DailyQuotaRecord{AccountID: accountId})
	common.HandleErrorReturn(err)
	return nil
}

func (d DailyQuotasEntityImpl) UpdateIncreaseSwipeCount(ctx context.Context, tx *sql.Tx, accountId int64) error {
	err := d.DailyQuotaRepository.UpdateIncreaseSwipeCount(ctx, tx, record.DailyQuotaRecord{AccountID: accountId})
	common.HandleErrorReturn(err)
	return nil
}

func (d DailyQuotasEntityImpl) UpdateTotalQuotasInPremiumAccount(ctx context.Context, tx *sql.Tx, accountId int64) error {
	err := d.DailyQuotaRepository.UpdateTotalQuotaInPremiumAccount(ctx, tx, record.DailyQuotaRecord{AccountID: accountId})
	common.HandleErrorReturn(err)
	return nil
}
