package users

import (
	"context"
)

type InputUserBoundary interface {
	ExecuteUserViewsUsecase(ctx context.Context, token string, boundary OutputUserBoundary) error
}
