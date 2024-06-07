package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"godating-dealls/conf"
	"godating-dealls/internal/common"
	authEntities "godating-dealls/internal/core/entities/auths"
	userEntities "godating-dealls/internal/core/entities/users"
	authUsecase "godating-dealls/internal/core/usecase/auths"
	authHandler "godating-dealls/internal/handler"
	repo "godating-dealls/internal/infra/mysql/repo"
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
	conf.InitializeRedisClient(ctx)

	// Initiate validator
	validate := validator.New()

	// Initiate repo
	repoAuth := repo.NewAccountsRepositoryImpl(validate)
	repoUser := repo.NewUsersRepositoryImpl(validate)

	// Call business rules
	entitiesAuth := authEntities.NewAccountsEntitiesImpl(repoAuth, validate)
	entitiesUser := userEntities.NewUserEntitiesImpl(repoUser, validate)

	// Create the use case with entities
	usecaseAuth := authUsecase.NewAuthUsecase(DB, entitiesAuth, entitiesUser)
	// Create the handler with the use case
	handlerAuth := authHandler.NewAuthHandler(usecaseAuth)

	// Set up the router
	r := http.NewServeMux()
	r.HandleFunc("POST /godating-dealls/api/authenticate/register", handlerAuth.RegisterUserHandler)
	r.HandleFunc("POST /godating-dealls/api/authenticate/login", handlerAuth.LoginUserHandler)

	err = http.ListenAndServe(":8000", r)
	common.HandleErrorReturn(err)
}
