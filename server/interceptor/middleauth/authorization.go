package middleauth

import (
  "log"
  "context"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
)

const (
  PermissionAdmin = "ADMIN"
  PermissionTester = "TESTER"
	PermissionPlayer = "PLAYER"
)

var routes = map[string][]string {
  "/helloworld.Greeter/Hello": { PermissionHello },
  "/helloworld.Greeter/TellMeSecret": { PermissionSecret },
}

type User struct {
  permissions []string
}

func AuthorizationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
  return func(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
  )(interface{}, error) {
    if canAccess(info.FullMethod, getUser(GetToken(ctx).Subject)) {
      return handler(ctx, req)
    }
    return nil, status.Error(
      codes.PermissionDenied,
      "could not access to specified method",
    )
  }
}

func getUser(id string) *User {
  switch id {
  case "alice":
    return &User{permissions: []string{PermissionHello}}
  case "bob":
    return &User{permissions: []string{PermissionHello, PermissionSecret}}
  }
  return &User{}
}

func canAccess(method string, user *User) bool {
  r, ok := routes[method]
  log.Print("Access: %v %v",r,ok)
  if !ok {
    return false
  }
  permissions := map[string]bool{}
  for _, p := range user.permissions {
    permissions[p] = true
  }

  for _, p := range r {
    log.Print("AccessInternal: %+v %v",permissions, p)
    if !permissions[p] {
      return false
    }
  }
  return true
}

