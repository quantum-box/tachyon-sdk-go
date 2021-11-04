#!/bin/bash

# https://grpc.io/docs/languages/go/quickstart/
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/cms.proto
