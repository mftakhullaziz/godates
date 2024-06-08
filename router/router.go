package router

import (
	md "godating-dealls/internal/common"
	"godating-dealls/internal/handler"
	"net/http"
)

func InitializeRouter(authHandler *handler.AuthHandler, userHandler *handler.UsersHandler) *http.ServeMux {
	r := http.NewServeMux()

	// Without middleware
	r.HandleFunc("POST /godating-dealls/api/authenticate/register", authHandler.RegisterUserHandler)
	r.HandleFunc("POST /godating-dealls/api/authenticate/login", authHandler.LoginUserHandler)

	// Using middleware authenticate
	r.Handle("POST /godating-dealls/api/authenticate/logout", md.AuthMiddleware(http.HandlerFunc(authHandler.LogoutUserHandler)))
	r.Handle("GET /godating-dealls/api/daily-accounts", md.AuthMiddleware(http.HandlerFunc(userHandler.UserViewsHandler)))
	r.Handle("POST /godating-dealls/api/swipe-left", md.AuthMiddleware(http.HandlerFunc(nil)))
	r.Handle("POST /godating-dealls/api/swipe-right", md.AuthMiddleware(http.HandlerFunc(nil)))
	r.Handle("GET /godating-dealls/api/quota", md.AuthMiddleware(http.HandlerFunc(nil)))
	r.Handle("POST /godating-dealls/api/purchase-package", md.AuthMiddleware(http.HandlerFunc(nil)))
	r.Handle("GET /godating-dealls/api/packages", md.AuthMiddleware(http.HandlerFunc(nil)))
	r.Handle("GET /godating-dealls/api/track-view", md.AuthMiddleware(http.HandlerFunc(nil)))

	return r
}
