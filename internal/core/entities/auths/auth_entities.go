package auths

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain/auths"
)

type AuthEntities interface {
	SaveAccountEntities(ctx context.Context, tx *sql.Tx, dto auths.AccountDto) (auths.Accounts, error)
	AuthenticateAccount(ctx context.Context, tx *sql.Tx, dto auths.AccountDto) (auths.Accounts, error)
}
