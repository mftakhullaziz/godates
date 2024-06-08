package users

import (
	"context"
)

type InputUserBoundary interface {
	ExecuteUserViewsUsecase(ctx context.Context, boundary OutputUserBoundary) error
}
