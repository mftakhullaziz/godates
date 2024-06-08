package presenters

import (
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/users"
	"godating-dealls/internal/domain"
	"net/http"
)

type UserPresenter struct {
	w http.ResponseWriter
}

func NewUserPresenter(w http.ResponseWriter) users.OutputUserBoundary {
	return &UserPresenter{w: w}
}

func (u UserPresenter) UserViewsResponse(response []domain.UserViewsResponse, err error) {
	common.HandleInternalServerError(err, u.w)
	common.WriteJSONResponse(u.w, http.StatusOK, "Get users view successfully", response, int64(len(response)))
}
