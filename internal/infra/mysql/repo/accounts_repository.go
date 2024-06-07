package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type AccountRepository interface {
	CreateAccountToDB(ctx context.Context, tx *sql.Tx, accountRecord record.AccountRecord) (record.AccountRecord, error)
	FindAccountByAccountIDFromDB(ctx context.Context, tx *sql.Tx, id string) (record.AccountRecord, error)
	IsExistAccountByEmailFromDB(ctx context.Context, tx *sql.Tx, email string) bool
	IsExistAccountByUsernameFromDB(ctx context.Context, tx *sql.Tx, username string) bool
}
