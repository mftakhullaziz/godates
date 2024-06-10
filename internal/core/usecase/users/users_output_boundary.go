package users

import (
	res "godating-dealls/internal/domain"
)

type OutputUserBoundary interface {
	UserViewsResponse(response []res.UserViewsResponse, err error)
	PatchUserResponse(response res.PatchUserResponse, err error)
}
