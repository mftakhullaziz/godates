package views

import (
	"context"
	"database/sql"
	"godating-dealls/internal/domain"
)

type ViewEntity interface {
	InsertIntoViewAccountEntity(ctx context.Context, tx *sql.Tx, account domain.ViewedAccount) error
	FindTotalViewByAccountID(ctx context.Context, tx *sql.Tx, accountIDView int64) (int, error)
}
