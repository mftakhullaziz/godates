package redisclient

type RedisInterface interface {
	StoreToRedis(key string, data interface{}) error
	LoadFromRedis(key string) (interface{}, error)
}
