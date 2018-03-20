```bash
docker build -t grpc-example -f Dockerfile .
docker run -it --rm --read-only \
        --volume $(pwd)/docker-tmp:/tmp \
        --publish 50051:50051 \
        grpc-example
```