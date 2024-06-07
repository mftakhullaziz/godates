package handler

import (
	"encoding/json"
	"fmt"
	"godating-dealls/internal/common"
	input "godating-dealls/internal/core/usecase/auths"
	"godating-dealls/internal/domain/auths"
	presenters "godating-dealls/internal/presenter"
	"log"
	"net/http"
)

type AuthHandler struct {
	usecase input.InputAuthBoundary
}

func NewAuthHandler(usecase input.InputAuthBoundary) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
	}
}

func (ah *AuthHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var request auths.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Println(request)

	ctx := r.Context()
	// Instantiate the presenter
	presenter := presenters.NewAuthPresenter(w)

	// Call the use case method passing the presenter
	err := ah.usecase.ExecuteRegisterUsecase(ctx, request, presenter)
	common.HandleInternalServerError(err, w)
}

func (ah *AuthHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var request auths.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Println(request)

	ctx := r.Context()

	// Instantiate the presenter
	presenter := presenters.NewAuthPresenter(w)

	// Call the use case method passing the presenter
	err := ah.usecase.ExecuteLoginUsecase(ctx, request, presenter)
	common.HandleInternalServerError(err, w)
}

func (ah *AuthHandler) LogoutUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, ok := ctx.Value("token").(string)
	if !ok || token == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	presenter := presenters.NewAuthPresenter(w)
	fmt.Println(token)

	// Call the use case method passing the presenter
	err := ah.usecase.ExecuteLogoutUsecase(ctx, &token, presenter)
	common.HandleInternalServerError(err, w)
}
