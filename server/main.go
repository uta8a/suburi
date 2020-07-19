package main

import (
  "log"
  "net"
  pbUser "github.com/uta8a/suburi/server/proto/user"
  handler "github.com/uta8a/suburi/server/handler"

  _ "github.com/grpc-ecosystem/go-grpc-middleware"
  _ "github.com/grpc-ecosystem/go-grpc-middleware/auth"

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
  server := grpc.NewServer()
  pbUser.RegisterUserServer(server, handler.NewApp())
  reflection.Register(server)
  if err := server.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
