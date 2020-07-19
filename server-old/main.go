package main

import (
	"log"
	"net"
	server "suburi/server/handler"
	pb "suburi/server/helloworld"

  "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"

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
	s := grpc.NewServer(
    grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
      grpc_auth.UnaryServerInterceptor(server.AuthFunc),
      server.AuthorizationUnaryServerInterceptor(),
    )),
  )
	pb.RegisterGreeterServer(s, server.NewApp())
	// reflection
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}