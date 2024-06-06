package users

import (
	"context"
	"godating-dealls/internal/domain/users"
)

type UserEntities interface {
	SaveUserEntities(ctx context.Context, dto users.UserDto) error
}
