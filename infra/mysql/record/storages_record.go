package record

import "time"

// StorageRecord represents a user in the system
type StorageRecord struct {
	StorageID  int64     `db:"storage_id"`
	AccountID  int64     `db:"account_id"`
	UserID     string    `db:"user_id"`
	VideosPath string    `db:"videos_path"`
	PhotosPath string    `db:"photos_path"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (StorageRecord) TableName() string {
	return "storages"
}
