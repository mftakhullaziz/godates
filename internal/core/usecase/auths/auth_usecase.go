package auths

import (
	"context"
	"database/sql"
	"godating-dealls/internal/common"
	ae "godating-dealls/internal/core/entities/auths"
	ue "godating-dealls/internal/core/entities/users"
	payload "godating-dealls/internal/domain/auths"
	"godating-dealls/internal/domain/users"
	"log"
)

type AuthUsecase struct {
	db *sql.DB
	ae ae.AuthEntities
	ue ue.UserEntities
}

func NewAuthUsecase(db *sql.DB, ae ae.AuthEntities, ue ue.UserEntities) InputAuthBoundary {
	return &AuthUsecase{
		db: db,
		ae: ae,
		ue: ue,
	}
}

func (au *AuthUsecase) ExecuteLoginUsecase(ctx context.Context, request payload.LoginRequest, boundary OutputAuthBoundary) error {
	// TODO: Implement login logic
	return nil
}

func (au *AuthUsecase) ExecuteRegisterUsecase(ctx context.Context, request payload.RegisterRequest, boundary OutputAuthBoundary) error {
	fn := func(tx *sql.Tx) error {
		accountDTO := payload.AccountDto{
			Username: request.Username,
			Password: request.Password,
			Email:    request.Email,
		}
		account, err := au.ae.SaveAccountEntities(ctx, tx, accountDTO)
		if err != nil {
			return err
		}

		userDto := users.UserDto{
			AccountID: account.AccountId,
		}
		if err := au.ue.SaveUserEntities(ctx, tx, userDto); err != nil {
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

	err := common.WithTransactionalManager(ctx, au.db, fn)
	if err != nil {
		log.Println("Transaction failed:", err)
	}
	return err
}

func (au *AuthUsecase) ExecuteLogoutUsecase(ctx context.Context, request payload.LoginRequest, boundary OutputAuthBoundary) error {
	// TODO: Implement logout logic
	return nil
}
