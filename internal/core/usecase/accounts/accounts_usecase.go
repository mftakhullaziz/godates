package accounts

import (
	"context"
	"database/sql"
	"errors"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/accounts"
	"godating-dealls/internal/core/entities/swipes"
	"godating-dealls/internal/core/entities/users"
	"godating-dealls/internal/core/usecase/views"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/jsonwebtoken"
)

type AccountUsecase struct {
	Db            *sql.DB
	AccountEntity accounts.AccountEntity
	SwipeEntity   swipes.SwipeEntity
	UserEntity    users.UserEntity
	ViewEntity    views.ViewEntity
}

func NewAccountsUsecase(
	db *sql.DB,
	accountEntity accounts.AccountEntity,
	swipeEntity swipes.SwipeEntity,
	userEntity users.UserEntity,
	viewEntity views.ViewEntity) InputAccountBoundary {
	return &AccountUsecase{
		Db:            db,
		AccountEntity: accountEntity,
		SwipeEntity:   swipeEntity,
		UserEntity:    userEntity,
		ViewEntity:    viewEntity,
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

		swipe, err := a.SwipeEntity.FindTotalSwipeActionEntity(ctx, tx, claims.AccountId)
		if err != nil {
			return errors.New("invalid fetch swipes")
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

		accountView := domain.AccountViewResponse{
			TotalSwipeLike:   swipe.TotalSwipeLike,
			TotalSwipePassed: swipe.TotalSwipePassed,
		}

		res := domain.AccountResponse{
			AccountData: accountData,
			AccountView: accountView,
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

func (a AccountUsecase) ExecuteViewAccountDetail(ctx context.Context, token string, request domain.ViewedAccountRequest, boundary OutputAccountBoundary) error {
	fn := func(tx *sql.Tx) error {
		_, err := jsonwebtoken.VerifyJWTToken(token)
		if err != nil {
			return errors.New("invalid token")
		}

		account, err := a.AccountEntity.FindAccountDetails(ctx, tx, request.AccountIDView)
		if err != nil {
			return errors.New("invalid fetch accounts")
		}

		user, err := a.UserEntity.FindUserDetailEntity(ctx, tx, request.AccountIDView)

		// update to account view
		rec := domain.ViewedAccount{
			AccountIDView: request.AccountIDView,
			UserID:        user.UserID,
		}
		err = a.ViewEntity.InsertIntoViewAccountEntity(ctx, tx, rec)
		if err != nil {
			return errors.New("invalid post views account")
		}

		res := domain.ViewedAccountResponse{
			AccountID: account.AccountId,
			UserName:  account.Username,
			Email:     account.Email,
			Verified:  account.Verified,
			FullName:  user.FullName,
		}
		boundary.ViewAccountResponse(res, nil)

		return nil
	}

	err := common.WithExecuteTransactionalManager(ctx, a.Db, fn)
	if err != nil {
		return errors.New("error executing transaction")
	}
	return nil
}
