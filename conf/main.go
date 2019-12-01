package conf

import "go-snowflake-service/cache"

func init() {
	cache.InitRedis()
}
