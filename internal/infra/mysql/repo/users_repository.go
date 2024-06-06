package repo

import (
	"context"
	"godating-dealls/internal/infra/mysql/record"
)

type UserRepository interface {
	CreateUserToDB(ctx context.Context, userRecord record.UserRecord) (record.UserRecord, error)
	FindUserByUserIDFromDB(ctx context.Context, id string) (record.UserRecord, error)
}
