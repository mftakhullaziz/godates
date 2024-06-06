package auths

import (
	"context"
	"godating-dealls/internal/domain/auths"
)

type AuthUsecase struct {
	input  InputAuthBoundary
	output OutputAuthBoundary
}

func NewAuthUsecase(input InputAuthBoundary, output OutputAuthBoundary) *AuthUsecase {
	return &AuthUsecase{
		input:  input,
		output: output,
	}
}

//func (au *AuthUsecase) LoginUser(username string, password string) (bool, error) {}

//func (au *AuthUsecase) LogoutUser(username string) (bool, error) {}

func (au *AuthUsecase) RegisterUser(ctx context.Context, request auths.RegisterRequest) {
	err := au.input.ExecuteRegister(ctx, request)
	if err != nil {
		panic(err.Error())
	}
	au.output.RegisterResponse(err)
}
