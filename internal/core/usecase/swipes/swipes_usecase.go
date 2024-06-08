package swipes

import (
	"context"
	"database/sql"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/swipes"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/jsonwebtoken"
	"log"
)

type SwipeUsecase struct {
	DB          *sql.DB
	SwipeEntity swipes.SwipeEntity
}

func NewSwipeUsecase(db *sql.DB, swipeEntity swipes.SwipeEntity) InputSwipeBoundary {
	return &SwipeUsecase{DB: db, SwipeEntity: swipeEntity}
}

func (s SwipeUsecase) ExecuteSwipes(ctx context.Context, token string, request domain.SwipeRequest, boundary OutputSwipesBoundary) error {
	fn := func(tx *sql.Tx) error {
		// Verify token is not expired
		claims, err := jsonwebtoken.VerifyJWTToken(token)
		common.HandleErrorReturn(err)

		err = s.SwipeEntity.InsertSwipeActionEntity(ctx, tx, claims.AccountId, claims.UserId, request.ActionType, request.AccountIdSwipe)
		common.HandleErrorReturn(err)

		// if swipe success decrease total quota and remove user from list

		var message string
		if request.ActionType == "left" {
			message = "Account Passed!"
		} else {
			message = "Account Liked!"
		}
		boundary.SwipeResponse(domain.SwipeResponse{
			Message: message,
		}, nil)

		return nil
	}

	err := common.WithExecuteTransactionalManager(ctx, s.DB, fn)
	if err != nil {
		log.Println("Transaction failed:", err)
	}
	return err
}
