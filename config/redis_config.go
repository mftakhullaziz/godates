package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

// RedisClient is a global variable to hold the Redis client
var RedisClient *redis.Client

// InitializeRedisClient initializes the Redis client
func InitializeRedisClient(ctx context.Context) *redis.Client {
	// Load environment variables from .env file only in development
	env := os.Getenv("ENV")
	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	rdsHost := os.Getenv("REDIS_HOST")
	rdsPort := os.Getenv("REDIS_PORT")
	uri := fmt.Sprintf("%s:%s", rdsHost, rdsPort)

	options := &redis.Options{
		Addr:     uri,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
		Username: os.Getenv("REDIS_USER"),
	}

	// Enable TLS for non-local environments
	if env != "development" {
		options.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	RedisClient = redis.NewClient(options)

	// Test the connection
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to redisclient: %v", err)
	}
	log.Println("Connected to Redis successfully")

	return RedisClient
}
