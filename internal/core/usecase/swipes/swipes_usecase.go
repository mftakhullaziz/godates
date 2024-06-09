package swipes

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/auths"
	"godating-dealls/internal/core/entities/daily_quotas"
	"godating-dealls/internal/core/entities/swipes"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/jsonwebtoken"
	"log"
)

type SwipeUsecase struct {
	DB                *sql.DB
	SwipeEntity       swipes.SwipeEntity
	DailyQuotasEntity daily_quotas.DailyQuotasEntity
	AuthEntity        auths.AuthEntities
}

func NewSwipeUsecase(db *sql.DB, swipeEntity swipes.SwipeEntity, dailyQuotasEntity daily_quotas.DailyQuotasEntity, authEntity auths.AuthEntities) InputSwipeBoundary {
	return &SwipeUsecase{DB: db, SwipeEntity: swipeEntity, DailyQuotasEntity: dailyQuotasEntity, AuthEntity: authEntity}
}

func (s SwipeUsecase) ExecuteSwipes(ctx context.Context, token string, request domain.SwipeRequest, boundary OutputSwipesBoundary) error {
	fn := func(tx *sql.Tx) error {
		// Verify token is not expired
		claims, err := jsonwebtoken.VerifyJWTToken(token)
		if err != nil {
			return errors.New("invalid token")
		}

		accountIdIdentifier := claims.AccountId
		verifiedAccount, err := s.AuthEntity.FindAccountVerifiedEntities(ctx, tx, accountIdIdentifier)

		var message string

		if verifiedAccount {
			err := s.SwipeEntity.InsertSwipeActionEntity(ctx, tx, claims.AccountId, claims.UserId, request.ActionType, request.AccountIdSwipe)
			if err != nil {
				return errors.New("failed to insert swipe action entity")
			}

			// just increase swipe count
			err = s.DailyQuotasEntity.UpdateIncreaseSwipeCount(ctx, tx, accountIdIdentifier)
			if err != nil {
				return errors.New("failed to update swipe count")
			}
		} else {
			// before swipe check total quota is not limited
			totalQuotaSwipe, err := s.DailyQuotasEntity.FetchTotalDailyQuotas(ctx, tx, accountIdIdentifier)
			if err != nil {
				return errors.New("failed to fetch total swipe count")
			}
			// if user is limited cannot be swipes
			if totalQuotaSwipe > 0 && totalQuotaSwipe <= 10 {
				err := s.DailyQuotasEntity.UpdateIncreaseSwipeCountAndDecreaseTotalQuota(ctx, tx, accountIdIdentifier)
				if err != nil {
					return errors.New("failed to update swipe count and total count")
				}

				err = s.SwipeEntity.InsertSwipeActionEntity(ctx, tx, claims.AccountId, claims.UserId, request.ActionType, request.AccountIdSwipe)
				if err != nil {
					return errors.New("failed to insert swipe action entity")
				}
			}
			message = "The total quota for swipe users is limited"
		}

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
