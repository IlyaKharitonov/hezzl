package redis

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
)

func (r *RedisConfig) Connect() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
		Password: r.Password,
		DB:       r.NameDB,
	})

	pong, err := client.Ping().Result()
	if err == nil && pong != "PONG" || err != nil {
		return nil, errors.New("No connect redis")
	}
	return client, nil
}
