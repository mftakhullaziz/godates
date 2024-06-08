package repo

import (
	"context"
	"database/sql"
)

// TaskHistoryRepository defines the methods for interacting with task history data.
type TaskHistoryRepository interface {
	GetLastRunTimestamp(ctx context.Context, taskName string, tx *sql.Tx) (int64, error)
	UpdateLastRunTimestamp(ctx context.Context, taskName string, timestamp int64, tx *sql.Tx) error
}
