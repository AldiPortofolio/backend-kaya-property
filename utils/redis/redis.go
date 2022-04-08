package redis

import (
	"kaya-backend/utils"

	"github.com/go-redis/redis/v7"
)

func RedisStore() *redis.Client {
	//Initializing redis
	redisEnpoint := utils.GetEnv("REDIS_ENDPOINT", "127.0.0.1:6379")

	client := redis.NewClient(&redis.Options{
		Addr: redisEnpoint, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
