package redisclient

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type RdsImpl struct {
	Redis  RedisInterface
	Client *redis.Client
}

func NewRedisService(Client *redis.Client) RedisInterface {
	return &RdsImpl{
		Client: Client,
	}
}

func (r RdsImpl) StoreToRedis(key string, data interface{}) error {
	ctx := context.Background()

	serializedData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = r.Client.Set(ctx, key, serializedData, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r RdsImpl) LoadFromRedis(key string) (interface{}, error) {
	ctx := context.Background()

	data, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errors.New("key does not exist")
		}
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
