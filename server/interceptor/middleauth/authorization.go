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
  "/check.Routes/HealthCheck": { PermissionPlayer },
  "/check.Routes/TesterCheck": { PermissionTester },
  "/check.Routes/SecretCheck": { PermissionAdmin },
}

type Role struct {
  permissions []string
}

// TODO Subject.usertype
func AuthorizationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
  return func(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
  )(interface{}, error) {
    if canAccess(info.FullMethod, getRole(GetToken(ctx).Subject)) {
      return handler(ctx, req)
    }
    return nil, status.Error(
      codes.PermissionDenied,
      "could not access to specified method",
    )
  }
}

func getRole(usertype string) *Role {
  switch usertype {
  case "player":
    return &User{permissions: []string{PermissionPlayer}}
  case "tester":
    return &User{permissions: []string{PermissionTester}}
  case "admin":
    return &User{permissions: []string{PermissionPlayer,PermissionTester, PermissionAdmin}}
  }
  return &User{}
}
// TODO modify
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

