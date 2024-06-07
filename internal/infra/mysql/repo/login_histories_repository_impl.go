package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"godating-dealls/internal/infra/mysql/queries"
	"godating-dealls/internal/infra/mysql/record"
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

func (l LoginRepositoryImpl) UpdateLoginHistoryDB(ctx context.Context, tx *sql.Tx, record record.LoginHistoriesRecord) (record.LoginHistoriesRecord, error) {
	//TODO implement me
	panic("implement me")
}
