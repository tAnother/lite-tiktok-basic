package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func RedisInit() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 如果没有密码则为空
		DB:       0,  // 默认 DB
	})
	client.Set(context.Background(), "", "", 0) // empty token - invalid userid
}

func RedisClient() *redis.Client {
	return client
}
