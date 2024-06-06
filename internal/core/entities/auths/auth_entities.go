package auths

import (
	"context"
	"godating-dealls/internal/domain/auths"
)

type AuthEntities interface {
	SaveAccountEntities(ctx context.Context, dto auths.AccountDto) (auths.Accounts, error)
}
