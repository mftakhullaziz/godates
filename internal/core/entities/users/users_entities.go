package users

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain/users"
)

type UserEntities interface {
	SaveUserEntities(ctx context.Context, tx *sql.Tx, dto users.UserDto) error
	FindUserEntities(ctx context.Context, tx *sql.Tx, accountId int64) (users.Users, error)
	FindAllUserEntities(ctx context.Context, tx *sql.Tx) ([]users.AllUsers, error)
}
