package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

// SelectionHistoriesRepositoryImpl is the implementation of SelectionHistoriesRepository
type SelectionHistoriesRepositoryImpl struct {
	SelectionHistoriesRepository SelectionHistoriesRepository
}

// NewSelectionHistoriesRepositoryImpl creates a new instance of SelectionHistoriesRepositoryImpl
func NewSelectionHistoriesRepositoryImpl() SelectionHistoriesRepository {
	return &SelectionHistoriesRepositoryImpl{}
}

// InsertIntoSelectionHistories inserts a new record into the selection_histories table
func (s *SelectionHistoriesRepositoryImpl) InsertIntoSelectionHistories(ctx context.Context, tx *sql.Tx, record record.SelectionHistoryRecord) error {
	query := `INSERT INTO selection_histories (account_id, selection_date) VALUES (?, CURDATE())`
	_, err := tx.ExecContext(ctx, query, record.AccountID)
	if err != nil {
		return err
	}
	return nil
}
