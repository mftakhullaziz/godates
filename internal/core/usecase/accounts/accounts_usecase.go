package accounts

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/accounts"
	"godating-dealls/internal/core/entities/swipes"
	"godating-dealls/internal/core/entities/users"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/jsonwebtoken"
)

type AccountUsecase struct {
	Db            *sql.DB
	AccountEntity accounts.AccountEntity
	SwipeEntity   swipes.SwipeEntity
	UserEntity    users.UserEntity
}

func NewAccountsUsecase(
	db *sql.DB,
	accountEntity accounts.AccountEntity,
	swipeEntity swipes.SwipeEntity,
	userEntity users.UserEntity) InputAccountBoundary {
	return &AccountUsecase{
		Db:            db,
		AccountEntity: accountEntity,
		SwipeEntity:   swipeEntity,
		UserEntity:    userEntity,
	}
}

func (a AccountUsecase) ExecuteFetchAccountDetail(ctx context.Context, token string, boundary OutputAccountBoundary) error {
	fn := func(tx *sql.Tx) error {
		claims, err := jsonwebtoken.VerifyJWTToken(token)
		if err != nil {
			return errors.New("invalid token")
		}

		account, err := a.AccountEntity.FindAccountDetails(ctx, tx, claims.AccountId)
		if err != nil {
			return errors.New("invalid fetch account")
		}

		user, err := a.UserEntity.FindUserDetailEntity(ctx, tx, claims.AccountId)
		if err != nil {
			return errors.New("invalid fetch users")
		}

		accountData := domain.AccountDataResponse{
			UserID:      user.UserID,
			AccountID:   account.AccountId,
			Username:    account.Username,
			Email:       account.Email,
			Verified:    account.Verified,
			Age:         user.Age,
			Gender:      user.Gender,
			Address:     user.Address,
			Bio:         user.Bio,
			FullName:    user.FullName,
			DateOfBirth: user.DateOfBirth,
		}

		res := domain.AccountResponse{
			AccountData: accountData,
		}
		boundary.AccountDetailResponse(res, nil)
		return nil
	}

	err := common.WithReadOnlyTransactionManager(ctx, a.Db, fn)
	if err != nil {
		return errors.New("error executing transaction")
	}
	return nil
}
