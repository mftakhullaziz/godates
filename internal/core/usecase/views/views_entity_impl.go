package views

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/mysql/record"
	"godating-dealls/internal/infra/mysql/repo"
)

type ViewEntityImpl struct {
	ViewAccountsRepository repo.ViewAccountsRepository
}

func NewViewEntityImpl(viewAccountsRepository repo.ViewAccountsRepository) ViewEntity {
	return &ViewEntityImpl{
		ViewAccountsRepository: viewAccountsRepository,
	}
}

func (v ViewEntityImpl) InsertIntoViewAccountEntity(ctx context.Context, tx *sql.Tx, account domain.ViewedAccount) error {
	rec := record.ViewAccountRecord{
		AccountID:       account.AccountID,
		AccountIDViewed: account.AccountIDView,
		UserIDViewed:    account.UserIDView,
	}
	err := v.ViewAccountsRepository.InsertIntoViewAccount(ctx, tx, rec)
	if err != nil {
		return errors.New("failed to insert into view account")
	}
	return nil
}

func (v ViewEntityImpl) FindTotalViewByAccountID(ctx context.Context, tx *sql.Tx, accountIDView int64) (int, error) {
	totalView, err := v.ViewAccountsRepository.FindTotalViewByAccountID(ctx, tx, accountIDView)
	if err != nil {
		return 0, errors.New("failed to find total view by account id")
	}
	return totalView, nil
}
