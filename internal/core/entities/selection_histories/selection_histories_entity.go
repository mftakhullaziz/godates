package selection_histories

import (
	"context"
	"database/sql"
)

type SelectionHistoryEntity interface {
	InsertSelectionHistoryEntity(ctx context.Context, tx *sql.Tx, accountIdIdentifier int64, accountId int64) error
}
