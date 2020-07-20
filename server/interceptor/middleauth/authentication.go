package middleauth

import (
  "context"
  "github.com/dgrijalva/jwt-go"
  "github.com/grpc-ecosystem/go-grpc-middleware/auth"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
)


func AuthFunc(ctx context.Context) (context.Context, error) {
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
  parsedToken, _, err := parser.ParseUnverified(token, &jwt.StandardClaims{})
  if err != nil {
    return nil, status.Errorf(
      codes.Unauthenticated,
      "could not parsed auth token: %v",
      err,
    )
  }

  return setToken(ctx, parsedToken.Claims.(*jwt.StandardClaims), key), nil
}

func setToken(ctx context.Context, token *jwt.StandardClaims, key string) context.Context {
  return context.WithValue(ctx, key, token)
}

