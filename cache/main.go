package cache

import (
	"go-snowflake-service/conf"
	"log"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.ConfigObject.RedisAddr,
		Password: conf.ConfigObject.RedisPassword,
		DB:       conf.ConfigObject.RedisDatabase,
	})

	if err := RedisClient.Ping().Err(); err != nil {
		log.Fatal(err)
	}
}
