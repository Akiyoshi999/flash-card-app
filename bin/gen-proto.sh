#!/bin/bash

protoc \
  -I ../backend/proto \
  --go_out=../backend/proto/protobuf \
  --go_opt=paths=source_relative \
  --go-grpc_out=../backend/proto/protobuf \
  --go-grpc_opt=paths=source_relative \
  deck.proto
