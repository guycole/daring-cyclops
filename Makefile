GO ?= go
GOPATH ?= $(shell go env GOPATH)
export PATH := $(GOPATH)/bin:$(PATH)

.PHONY: proto

proto:
	protoc \
		--go_out=. \
		--go_opt=module=github.com/guycole/daring-cyclops \
		--go-grpc_out=. \
		--go-grpc_opt=module=github.com/guycole/daring-cyclops \
		proto/ping/v1/ping.proto \
		proto/session/v1/session.proto