package cache

import "github.com/go-redis/redis"

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
}
