package users

import (
	"context"
	"database/sql"
	"godating-dealls/internal/common"
	ae "godating-dealls/internal/core/entities/auths"
	ue "godating-dealls/internal/core/entities/users"
	res "godating-dealls/internal/domain"
	"godating-dealls/internal/infra/jsonwebtoken"
	"log"
)

type UserUsecase struct {
	DB *sql.DB
	Ue ue.UserEntities
	Ae ae.AuthEntities
}

func NewUserUsecase(db *sql.DB, ue ue.UserEntities, ae ae.AuthEntities) InputUserBoundary {
	return &UserUsecase{DB: db, Ue: ue, Ae: ae}
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

		// insert to historical selection users

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

	err := common.WithReadOnlyTransactionManager(ctx, u.DB, fn)
	if err != nil {
		log.Println("Transaction failed:", err)
	}
	return err
}
