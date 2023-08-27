# Launch GRPC-API-REFLECTION
```
docker-compose -f docker/docker-compose.yml up --build

or

docker run -p 50051:50051 nightmareze1/pacman:grpc-api-reflection
```

# Test GRPC API in LOCAL(without TLS)

```
grpcurl -plaintext -import-path . -proto hello/hello.proto -d '{"name":"world 123123"}' localhost:50051 hello.Greeter/SayHello
```
