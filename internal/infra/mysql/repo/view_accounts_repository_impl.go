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

func (v ViewAccountsRepositoryImpl) FindTotalViewByAccountID(ctx context.Context, tx *sql.Tx, accountID int64) (int, error) {
	query := "SELECT COUNT(*) as total_view FROM view_accounts WHERE account_id_viewed = ?"
	row := tx.QueryRowContext(ctx, query, accountID)
	var totalViews int
	err := row.Scan(&totalViews)
	if err != nil {
		return 0, errors.New("error fetching total views")
	}

	return totalViews, nil
}
