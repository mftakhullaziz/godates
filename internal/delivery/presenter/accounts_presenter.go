package presenters

import (
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/accounts"
	"godating-dealls/internal/domain"
	"net/http"
)

type AccountPresenter struct {
	w http.ResponseWriter
}

// NewAccountPresenter creates a new AuthPresenter
func NewAccountPresenter(w http.ResponseWriter) accounts.OutputAccountBoundary {
	return &AccountPresenter{w: w}
}

func (a AccountPresenter) AccountDetailResponse(response domain.AccountResponse, err error) {
	common.HandleInternalServerError(err, a.w)
	common.WriteJSONResponse(a.w, http.StatusOK, "Fetch account detail successfully", response, 1)
}

func (a AccountPresenter) ViewAccountResponse(response domain.ViewedAccountResponse, err error) {
	common.HandleInternalServerError(err, a.w)
	common.WriteJSONResponse(a.w, http.StatusOK, "View account successfully", response, 1)
}
