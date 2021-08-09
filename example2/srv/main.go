package main

import (
	_ "github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/lpxxn/gomicrorpc/example2/common"
	"github.com/lpxxn/gomicrorpc/example2/handler"
	"github.com/lpxxn/gomicrorpc/example2/proto/rpcapi"
	"github.com/lpxxn/gomicrorpc/example2/subscriber"
	"time"
)

func main() {

	// 初始化服务
	service := micro.NewService(
		micro.Name(common.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
	)

	service.Init()
	// 注册 Handler
	rpcapi.RegisterSayHandler(service.Server(), new(handler.Say))

	// Register Subscribers
	if err := server.Subscribe(server.NewSubscriber(common.Topic1, subscriber.Handler)); err != nil {
		panic(err)
	}

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
