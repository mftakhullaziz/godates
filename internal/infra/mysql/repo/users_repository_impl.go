package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/infra/mysql/queries"
	"godating-dealls/internal/infra/mysql/record"
	"log"
)

type UserRepositoryImpl struct {
	UsersRepository UserRepository
	validate        *validator.Validate
}

func NewUsersRepositoryImpl(validate *validator.Validate) UserRepository {
	return &UserRepositoryImpl{validate: validate}
}

func (u UserRepositoryImpl) CreateUserToDB(ctx context.Context, tx *sql.Tx, userRecord record.UserRecord) (record.UserRecord, error) {
	// Execute the query within the transaction
	result, err := tx.ExecContext(ctx, queries.SaveToUserRecord,
		userRecord.AccountID,
		userRecord.DateOfBirth,
		userRecord.Age,
		userRecord.Gender,
		userRecord.Address,
		userRecord.Bio,
	)

	if err != nil {
		return record.UserRecord{}, fmt.Errorf("could not insert users: %v", err)
	}

	// Log the query result
	log.Printf("Query result: %v", result)

	// Get the last inserted ID
	userId, err := result.LastInsertId()
	if err != nil {
		return record.UserRecord{}, fmt.Errorf("could not retrieve last insert id: %v", err)
	}

	// Set the UserID in the userRecord
	userRecord.UserID = userId
	return userRecord, nil
}

func (u UserRepositoryImpl) FindUserByUserIDFromDB(ctx context.Context, tx *sql.Tx, accountId int64) bool {
	result, err := tx.ExecContext(ctx, queries.FindByAccountIdUserRecord, accountId)
	if err != nil {
		return false
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return false
	}
	return rowCount > 0
}

func (u UserRepositoryImpl) GetUserByAccountIdFromDB(ctx context.Context, tx *sql.Tx, accountId int64) (record.UserRecord, error) {
	// Query the database to find the account record by username
	row := tx.QueryRowContext(ctx, queries.GetUserByAccountIdUserRecord, accountId)

	// Initialize a new AccountRecord to store the result
	var userRecord record.UserRecord
	// Scan the row into the AccountRecord fields
	err := row.Scan(
		&userRecord.UserID,
		&userRecord.AccountID,
		&userRecord.DateOfBirth,
		&userRecord.Age,
		&userRecord.Gender,
		&userRecord.Address,
		&userRecord.Bio,
		&userRecord.CreatedAt,
		&userRecord.UpdatedAt,
	)

	if err != nil {
		// Handle any error that occurred during scanning
		if errors.Is(err, sql.ErrNoRows) {
			// If no rows were found, return a specific error
			return record.UserRecord{}, fmt.Errorf("no account found with accountId: %s", accountId)
		}
		// Return any other error encountered during scanning
		return record.UserRecord{}, fmt.Errorf("error scanning account record: %v", err)
	}

	// Return the retrieved account record
	return userRecord, nil
}
