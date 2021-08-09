module github.com/wuyazi/gomicrorpc

go 1.16

require (
	github.com/asim/go-micro/plugins/client/grpc/v3 v3.0.0-20210804083901-3e0411a3f61b
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.0.0-20210804083901-3e0411a3f61b
	github.com/asim/go-micro/plugins/server/grpc/v3 v3.0.0-20210804083901-3e0411a3f61b
	github.com/asim/go-micro/v3 v3.5.2
	google.golang.org/grpc/examples v0.0.0-20210806175644-574137db7de3 // indirect
	google.golang.org/protobuf v1.27.1
)

replace github.com/asim/go-micro/plugins/client/grpc/v3 => ../go-micro/plugins/client/grpc

replace github.com/asim/go-micro/plugins/registry/consul/v3 => ../go-micro/plugins/registry/consul

replace github.com/asim/go-micro/plugins/server/grpc/v3 => ../go-micro/plugins/server/grpc

replace github.com/asim/go-micro/v3 => ../go-micro

replace github.com/wuyazi/gomicrorpc => ../gomicrorpc
