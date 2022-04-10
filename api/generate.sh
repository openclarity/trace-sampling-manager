#!/bin/sh

set -euo pipefail

alias goswagger="docker run --rm -it --user $(id -u):$(id -g) -e GOPATH=$GOPATH:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger:v0.27.0"

cp server/restapi/configure_trace_sampling_manager.go /tmp/configure_trace_sampling_manager.go

rm -rf client/*
goswagger generate client -f swagger.yaml -t ./client

rm -rf server/*
goswagger generate server -f swagger.yaml -t ./server
cp /tmp/configure_trace_sampling_manager.go server/restapi/configure_trace_sampling_manager.go

rm -rf ./grpc_gen/*
mkdir -p ./grpc_gen/trace-sampling-manager-service
go generate ./grpc_proto/