#!/bin/bash

# protoc greet/greetpb/greet.proto --go_out=plugins=grps:.ß
echo $PWD

SRC_DIR="$PWD/greet/greetpb"
DST_DIR="$PWD/greet"

echo "dst -  $DST_DIR"

# protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/greet.proto

#  protoc --go_out=. --go_opt=paths=$SRC_DIR --go-grpc_out=. --go-grpc_opt=paths=$SRC_DIR  "$SRC_DIR/greet.proto"

# protoc --go-grpc_out=. --go_opt=paths=source_relative \. "greet/greetpb/greet.proto"

# v
# protoc --go_out=paths=source_relative:.   -I. greet/greetpb/*.proto
# v
# protoc  --go_out=$DST_DIR --go-grpc_out=$DST_DIR  -I. greet/greetpb/*.proto

protoc greet/greetpb/*.proto \
    --go_out=. \
    --go_opt=paths=source_relative \
    --proto_path=. \
    --go-grpc_out=$DST_DIR
    # --go-grpc_out=require_unimplemented_servers=false:. 


