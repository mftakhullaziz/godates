package auths

import (
	"context"
	"godating-dealls/internal/domain/auths"
)

type InputAuthBoundary interface {
	ExecuteLogin(ctx context.Context, request auths.LoginRequest) error
	ExecuteRegister(ctx context.Context, request auths.RegisterRequest) error
	ExecuteLogout(ctx context.Context, request auths.LoginRequest) error
}
