package repo

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/infra/mysql/record"
)

type ViewAccountsRepositoryImpl struct {
	ViewAccountsRepository ViewAccountsRepository
}

func NewViewAccountsRepositoryImpl() ViewAccountsRepository {
	return &ViewAccountsRepositoryImpl{}
}

func (v ViewAccountsRepositoryImpl) InsertIntoViewAccount(ctx context.Context, tx *sql.Tx, record record.ViewAccountRecord) error {
	query := "INSERT INTO view_accounts (account_id, user_id_viewed, account_id_viewed) VALUES (?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, record.AccountID, record.UserIDViewed, record.AccountIDViewed)
	if err != nil {
		return errors.New("error while executing insert into view_accounts")
	}
	return nil
}
