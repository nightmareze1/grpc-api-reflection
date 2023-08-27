package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "myproject.com/hello"
)

type server struct {
    hello.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
    log.Printf("Received request from %s", in.Name)
    return &hello.HelloReply{Message: "Hello dtp-pacman grpc" + in.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    log.Println("Server started, listening on :50051")

    s := grpc.NewServer()
    hello.RegisterGreeterServer(s, &server{})

    // Register reflection service on gRPC server.
    reflection.Register(s)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

