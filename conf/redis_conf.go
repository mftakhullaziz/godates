package conf

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"godating-dealls/internal/common"
	"godating-dealls/internal/infra/redisclient"
	"log"
	"os"
)

// RedisClient is a global variable to hold the Redis client
var RedisClient *redis.Client

// InitializeRedisClient initializes the Redis client
func InitializeRedisClient(ctx context.Context) {
	err := godotenv.Load()
	common.HandleErrorWithParam(err, "Error loading .env file")

	rdsHost := os.Getenv("REDIS_HOST")
	rdsPort := os.Getenv("REDIS_PORT")
	uri := fmt.Sprintf("%s:%s", rdsHost, rdsPort)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     uri, // Replace with your Redis server address
		Password: "",  // No password set
		DB:       0,   // Use default DB
	})

	// Test the connection
	_, err = RedisClient.Ping(ctx).Result()
	common.HandleErrorWithParam(err, "Error connecting to redisclient")
	log.Println("Connected to Redis successfully")

	// Instantiate a new Redis service with the Redis client
	_ = redisclient.NewRedisService(RedisClient)
}
