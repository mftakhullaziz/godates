package repo

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/common"
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

func (s SwipesRepositoryImpl) FindTotalSwipes(ctx context.Context, tx *sql.Tx, accountIdSwipe int64) (record.SwipeActionsRecord, error) {
	query := "SELECT SUM(CASE WHEN s.action = 'LIKED' THEN 1 ELSE 0 END) as total_swipe_like, SUM(CASE WHEN s.action = 'PASSED' THEN 1 ELSE 0 END) as total_swipe_pass FROM swipes s WHERE s.account_id_swipe = ?"
	common.PrintJSON("find total swipes", query)

	rows, err := tx.QueryContext(ctx, query, accountIdSwipe)
	if err != nil {
		return record.SwipeActionsRecord{}, errors.New("error fetching total swipes")
	}
	defer rows.Close()

	var swipeActions record.SwipeActionsRecord
	if rows.Next() {
		err := rows.Scan(
			&swipeActions.TotalSwipeLike,
			&swipeActions.TotalSwipePass,
		)

		if err != nil {
			return record.SwipeActionsRecord{}, errors.New("error scanning total swipes")
		}
	}

	return swipeActions, nil
}
