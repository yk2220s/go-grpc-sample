# gRPC implementation sample by golang

## requirement

see official [quickstart](https://grpc.io/docs/quickstart/go/)

- grpc
  - `go get -u google.golang.org/grpc`
- protoc-gen-go
  - install protoc compiler [url](https://github.com/google/protobuf/releases) or `brew install protobuf`
  - `go get -u github.com/golang/protobuf/protoc-gen-go`

## compile pb

`$ protoc -I blog/ blog/blog.proto --go_out=plugins=grpc:blog`

## using grpcurl

### list methods

`grpcurl -plaintext localhost:50001 list blog.Blog`

### call a method

`grpcurl -plaintext -d '{"param":"jsontext"}' localhost:50001 blog.Blog/GetPost`
