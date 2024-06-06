package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/common"
	"godating-dealls/internal/infra/mysql/queries"
	"godating-dealls/internal/infra/mysql/record"
)

type AccountRepositoryImpl struct {
	AccountsRepository AccountRepository
	query              *sql.DB
	validate           *validator.Validate
}

func NewAccountsRepositoryImpl(query *sql.DB, validate *validator.Validate) AccountRepository {
	return &AccountRepositoryImpl{query: query, validate: validate}
}

func (a AccountRepositoryImpl) CreateAccountToDB(ctx context.Context, accountRecord record.AccountRecord) (record.AccountRecord, error) {
	// Begin a new transaction
	tx, err := a.query.BeginTx(ctx, nil)
	if err != nil {
		return record.AccountRecord{}, fmt.Errorf("could not begin transaction: %v", err)
	}

	// Ensure to rollback the transaction in case of an error
	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}()

	// Execute the query within the transaction
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

	// Commit the transaction if no errors occur
	if err := tx.Commit(); err != nil {
		return record.AccountRecord{}, fmt.Errorf("could not commit transaction: %v", err)
	}

	// Set the AccountID in the accountRecord
	accountRecord.AccountID = accountId

	return accountRecord, nil
}

func (a AccountRepositoryImpl) FindAccountByAccountIDFromDB(ctx context.Context, id string) (record.AccountRecord, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountRepositoryImpl) IsExistAccountByEmailFromDB(ctx context.Context, email string) bool {
	tx, err := a.query.BeginTx(ctx, nil)
	common.HandleErrorWithParam(err, "Could not begin transaction")
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

func (a AccountRepositoryImpl) IsExistAccountByUsernameFromDB(ctx context.Context, username string) bool {
	tx, err := a.query.BeginTx(ctx, nil)
	common.HandleErrorWithParam(err, "Could not begin transaction")
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
