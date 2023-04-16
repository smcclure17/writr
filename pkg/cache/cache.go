package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client  *redis.Client
	context context.Context
}

func NewCache() *Cache {
	redisDatabase := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Cache{
		Client:  redisDatabase,
		context: context.Background(),
	}
}

func (cache *Cache) SetCacheMessage(key string, value string) {
	err := cache.Client.Set(cache.context, key, value, 10000).Err()
	if err != nil {
		panic(err)
	}
}

func (cache *Cache) GetCacheMessage(key string) string {
	val, err := cache.Client.Get(cache.context, key).Result()
	if err != nil {
		if err == redis.Nil {
			return ""
		}
	}
	return val
}
