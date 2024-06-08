package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type SelectionHistoriesRepository interface {
	InsertIntoSelectionHistories(ctx context.Context, tx *sql.Tx, record record.SelectionHistoryRecord) error
}
