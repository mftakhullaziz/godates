package task_history

import (
	"context"
	"database/sql"
)

type TaskHistoryEntity interface {
	InsertTaskHistoryEntity(ctx context.Context, tx *sql.Tx, taskName string, timestamp int64) error
	GetLatestTaskHistoryEntity(ctx context.Context, tx *sql.Tx, taskName string) (int64, error)
}
