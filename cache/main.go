package cache

import (
	"go-snowflake-service/conf"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.ConfigObject.Redis.Addr,
		Password: conf.ConfigObject.Redis.Password,
		DB:       conf.ConfigObject.Redis.Database,
	})
}
