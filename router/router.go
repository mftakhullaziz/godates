package router

import (
	md "godating-dealls/internal/common"
	"godating-dealls/internal/delivery/handler"
	"net/http"
)

func InitializeRouter(
	authHandler *handler.AuthHandler,
	userHandler *handler.UsersHandler,
	swipeHandler *handler.SwipeHandler,
	packageHandler *handler.PackageHandler,
	quotaHandler *handler.QuotaHandler,
	accountHandler *handler.AccountHandler) *http.ServeMux {

	r := http.NewServeMux()

	// Without middleware
	r.HandleFunc("POST /godating-dealls/api/authenticate/register", authHandler.RegisterUserHandler)
	r.HandleFunc("POST /godating-dealls/api/authenticate/login", authHandler.LoginUserHandler)

	// Using middleware authenticate
	r.Handle("POST /godating-dealls/api/authenticate/logout", md.AuthMiddleware(http.HandlerFunc(authHandler.LogoutUserHandler)))
	r.Handle("POST /godating-dealls/api/daily-accounts", md.AuthMiddleware(http.HandlerFunc(userHandler.UserViewsHandler)))
	r.Handle("POST /godating-dealls/api/swipes", md.AuthMiddleware(http.HandlerFunc(swipeHandler.SwipeHandler)))
	r.Handle("GET /godating-dealls/api/quota", md.AuthMiddleware(http.HandlerFunc(quotaHandler.CheckQuotaAccountHandler)))
	r.Handle("POST /godating-dealls/api/purchase-package", md.AuthMiddleware(http.HandlerFunc(packageHandler.PurchasePackages)))
	r.Handle("GET /godating-dealls/api/packages", md.AuthMiddleware(http.HandlerFunc(packageHandler.GetPackageHandler)))
	r.Handle("GET /godating-dealls/api/account-details", md.AuthMiddleware(http.HandlerFunc(accountHandler.FetchAccountDetailsHandler)))
	r.Handle("POST /godating-dealls/api/account-view", md.AuthMiddleware(http.HandlerFunc(accountHandler.AccountViewHandler)))

	return r
}
