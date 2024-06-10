package users

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain"
)

type UserEntity interface {
	SaveUserEntities(ctx context.Context, tx *sql.Tx, dto domain.UserDto) error
	FindUserEntities(ctx context.Context, tx *sql.Tx, accountId int64) (domain.Users, error)
	FindAllUserEntities(ctx context.Context, tx *sql.Tx) ([]domain.AllUsers, error)
	FindAllUserViewsEntities(ctx context.Context, tx *sql.Tx, verified bool, shouldNext bool, accountIdIdentifier int64) ([]domain.AllUserViews, error)
	FindUserDetailEntity(ctx context.Context, tx *sql.Tx, accountId int64) (domain.Users, error)
}
