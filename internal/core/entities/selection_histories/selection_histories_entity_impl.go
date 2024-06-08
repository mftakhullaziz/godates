package selection_histories

import (
	"context"
	"database/sql"
	"godating-dealls/internal/common"
	"godating-dealls/internal/infra/mysql/record"
	"godating-dealls/internal/infra/mysql/repo"
)

type SelectionHistoryEntityImpl struct {
	SelectionHistoriesRepository repo.SelectionHistoriesRepository
}

func NewSelectionHistoryEntityImpl(selectionHistoriesRepository repo.SelectionHistoriesRepository) SelectionHistoryEntity {
	return &SelectionHistoryEntityImpl{SelectionHistoriesRepository: selectionHistoriesRepository}
}

func (s SelectionHistoryEntityImpl) InsertSelectionHistoryEntity(ctx context.Context, tx *sql.Tx, accountIdIdentifier int64, accountId int64) error {
	err := s.SelectionHistoriesRepository.InsertIntoSelectionHistories(ctx, tx,
		record.SelectionHistoryRecord{
			AccountIdIdentifier: accountIdIdentifier,
			AccountID:           accountId,
		})
	common.HandleErrorReturn(err)
	return nil
}
