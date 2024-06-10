package handler

import (
	"encoding/json"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/users"
	"godating-dealls/internal/delivery/presenter"
	"godating-dealls/internal/domain"
	"net/http"
)

type UsersHandler struct {
	UserInput users.InputUserBoundary
}

func NewUsersHandler(userInput users.InputUserBoundary) *UsersHandler {
	return &UsersHandler{UserInput: userInput}
}

func (uh *UsersHandler) UserViewsHandler(w http.ResponseWriter, r *http.Request) {
	// If user premium is unlimited, if not is just 10 data
	ctx := r.Context()
	token, ok := ctx.Value("token").(string)
	if !ok || token == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Instantiate the presenter
	presenter := presenters.NewUserPresenter(w)

	// Call the use case method passing the presenter
	err := uh.UserInput.ExecuteUserViewsUsecase(ctx, token, presenter)
	common.HandleInternalServerError(err, w)
}

func (uh *UsersHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token, ok := ctx.Value("token").(string)
	if !ok || token == "" {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	var request domain.PatchUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	presenter := presenters.NewUserPresenter(w)

	// Call the use case method passing the presenter
	err := uh.UserInput.ExecutePatchUserUsecase(ctx, token, request, presenter)
	common.HandleInternalServerError(err, w)
}
