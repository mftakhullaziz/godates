package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/infra/mysql/queries"
	"godating-dealls/internal/infra/mysql/record"
	"time"
)

type LoginRepositoryImpl struct {
	LoginHistoryRepository LoginHistoriesRepository
	validate               *validator.Validate
}

func NewLoginHistoriesRepositoryImpl(validate *validator.Validate) LoginHistoriesRepository {
	return &LoginRepositoryImpl{validate: validate}
}

func (l LoginRepositoryImpl) CreateLoginHistoryDB(ctx context.Context, tx *sql.Tx, loginRecord record.LoginHistoriesRecord) (record.LoginHistoriesRecord, error) {
	result, err := tx.ExecContext(ctx, queries.SaveLoginHistoryRecord,
		loginRecord.UserID,
		loginRecord.AccountID,
	)

	if err != nil {
		return record.LoginHistoriesRecord{}, fmt.Errorf("could not insert account: %v", err)
	}

	// Get the last inserted ID
	loginId, err := result.LastInsertId()
	if err != nil {
		return record.LoginHistoriesRecord{}, fmt.Errorf("could not retrieve last insert id: %v", err)
	}

	// Set the AccountID in the accountRecord
	loginRecord.LoginHistoriesID = loginId
	return loginRecord, nil
}

func (l LoginRepositoryImpl) UpdateLoginHistoryDB(ctx context.Context, tx *sql.Tx, loginRecord record.LoginHistoriesRecord) (record.LoginHistoriesRecord, error) {
	// Find the login history record by user_id and account_id
	query := queries.FindByUserIdAndAccountIdLoginHistoryRecord
	row := tx.QueryRowContext(ctx, query, loginRecord.UserID, loginRecord.AccountID)

	var existingRecord record.LoginHistoriesRecord
	err := row.Scan(
		&existingRecord.LoginHistoriesID,
		&existingRecord.UserID,
		&existingRecord.AccountID,
		&existingRecord.LoginAt,
		&existingRecord.LogoutAt,
		&existingRecord.DurationInSeconds,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return record.LoginHistoriesRecord{}, fmt.Errorf("no login history record found for user_id: %d and account_id: %d", loginRecord.UserID, loginRecord.AccountID)
		}
		return record.LoginHistoriesRecord{}, fmt.Errorf("could not find login history record: %v", err)
	}

	duration := time.Since(*existingRecord.LoginAt)
	fmt.Println(duration.Seconds())

	// Update the logout_at and user_active_duration fields
	query = queries.UpdateLoginHistoryRecord
	_, err = tx.ExecContext(ctx, query, time.Now(), duration.Seconds(), existingRecord.LoginHistoriesID)
	if err != nil {
		return record.LoginHistoriesRecord{}, fmt.Errorf("could not update login history record: %v", err)
	}

	// Return the updated record
	existingRecord.LogoutAt = loginRecord.LogoutAt
	existingRecord.DurationInSeconds = loginRecord.DurationInSeconds
	return existingRecord, nil
}
