package generator_pool

import (
	"go-snowflake-service/snowflake"
	"sync"
)

var (
	GeneratorPool = make(map[int64]*snowflake.SnowFlake)
	mutex         = sync.Mutex{}
)

func GetSnowFlake(nodeId int64) *snowflake.SnowFlake {
	mutex.Lock()

	if _, err := GeneratorPool[nodeId]; err != true {
		GeneratorPool[nodeId] = snowflake.NewSnowFlake(nodeId)
	}

	mutex.Unlock()
	return GeneratorPool[nodeId]
}
