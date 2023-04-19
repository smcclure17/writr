package cache

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client  *redis.Client
	context context.Context
}

func NewCache() *Cache {
	redisDatabase := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})

	err := redisDatabase.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("error with redis database: ", err)
	}
	fmt.Println("Connected to redis at: ", redisDatabase.Options().Addr)
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
