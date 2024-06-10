package daily_quotas

import "godating-dealls/internal/domain"

type DailyQuotasOutputBoundary interface {
	DailyQuotaResponse(response domain.DailyQuotaResponse, err error)
}
