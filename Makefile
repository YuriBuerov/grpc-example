default: build

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "cmd/grpc-example/grpc-example" cmd/grpc-example/main.go
test:
	@go test -cover ./cmd/... ./pkg/...