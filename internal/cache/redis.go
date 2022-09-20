package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"

	"hezzlTestTask/config"
)

type RedisStruct struct {
	Key      string
	Expires  time.Duration
}


func(r *RedisStruct)Get(data interface{})error{
	b, err := config.Config.Redis.Client.Get(r.Key).Bytes()
	switch err {
	case nil:
	case redis.Nil:
		return err
	default:
		return fmt.Errorf("cache.(r *redisStruct) Get #1:\nError: %w", err)
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return fmt.Errorf("cache.(r *redisStruct) Get #2:\nError: %w", err)
	}

	return nil
}

func(r *RedisStruct)Load(data interface{})error{
	b,err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("cache.(r *redisStruct) Load #1:\nError: %w", err)
	}
	err = config.Config.Redis.Client.Set(r.Key, b, r.Expires).Err()
	if err != nil {
		return fmt.Errorf("cache.(r *redisStruct) Load #1:\nError: %w", err)
	}
	return nil
}

func(r *RedisStruct)Delete()error{
	err := config.Config.Redis.Client.Del(r.Key).Err()
	if err != nil {
		return fmt.Errorf("cache.(r *redisStruct) Load #1:\nError: %w", err)
	}
	return nil
}