package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"godating-dealls/conf"
	authEntities "godating-dealls/internal/core/entities/auths"
	authUsecase "godating-dealls/internal/core/usecase/auths"
	authHandler "godating-dealls/internal/handler"
	authRepo "godating-dealls/internal/infra/mysql/repo"
	"net/http"
)

func main() {
	// Init context before run application
	ctx := context.Background()
	// Ensure to close the database connection when the application exits
	defer conf.CloseDBConnection()
	// Create of the database connection
	sqlConnection := conf.CreateDBConnection(ctx)
	// Initiate validator
	validate := validator.New()

	// initiate repo
	repoAuth := authRepo.NewAccountsRepositoryImpl(sqlConnection, validate)
	// call business rules
	entitiesAuth := authEntities.NewAccountsEntitiesImpl(repoAuth, validate)
	// Create the use case with entities
	usecaseAuth := authUsecase.NewAuthUsecase(entitiesAuth)
	// Create the handler with the use case
	handlerAuth := authHandler.NewAuthHandler(usecaseAuth)

	// Set up the router
	r := http.NewServeMux()
	r.HandleFunc("POST /godating-dealls/api/authenticate/register", handlerAuth.RegisterUserHandler)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		return
	}
}
