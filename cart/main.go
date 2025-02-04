package main

import (
	"cart/conf"
	"cart/core"
	"cart/rpc_server/etcd"
	protos "cart/services"
	"github.com/micro/go-micro/v2"
)

func main() {
	// 初始化配置 & 数据库
	conf.Init()

	// 注册 Etcd
	microService := micro.NewService(
		micro.Name("rpcCartService"),
		micro.Registry(etcd.EtcdReg),
	)

	microService.Init()

	// 绑定 CartService 处理逻辑
	_ = protos.RegisterCartServiceHandler(microService.Server(), new(core.CartService))

	// 启动服务
	_ = microService.Run()
}
