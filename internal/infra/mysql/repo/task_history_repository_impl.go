package repo

import (
	"context"
	"database/sql"
	"errors"
)

// TaskHistorySQLRepositoryImpl represents the repository implementation for task history data using SQL database.
type TaskHistorySQLRepositoryImpl struct {
	TaskHistoryRepository TaskHistoryRepository
}

// NewTaskHistorySQLRepository creates a new instance of TaskHistorySQLRepository.
func NewTaskHistorySQLRepository() TaskHistoryRepository {
	return &TaskHistorySQLRepositoryImpl{}
}

// GetLastRunTimestamp retrieves the last run timestamp for a task from the storage within the provided transaction.
func (r *TaskHistorySQLRepositoryImpl) GetLastRunTimestamp(ctx context.Context, taskName string, tx *sql.Tx) (int64, error) {
	var lastRunTimestamp int64
	err := tx.QueryRowContext(ctx, "SELECT last_run_timestamp FROM task_histories WHERE task_name = ?", taskName).Scan(&lastRunTimestamp)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}
	return lastRunTimestamp, nil
}

// UpdateLastRunTimestamp updates the last run timestamp for a task in the storage within the provided transaction.
func (r *TaskHistorySQLRepositoryImpl) UpdateLastRunTimestamp(ctx context.Context, taskName string, timestamp int64, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "INSERT INTO task_histories (task_name, last_run_timestamp) VALUES (?, ?) ON DUPLICATE KEY UPDATE last_run_timestamp = ?", taskName, timestamp, timestamp)
	return err
}
