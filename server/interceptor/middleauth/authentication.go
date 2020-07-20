package middleauth

import (
  "os"
  "log"
  "context"
  "github.com/dgrijalva/jwt-go"
  "github.com/grpc-ecosystem/go-grpc-middleware/auth"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"

  "github.com/uta8a/suburi/server/auth"
)


func AuthMiddleware(ctx context.Context) (context.Context, error) {
  key := os.Getenv("TOKEN_SECRET")
  if key == "" {
    log.Fatal("cannot read TOKEN_SECRET from .env")
  }
  token , err := grpc_auth.AuthFromMD(ctx, "bearer")
  if err != nil {
    return nil, status.Errorf(
      codes.Unauthenticated,
      "could not read auth token: %v",
      err,
    )
  }

  // TODO demo, so skip signature
  parser := new(jwt.Parser)
  parsedToken, _, err := parser.ParseUnverified(token, &auth.UserClaims{})
  if err != nil {
    return nil, status.Errorf(
      codes.Unauthenticated,
      "could not parsed auth token: %v",
      err,
    )
  }

  return setToken(ctx, parsedToken.Claims.(*auth.UserClaims), key), nil
}

func setToken(ctx context.Context, token *auth.UserClaims, key string) context.Context {
  return context.WithValue(ctx, key, token)
}
func GetToken(ctx context.Context, key string) *auth.UserClaims {
  return ctx.Value(key).(*auth.UserClaims)
}
