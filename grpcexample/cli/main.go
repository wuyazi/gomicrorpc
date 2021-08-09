package main

import (
	"context"
	"fmt"
	"time"

	"github.com/asim/go-micro/plugins/client/grpc/v3"
	_ "github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/lpxxn/gomicrorpc/grpcexample/common"
	"github.com/lpxxn/gomicrorpc/grpcexample/proto"
)

func main() {
	// 初始化服务
	service := micro.NewService(
		micro.Client(grpc.NewClient()),
	)
	service.Init()
	service.Options().Registry.Init(func(options *registry.Options) {
		options.Timeout = time.Second * 2
	})

	sayClent := proto.NewSayService(common.GrpcExampleName, service.Client())

	rsp, err := sayClent.Hello(context.Background(), &proto.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}
