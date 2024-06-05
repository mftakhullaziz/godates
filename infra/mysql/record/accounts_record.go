package record

import "time"

// AccountRecord represents a user in the system
type AccountRecord struct {
	AccountID    int64     `db:"account_id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	Email        string    `db:"email"`
	PhoneNumber  string    `db:"phone_number"`
	Verified     bool      `db:"verified"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (AccountRecord) TableName() string {
	return "accounts"
}
