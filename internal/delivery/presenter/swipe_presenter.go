package presenters

import (
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/swipes"
	"godating-dealls/internal/domain"
	"net/http"
)

type SwipePresenter struct {
	w http.ResponseWriter
}

func NewSwipePresenter(w http.ResponseWriter) swipes.OutputSwipesBoundary {
	return &SwipePresenter{w: w}
}

func (u SwipePresenter) SwipeResponse(response domain.SwipeResponse, err error) {
	common.HandleInternalServerError(err, u.w)
	common.WriteJSONResponse(u.w, http.StatusOK, "Get users view successfully", response, 1)
}
