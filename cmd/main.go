package main

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/robfig/cron/v3"
	"godating-dealls/conf"
	"godating-dealls/internal/common"
	"godating-dealls/internal/core/entities/accounts"
	dailyquotaentity "godating-dealls/internal/core/entities/daily_quotas"
	loginhistoryentity "godating-dealls/internal/core/entities/login_histories"
	"godating-dealls/internal/core/entities/packages"
	"godating-dealls/internal/core/entities/selection_histories"
	"godating-dealls/internal/core/entities/swipes"
	"godating-dealls/internal/core/entities/task_history"
	usersentity "godating-dealls/internal/core/entities/users"
	accountusecase "godating-dealls/internal/core/usecase/auths"
	dailyquotausecase "godating-dealls/internal/core/usecase/daily_quotas"
	packageusecase "godating-dealls/internal/core/usecase/packages"
	swipeusecase "godating-dealls/internal/core/usecase/swipes"
	"godating-dealls/internal/core/usecase/users"
	handler2 "godating-dealls/internal/delivery/handler"
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
	// Init context before run application
	ctx := context.Background()

	// Set up logging
	logs := InitializeLogger()
	defer logs.Close()

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
	purchaseRepository := repo.NewPurchasePackagesRepositoryImpl()

	// Entities represented of enterprise business rules for that self of entity
	accountEntity := accounts.NewAccountsEntityImpl(accountRepository, val)
	userEntity := usersentity.NewUserEntityImpl(userRepository, val)
	loginHistoryEntity := loginhistoryentity.NewLoginHistoriesEntityImpl(val, loginHistoryRepository)
	dailyQuotasEntity := dailyquotaentity.NewDailyQuotasEntityImpl(val, dailyQuotaRepository)
	selectionHistoryEntity := selection_histories.NewSelectionHistoryEntityImpl(selectionHistoryRepository)
	taskHistoryEntity := task_history.NewTaskHistoryEntityImpl(taskHistoryRepository)
	swipeEntity := swipes.NewSwipeEntityImpl(swipeRepository)
	packageEntity := packages.NewPackageEntityImpl(packageRepository, purchaseRepository)

	// Usecase
	authenticateUsecase := accountusecase.NewAuthUsecase(DB, accountEntity, userEntity, RS, loginHistoryEntity)
	dailyQuotasUsecase := dailyquotausecase.NewDailyQuotasUsecase(DB, dailyQuotasEntity, userEntity, accountEntity, packageEntity)
	InitializeCronJobDailyQuota(ctx, dailyQuotasUsecase)
	usersUsecase := users.NewUserUsecase(DB, userEntity, accountEntity, selectionHistoryEntity, taskHistoryEntity)
	swipeUsecase := swipeusecase.NewSwipeUsecase(DB, swipeEntity, dailyQuotasEntity, accountEntity)
	packageUsecase := packageusecase.NewPackageUsecase(DB, packageEntity, accountEntity, dailyQuotasEntity)

	// Create the handler with the use case
	authenticateHandler := handler2.NewAuthHandler(authenticateUsecase)
	usersHandler := handler2.NewUsersHandler(usersUsecase)
	swipeHandler := handler2.NewSwipeHandler(swipeUsecase)
	packageHandler := handler2.NewPackageHandler(packageUsecase)
	quotaHandler := handler2.NewQuotaHandler(dailyQuotasUsecase)

	// Set up the router
	r := router.InitializeRouter(authenticateHandler, usersHandler, swipeHandler, packageHandler, quotaHandler)

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

func InitializeCronJobDailyQuota(ctx context.Context, boundary dailyquotausecase.InputDailyQuotaBoundary) {
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
