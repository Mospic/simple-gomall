package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var Ctx = context.Background()
var RdbJwt *redis.Client

func InitRedis() {
	RdbJwt = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "292023",
		DB:       0, // jwt 信息存入
	})
	ctx := context.Background()
	pong, err := RdbJwt.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("无法连接到Redis: %v", err)
	}
	fmt.Println("连接成功:", pong)
}
