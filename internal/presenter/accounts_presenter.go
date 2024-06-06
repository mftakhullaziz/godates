package presenters

import (
	"encoding/json"
	"godating-dealls/internal/common"
	usecase "godating-dealls/internal/core/usecase/auths"
	"godating-dealls/internal/domain/auths"
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
func (ap *AuthPresenter) RegisterResponse(response auths.RegisterResponse, err error) {
	common.HandleInternalServerError(err, ap.w)

	// Set content type to JSON
	ap.w.Header().Set("Content-Type", "application/json")
	// Set status code to 200 OK
	ap.w.WriteHeader(http.StatusOK)
	// Encode the response as JSON and write to the response writer
	json.NewEncoder(ap.w).Encode(response)
}

func (ap *AuthPresenter) LoginResponse(err error) {
	//TODO implement me
	panic("implement me")
}

func (ap *AuthPresenter) LogoutResponse(err error) {
	//TODO implement me
	panic("implement me")
}
