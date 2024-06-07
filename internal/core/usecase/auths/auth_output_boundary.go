package auths

import res "godating-dealls/internal/domain/auths"

type OutputAuthBoundary interface {
	LoginResponse(response res.LoginResponse, err error)
	RegisterResponse(response res.RegisterResponse, err error)
	LogoutResponse(response res.LogoutResponse, err error)
}
