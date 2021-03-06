### Makefile for micro
GO        := GO111MODULE=on go
GOBUILD   := GO111MODULE=on CGO_ENABLED=0 $(GO) build

ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"


.PHONY: test proto

test:
	go test -v ./...

proto:
	@cd ./internal/transport/proto/ && protoc --go_out=plugins=grpc:./ *.proto
