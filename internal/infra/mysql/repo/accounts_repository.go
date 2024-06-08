package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type AccountRepository interface {
	CreateAccountToDB(ctx context.Context, tx *sql.Tx, accountRecord record.AccountRecord) (record.AccountRecord, error)
	FindAccountByUsernameFromDB(ctx context.Context, tx *sql.Tx, username string) (record.AccountRecord, error)
	FindAccountByEmailFromDB(ctx context.Context, tx *sql.Tx, email string) (record.AccountRecord, error)
	FindAccountByUsernameAndEmailFromDB(ctx context.Context, tx *sql.Tx, email string, username string) (record.AccountRecord, error)
	IsExistAccountByEmailFromDB(ctx context.Context, tx *sql.Tx, email string) bool
	IsExistAccountByUsernameFromDB(ctx context.Context, tx *sql.Tx, username string) bool
	FindAccountVerifiedByAccountIdFromDB(ctx context.Context, tx *sql.Tx, accountId int64) (bool, error)
}
