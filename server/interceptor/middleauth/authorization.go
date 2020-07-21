package middleauth

import (
  "os"
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

var authzRoutes = map[string][]string {
  "/check.Routes/HealthCheck": { PermissionPlayer },
  "/check.Routes/TesterCheck": { PermissionTester },
  "/check.Routes/SecretCheck": { PermissionAdmin },
}

type Role struct {
  permissions []string
}

func AuthorizationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
  key := os.Getenv("TOKEN_SECRET")
  if key == "" {
    log.Fatal("cannot read TOKEN_SECRET from .env")
  }
  return func(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
  )(interface{}, error) {
    if _, ok := authzRoutes[info.FullMethod]; ok != true {
      return handler(ctx, req)
    }
    log.Printf("Username: %v", GetTokenString(ctx, key))
    log.Printf("Usertype: %+v", GetTokenString(ctx, key).Usertype)
    if canAccess(info.FullMethod, getRole(GetTokenString(ctx, key).Usertype)) {
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
    return &Role{permissions: []string{PermissionPlayer}}
  case "tester":
    return &Role{permissions: []string{PermissionTester}}
  case "admin":
    return &Role{permissions: []string{PermissionPlayer,PermissionTester, PermissionAdmin}}
  }
  return &Role{}
}
func canAccess(method string, role *Role) bool {
  r, ok := authzRoutes[method]
  log.Printf("Access: %v %v",r,ok)
  if !ok {
    return false
  }
  permissions := map[string]bool{}
  for _, p := range role.permissions {
    permissions[p] = true
  }

  for _, p := range r {
    log.Printf("AccessInternal: %+v %v",permissions, p)
    if !permissions[p] {
      return false
    }
  }
  return true
}

