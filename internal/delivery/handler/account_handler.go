package handler

import (
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/accounts"
	presenters "godating-dealls/internal/delivery/presenter"
	"net/http"
)

type AccountHandler struct {
	InputAccountBoundary accounts.InputAccountBoundary
}

func NewAccountHandler(inputAccountBoundary accounts.InputAccountBoundary) *AccountHandler {
	return &AccountHandler{InputAccountBoundary: inputAccountBoundary}
}

func (ac *AccountHandler) FetchAccountDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// If user premium is unlimited, if not is just 10 data
	ctx := r.Context()
	token, ok := ctx.Value("token").(string)
	if !ok || token == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	presenter := presenters.NewAccountPresenter(w)

	err := ac.InputAccountBoundary.ExecuteFetchAccountDetail(ctx, token, presenter)
	common.HandleErrorReturn(err)
}
