package auths

import (
	"context"
	"godating-dealls/internal/domain/auths"
)

type InputAuthBoundary interface {
	ExecuteLoginUsecase(ctx context.Context, request auths.LoginRequest, boundary OutputAuthBoundary) error
	ExecuteRegisterUsecase(ctx context.Context, request auths.RegisterRequest, boundary OutputAuthBoundary) error
	ExecuteLogoutUsecase(ctx context.Context, request auths.LoginRequest, boundary OutputAuthBoundary) error
}
