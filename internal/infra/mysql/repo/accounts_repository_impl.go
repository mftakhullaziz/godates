package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/infra/mysql/queries"
	"godating-dealls/internal/infra/mysql/record"
)

type AccountRepositoryImpl struct {
	AccountsRepository AccountRepository
	validate           *validator.Validate
}

func NewAccountsRepositoryImpl(validate *validator.Validate) AccountRepository {
	return &AccountRepositoryImpl{validate: validate}
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

func (a AccountRepositoryImpl) FindAccountByAccountIDFromDB(ctx context.Context, tx *sql.Tx, id string) (record.AccountRecord, error) {
	//TODO implement me
	panic("implement me")
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
