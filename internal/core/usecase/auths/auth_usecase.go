package auths

import (
	"context"
	entities "godating-dealls/internal/core/entities/auths"
	payload "godating-dealls/internal/domain/auths"
	"log"
)

type AuthUsecase struct {
	entities entities.AuthEntities
}

func NewAuthUsecase(entities entities.AuthEntities) InputAuthBoundary {
	return &AuthUsecase{
		entities: entities,
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

	account, err := au.entities.SaveAccountEntities(ctx, accountDTO)
	if err != nil {
		log.Printf(err.Error())
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
