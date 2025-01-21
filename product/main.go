package main

import (
	"time"
	//"user/utils/redis"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"product/conf"
	"product/core"
	services "product/services"
)

func main() {
	conf.Init()
	// 如果需要用到Redis，在此处初始化Redis连接实例
	//redis.InitRedis()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	microService := micro.NewService(
		micro.Name("rpcProductService"), // 微服务名字
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg),         // etcd注册件
		micro.RegisterTTL(24*time.Hour), // TTL时间
		micro.Metadata(map[string]string{"protocol": "http"}),
	)

	microService.Init()

	_ = services.RegisterProductServiceHandler(microService.Server(), new(core.ProductService))

	_ = microService.Run()
}
