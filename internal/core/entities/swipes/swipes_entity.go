package swipes

import (
	"context"
	"database/sql"
)

type SwipeEntity interface {
	InsertSwipeActionEntity(ctx context.Context, tx *sql.Tx, accountId int64, userId int64, action string, accountIdSwipe int64) error
}
