package queries

const (
	SaveToAccountsRecord        = `INSERT INTO accounts (username, password_hash, email, verified) VALUES(?, ?, ?, ?);`
	FindByEmailAccountRecord    = `SELECT EXISTS(SELECT 1 FROM accounts WHERE email = ?);`
	FindByUsernameAccountRecord = `SELECT EXISTS(SELECT 1 FROM accounts WHERE username = ?)`
)
