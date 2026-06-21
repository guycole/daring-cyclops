# daring-cyclops

Daring Cyclops is a Go command-line client and Go gRPC server for a multiplayer
space strategy game inspired by [MegaWars](https://en.wikipedia.org/wiki/MegaWars).

## Ping

The first implemented feature is a unary gRPC ping endpoint. The server returns
its version and current time, and the CLI client prints both values in
human-readable output.

## Requirements

- Go 1.26+
- `protoc`
- `protoc-gen-go`
- `protoc-gen-go-grpc`

## Generate protobuf code

```bash
source ~/.bash_aliases
make proto
```

## Run tests

```bash
go test ./...
```

## Start the server

```bash
go run ./cmd/cyclopsd -listen 127.0.0.1:50051
```

## Ping the server

```bash
go run ./cmd/cyclops ping -server 127.0.0.1:50051
```

Example output:

```text
Server Version: dev
Server Time: 2026-06-15T12:34:56Z
```
