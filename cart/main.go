package main

import (
	"cart/conf"
	"cart/core"
	"cart/rpc_server/etcd"
	protos "cart/services"
	"github.com/micro/go-micro/v2"
)

func main() {
	conf.Init()

	microService := micro.NewService(
		micro.Name("rpcCartService"),
		micro.Registry(etcd.EtcdReg),
	)

	microService.Init()

	_ = protos.RegisterCartServiceHandler(microService.Server(), new(core.CartService))
	_ = microService.Run()
}
