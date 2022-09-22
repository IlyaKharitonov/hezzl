package internal

import (
	"encoding/json"
	"io/ioutil"

	"hezzlTestTask/infrastructure/databases"
	"hezzlTestTask/infrastructure/nats"
	"hezzlTestTask/infrastructure/redis"
)

type ConfigJSON struct {
	Server     Server             `json:"server"`
	Postgres   databases.DBConfig `json:"postgres"`
	ClickHouse databases.DBConfig `json:"clickHouse"`
	Redis      redis.RedisConfig  `json:"redis"`
	Nats       nats.NatsConfig    `json:"nats"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func Parse(path string) (*ConfigJSON, error) {
	var config = &ConfigJSON{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil
	}
	json.Unmarshal(file, &config)

	return config, err
}
