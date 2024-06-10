package config

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"godating-dealls/internal/common"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// initMySQLDB initializes the database connection
func initMySQLDB(ctx context.Context) *sql.DB {
	// Load .env file
	err := godotenv.Load()
	common.HandleErrorWithParam(err, "Error loading .env file")

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	common.HandleErrorWithParam(err, "Could not opn DB connection, DB connection is failed")

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	common.HandleErrorWithParam(err, "Could not connect to DB MySQL")
	log.Println("Connected to DB MySQL success!")

	return db
}

// CreateDBConnection returns the singleton database instance
func CreateDBConnection(ctx context.Context) *sql.DB {
	if db == nil {
		db = initMySQLDB(ctx)
	}
	return db
}

// CloseDBConnection closes the database connection
func CloseDBConnection() {
	if db != nil {
		err := db.Close()
		common.HandleErrorReturn(err)
	}
}
