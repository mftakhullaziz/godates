package swipes

import res "godating-dealls/internal/domain"

type OutputSwipesBoundary interface {
	SwipeResponse(response res.SwipeResponse, err error)
}
