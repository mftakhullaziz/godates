package swipes

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain"
)

type SwipeEntity interface {
	InsertSwipeActionEntity(ctx context.Context, tx *sql.Tx, accountId int64, userId int64, action string, accountIdSwipe int64) error
	FindTotalSwipeActionEntity(ctx context.Context, tx *sql.Tx, accountIdSwipe int64) (domain.TotalSwipeAction, error)
}
