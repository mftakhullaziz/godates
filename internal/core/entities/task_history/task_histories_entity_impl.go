package task_history

import (
	"context"
	"database/sql"
	"godating-dealls/internal/common"
	"godating-dealls/internal/infra/mysql/repo"
)

type TaskHistoryEntityImpl struct {
	TaskHistoriesRepository repo.TaskHistoryRepository
}

func NewTaskHistoryEntityImpl(taskHistoriesRepository repo.TaskHistoryRepository) TaskHistoryEntity {
	return &TaskHistoryEntityImpl{TaskHistoriesRepository: taskHistoriesRepository}
}

func (t TaskHistoryEntityImpl) InsertTaskHistoryEntity(ctx context.Context, tx *sql.Tx, taskName string, timestamp int64) error {
	err := t.TaskHistoriesRepository.UpdateLastRunTimestamp(ctx, taskName, timestamp, tx)
	return err
}

func (t TaskHistoryEntityImpl) GetLatestTaskHistoryEntity(ctx context.Context, tx *sql.Tx, taskName string) (int64, error) {
	lastRunTime, err := t.TaskHistoriesRepository.GetLastRunTimestamp(ctx, taskName, tx)
	common.HandleErrorReturn(err)
	return lastRunTime, nil
}
