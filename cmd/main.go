package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"godating-dealls/conf"
	"godating-dealls/internal/common"
	authEntities "godating-dealls/internal/core/entities/auths"
	"godating-dealls/internal/core/entities/login_histories"
	userEntities "godating-dealls/internal/core/entities/users"
	authUsecase "godating-dealls/internal/core/usecase/auths"
	authHandler "godating-dealls/internal/handler"
	"godating-dealls/internal/infra/mysql/repo"
	"godating-dealls/internal/infra/redisclient"
	"net/http"
)

func main() {
	// Set up logging
	setupLogger, err := common.SetupLogger()
	common.HandleErrorWithParam(err, "Setup Logger Failed")
	defer setupLogger.Close()

	// Init context before run application
	ctx := context.Background()
	// Ensure to close the database connection when the application exits
	defer conf.CloseDBConnection()
	// Create of the database connection
	DB := conf.CreateDBConnection(ctx)
	// Create redis client connection
	RedisClient := conf.InitializeRedisClient(ctx)
	RedisService := redisclient.NewRedisService(RedisClient)

	// Initiate validator
	validate := validator.New()

	// Initiate repo
	repoAuth := repo.NewAccountsRepositoryImpl(validate)
	repoUser := repo.NewUsersRepositoryImpl(validate)
	repoLoginHistory := repo.NewLoginHistoriesRepositoryImpl(validate)

	// Call business rules
	entitiesAuth := authEntities.NewAccountsEntitiesImpl(repoAuth, validate)
	entitiesUser := userEntities.NewUserEntitiesImpl(repoUser, validate)
	entitiesLoginHistory := login_histories.NewLoginHistoriesEntitiesImpl(validate, repoLoginHistory)

	// Create the use case with entities
	usecaseAuth := authUsecase.NewAuthUsecase(DB, entitiesAuth, entitiesUser, RedisService, entitiesLoginHistory)
	// Create the handler with the use case
	handlerAuth := authHandler.NewAuthHandler(usecaseAuth)

	// Set up the router
	r := setupRouter(handlerAuth)

	err = http.ListenAndServe(":8000", r)
	common.HandleErrorReturn(err)
}

func setupRouter(handlerAuth *authHandler.AuthHandler) *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("POST /godating-dealls/api/authenticate/register", handlerAuth.RegisterUserHandler)
	r.HandleFunc("POST /godating-dealls/api/authenticate/login", handlerAuth.LoginUserHandler)
	// Using middleware authenticate
	r.Handle("POST /godating-dealls/api/authenticate/logout", common.AuthMiddleware(http.HandlerFunc(handlerAuth.LogoutUserHandler)))
	return r
}
