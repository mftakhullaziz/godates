package queries

const (
	SaveToAccountsRecord = `
		INSERT INTO accounts
		(username, password_hash, email, verified)
		VALUES(?, ?, ?, ?);`

	selectQuery    = ``
	selectQueryAll = ``
)
