package grpc_proto

//go:generate protoc -I $GOPATH/src/github.com/apiclarity/trace-sampling-manager/api/grpc_proto/trace-sampling-manager-service trace-sampling-manager.proto --go_out=plugins=grpc:$GOPATH/src/github.com/apiclarity/trace-sampling-manager/api/grpc_gen/trace-sampling-manager-service --go_opt=paths=source_relative
