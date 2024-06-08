package selection_histories

import (
	"context"
	"database/sql"
)

type SelectionHistoryEntity interface {
	InsertSelectionHistoryEntity(ctx context.Context, tx *sql.Tx, accountId int64) error
}
