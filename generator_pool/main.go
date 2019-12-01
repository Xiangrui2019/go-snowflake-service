package generator_pool

import "go-snowflake-service/snowflake"

var GeneratorPool map[int64]*snowflake.SnowFlake

func GetSnowFlake(nodeId int64) *snowflake.SnowFlake {
	if _, err := GeneratorPool[nodeId]; err != true {
		GeneratorPool[nodeId] = snowflake.NewSnowFlake(nodeId)
	}

	return GeneratorPool[nodeId]
}
