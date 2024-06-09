package queries

import (
	"context"
	"database/sql"
	"log"
)

const (
	SaveToAccountsRecord                             = `INSERT INTO accounts (username, password_hash, email, verified) VALUES(?, ?, ?, ?);`
	FindByEmailAccountRecord                         = `SELECT EXISTS(SELECT 1 FROM accounts WHERE email = ?);`
	FindByUsernameAccountRecord                      = `SELECT EXISTS(SELECT 1 FROM accounts WHERE username = ?)`
	SaveToUserRecord                                 = `INSERT INTO users (account_id, date_of_birth, full_name, age, gender, address, bio) VALUES(?, ?, ?, ?, ?, ?, ?);`
	FindByAccountIdUserRecord                        = `SELECT EXISTS(SELECT 1 FROM users WHERE account_id = ?);`
	GetByUsernameAccountRecord                       = `SELECT * FROM accounts WHERE username = ?;`
	GetByEmailAccountRecord                          = `SELECT * FROM accounts WHERE email = ?;`
	GetByUsernameAndEmailAccountRecord               = `SELECT * FROM accounts WHERE username = ? AND email = ?;`
	GetUserByAccountIdUserRecord                     = `SELECT * FROM users WHERE account_id = ?`
	SaveLoginHistoryRecord                           = `INSERT INTO login_histories (user_id, account_id) VALUES(?, ?);`
	FindByUserIdAndAccountIdLoginHistoryRecord       = `SELECT * FROM login_histories WHERE user_id = ? AND account_id = ? AND logout_at IS NULL`
	UpdateLoginHistoryRecord                         = `UPDATE login_histories SET logout_at = ?, duration_in_seconds = ? WHERE login_histories_id = ?`
	InsertIntoDailyQuotaRecord                       = `INSERT INTO daily_quotas (account_id, swipe_count, total_quota) VALUES (?, ?, ?)`
	FindAllUserAccountsListRecord                    = `SELECT a.account_id, u.user_id, a.verified FROM users u INNER JOIN accounts a ON u.account_id = a.account_id`
	FindAllUserAccountsViewInPremiumFirstListRecord  = `SELECT a.account_id, u.user_id, a.verified, a.username, u.full_name, u.gender, u.bio, u.age, u.address FROM users u INNER JOIN accounts a ON u.account_id = a.account_id WHERE a.account_id != ?`
	FindAllUserAccountsViewInPremiumSecondListRecord = `SELECT a.account_id, u.user_id, a.verified, a.username, u.full_name, u.gender, u.bio, u.age, u.address FROM selection_histories sh INNER JOIN users u ON sh.account_id = u.account_id INNER JOIN accounts a ON sh.account_id = a.account_id WHERE sh.account_id_identifier = ? AND a.account_id != ? AND a.account_id NOT IN ( SELECT s.account_id_swipe from swipes s WHERE s.account_id = ? ) ORDER BY RAND();`
	FindAllUserAccountsView10InFirstHitListRecord    = `SELECT a.account_id, u.user_id, a.verified, a.username, u.full_name, u.gender, u.bio, u.age, u.address FROM users u INNER JOIN accounts a ON u.account_id = a.account_id WHERE a.verified = FALSE AND a.account_id != ? AND a.account_id NOT IN (SELECT DISTINCT sh2.account_id_identifier FROM selection_histories sh2 WHERE sh2.selection_date = CURDATE()) ORDER BY RAND() LIMIT 10;`
	FindAllUserAccountsView10InSecondHitListRecord   = `SELECT a.account_id, u.user_id, a.verified, a.username, u.full_name, u.gender, u.bio, u.age, u.address FROM users u INNER JOIN accounts a ON u.account_id = a.account_id INNER JOIN selection_histories sh ON a.account_id = sh.account_id AND u.account_id = sh.account_id AND sh.selection_date = CURDATE() WHERE a.verified = FALSE AND sh.account_id_identifier = ? AND a.account_id != ? AND a.account_id NOT IN (SELECT s.account_id_swipe from swipes s WHERE s.account_id = ?) ORDER BY RAND() LIMIT 10;`
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
