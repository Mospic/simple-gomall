package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

var Ctx = context.Background()
var Rdb *redis.Client

var (
	Redisaddr     string
	RedisPassWord string
	RedisName     int
)

// LoadMysqlData 获取 Redis 配置
func LoadMysqlData(file *ini.File) {
	Redisaddr = file.Section("redis").Key("addr").String()
	RedisPassWord = file.Section("redis").Key("password").String()
	RedisName, _ = file.Section("redis").Key("db").Int()
}

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Redisaddr,
		Password: RedisPassWord,
		DB:       RedisName, // 用户信息存入 DB0.
	})
}
