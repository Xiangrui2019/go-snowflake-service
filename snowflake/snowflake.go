package snowflake

import (
	"go-snowflake-service/cache"
	"go-snowflake-service/conf"
	"go-snowflake-service/consts"
	"strconv"
	"time"
)

type ID int64

type SnowFlake struct {
	timestamp int64
	node      int64
	step      int64
}

func NewSnowFlake(node int64) *SnowFlake {
	return &SnowFlake{
		timestamp: 0,
		node:      node,
		step:      0,
	}
}

func (gen *SnowFlake) GenerateId() (ID, error) {
	for {
		if qs, _ := cache.RedisClient.SetNX("lock"+strconv.Itoa(int(gen.node)), "1", time.Minute*5).Result(); qs == true {
			break
		}

		time.Sleep(time.Nanosecond)
	}

	// 获取当前时间戳
	now := time.Now().UnixNano() / 1e6

	// 获取Redis中存储的时间戳
	gen.timestamp, _ = cache.RedisClient.Get("timestamp" + strconv.Itoa(int(gen.node))).Int64()

	if gen.timestamp == now {
		// 1. step步进
		gen.step++

		// 2. 写入步进
		if err := cache.RedisClient.Incr("step" + strconv.Itoa(int(gen.node))).Err(); err != nil {
			return 0, err
		}

		// 3. 检查当前 step 是否用完
		if gen.step > consts.StepMax {
			// 等待本毫秒结束
			for now <= gen.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}

	} else {
		// 本毫秒内 step 用完
		gen.step = 0
		// 写入Step
		if err := cache.RedisClient.Set(
			"step"+strconv.Itoa(int(gen.node)),
			0,
			0).Err(); err != nil {
			return 0, err
		}
	}

	// 更新当前timestamp
	gen.timestamp = now

	// 写入timestamp
	if err := cache.RedisClient.Set(
		"timestamp"+strconv.Itoa(int(gen.node)),
		gen.timestamp, 0).Err(); err != nil {
		return 0, err
	}

	cache.RedisClient.Del("lock" + strconv.Itoa(int(gen.node)))

	// 移位运算，生产最终 ID
	result := ID((now-conf.ConfigObject.EpochTimeStamp)<<consts.TimeShift | (gen.node << consts.NodeShift) | (gen.step))

	return result, nil
}
