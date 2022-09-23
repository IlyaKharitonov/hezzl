package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func (r *RedisConfig) Connect() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
		Password: r.Password,
		DB:       r.NameDB,
	})

	_, err := client.Ping().Result()
	if err != nil{
		return nil, err
	}
	//if err == nil && pong != "PONG"{
	//	return nil, errors.New("No connect redis")
	//}

	return client, nil
}
