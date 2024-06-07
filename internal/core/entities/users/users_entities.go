package users

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain/users"
)

type UserEntities interface {
	SaveUserEntities(ctx context.Context, tx *sql.Tx, dto users.UserDto) error
}
