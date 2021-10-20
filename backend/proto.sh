#!/bin/bash

for t in rbac user index admin; do \
  rm -f rm -f api/protos/$t/*.go api/protos/$t/*.json
  protoc -Iapi/protos -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I $GOPATH/src/github.com/envoyproxy/protoc-gen-validate \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    --go_out=plugins=grpc,paths=source_relative:api/protos --validate_out=lang=go,paths=source_relative:api/protos api/protos/$t/*.proto; \
  if ls api/protos/$t/public*.proto 1> /dev/null 2>&1;  then protoc -Iapi/protos -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    -I $GOPATH/src/github.com/envoyproxy/protoc-gen-validate \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    --grpc-gateway_out=logtostderr=true,paths=source_relative:api/protos --swagger_out=logtostderr=true:api/protos \
     api/protos/$t/public*.proto; fi \
done