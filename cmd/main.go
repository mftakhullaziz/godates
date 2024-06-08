package main

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/robfig/cron/v3"
	"godating-dealls/conf"
	"godating-dealls/internal/common"
	authEntities "godating-dealls/internal/core/entities/auths"
	dailyQuotaEntities "godating-dealls/internal/core/entities/daily_quotas"
	loginHistories "godating-dealls/internal/core/entities/login_histories"
	userEntities "godating-dealls/internal/core/entities/users"
	authUsecase "godating-dealls/internal/core/usecase/auths"
	dailyQuotaUsecase "godating-dealls/internal/core/usecase/daily_quotas"
	"godating-dealls/internal/core/usecase/users"
	"godating-dealls/internal/handler"
	"godating-dealls/internal/infra/mysql/repo"
	"godating-dealls/internal/infra/redisclient"
	"godating-dealls/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Set up logging
	logs := InitializeLogger()
	defer logs.Close()

	// Init context before run application
	ctx := context.Background()

	DB := InitializeDB(ctx)
	defer conf.CloseDBConnection()

	RS := InitializeRedis(ctx)

	// Initiate validator
	val := validator.New()

	// Initiate repo
	accountRepository := repo.NewAccountsRepositoryImpl()
	userRepository := repo.NewUsersRepositoryImpl()
	loginHistoryRepository := repo.NewLoginHistoriesRepositoryImpl()
	dailyQuotaRepository := repo.NewDailyQuotasRepositoryImpl()

	// Call business rules
	ea := authEntities.NewAccountsEntitiesImpl(accountRepository, val)
	eu := userEntities.NewUserEntitiesImpl(userRepository, val)
	elh := loginHistories.NewLoginHistoriesEntitiesImpl(val, loginHistoryRepository)
	edq := dailyQuotaEntities.NewDailyQuotasEntitiesImpl(val, dailyQuotaRepository)

	// Create the use case with entities
	ua := authUsecase.NewAuthUsecase(DB, ea, eu, RS, elh)

	udq := dailyQuotaUsecase.NewDailyQuotasUsecase(DB, edq, eu)
	InitializeCronJobDailyQuota(ctx, udq)

	uu := users.NewUserUsecase(eu, DB)

	// Create the handler with the use case
	ha := handler.NewAuthHandler(ua)
	hu := handler.NewUsersHandler(uu)

	// Set up the router
	r := router.InitializeRouter(ha, hu)

	// Create a channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		err := http.ListenAndServe(":8000", r)
		common.HandleErrorReturn(err)
	}()

	// Block until a signal is received
	<-stop

	log.Println("Shutting down the server...")
}

func InitializeLogger() *os.File {
	// Set up logging
	setupLogger, err := common.SetupLogger()
	common.HandleErrorWithParam(err, "Setup Logger Failed")
	return setupLogger
}

func InitializeDB(ctx context.Context) *sql.DB {
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

func InitializeCronJobDailyQuota(ctx context.Context, boundary dailyQuotaUsecase.InputDailyQuotaBoundary) {
	c := cron.New()
	// Run every 24 hours
	_, err := c.AddFunc("@every 24h", func() { // Changed to run every minute for testing
		log.Println("Executing daily quota update usecase")
		err := boundary.ExecuteAutoUpdateDailyQuotaUsecase(ctx)
		if err != nil {
			log.Printf("Error executing daily quota update usecase: %v", err)
		} else {
			log.Println("Successfully executed daily quota update usecase")
		}
	})
	if err != nil {
		log.Printf("Error adding cron job: %v", err)
	}
	c.Start()
	log.Println("Cron job started")
}
