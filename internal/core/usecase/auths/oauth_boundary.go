package auths

import res "godating-dealls/internal/domain/auths"

type OutputAuthBoundary interface {
	LoginResponse(err error)
	RegisterResponse(response res.RegisterResponse, err error)
	LogoutResponse(err error)
}
