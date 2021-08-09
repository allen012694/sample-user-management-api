package context

import (
	"github.com/allen012694/usersystem/config"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis() (*redis.Client, error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:        config.REDIS_ADDRESS,
		Password:    config.REDIS_PASSWORD,
		DB:          0,
		MaxRetries:  0,
		ReadTimeout: 0,
	})

	return redisClient, nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}
