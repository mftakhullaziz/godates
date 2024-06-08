package record

import "time"

// UserRecord represents a user profile in the system
type UserRecord struct {
	UserID      int64      `db:"user_id"`
	AccountID   int64      `db:"account_id"`
	FullName    *string    `db:"full_name"`
	DateOfBirth *time.Time `db:"date_of_birth"`
	Age         int        `db:"age"`
	Gender      string     `db:"gender"`
	Address     string     `db:"address"`
	Bio         string     `db:"bio"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
}

func (UserRecord) TableName() string {
	return "users"
}

// UserAccountRecord represents a user profile with additional verified field
type UserAccountRecord struct {
	UserRecord
	Verified bool
	Username string
	FullName string
}
