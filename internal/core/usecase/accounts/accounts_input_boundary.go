package accounts

import (
	"context"
	"godating-dealls/internal/domain"
)

type InputAccountBoundary interface {
	ExecuteFetchAccountDetail(ctx context.Context, token string, boundary OutputAccountBoundary) error
	ExecuteViewAccountDetail(ctx context.Context, token string, request domain.ViewedAccountRequest, boundary OutputAccountBoundary) error
}
