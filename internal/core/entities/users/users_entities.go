package users

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain"
)

type UserEntities interface {
	SaveUserEntities(ctx context.Context, tx *sql.Tx, dto domain.UserDto) error
	FindUserEntities(ctx context.Context, tx *sql.Tx, accountId int64) (domain.Users, error)
	FindAllUserEntities(ctx context.Context, tx *sql.Tx) ([]domain.AllUsers, error)
	FindAllUserViewsEntities(ctx context.Context, tx *sql.Tx) ([]domain.AllUserViews, error)
}
