package main

import (
	"context"
	"fmt"
	_ "github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/wuyazi/gomicrorpc/example1/proto"
)

func main() {

	// 初始化服务
	service := micro.NewService()

	service.Init()

	sayClent := proto.NewSayService("lp.srv.eg1", service.Client())

	rsp, err := sayClent.Hello(context.Background(), &proto.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)

}
