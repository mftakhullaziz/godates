package redisclient

import "context"

type RedisInterface interface {
	StoreToRedis(ctx context.Context, key string, data interface{}) error
	LoadFromRedis(ctx context.Context, key string) (interface{}, error)
	ClearFromRedis(ctx context.Context, key string) error
}
