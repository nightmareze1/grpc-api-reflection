FROM golang:1.17-alpine AS build-env

RUN apk add --no-cache git protobuf
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. hello/hello.proto

RUN CGO_ENABLED=0 go build -o server main.go

FROM debian:bullseye-slim
COPY --from=build-env /app/server /server

RUN apt-get update
RUN apt-get install -y ca-certificates
RUN sed -i 's|http://|https://|g' /etc/apt/sources.list
RUN apt-get update && \
    apt-get install --no-install-recommends -y ca-certificates openssl libssl-dev libpq-dev wget && \
    rm -rf /var/lib/apt/lists/* && \
    wget https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.18/grpc_health_probe-linux-amd64 && \
    chmod +x grpc_health_probe-linux-amd64

EXPOSE 50051/tcp

ENTRYPOINT ["/server"]
