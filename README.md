# grpc-example
1. Install `dep` tool for dependency management:
```sh
brew install dep
```
```sh
dep ensure
```
# Run Tests
```bash
make test
```

# Run App
```bash
make
docker build -t grpc-example -f ./cmd/grpc-example/Dockerfile cmd/grpc-example
docker run -it --rm --read-only \
        --volume $(pwd)/docker-tmp:/tmp \
        --publish 50051:50051 \
        grpc-example
```
* Connection test
```bash
curl -L https://127.0.0.1:50051  -XGET
```

* Client example
```bash
go run ./cmd/testclient/main.go -port 50051 -limit 10
```
