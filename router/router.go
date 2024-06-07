package router

import (
	"godating-dealls/internal/common"
	authHandler "godating-dealls/internal/handler"
	"net/http"
)

func SetupRouter(authHandler *authHandler.AuthHandler) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /godating-dealls/api/authenticate/register", authHandler.RegisterUserHandler)
	r.HandleFunc("POST /godating-dealls/api/authenticate/login", authHandler.LoginUserHandler)

	// Using middleware authenticate
	r.Handle("POST /godating-dealls/api/authenticate/logout",
		common.AuthMiddleware(http.HandlerFunc(authHandler.LogoutUserHandler)))

	return r
}
