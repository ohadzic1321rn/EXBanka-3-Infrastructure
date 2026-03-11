.PHONY: all build run seed proto test docker-up docker-down tidy clean

all: proto build

# Generate Go code from proto files (requires buf installed: https://buf.build/docs/installation)
proto:
	buf generate

build: proto
	go build -o bin/server ./cmd/server
	go build -o bin/seed ./cmd/seed

run:
	go run ./cmd/server

seed:
	go run ./cmd/seed

tidy:
	go mod tidy

test:
	go test ./...

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

clean:
	rm -rf bin/ gen/proto/

# Setup: install buf CLI (run once)
install-buf:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
