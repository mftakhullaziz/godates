package presenters

import (
	"godating-dealls/internal/common"
	usecase "godating-dealls/internal/core/usecase/auths"
	"godating-dealls/internal/domain"
	"net/http"
)

// AuthPresenter is responsible for presenting authentication responses
type AuthPresenter struct {
	w http.ResponseWriter
}

// NewAuthPresenter creates a new AuthPresenter
func NewAuthPresenter(w http.ResponseWriter) usecase.OutputAuthBoundary {
	return &AuthPresenter{w: w}
}

// RegisterResponse sends the registration response to the client
func (ap *AuthPresenter) RegisterResponse(response domain.RegisterResponse, err error) {
	common.HandleInternalServerError(err, ap.w)
	common.WriteJSONResponse(ap.w, http.StatusCreated, "Created account successfully", response)
}

// LoginResponse sends the login response to the client
func (ap *AuthPresenter) LoginResponse(response domain.LoginResponse, err error) {
	common.HandleInternalServerError(err, ap.w)
	common.WriteJSONResponse(ap.w, http.StatusOK, "Login account successfully", response)
}

func (ap *AuthPresenter) LogoutResponse(response domain.LogoutResponse, err error) {
	common.HandleInternalServerError(err, ap.w)
	common.WriteJSONResponse(ap.w, http.StatusOK, "Logout account successfully", response)
}
