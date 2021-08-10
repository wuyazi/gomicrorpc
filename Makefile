build-proto:
	protoc --proto_path=$(GOPATH)/src:./example2/proto/model --go_out=paths=source_relative:./example2/proto/model/ ./example2/proto/model/*.proto
	protoc --proto_path=$(GOPATH)/src:./example2/proto/model:./example2/proto/rpcapi --micro_out=paths=source_relative:./example2/proto/rpcapi/ ./example2/proto/rpcapi/*.proto