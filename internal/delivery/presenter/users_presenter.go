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
	if response == nil {
		common.WriteJSONResponse(u.w, http.StatusOK, "Get users view successfully", domain.UserViewNilResponse{
			Message: "Your is swipe in maximum 10",
		}, int64(len(response)))
	} else {
		common.WriteJSONResponse(u.w, http.StatusOK, "Get users view successfully", response, int64(len(response)))
	}
}
