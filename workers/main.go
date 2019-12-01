package workers

import (
	"fmt"
	"go-snowflake-service/cache"
	"go-snowflake-service/conf"
	"go-snowflake-service/generator_pool"
	"log"
	"strconv"
	"time"
)

func StartAutoGenerator() {
	go autoGenerator()
}

func autoGenerator() {
	for {
		node_list, err := cache.RedisClient.LRange("node-ids", 0, -1).Result()

		if err != nil {
			log.Println("[Error] ", err)
		}

		go gen(node_list)

		time.Sleep(time.Second * time.Duration(conf.ConfigObject.AutoGeneratorIdInterval))
	}
}

func gen(node_list []string) {
	for _, nodeid := range node_list {
		node_id, _ := strconv.Atoi(nodeid)

		flake := generator_pool.GetSnowFlake(int64(node_id))

		data, err := flake.GenerateId()

		if err != nil {
			log.Println("[Error] ", err)
		}

		cache.RedisClient.LPush(fmt.Sprintf("%s-id-list", nodeid), int64(data))
	}
}
