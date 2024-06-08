package auths

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"godating-dealls/internal/common"
	ae "godating-dealls/internal/core/entities/auths"
	lh "godating-dealls/internal/core/entities/login_histories"
	ue "godating-dealls/internal/core/entities/users"
	"godating-dealls/internal/domain"
	payload "godating-dealls/internal/domain/auths"
	"godating-dealls/internal/domain/users"
	"godating-dealls/internal/infra/jsonwebtoken"
	"godating-dealls/internal/infra/redisclient"
	"log"
)

type AuthUsecase struct {
	DB  *sql.DB
	AE  ae.AuthEntities
	UE  ue.UserEntities
	Rds redisclient.RedisInterface
	LH  lh.LoginHistoriesEntities
}

func NewAuthUsecase(
	db *sql.DB,
	ae ae.AuthEntities,
	ue ue.UserEntities,
	rds redisclient.RedisInterface,
	lh lh.LoginHistoriesEntities) InputAuthBoundary {
	return &AuthUsecase{
		DB:  db,
		AE:  ae,
		UE:  ue,
		Rds: rds,
		LH:  lh,
	}
}

func (au *AuthUsecase) ExecuteLoginUsecase(ctx context.Context, request payload.LoginRequest, boundary OutputAuthBoundary) error {
	fn := func(tx *sql.Tx) error {
		accountDTO := payload.AccountDto{
			Username: &request.Username,
			Password: request.Password,
			Email:    &request.Email,
		}

		account, err := au.AE.AuthenticateAccount(ctx, tx, accountDTO)
		if err != nil {
			return err
		}

		passwordIsValid, err := common.ComparedPassword(account.Password, []byte(request.Password))
		if err != nil {
			return err
		}

		if !passwordIsValid {
			return errors.New("invalid password")
		}

		user, err := au.UE.FindUserEntities(ctx, tx, account.AccountId)
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
		err = au.LH.SaveLoginHistoriesEntities(ctx, tx, loginDto)
		common.HandleErrorWithParam(err, "Failed to save login history")

		// Store token to redis
		redisKey := BuildRedisKey(fmt.Sprintf("access_token:%s:%s", account.AccountId, account.Email))
		err = au.Rds.StoreToRedis(ctx, redisKey, token)
		if err != nil {
			return errors.New("failed to save token")
		}

		res := payload.LoginResponse{
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

func (au *AuthUsecase) ExecuteRegisterUsecase(ctx context.Context, request payload.RegisterRequest, boundary OutputAuthBoundary) error {
	fn := func(tx *sql.Tx) error {
		accountDTO := payload.AccountDto{
			Username: &request.Username,
			Password: request.Password,
			Email:    &request.Email,
		}
		account, err := au.AE.SaveAccountEntities(ctx, tx, accountDTO)
		if err != nil {
			return err
		}

		userDto := users.UserDto{
			AccountID: account.AccountId,
		}
		if err := au.UE.SaveUserEntities(ctx, tx, userDto); err != nil {
			return err
		}

		res := payload.RegisterResponse{
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

		err = au.LH.UpdateLoginHistoriesEntities(ctx, tx, domain.LoginHistoriesDto{
			UserID:    verify.UserId,
			AccountID: verify.AccountId,
		})
		common.HandleErrorReturn(err)

		redisKey := BuildRedisKey(fmt.Sprintf("access_token:%s:%s", verify.AccountId, verify.Email))
		err = au.Rds.ClearFromRedis(ctx, redisKey)
		common.HandleErrorReturn(err)

		res := payload.LogoutResponse{
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

func BuildRedisKey(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}
