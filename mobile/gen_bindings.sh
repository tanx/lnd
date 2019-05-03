#!/bin/sh

mkdir -p build

# Check promobile version.
promobileVersion="0.1"
promobile=$(which promobile)
if [ $promobile ]
then
        version=$($promobile -v)
        if [ $version != $promobileVersion ]
        then
                echo "promobile version $promobileVersion required"
                exit 1
        fi
        echo "Using plugin $promobile $version"
else
        echo "promobile not found"
        exit 1
fi

pkg="lndmobile"
target_pkg="github.com/lightningnetwork/lnd/lnrpc"

# Generate APIs by passing the parsed protos to ./gen
protoc -I/usr/local/include -I. \
       -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --plugin=protoc-gen-custom=$promobile\
       --custom_out=./build \
       --custom_opt="package_name=$pkg,target_package=$target_pkg,listeners=lightning=lightningLis walletunlocker=walletUnlockerLis" \
       --proto_path=../lnrpc \
       rpc.proto

for file in ../lnrpc/**/*.proto
do
    DIRECTORY=$(dirname ${file})
    tag=$(basename ${DIRECTORY})
    filename=$(basename ${file})
    service=${filename%.proto}
    build_tags="// +build $tag"
    use_prefix="0"
    lis="$service=lightningLis"

    opts="package_name=$pkg,target_package=$target_pkg/$tag,build_tags=$build_tags,api_prefix=$use_prefix,listeners=$lis"

    echo "Generating mobile protos from ${file}, with build tag ${tag}"

    protoc -I/usr/local/include -I. \
           -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
           -I../lnrpc \
           --plugin=protoc-gen-custom=$promobile \
           --custom_out=./build \
           --custom_opt="$opts" \
           --proto_path=${DIRECTORY} \
           ${file}
done
