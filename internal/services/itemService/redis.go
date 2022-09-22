package itemService

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type redisCache struct {
	*redis.Client
}

func NewRedisCache(rc *redis.Client) *redisCache {
	return &redisCache{rc}
}

func (rc *redisCache) GetList(key string) ([]*Item, error) {
	items := make([]*Item, 0)
	b, err := rc.Get(key).Bytes()
	switch err {
	case nil:
	case redis.Nil:
		return nil, err
	default:
		return nil, fmt.Errorf("(rc *redisCache) Get #1:\nError: %w", err)
	}

	err = json.Unmarshal(b, &items)
	if err != nil {
		return nil, fmt.Errorf("(rc *redisCache) Get #2:\nError: %w", err)
	}

	return items, nil
}

func (rc *redisCache) Load(data interface{}, key string, expires time.Duration) error {
	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("(rc *redisCache)Load #1:\nError: %w", err)
	}
	err = rc.Set(key, b, expires).Err()
	if err != nil {
		return fmt.Errorf("(rc *redisCache)Load #1:\nError: %w", err)
	}
	return nil
}

func (rc *redisCache) Delete(key string) error {
	err := rc.Del(key).Err()
	if err != nil {
		return fmt.Errorf("(rc *redisCache)Delete #1:\nError: %w", err)
	}
	return nil
}
