package swipes

import (
	"context"
	"database/sql"
	"godating-dealls/internal/infra/mysql/record"
	"godating-dealls/internal/infra/mysql/repo"
)

type SwipeEntityImpl struct {
	SwipesRepository repo.SwipesRepository
}

func NewSwipeEntityImpl(swipesRepository repo.SwipesRepository) SwipeEntity {
	return &SwipeEntityImpl{SwipesRepository: swipesRepository}
}

func (s SwipeEntityImpl) InsertSwipeActionEntity(ctx context.Context, tx *sql.Tx, accountId int64, userId int64, action string, accountIdSwipe int64) error {
	var actionType string
	if action == "left" {
		actionType = "PASSED"
	} else {
		actionType = "LIKED"
	}

	err := s.SwipesRepository.InsertSwipesToDB(ctx, tx, record.SwipeRecord{
		AccountID:      accountId,
		UserID:         userId,
		Action:         actionType,
		AccountIDSwipe: accountIdSwipe,
	})
	if err != nil {
		return err
	}
	return nil
}
