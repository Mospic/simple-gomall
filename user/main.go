package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"time"
	"user/conf"
	"user/core"
	services "user/services"
	"user/utils/redis"
)

func main() {
	conf.Init()
	// 初始化redis-DB0的连接，follow选择的DB0.
	redis.InitRedis()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	microService := micro.NewService(
		micro.Name("rpcUserService"), // 微服务名字
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg),         // etcd注册件
		micro.RegisterTTL(24*time.Hour), // TTL时间
		micro.Metadata(map[string]string{"protocol": "http"}),
	)

	microService.Init()

	_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))

	_ = microService.Run()
}
