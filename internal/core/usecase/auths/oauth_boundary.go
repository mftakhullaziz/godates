package auths

type OutputAuthBoundary interface {
	LoginResponse(err error)
	RegisterResponse(err error)
	LogoutResponse(err error)
}
