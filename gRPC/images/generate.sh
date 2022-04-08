#!/bin/bash


echo $PWD

SRC_DIR="$PWD/imagepb"
DST_DIR="$PWD"

echo "dst -  $DST_DIR"

protoc imagepb/*.proto \
    --go_out=. \
    --go_opt=paths=source_relative \
    --proto_path=. \
    --go-grpc_out=$DST_DIR



