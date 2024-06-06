package auths

import (
	"context"
	ae "godating-dealls/internal/core/entities/auths"
	ue "godating-dealls/internal/core/entities/users"
	payload "godating-dealls/internal/domain/auths"
	"godating-dealls/internal/domain/users"
	"log"
)

type AuthUsecase struct {
	ae ae.AuthEntities
	ue ue.UserEntities
}

func NewAuthUsecase(ae ae.AuthEntities, ue ue.UserEntities) InputAuthBoundary {
	return &AuthUsecase{
		ae: ae,
		ue: ue,
	}
}

func (au *AuthUsecase) ExecuteLoginUsecase(ctx context.Context, request payload.LoginRequest, boundary OutputAuthBoundary) error {
	// TODO: Implement login logic
	return nil
}

func (au *AuthUsecase) ExecuteRegisterUsecase(ctx context.Context, request payload.RegisterRequest, boundary OutputAuthBoundary) error {
	accountDTO := payload.AccountDto{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}

	account, err := au.ae.SaveAccountEntities(ctx, accountDTO)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	userDto := users.UserDto{
		AccountID: account.AccountId,
	}

	// Save to user record again
	err = au.ue.SaveUserEntities(ctx, userDto)
	if err != nil {
		log.Fatalln(err.Error())
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
