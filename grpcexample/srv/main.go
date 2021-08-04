package main

import (
	"context"
	"fmt"

	"github.com/lpxxn/gomicrorpc/grpcexample/common"
	model "github.com/lpxxn/gomicrorpc/grpcexample/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server/grpc"
	_ "github.com/micro/go-plugins/registry/consul"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *model.SayParam, rsp *model.SayResponse) error {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*model.Pair)
	rsp.Header["name"] = &model.Pair{Key: 1, Values: "abc"}

	rsp.Msg = "hello world"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = model.RespType_DESCEND

	return nil
}

func main() {
	// 初始化服务
	service := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Name(common.GrpcExampleName),
	)

	service.Init()

	// 注册 Handler
	if err := model.RegisterSayHandler(service.Server(), new(Say)); err != nil {
		panic(err)
	}

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
