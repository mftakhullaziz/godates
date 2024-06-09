package main

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/robfig/cron/v3"
	"godating-dealls/conf"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/auths"
	dailyquotaentity "godating-dealls/internal/core/entities/daily_quotas"
	loginhistoryentity "godating-dealls/internal/core/entities/login_histories"
	"godating-dealls/internal/core/entities/packages"
	"godating-dealls/internal/core/entities/selection_histories"
	"godating-dealls/internal/core/entities/swipes"
	"godating-dealls/internal/core/entities/task_history"
	usersentity "godating-dealls/internal/core/entities/users"
	authUsecase "godating-dealls/internal/core/usecase/auths"
	dailyQuotaUsecase "godating-dealls/internal/core/usecase/daily_quotas"
	packages2 "godating-dealls/internal/core/usecase/packages"
	swipes2 "godating-dealls/internal/core/usecase/swipes"
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
	selectionHistoryRepository := repo.NewSelectionHistoriesRepositoryImpl()
	taskHistoryRepository := repo.NewTaskHistorySQLRepository()
	swipeRepository := repo.NewSwipesRepositoryImpl()
	packageRepository := repo.NewPackagesRepositoryImpl()

	// Entities represented of enterprise business rules for that self of entity
	authEntity := auths.NewAccountsEntitiesImpl(accountRepository, val)
	userEntity := usersentity.NewUserEntitiesImpl(userRepository, val)
	loginHistoryEntity := loginhistoryentity.NewLoginHistoriesEntitiesImpl(val, loginHistoryRepository)
	dailyQuotasEntity := dailyquotaentity.NewDailyQuotasEntityImpl(val, dailyQuotaRepository)
	selectionHistoryEntity := selection_histories.NewSelectionHistoryEntityImpl(selectionHistoryRepository)
	taskHistoryEntity := task_history.NewTaskHistoryEntityImpl(taskHistoryRepository)
	swipeEntity := swipes.NewSwipeEntityImpl(swipeRepository)
	packageEntity := packages.NewPackageEntityImpl(packageRepository)

	// Usecase
	authenticateUsecase := authUsecase.NewAuthUsecase(DB, authEntity, userEntity, RS, loginHistoryEntity)
	dailyQuotasUsecase := dailyQuotaUsecase.NewDailyQuotasUsecase(DB, dailyQuotasEntity, userEntity)
	InitializeCronJobDailyQuota(ctx, dailyQuotasUsecase)
	usersUsecase := users.NewUserUsecase(DB, userEntity, authEntity, selectionHistoryEntity, taskHistoryEntity)
	swipeUsecase := swipes2.NewSwipeUsecase(DB, swipeEntity, dailyQuotasEntity, authEntity)
	packageUsease := packages2.NewPackageUsecase(DB, packageEntity)

	// Create the handler with the use case
	authenticateHandler := handler.NewAuthHandler(authenticateUsecase)
	usersHandler := handler.NewUsersHandler(usersUsecase)
	swipeHandler := handler.NewSwipeHandler(swipeUsecase)
	packageHandler := handler.NewPackageHandler(packageUsease)

	// Set up the router
	r := router.InitializeRouter(authenticateHandler, usersHandler, swipeHandler, packageHandler)

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
