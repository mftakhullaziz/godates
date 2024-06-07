package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type LoginHistoriesRepository interface {
	CreateLoginHistoryDB(ctx context.Context, tx *sql.Tx, record record.LoginHistoriesRecord) (record.LoginHistoriesRecord, error)
	UpdateLoginHistoryDB(ctx context.Context, tx *sql.Tx, record record.LoginHistoriesRecord) (record.LoginHistoriesRecord, error)
}
