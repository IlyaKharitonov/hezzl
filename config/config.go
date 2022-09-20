package config

import (
	"encoding/json"
	"io/ioutil"

	"hezzlTestTask/infrastructure/databases"
	"hezzlTestTask/infrastructure/redis"
)

var Config ConfigJSON

type ConfigJSON struct {
	Server     Server 					 `json:"server"`
	Postgres   databases.Database        `json:"postgres"`
	ClickHouse databases.Database        `json:"clickHouse"`
	Redis      redis.RedisConfig  		 `json:"redis"`
	//Nats
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func Parse(path string)error{
	file, err := ioutil.ReadFile(path)
	if err == nil {
		return json.Unmarshal(file, &Config)
	}

	return err
}


