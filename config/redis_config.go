package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"godating-dealls/internal/common"
	"log"
	"os"
)

// RedisClient is a global variable to hold the Redis client
var RedisClient *redis.Client

// InitializeRedisClient initializes the Redis client
func InitializeRedisClient(ctx context.Context) *redis.Client {
	err := godotenv.Load()
	common.HandleErrorWithParam(err, "Error loading .env file")

	rdsHost := os.Getenv("REDIS_HOST")
	rdsPort := os.Getenv("REDIS_PORT")
	uri := fmt.Sprintf("%s:%s", rdsHost, rdsPort)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
		Username: os.Getenv("REDIS_USER"),
		//TLSConfig: &tls.Config{
		//	InsecureSkipVerify: false, // if run from localhost set false
		//},
	})

	// Test the connection
	_, err = RedisClient.Ping(ctx).Result()
	common.HandleErrorWithParam(err, "Error connecting to redisclient")
	log.Println("Connected to Redis successfully")

	return RedisClient
}
