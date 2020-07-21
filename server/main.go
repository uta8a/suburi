package main

import (
  "log"
  "net"
  pbUser "github.com/uta8a/suburi/server/proto/user"
  pbCheck "github.com/uta8a/suburi/server/proto/check"
  handler "github.com/uta8a/suburi/server/handler"
  middleauth "github.com/uta8a/suburi/server/interceptor/middleauth"

  "github.com/grpc-ecosystem/go-grpc-middleware"

  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
)

const (
  port = ":9090"
)

func main() {
  con,err := handler.InitDB()
  if err != nil {
    log.Fatalf("failed to connect to DB: %v", err)
  }
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  server := grpc.NewServer(
    grpc_middleware.WithUnaryServerChain(
      middleauth.AuthenticationUnaryServerInterceptor(middleauth.ApiAuth()),
      middleauth.AuthorizationUnaryServerInterceptor(),
    ),
  )
  pbUser.RegisterUserServer(server, handler.NewApp(con))
  pbCheck.RegisterRoutesServer(server, handler.NewApp(con))
  reflection.Register(server)
  if err := server.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
