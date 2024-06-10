package handler

import (
	"encoding/json"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/accounts"
	presenters "godating-dealls/internal/delivery/presenter"
	"godating-dealls/internal/domain"
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

func (ac *AccountHandler) AccountViewHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, ok := ctx.Value("token").(string)
	if !ok || token == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	var request domain.ViewedAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	presenter := presenters.NewAccountPresenter(w)

	err := ac.InputAccountBoundary.ExecuteViewAccountDetail(ctx, token, request, presenter)
	common.HandleErrorReturn(err)
}
