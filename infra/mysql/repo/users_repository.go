package repo

import (
	"context"
	"godating-dealls/infra/mysql/record"
)

type UserRepository interface {
	CreateUserToDB(ctx context.Context, accountRecord record.AccountRecord) (record.AccountRecord, error)
	FindUserByUserIDFromDB(ctx context.Context, id string) (record.AccountRecord, error)
}
