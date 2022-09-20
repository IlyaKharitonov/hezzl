package redis

import "github.com/go-redis/redis"

type RedisConfig struct {
	Client   *redis.Client
	Host     string 	`json:"host"`
	Port     string     `json:"port"`
	Password string 	`json:"password"`
	NameDB     int    	`json:"name"`
}
