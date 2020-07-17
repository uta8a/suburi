package main

import (
  pb "suburi/server/helloworld"
  server "suburi/server/handler"
  "log"
  "net"

  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
)

const (
  port = ":9090"
)

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterGreeterServer(s, server.NewApp())
  // reflection
  reflection.Register(s)
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
