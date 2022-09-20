package redis

import (
	"fmt"

	"errors"
	"github.com/go-redis/redis"
)

func(r *RedisConfig)Connect()error{
	r.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
		Password: r.Password,
		DB:       r.NameDB,
	})

	pong, err := r.Client.Ping().Result()
	if err == nil && pong != "PONG" {
		return errors.New("No connect redis")
	}
	return err
}
