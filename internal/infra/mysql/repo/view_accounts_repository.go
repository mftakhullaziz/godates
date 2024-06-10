package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type ViewAccountsRepository interface {
	InsertIntoViewAccount(ctx context.Context, tx *sql.Tx, record record.ViewAccountRecord) error
}
