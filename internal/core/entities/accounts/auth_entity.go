package accounts

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain"
)

type AccountEntity interface {
	SaveAccountEntities(ctx context.Context, tx *sql.Tx, dto domain.AccountDto) (domain.Accounts, error)
	AuthenticateAccount(ctx context.Context, tx *sql.Tx, dto domain.AccountDto) (domain.Accounts, error)
	FindAccountVerifiedEntities(ctx context.Context, tx *sql.Tx, accountId int64) (bool, error)
}
