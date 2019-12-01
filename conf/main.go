package conf

import (
	"go-snowflake-service/cache"
	"log"
	"os"

	"github.com/hashicorp/consul/api"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
}

var ConfigObject *Config

func init() {
	// 加载环境变量
	godotenv.Load()

	// 加载Consul的YAML配置
	config := api.DefaultConfig()

	config.Address = os.Getenv("CONSUL_CONFIG_CENTER_URL")

	client, err := api.NewClient(config)

	if err != nil {
		log.Fatal(err)
	}

	data, _, err := client.KV().Get(os.Getenv("CONSUL_CONFIG_PATH"), nil)

	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data.Value, &ConfigObject)

	if err != nil {
		log.Fatal(err)
	}

	// 初始化所有的Client
	cache.InitRedis()
}
