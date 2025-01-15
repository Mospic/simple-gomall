package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var RdbUserId *redis.Client

func InitRedis() {
	RdbUserId = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "292023",
		DB:       0, // 用户信息存入 DB0.
	})
}
