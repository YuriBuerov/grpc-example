# grpc-example
1. Install `trash` tool for dependency management:
```sh
go get -u github.com/rancher/trash
```
```sh
mkdir vendor && trash
```
# Run Tests
```bash
make test
```

# Run App
```bash
make
docker build -t grpc-example -f ./cmd/grpc-example/Dockerfile .
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
go run ./cmd/testclient/main.go -port 50051 -domain-name stackoverflow.com
```
