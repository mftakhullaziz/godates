package record

// TaskHistories represents the task_histories table.
type TaskHistories struct {
	TaskID           int64  `json:"task_id"`
	TaskName         string `json:"task_name"`
	LastRunTimestamp int64  `json:"last_run_timestamp"`
}

func (TaskHistories) TableName() string {
	return "task_histories"
}
