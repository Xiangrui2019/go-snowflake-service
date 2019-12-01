package snowflake

import (
	"go-snowflake-service/conf"
	"go-snowflake-service/consts"
	"sync"
	"time"
)

type ID int64

type SnowFlake struct {
	mutex     *sync.Mutex
	timestamp int64
	node      int64
	step      int64
}

func NewSnowFlake(node int64) *SnowFlake {
	return &SnowFlake{
		mutex:     &sync.Mutex{},
		timestamp: 0,
		node:      node + conf.ConfigObject.Offset,
		step:      0,
	}
}

func (flake *SnowFlake) NextId() (ID, error) {
	flake.mutex.Lock()         // 保证并发安全, 加锁
	defer flake.mutex.Unlock() // 方法运行完毕后解锁

	// 获取当前时间的时间戳 (毫秒数显示)
	now := time.Now().UnixNano() / 1e6

	if flake.timestamp == now {
		// step 步进 1
		flake.step++

		// 当前 step 用完
		if flake.step > consts.StepMax {
			// 等待本毫秒结束
			for now <= flake.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}

	} else {
		// 本毫秒内 step 用完
		flake.step = 0
	}

	flake.timestamp = now
	// 移位运算，生产最终 ID
	result := ID((now-conf.ConfigObject.EpochTimeStamp)<<consts.TimeShift | (flake.node << consts.NodeShift) | (flake.step))

	return result, nil
}
