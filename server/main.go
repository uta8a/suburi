package main

import (
  pb "suburi/server/helloworld"
  "log"
  "net"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
)

const (
  port = ":9090"
)

// server for helloworld.GreeterServer
type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
  return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterGreeterServer(s, &server{})
  // reflection
  reflection.Register(s)
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
