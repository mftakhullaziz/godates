package repo

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
)

// SwipesRepositoryImpl struct
type SwipesRepositoryImpl struct {
	SwipesRepository SwipesRepository
}

// NewSwipesRepositoryImpl function to create a new instance of SwipesRepositoryImpl
func NewSwipesRepositoryImpl() SwipesRepository {
	return &SwipesRepositoryImpl{}
}

func (s SwipesRepositoryImpl) InsertSwipesToDB(ctx context.Context, tx *sql.Tx, record record.SwipeRecord) error {
	query := `
		INSERT INTO swipes (account_id, user_id, action, account_id_swipe) 
		VALUES (?, ?, ?, ?)
	`
	_, err := tx.ExecContext(ctx, query, record.AccountID, record.UserID, record.Action, record.AccountIDSwipe)
	return err
}
