package auths

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/accounts"
	"godating-dealls/internal/core/entities/login_histories"
	"godating-dealls/internal/core/entities/users"
	"godating-dealls/internal/domain"
	"godating-dealls/internal/infra/jsonwebtoken"
	"godating-dealls/internal/infra/redisclient"
	"log"
)

type AuthUsecase struct {
	DB                   *sql.DB
	AccountEntity        accounts.AccountEntity
	UserEntity           users.UserEntity
	Rds                  redisclient.RedisInterface
	LoginHistoriesEntity login_histories.LoginHistoriesEntity
}

func NewAuthUsecase(
	db *sql.DB,
	accountEntity accounts.AccountEntity,
	userEntity users.UserEntity,
	rds redisclient.RedisInterface,
	loginHistoriesEntity login_histories.LoginHistoriesEntity) InputAuthBoundary {
	return &AuthUsecase{
		DB:                   db,
		AccountEntity:        accountEntity,
		UserEntity:           userEntity,
		Rds:                  rds,
		LoginHistoriesEntity: loginHistoriesEntity,
	}
}

func (au *AuthUsecase) ExecuteLoginUsecase(ctx context.Context, request domain.LoginRequest, boundary OutputAuthBoundary) error {
	fn := func(tx *sql.Tx) error {
		accountDTO := domain.AccountDto{
			Username: &request.Username,
			Password: request.Password,
			Email:    &request.Email,
		}

		account, err := au.AccountEntity.AuthenticateAccount(ctx, tx, accountDTO)
		if err != nil {
			return errors.New("failed to authenticate account")
		}

		passwordIsValid, err := common.ComparedPassword(account.Password, []byte(request.Password))
		if err != nil {
			return errors.New("failed to compare password")
		}

		if !passwordIsValid {
			return errors.New("invalid password")
		}

		user, err := au.UserEntity.FindUserEntities(ctx, tx, account.AccountId)
		if err != nil {
			return errors.New("failed to find user")
		}

		token, err := jsonwebtoken.GenerateJWTToken(user.UserID, account.AccountId, account.Email)
		if err != nil {
			return errors.New("failed to generate JWT token")
		}

		// Store to logins history
		loginDto := domain.LoginHistoriesDto{
			UserID:    user.UserID,
			AccountID: account.AccountId,
		}
		err = au.LoginHistoriesEntity.SaveLoginHistoriesEntities(ctx, tx, loginDto)
		common.HandleErrorWithParam(err, "Failed to save login history")

		// Store token to redis
		redisKey := common.StringEncoder(fmt.Sprintf("access_token:%s:%s", account.AccountId, account.Email))
		err = au.Rds.StoreToRedis(ctx, redisKey, token)
		if err != nil {
			return errors.New("failed to save token")
		}

		res := domain.LoginResponse{
			Username:    account.Username,
			Email:       account.Email,
			AccessToken: token,
		}
		boundary.LoginResponse(res, nil)
		return nil
	}

	err := common.WithExecuteTransactionalManager(ctx, au.DB, fn)
	if err != nil {
		log.Println("Transaction failed:", err)
	}
	return err
}

func (au *AuthUsecase) ExecuteRegisterUsecase(ctx context.Context, request domain.RegisterRequest, boundary OutputAuthBoundary) error {
	fn := func(tx *sql.Tx) error {
		accountDTO := domain.AccountDto{
			Username: &request.Username,
			Password: request.Password,
			Email:    &request.Email,
		}
		common.PrintJSON("auth usecase | account dto", accountDTO)
		account, err := au.AccountEntity.SaveAccountEntities(ctx, tx, accountDTO)
		if err != nil {
			return err
		}

		userDto := domain.UserDto{
			AccountID: account.AccountId,
			FullName:  &request.FullName,
		}
		common.PrintJSON("auth usecase | user dto", userDto)
		if err := au.UserEntity.SaveUserEntities(ctx, tx, userDto); err != nil {
			return err
		}

		res := domain.RegisterResponse{
			AccountId: account.AccountId,
			Email:     account.Email,
			Username:  account.Username,
			Password:  account.Password,
		}
		boundary.RegisterResponse(res, nil)

		return nil
	}

	err := common.WithExecuteTransactionalManager(ctx, au.DB, fn)
	if err != nil {
		log.Println("Transaction failed:", err)
	}
	return err
}

func (au *AuthUsecase) ExecuteLogoutUsecase(ctx context.Context, accessToken *string, boundary OutputAuthBoundary) error {
	fn := func(tx *sql.Tx) error {
		verify, err := jsonwebtoken.VerifyJWTToken(*accessToken)
		common.HandleErrorReturn(err)

		err = au.LoginHistoriesEntity.UpdateLoginHistoriesEntities(ctx, tx, domain.LoginHistoriesDto{
			UserID:    verify.UserId,
			AccountID: verify.AccountId,
		})
		common.HandleErrorReturn(err)

		redisKey := common.StringEncoder(fmt.Sprintf("access_token:%s:%s", verify.AccountId, verify.Email))
		err = au.Rds.ClearFromRedis(ctx, redisKey)
		common.HandleErrorReturn(err)

		res := domain.LogoutResponse{
			Message: "User successfully logged out",
		}
		boundary.LogoutResponse(res, nil)

		return nil
	}

	err := common.WithExecuteTransactionalManager(ctx, au.DB, fn)
	if err != nil {
		log.Println("Transaction failed:", err)
	}
	return err
}
