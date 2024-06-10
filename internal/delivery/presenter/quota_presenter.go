package presenters

import (
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/daily_quotas"
	"godating-dealls/internal/domain"
	"net/http"
)

type QuotaPresenter struct {
	w http.ResponseWriter
}

func NewQuotaPresenter(w http.ResponseWriter) daily_quotas.DailyQuotasOutputBoundary {
	return &QuotaPresenter{w: w}
}

func (q QuotaPresenter) DailyQuotaResponse(response domain.DailyQuotaResponse, err error) {
	common.HandleInternalServerError(err, q.w)
	common.WriteJSONResponse(q.w, http.StatusOK, "Fetch quota successfully", response, 1)
}
