PROTO_DIR=proto
PB_DIR=pkg/pb

PROTOC=protoc
PROTOC_GEN_GO=protoc-gen-go
PROTOC_GEN_GO_GRPC=protoc-gen-go-grpc

.PHONY: all proto build server client implant clean

all: proto build

proto:
	@mkdir -p $(PB_DIR)
	$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(PB_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PB_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto

build: server admin-client implant

server: proto
	go build -o bin/server cmd/server/main.go

admin-client: proto
	go build -o bin/admin-client cmd/client/main.go

implant: proto
	go build -o bin/implant cmd/implant/main.go

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

test:
	go test ./...

clean:
	rm -rf bin
	rm -rf $(PB_DIR)/*.pb.go

