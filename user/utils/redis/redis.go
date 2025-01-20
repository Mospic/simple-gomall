package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var RdbUserId *redis.Client

func InitRedis() {
	RdbUserId = redis.NewClient(&redis.Options{
		Addr:     "111.229.215.58:6379",
		Password: "",
		DB:       0, // 用户信息存入 DB0.
	})
}
