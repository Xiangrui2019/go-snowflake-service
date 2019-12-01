package conf

import "go-snowflake-service/cache"

type Config struct {
}

type ConfigObject *Config

func init() {
	cache.InitRedis()
}
