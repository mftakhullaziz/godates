package swipes

import (
	"context"
	"godating-dealls/internal/domain"
)

type InputSwipeBoundary interface {
	ExecuteSwipes(ctx context.Context, token string, request domain.SwipeRequest, boundary OutputSwipesBoundary) error
}
