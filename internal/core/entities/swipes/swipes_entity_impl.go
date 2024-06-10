package swipes

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/domain"
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

func (s SwipeEntityImpl) FindTotalSwipeActionEntity(ctx context.Context, tx *sql.Tx, accountIdSwipe int64) (domain.TotalSwipeAction, error) {
	swipeTotal, err := s.SwipesRepository.FindTotalSwipes(ctx, tx, accountIdSwipe)
	if err != nil {
		return domain.TotalSwipeAction{}, errors.New("swipe not found")
	}
	res := domain.TotalSwipeAction{
		TotalSwipeLike:   0,
		TotalSwipePassed: 0,
	}

	if swipeTotal.TotalSwipeLike != nil || swipeTotal.TotalSwipePass != nil {
		res.TotalSwipeLike = *swipeTotal.TotalSwipeLike
		res.TotalSwipePassed = *swipeTotal.TotalSwipePass
	}

	return res, nil
}
