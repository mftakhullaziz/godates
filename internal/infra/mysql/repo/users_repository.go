package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

type UserRepository interface {
	CreateUserToDB(ctx context.Context, tx *sql.Tx, userRecord record.UserRecord) (record.UserRecord, error)
	FindUserByUserIDFromDB(ctx context.Context, tx *sql.Tx, id int64) bool
	GetUserByAccountIdFromDB(ctx context.Context, tx *sql.Tx, accountId int64) (record.UserRecord, error)
	GetAllUsersFromDB(ctx context.Context, tx *sql.Tx) ([]record.UserAccountRecord, error)
	GetAllUsersViewsFromDB(ctx context.Context, verifiedUser bool, accountIdIdentifier int64, tx *sql.Tx) ([]record.UserAccountRecord, error)
	GetAllUsersNextViewsFromDB(ctx context.Context, verifiedUser bool, accountId int64, tx *sql.Tx) ([]record.UserAccountRecord, error)
	UpdateUserToDB(ctx context.Context, tx *sql.Tx, userRecord record.UserRecord) (record.UserRecord, error)
}
