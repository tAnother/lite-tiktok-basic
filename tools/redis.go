package tools

import (
	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func RedisInit() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 如果没有密码则为空
		DB:       0,  // 默认 DB
	})
}

func GetClient() *redis.Client {
	return Client
}
