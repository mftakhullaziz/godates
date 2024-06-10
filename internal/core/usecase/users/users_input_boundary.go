package users

import (
	"context"
	"godating-dealls/internal/domain"
)

type InputUserBoundary interface {
	ExecuteUserViewsUsecase(ctx context.Context, token string, boundary OutputUserBoundary) error
	ExecutePatchUserUsecase(ctx context.Context, token string, request domain.PatchUserRequest, boundary OutputUserBoundary) error
}
