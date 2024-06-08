package users

import (
	"context"
	"database/sql"
	"godating-dealls/internal/common"
	ue "godating-dealls/internal/core/entities/users"
	res "godating-dealls/internal/domain"
	"log"
)

type UserUsecase struct {
	Ue ue.UserEntities
	DB *sql.DB
}

func NewUserUsecase(ue ue.UserEntities, db *sql.DB) InputUserBoundary {
	return &UserUsecase{Ue: ue, DB: db}
}

func (u UserUsecase) ExecuteUserViewsUsecase(ctx context.Context, boundary OutputUserBoundary) error {
	fn := func(tx *sql.Tx) error {
		users, err := u.Ue.FindAllUserViewsEntities(ctx, tx)
		common.HandleErrorReturn(err)

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
			})
		}
		boundary.UserViewsResponse(userViews, nil)

		return nil
	}

	err := common.WithReadOnlyTransactionManager(ctx, u.DB, fn)
	if err != nil {
		log.Println("Transaction failed:", err)
	}
	return err
}
