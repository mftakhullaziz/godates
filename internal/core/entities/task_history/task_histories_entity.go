package task_history

import (
	"context"
	"database/sql"
)

type TaskHistoryEntity interface {
	InsertTaskHistoryEntity(ctx context.Context, tx *sql.Tx, taskName string, timestamp int64, accountIdIdentifier int64) error
	GetLatestTaskHistoryEntity(ctx context.Context, tx *sql.Tx, taskName string, accountIdIdentifier int64) (int64, error)
}
