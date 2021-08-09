package handler

import (
	"context"
	"fmt"
	"github.com/wuyazi/gomicrorpc/example2/lib"
	"github.com/wuyazi/gomicrorpc/example2/proto/model"
	"github.com/wuyazi/gomicrorpc/example2/proto/rpcapi"
	"io"
	"time"
)

type Say struct{}

var _ rpcapi.SayHandler = (*Say)(nil)

func (s *Say) Hello(ctx context.Context, req *model.SayParam, rsp *model.SayResponse) error {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*model.Pair)
	rsp.Header["name"] = &model.Pair{Key: 1, Values: "abc"}

	rsp.Msg = "hello world"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = model.RespType_DESCEND

	return nil
}

func (s *Say) MyName(ctx context.Context, req *model.SayParam, rsp *model.SayParam) error {
	rsp.Msg = "lp"
	return nil
}

/*
 模拟得到一些数据
*/
func (s *Say) Stream(ctx context.Context, req *model.SRequest, stream rpcapi.Say_StreamStream) error {

	for i := 0; i < int(req.Count); i++ {
		rsp := &model.SResponse{}
		for j := lib.Random(3, 5); j < 10; j++ {
			rsp.Value = append(rsp.Value, lib.RandomStr(lib.Random(3, 10)))
		}
		if err := stream.Send(rsp); err != nil {
			return err
		}
		// 模拟处理过程
		time.Sleep(time.Microsecond * 50)
	}
	return nil
}

/*
 模拟数据
*/
func (s *Say) BidirectionalStream(ctx context.Context, stream rpcapi.Say_BidirectionalStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(req.Count)
		if err := stream.Send(&model.SResponse{Value: []string{lib.RandomStr(lib.Random(3, 6))}}); err != nil {
			return err
		}
	}
	fmt.Println("end BidirectionalStream")
	return nil
}
