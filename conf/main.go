package conf

import (
	"go-snowflake-service/cache"
	"log"
	"os"

	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"
)

type Config struct {
}

var ConfigObject *Config

func init() {
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

	cache.InitRedis()
}
