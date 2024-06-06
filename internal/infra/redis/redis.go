package redis

type Redis interface {
	SaveDataToRedis()
	LoadDataFromRedis()
}
