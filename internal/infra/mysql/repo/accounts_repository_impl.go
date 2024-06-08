package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"godating-dealls/internal/infra/mysql/queries"
	"godating-dealls/internal/infra/mysql/record"
)

type AccountRepositoryImpl struct {
	AccountsRepository AccountRepository
}

func NewAccountsRepositoryImpl() AccountRepository {
	return &AccountRepositoryImpl{}
}

func (a AccountRepositoryImpl) CreateAccountToDB(ctx context.Context, tx *sql.Tx, accountRecord record.AccountRecord) (record.AccountRecord, error) {
	// Execute the query within the provided transaction
	result, err := tx.ExecContext(ctx, queries.SaveToAccountsRecord,
		accountRecord.Username,
		accountRecord.PasswordHash,
		accountRecord.Email,
		accountRecord.Verified,
	)
	if err != nil {
		return record.AccountRecord{}, fmt.Errorf("could not insert account: %v", err)
	}

	// Get the last inserted ID
	accountId, err := result.LastInsertId()
	if err != nil {
		return record.AccountRecord{}, fmt.Errorf("could not retrieve last insert id: %v", err)
	}

	// Set the AccountID in the accountRecord
	accountRecord.AccountID = accountId

	return accountRecord, nil
}

func (a AccountRepositoryImpl) IsExistAccountByEmailFromDB(ctx context.Context, tx *sql.Tx, email string) bool {
	result, err := tx.ExecContext(ctx, queries.FindByEmailAccountRecord, email)
	if err != nil {
		return false
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return false
	}
	return rowCount > 0
}

func (a AccountRepositoryImpl) IsExistAccountByUsernameFromDB(ctx context.Context, tx *sql.Tx, username string) bool {
	result, err := tx.ExecContext(ctx, queries.FindByUsernameAccountRecord, username)
	if err != nil {
		return false
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return false
	}
	return rowCount > 0
}

func (a AccountRepositoryImpl) FindAccountByUsernameFromDB(ctx context.Context, tx *sql.Tx, username string) (record.AccountRecord, error) {
	// Query the database to find the account record by username
	row := tx.QueryRowContext(ctx, queries.GetByUsernameAccountRecord, username)
	// Initialize a new AccountRecord to store the result
	var accountRecord record.AccountRecord
	// Scan the row into the AccountRecord fields
	err := row.Scan(
		&accountRecord.AccountID,
		&accountRecord.Username,
		&accountRecord.PasswordHash,
		&accountRecord.Email,
		&accountRecord.Verified,
		&accountRecord.CreatedAt,
		&accountRecord.UpdatedAt,
	)

	if err != nil {
		// Handle any error that occurred during scanning
		if errors.Is(err, sql.ErrNoRows) {
			// If no rows were found, return a specific error
			return record.AccountRecord{}, fmt.Errorf("no account found with username: %s", username)
		}
		// Return any other error encountered during scanning
		return record.AccountRecord{}, fmt.Errorf("error scanning account record: %v", err)
	}

	// Return the retrieved account record
	return accountRecord, nil
}

func (a AccountRepositoryImpl) FindAccountByEmailFromDB(ctx context.Context, tx *sql.Tx, email string) (record.AccountRecord, error) {
	// Query the database to find the account record by username
	row := tx.QueryRowContext(ctx, queries.GetByEmailAccountRecord, email)
	// Initialize a new AccountRecord to store the result
	var accountRecord record.AccountRecord
	// Scan the row into the AccountRecord fields
	err := row.Scan(
		&accountRecord.AccountID,
		&accountRecord.Username,
		&accountRecord.PasswordHash,
		&accountRecord.Email,
		&accountRecord.Verified,
		&accountRecord.CreatedAt,
		&accountRecord.UpdatedAt,
	)

	if err != nil {
		// Handle any error that occurred during scanning
		if errors.Is(err, sql.ErrNoRows) {
			// If no rows were found, return a specific error
			return record.AccountRecord{}, fmt.Errorf("no account found with email: %s", email)
		}
		// Return any other error encountered during scanning
		return record.AccountRecord{}, fmt.Errorf("error scanning account record: %v", err)
	}

	// Return the retrieved account record
	return accountRecord, nil
}

func (a AccountRepositoryImpl) FindAccountByUsernameAndEmailFromDB(ctx context.Context, tx *sql.Tx, email string, username string) (record.AccountRecord, error) {
	// Query the database to find the account record by username
	row := tx.QueryRowContext(ctx, queries.GetByUsernameAndEmailAccountRecord, email, username)
	// Initialize a new AccountRecord to store the result
	var accountRecord record.AccountRecord
	// Scan the row into the AccountRecord fields
	err := row.Scan(
		&accountRecord.AccountID,
		&accountRecord.Username,
		&accountRecord.PasswordHash,
		&accountRecord.Email,
		&accountRecord.Verified,
		&accountRecord.CreatedAt,
		&accountRecord.UpdatedAt,
	)

	if err != nil {
		// Handle any error that occurred during scanning
		if errors.Is(err, sql.ErrNoRows) {
			// If no rows were found, return a specific error
			return record.AccountRecord{}, fmt.Errorf("no account found with email %s and username: %s", email, username)
		}
		// Return any other error encountered during scanning
		return record.AccountRecord{}, fmt.Errorf("error scanning account record: %v", err)
	}

	// Return the retrieved account record
	return accountRecord, nil
}
