package auths

import (
	"context"
	"database/sql"
	ae "godating-dealls/internal/core/entities/auths"
	ue "godating-dealls/internal/core/entities/users"
	payload "godating-dealls/internal/domain/auths"
	"godating-dealls/internal/domain/users"
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
	// Start a transaction
	//fn := func(tx *sql.Tx) error {
	//
	//}
	//
	//// Execute the business logic within a transaction
	//return common.WithTransactionalManager(ctx, au.db, fn)

	// Save account entity
	accountDTO := payload.AccountDto{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}
	account, err := au.ae.SaveAccountEntities(ctx, accountDTO)
	if err != nil {
		return err
	}

	// Save user entity
	userDto := users.UserDto{
		AccountID: account.AccountId,
	}
	if err := au.ue.SaveUserEntities(ctx, userDto); err != nil {
		return err
	}

	// Transform to response
	res := payload.RegisterResponse{
		AccountId: account.AccountId,
		Email:     account.Email,
		Username:  account.Username,
		Password:  account.Password,
	}
	boundary.RegisterResponse(res, nil)

	return nil
}

func (au *AuthUsecase) ExecuteLogoutUsecase(ctx context.Context, request payload.LoginRequest, boundary OutputAuthBoundary) error {
	// TODO: Implement logout logic
	return nil
}
