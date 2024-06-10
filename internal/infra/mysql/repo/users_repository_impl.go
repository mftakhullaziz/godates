package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"godating-dealls/internal/common"
	"godating-dealls/internal/infra/mysql/queries"
	"godating-dealls/internal/infra/mysql/record"
)

type UserRepositoryImpl struct {
	UsersRepository UserRepository
}

func NewUsersRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (u UserRepositoryImpl) CreateUserToDB(ctx context.Context, tx *sql.Tx, userRecord record.UserRecord) (record.UserRecord, error) {
	// Execute the query within the transaction
	result, err := tx.ExecContext(ctx, queries.SaveToUserRecord,
		userRecord.AccountID,
		userRecord.DateOfBirth,
		userRecord.FullName,
		userRecord.Age,
		userRecord.Gender,
		userRecord.Address,
		userRecord.Bio,
	)

	if err != nil {
		return record.UserRecord{}, fmt.Errorf("could not insert users: %v", err)
	}

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
		&userRecord.FullName,
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

func (u UserRepositoryImpl) GetAllUsersFromDB(ctx context.Context, tx *sql.Tx) ([]record.UserAccountRecord, error) {
	rows, err := tx.QueryContext(ctx, queries.FindAllUserAccountsListRecord)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	defer rows.Close()

	var users []record.UserAccountRecord
	for rows.Next() {
		var user record.UserAccountRecord
		if err := rows.Scan(
			&user.AccountID,
			&user.UserID,
			&user.Verified,
		); err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return users, nil
}

func (u UserRepositoryImpl) GetAllUsersViewsFromDB(ctx context.Context, verifiedUser bool, accountIdIdentifier int64, tx *sql.Tx) ([]record.UserAccountRecord, error) {
	var query string
	if verifiedUser {
		query = queries.FindAllUserAccountsViewInPremiumFirstListRecord
	} else {
		query = queries.FindAllUserAccountsView10InFirstHitListRecord
	}
	common.PrintJSON("printed query for daily views", query)

	rows, err := tx.QueryContext(ctx, query, accountIdIdentifier)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	defer rows.Close()

	var users []record.UserAccountRecord
	for rows.Next() {
		var user record.UserAccountRecord
		if err := rows.Scan(
			&user.AccountID,
			&user.UserID,
			&user.Verified,
			&user.Username,
			&user.FullName,
			&user.Gender,
			&user.Bio,
			&user.Age,
			&user.Address,
		); err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return users, nil
}

func fetchSecondAllUsersViewInPremiumUser(ctx context.Context, tx *sql.Tx, identifier int64) (*sql.Rows, error) {
	query := queries.FindAllUserAccountsViewInPremiumSecondListRecord
	rows, err := tx.QueryContext(ctx, query, identifier, identifier)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	return rows, nil
}

func fetchSecondAllUsersViewInRegularUser(ctx context.Context, tx *sql.Tx, identifier int64) (*sql.Rows, error) {
	query := queries.FindAllUserAccountsView10InSecondHitListRecord
	rows, err := tx.QueryContext(ctx, query, identifier, identifier, identifier)
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	return rows, nil
}

func (u UserRepositoryImpl) GetAllUsersNextViewsFromDB(ctx context.Context, verifiedUser bool, accountIdIdentifier int64, tx *sql.Tx) ([]record.UserAccountRecord, error) {
	var rows *sql.Rows
	var err error

	if verifiedUser {
		rows, err = fetchSecondAllUsersViewInPremiumUser(ctx, tx, accountIdIdentifier)
	} else {
		rows, err = fetchSecondAllUsersViewInRegularUser(ctx, tx, accountIdIdentifier)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []record.UserAccountRecord
	for rows.Next() {
		var user record.UserAccountRecord
		if err := rows.Scan(
			&user.AccountID,
			&user.UserID,
			&user.Verified,
			&user.Username,
			&user.FullName,
			&user.Gender,
			&user.Bio,
			&user.Age,
			&user.Address,
		); err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return users, nil
}

func (u UserRepositoryImpl) UpdateUserToDB(ctx context.Context, tx *sql.Tx, userRecord record.UserRecord) (record.UserRecord, error) {
	query := `
		UPDATE users
		SET full_name = ?, date_of_birth = ?, age = ?, gender = ?, address = ?, bio = ?, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = ?;
	`

	_, err := tx.ExecContext(ctx, query,
		userRecord.FullName,
		userRecord.DateOfBirth,
		userRecord.Age,
		userRecord.Gender,
		userRecord.Address,
		userRecord.Bio,
		userRecord.UserID,
	)

	if err != nil {
		return record.UserRecord{}, errors.New("failed to update user record: " + err.Error())
	}

	// Fetch the updated user record
	updatedUserRecord, err := u.findUserByID(ctx, tx, userRecord.UserID)
	if err != nil {
		return record.UserRecord{}, errors.New("failed to fetch updated user record: " + err.Error())
	}

	return updatedUserRecord, nil
}

func (u UserRepositoryImpl) findUserByID(ctx context.Context, tx *sql.Tx, userID int64) (record.UserRecord, error) {
	query := `
		SELECT user_id, account_id, full_name, date_of_birth, age, gender, address, bio, created_at, updated_at
		FROM users
		WHERE user_id = ?;
	`

	row := tx.QueryRowContext(ctx, query, userID)
	var userRecord record.UserRecord
	err := row.Scan(
		&userRecord.UserID,
		&userRecord.AccountID,
		&userRecord.FullName,
		&userRecord.DateOfBirth,
		&userRecord.Age,
		&userRecord.Gender,
		&userRecord.Address,
		&userRecord.Bio,
		&userRecord.CreatedAt,
		&userRecord.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record.UserRecord{}, errors.New("user not found")
		}
		return record.UserRecord{}, err
	}

	return userRecord, nil
}
