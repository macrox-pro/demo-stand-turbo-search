# Prepare

## Protobuf + gRPC 

For MacOs

```shell
brew install protobuf
```

## Install the protocol compiler plugins for Go

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## GraphQL

```shell
go run github.com/99designs/gqlgen generate
```