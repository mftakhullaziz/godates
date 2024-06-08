package queries

import (
	"context"
	"database/sql"
	"log"
)

const (
	SaveToAccountsRecord                       = `INSERT INTO accounts (username, password_hash, email, verified) VALUES(?, ?, ?, ?);`
	FindByEmailAccountRecord                   = `SELECT EXISTS(SELECT 1 FROM accounts WHERE email = ?);`
	FindByUsernameAccountRecord                = `SELECT EXISTS(SELECT 1 FROM accounts WHERE username = ?)`
	SaveToUserRecord                           = `INSERT INTO users (account_id, date_of_birth, full_name, age, gender, address, bio) VALUES(?, ?, ?, ?, ?, ?, ?);`
	FindByAccountIdUserRecord                  = `SELECT EXISTS(SELECT 1 FROM users WHERE account_id = ?);`
	GetByUsernameAccountRecord                 = `SELECT * FROM accounts WHERE username = ?;`
	GetByEmailAccountRecord                    = `SELECT * FROM accounts WHERE email = ?;`
	GetByUsernameAndEmailAccountRecord         = `SELECT * FROM accounts WHERE username = ? AND email = ?;`
	GetUserByAccountIdUserRecord               = `SELECT * FROM users WHERE account_id = ?`
	SaveLoginHistoryRecord                     = `INSERT INTO login_histories (user_id, account_id) VALUES(?, ?);`
	FindByUserIdAndAccountIdLoginHistoryRecord = `SELECT * FROM login_histories WHERE user_id = ? AND account_id = ? AND logout_at IS NULL`
	UpdateLoginHistoryRecord                   = `UPDATE login_histories SET logout_at = ?, duration_in_seconds = ? WHERE login_histories_id = ?`
	InsertIntoDailyQuotaRecord                 = `INSERT INTO daily_quotas (account_id, swipe_count, total_quota) VALUES (?, ?, ?)`
	FindAllUserAccountsListRecord              = `SELECT a.account_id, u.user_id, a.verified, a.username, u.full_name, u.gender, u.bio, u.age, u.address FROM users u INNER JOIN accounts a ON u.account_id = a.account_id`
)

func ExecuteQuery(ctx context.Context, db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	log.Printf("Executing query: %s", query)

	res, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Printf("Error executing query: %s", err)
		return nil, err
	}

	return res, nil
}
