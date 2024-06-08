package auths

import (
	"context"
	"godating-dealls/internal/domain"
)

type InputAuthBoundary interface {
	ExecuteLoginUsecase(ctx context.Context, request domain.LoginRequest, boundary OutputAuthBoundary) error
	ExecuteRegisterUsecase(ctx context.Context, request domain.RegisterRequest, boundary OutputAuthBoundary) error
	ExecuteLogoutUsecase(ctx context.Context, accessToken *string, boundary OutputAuthBoundary) error
}
