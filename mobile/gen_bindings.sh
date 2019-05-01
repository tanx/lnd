#!/bin/sh

mkdir -p build

pkg="lndmobile"
target_pkg="github.com/lightningnetwork/lnd/lnrpc"

# Generate APIs by passing the parsed protos to ./gen
protoc -I/usr/local/include -I. \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --plugin=protoc-gen-custom=$GOPATH/bin/promobile \
       --custom_out=./build \
       --custom_opt="package_name=$pkg,target_package=$target_pkg" \
       --proto_path=../lnrpc \
       rpc.proto
