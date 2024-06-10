package handler

import (
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/daily_quotas"
	"godating-dealls/internal/delivery/presenter"
	"net/http"
)

type QuotaHandler struct {
	InputDailyQuotaBoundary daily_quotas.InputDailyQuotaBoundary
}

func NewQuotaHandler(inputDailyQuotaBoundary daily_quotas.InputDailyQuotaBoundary) *QuotaHandler {
	return &QuotaHandler{InputDailyQuotaBoundary: inputDailyQuotaBoundary}
}

func (q *QuotaHandler) CheckQuotaAccountHandler(w http.ResponseWriter, r *http.Request) {
	// If user premium is unlimited, if not is just 10 data
	ctx := r.Context()
	token, ok := ctx.Value("token").(string)
	if !ok || token == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Instantiate the presenter
	presenter := presenters.NewQuotaPresenter(w)

	// Call the use case method passing the presenter
	err := q.InputDailyQuotaBoundary.ExecuteFindDailyQuotaUsecase(ctx, token, presenter)
	common.HandleInternalServerError(err, w)
}
