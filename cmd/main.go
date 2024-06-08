package main

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/robfig/cron/v3"
	"godating-dealls/conf"
	"godating-dealls/internal/common"
	authEntities "godating-dealls/internal/core/entities/auths"
	"godating-dealls/internal/core/entities/login_histories"
	userEntities "godating-dealls/internal/core/entities/users"
	authUsecase "godating-dealls/internal/core/usecase/auths"
	authHandler "godating-dealls/internal/handler"
	"godating-dealls/internal/infra/mysql/repo"
	"godating-dealls/internal/infra/redisclient"
	"godating-dealls/router"
	"net/http"
)

func main() {
	// Init context before run application
	ctx := context.Background()
	InitializeLogger()
	DB := InitializeDB(ctx)
	RS := InitializeRedis(ctx)
	InitializeCronJob()

	// Initiate validator
	val := validator.New()

	// Initiate repo
	ra := repo.NewAccountsRepositoryImpl(val)
	ru := repo.NewUsersRepositoryImpl(val)
	rlh := repo.NewLoginHistoriesRepositoryImpl(val)

	// Call business rules
	ea := authEntities.NewAccountsEntitiesImpl(ra, val)
	eu := userEntities.NewUserEntitiesImpl(ru, val)
	elh := login_histories.NewLoginHistoriesEntitiesImpl(val, rlh)

	// Create the use case with entities
	ua := authUsecase.NewAuthUsecase(DB, ea, eu, RS, elh)

	// Create the handler with the use case
	ha := authHandler.NewAuthHandler(ua)

	// Set up the router
	r := router.SetupRouter(ha)
	err := http.ListenAndServe(":8000", r)
	common.HandleErrorReturn(err)

	select {}
}

func InitializeLogger() {
	// Set up logging
	setupLogger, err := common.SetupLogger()
	common.HandleErrorWithParam(err, "Setup Logger Failed")
	defer setupLogger.Close()
}

func InitializeDB(ctx context.Context) *sql.DB {
	// Ensure to close the database connection when the application exits
	defer conf.CloseDBConnection()
	// Create of the database connection
	DB := conf.CreateDBConnection(ctx)
	return DB
}

func InitializeRedis(ctx context.Context) redisclient.RedisInterface {
	// Create redis client connection
	rdsClient := conf.InitializeRedisClient(ctx)
	rds := redisclient.NewRedisService(rdsClient)
	return rds
}

func InitializeCronJob() {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", func() {
		// resetSwipeQuotas()
	})
	common.HandleErrorReturn(err)
	c.Start()
}
