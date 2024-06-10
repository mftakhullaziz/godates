package daily_quotas

import "context"

type InputDailyQuotaBoundary interface {
	ExecuteAutoUpdateDailyQuotaUsecase(ctx context.Context) error
	ExecuteFindDailyQuotaUsecase(ctx context.Context, token string, boundary DailyQuotasOutputBoundary) error
}
