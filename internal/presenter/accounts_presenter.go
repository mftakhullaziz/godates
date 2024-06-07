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
	// Transform to default response
	res := common.DefaultResponse{
		StatusCode: http.StatusCreated,
		Message:    "Created account successfully",
		IsSuccess:  true,
		RequestAt:  common.FormatTime(),
		Data:       response,
	}
	// Set content type to JSON
	ap.w.Header().Set("Content-Type", "application/json")
	// Set status code to 200 OK
	ap.w.WriteHeader(http.StatusOK)
	// Encode the response as JSON and write to the response writer
	err = json.NewEncoder(ap.w).Encode(res)
	common.HandleErrorReturn(err)
}

// LoginResponse sends the login response to the client
func (ap *AuthPresenter) LoginResponse(response auths.LoginResponse, err error) {
	common.HandleInternalServerError(err, ap.w)
	// Transform to default response
	res := common.DefaultResponse{
		StatusCode: http.StatusCreated,
		Message:    "Created account successfully",
		IsSuccess:  true,
		RequestAt:  common.FormatTime(),
		Data:       response,
	}
	// Set content type to JSON
	ap.w.Header().Set("Content-Type", "application/json")
	// Set status code to 200 OK
	ap.w.WriteHeader(http.StatusOK)
	// Encode the response as JSON and write to the response writer
	err = json.NewEncoder(ap.w).Encode(res)
	common.HandleErrorReturn(err)
}

func (ap *AuthPresenter) LogoutResponse(err error) {
	//TODO implement me
	panic("implement me")
}
