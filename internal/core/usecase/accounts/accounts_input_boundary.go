package accounts

import "context"

type InputAccountBoundary interface {
	ExecuteFetchAccountDetail(ctx context.Context, token string, boundary OutputAccountBoundary) error
}
