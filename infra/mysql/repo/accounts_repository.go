package repo

import (
	"context"
	"godating-dealls/infra/mysql/record"
)

type AccountRepository interface {
	CreateAccountToDB(ctx context.Context, accountRecord record.AccountRecord) (record.AccountRecord, error)
	FindAccountByAccountIDFromDB(ctx context.Context, id string) (record.AccountRecord, error)
	IsExistAccountByEmailFromDB(ctx context.Context, email string) bool
	IsExistAccountByUsernameFromDB(ctx context.Context, username string) bool
}
