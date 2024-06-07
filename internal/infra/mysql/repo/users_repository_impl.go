package repo

import (
	"context"
	"database/sql"
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

func (u UserRepositoryImpl) FindUserByUserIDFromDB(ctx context.Context, tx *sql.Tx, id string) (record.UserRecord, error) {
	//TODO implement me
	panic("implement me")
}