package users

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/accounts"
	"godating-dealls/internal/core/entities/selection_histories"
	"godating-dealls/internal/core/entities/task_history"
	"godating-dealls/internal/core/entities/users"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/jsonwebtoken"
	"log"
	"time"
)

type UserUsecase struct {
	DB                     *sql.DB
	UserEntity             users.UserEntity
	AccountEntity          accounts.AccountEntity
	SelectionHistoryEntity selection_histories.SelectionHistoryEntity
	TaskHistoryEntity      task_history.TaskHistoryEntity
}

func NewUserUsecase(
	db *sql.DB,
	userEntity users.UserEntity,
	accountEntity accounts.AccountEntity,
	selectionHistoryEntity selection_histories.SelectionHistoryEntity,
	taskHistoryEntity task_history.TaskHistoryEntity) InputUserBoundary {
	return &UserUsecase{
		DB:                     db,
		UserEntity:             userEntity,
		AccountEntity:          accountEntity,
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
		accountIdIdentifier := claims.AccountId
		verifiedAccount, err := u.AccountEntity.FindAccountVerifiedEntities(ctx, tx, accountIdIdentifier)

		// Check if the historical selection task should run
		shouldRun, err := u.shouldRunHistoricalSelectionTask(ctx, tx, accountIdIdentifier)
		common.HandleErrorReturn(err)

		usersList, err := u.UserEntity.FindAllUserViewsEntities(ctx, tx, verifiedAccount, shouldRun, accountIdIdentifier)
		common.HandleErrorReturn(err)

		if shouldRun {
			if len(usersList) > 0 {
				for _, user := range usersList {
					err := u.SelectionHistoryEntity.InsertSelectionHistoryEntity(ctx, tx, accountIdIdentifier, user.AccountID)
					common.HandleErrorReturn(err)
				}
			}
			// Update the task history to indicate the task has run today
			err = u.TaskHistoryEntity.InsertTaskHistoryEntity(ctx, tx, "historical_selection_task", time.Now().Unix(), accountIdIdentifier)
			common.HandleErrorReturn(err)
		}

		// Build response
		var userViews []domain.UserViewsResponse
		for _, user := range usersList {
			userViews = append(userViews, domain.UserViewsResponse{
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
func (u UserUsecase) shouldRunHistoricalSelectionTask(ctx context.Context, tx *sql.Tx, accountIdIdentifier int64) (bool, error) {
	// Retrieve the last run timestamp from your storage.
	lastRunTimestamp, err := u.TaskHistoryEntity.GetLatestTaskHistoryEntity(ctx, tx, "historical_selection_task", accountIdIdentifier)
	if err != nil {
		return false, err
	}

	// Get the current date.
	today := time.Now().Truncate(24 * time.Hour)

	// If the last run timestamp is before today, return true indicating the task should run.
	return lastRunTimestamp < today.Unix(), nil
}

func (u UserUsecase) ExecutePatchUserUsecase(ctx context.Context, token string, request domain.PatchUserRequest, boundary OutputUserBoundary) error {
	fn := func(tx *sql.Tx) error {
		// Verify token is not expired
		claims, err := jsonwebtoken.VerifyJWTToken(token)
		if err != nil {
			return errors.New("invalid token")
		}
		userID := claims.UserId

		patch := domain.PatchUser{
			UserID:      userID,
			FullName:    request.FullName,
			Gender:      request.Gender,
			Bio:         request.Bio,
			Address:     request.Address,
			DateOfBirth: request.DateOfBirth,
		}
		res, err := u.UserEntity.UpdateUserEntities(ctx, tx, patch)
		boundary.PatchUserResponse(domain.PatchUserResponse{
			UserID:      res.UserID,
			AccountID:   res.AccountID,
			FullName:    res.FullName,
			Gender:      res.Gender,
			Bio:         res.Bio,
			Address:     res.Address,
			Age:         int(res.Age),
			DateOfBirth: common.FormatFromTimeToStr(res.DateOfBirth),
			UpdatedAt:   common.FormatTimeByParam(*res.UpdatedAt),
		}, nil)

		return nil
	}

	err := common.WithExecuteTransactionalManager(ctx, u.DB, fn)
	if err != nil {
		return errors.New("execute transactional failed: " + err.Error())
	}
	return err
}
