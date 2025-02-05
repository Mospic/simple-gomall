package main

import (
	"cart/conf"
	"cart/core"
	services "cart/services"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"time"
)

func main() {
	conf.Init()
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	microService := micro.NewService(
		micro.Name("rpcCartService"),
		micro.Address("127.0.0.1:8084"),
		micro.Registry(etcdReg),
		micro.RegisterTTL(24*time.Hour),
		micro.Metadata(map[string]string{"protocol": "http"}),
	)

	microService.Init()

	_ = services.RegisterCartServiceHandler(microService.Server(), new(core.CartService))

	_ = microService.Run()
}
