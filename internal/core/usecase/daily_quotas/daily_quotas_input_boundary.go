package daily_quotas

type InputDailyQuotaBoundary interface {
	ExecuteAutoUpdateDailyQuotaUsecase() error
}
