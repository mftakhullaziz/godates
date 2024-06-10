package accounts

import (
	"context"
	"database/sql"
	"godating-dealls/internal/core/entities/accounts"
	"godating-dealls/internal/core/entities/selection_histories"
)

type AccountsUsecase struct {
	Db                     *sql.DB
	AccountEntity          accounts.AccountEntity
	SelectionHistoryEntity selection_histories.SelectionHistoryEntity
}

func NewAccountsUsecase(
	db *sql.DB,
	accountEntity accounts.AccountEntity,
	selectionHistoryEntity selection_histories.SelectionHistoryEntity) InputAccountBoundary {
	return &AccountsUsecase{
		Db:                     db,
		AccountEntity:          accountEntity,
		SelectionHistoryEntity: selectionHistoryEntity,
	}
}

func (a AccountsUsecase) ExecuteFetchAccountDetail(ctx context.Context, token string, boundary OutputAccountBoundary) error {
	//TODO implement me
	panic("implement me")
}
