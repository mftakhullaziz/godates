package auths

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain"
)

type AuthEntities interface {
	SaveAccountEntities(ctx context.Context, tx *sql.Tx, dto domain.AccountDto) (domain.Accounts, error)
	AuthenticateAccount(ctx context.Context, tx *sql.Tx, dto domain.AccountDto) (domain.Accounts, error)
}
