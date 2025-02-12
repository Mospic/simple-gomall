package casbin

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	mlog "mall/log"
	"time"
)

func init() {
	log = mlog.NewLog("Auth", mlog.Info)
	RDB = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   1,
	})

	if err := RDB.Ping(context.Background()).Err(); err != nil {
		log.Warn("redis ping fail:" + err.Error())
		RDB = nil
	}

	dsn := "root:yzl222888@tcp(127.0.0.1:3306)/simple_mall?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Error(err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Warn(err.Error())
	} else {
		sqlDB.SetMaxOpenConns(16)
		sqlDB.SetMaxIdleConns(8)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

}

var Key = []byte("1-10 mall key")
var RDB *redis.Client
