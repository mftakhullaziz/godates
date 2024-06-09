package record

// TaskHistories represents the task_histories table.
type TaskHistories struct {
	TaskID              int64  `json:"task_id"`
	AccountIdIdentifier int64  `json:"account_id_identifier"`
	TaskName            string `json:"task_name"`
	LastRunTimestamp    int64  `json:"last_run_timestamp"`
}

func (TaskHistories) TableName() string {
	return "task_histories"
}
