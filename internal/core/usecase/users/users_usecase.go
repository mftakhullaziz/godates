package users

import (
	"context"
	"database/sql"
	"godating-dealls/internal/common"
	ae "godating-dealls/internal/core/entities/auths"
	"godating-dealls/internal/core/entities/selection_histories"
	"godating-dealls/internal/core/entities/task_history"
	ue "godating-dealls/internal/core/entities/users"
	res "godating-dealls/internal/domain"
	"godating-dealls/internal/infra/jsonwebtoken"
	"log"
	"time"
)

type UserUsecase struct {
	DB                     *sql.DB
	Ue                     ue.UserEntities
	Ae                     ae.AuthEntities
	SelectionHistoryEntity selection_histories.SelectionHistoryEntity
	TaskHistoryEntity      task_history.TaskHistoryEntity
}

func NewUserUsecase(
	db *sql.DB,
	ue ue.UserEntities,
	ae ae.AuthEntities,
	selectionHistoryEntity selection_histories.SelectionHistoryEntity,
	taskHistoryEntity task_history.TaskHistoryEntity) InputUserBoundary {
	return &UserUsecase{DB: db, Ue: ue, Ae: ae,
		SelectionHistoryEntity: selectionHistoryEntity,
		TaskHistoryEntity:      taskHistoryEntity,
	}
}

func (u UserUsecase) ExecuteUserViewsUsecase(ctx context.Context, token string, boundary OutputUserBoundary) error {
	fn := func(tx *sql.Tx) error {
		// Verify token is not expired
		claims, err := jsonwebtoken.VerifyJWTToken(token)
		common.HandleErrorReturn(err)

		// first find account type by claims if account verified return all, if not just 10 data
		accountId := claims.AccountId
		verifiedAccount, err := u.Ae.FindAccountVerifiedEntities(ctx, tx, accountId)

		users, err := u.Ue.FindAllUserViewsEntities(ctx, tx, verifiedAccount)
		common.HandleErrorReturn(err)

		// Check if the historical selection task should run
		shouldRun, err := u.shouldRunHistoricalSelectionTask(ctx, tx)
		common.HandleErrorReturn(err)

		if shouldRun {
			if len(users) > 0 {
				for _, user := range users {
					err := u.SelectionHistoryEntity.InsertSelectionHistoryEntity(ctx, tx, user.AccountID)
					common.HandleErrorReturn(err)
				}
			}
			// Update the task history to indicate the task has run today
			err = u.TaskHistoryEntity.InsertTaskHistoryEntity(ctx, tx, "historical_selection_task", time.Now().Unix())
			common.HandleErrorReturn(err)
		}

		// Build response
		var userViews []res.UserViewsResponse
		for _, user := range users {
			userViews = append(userViews, res.UserViewsResponse{
				UserID:    user.UserID,
				AccountID: user.AccountID,
				Username:  user.Username,
				FullName:  user.FullName,
				Age:       user.Age,
				Gender:    user.Gender,
				Bio:       user.Bio,
				Verified:  user.Verified,
				Videos:    make([]string, 0),
				Photos:    make([]string, 0),
			})
		}
		boundary.UserViewsResponse(userViews, nil)

		return nil
	}

	err := common.WithExecuteTransactionalManager(ctx, u.DB, fn)
	if err != nil {
		log.Println("Transaction failed:", err)
	}
	return err
}

// shouldRunHistoricalSelectionTask checks if the historical selection task should run today.
func (u UserUsecase) shouldRunHistoricalSelectionTask(ctx context.Context, tx *sql.Tx) (bool, error) {
	// Retrieve the last run timestamp from your storage.
	lastRunTimestamp, err := u.TaskHistoryEntity.GetLatestTaskHistoryEntity(ctx, tx, "historical_selection_task")
	if err != nil {
		return false, err
	}

	// Get the current date.
	today := time.Now().Truncate(24 * time.Hour)

	// If the last run timestamp is before today, return true indicating the task should run.
	return lastRunTimestamp < today.Unix(), nil
}
