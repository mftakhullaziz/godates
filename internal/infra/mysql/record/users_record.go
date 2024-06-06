package record

import "time"

// UserRecord represents a user profile in the system
type UserRecord struct {
	UserID    int64     `db:"user_id"`
	AccountID int64     `db:"account_id"`
	Age       int       `db:"age"`
	Gender    string    `db:"gender"`
	Address   string    `db:"address"`
	Bio       string    `db:"bio"`
	Photos    []string  `db:"photos"` // Adjust this according to your JSONB handling
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (UserRecord) TableName() string {
	return "users"
}
