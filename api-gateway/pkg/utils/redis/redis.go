package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var RdbJwt *redis.Client

func InitRedis() {
	RdbJwt = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "292023",
		DB:       0, // jwt 信息存入
	})
}
