package handler

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const Tokenkey = "token"

func AuthFunc(ctx context.Context) (context.Context, error) {
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

  return setToken(ctx, parsedToken.Claims.(*jwt.StandardClaims)), nil

}

func setToken(ctx context.Context, token *jwt.StandardClaims) context.Context {
  return context.WithValue(ctx, Tokenkey, token)
}

func GetToken(ctx context.Context) *jwt.StandardClaims {
  return ctx.Value(Tokenkey).(*jwt.StandardClaims)
}
