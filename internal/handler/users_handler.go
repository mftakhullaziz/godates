package handler

import (
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/usecase/users"
	presenters "godating-dealls/internal/presenter"
	"net/http"
)

type UsersHandler struct {
	UserInput users.InputUserBoundary
}

func NewUsersHandler(userInput users.InputUserBoundary) *UsersHandler {
	return &UsersHandler{UserInput: userInput}
}

func (uh *UsersHandler) UserViewsHandler(w http.ResponseWriter, r *http.Request) {
	//var request domain.LoginRequest
	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	http.Error(w, "Invalid request payload", http.StatusBadRequest)
	//	return
	//}
	//log.Println(request)

	ctx := r.Context()

	// Instantiate the presenter
	presenter := presenters.NewUserPresenter(w)

	// Call the use case method passing the presenter
	err := uh.UserInput.ExecuteUserViewsUsecase(ctx, presenter)
	common.HandleInternalServerError(err, w)
}
