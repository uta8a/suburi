package middleauth

import (
  "os"
  "fmt"
  "log"
  "context"
  "github.com/dgrijalva/jwt-go"
  "github.com/grpc-ecosystem/go-grpc-middleware/auth"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"

  "github.com/uta8a/suburi/server/auth"
)
type AuthFunc func(ctx context.Context, fullMethodName string) (context.Context, error)

const (
  NoAuthorization = "NOAUTH"
)
var authnRoutes = map[string][]string {
  "/user.User/GetToken": { NoAuthorization },
  "/user.User/Login": { NoAuthorization },
  "/user.User/Register": { NoAuthorization },
}
func ApiAuth() AuthFunc {
  return func(ctx context.Context, fullMethodName string) (context.Context, error) {
    if _, ok := authnRoutes[fullMethodName]; ok == true {
      return ctx,nil
    }
    ctx, err := AuthMiddleware(ctx)
    if err != nil {
      return nil, err
    }
    return ctx, nil
  }
}
func AuthenticationUnaryServerInterceptor(authFunc AuthFunc) grpc.UnaryServerInterceptor {
  return func(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
  )(interface{}, error) {
    newCtx, err := authFunc(ctx, info.FullMethod)
    if err != nil {
      return nil, status.Error(codes.Unauthenticated, err.Error())
    }
    res, err := handler(newCtx, req)
    if err != nil {
      return nil, err
    }
    log.Printf("AuthenticationUnaryServerInterceptor: %+v\n", res)
    return res, nil
  }
}

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

  tokenString, err := jwt.Parse(token, func(tokenString *jwt.Token) (interface{}, error) {
    if _, ok := tokenString.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Unexpected signing method: %v", tokenString.Header["alg"])
    }
    return []byte(key), nil
  })
  if err != nil {
    if errorType, ok := err.(*jwt.ValidationError); ok {
      if errorType.Errors & jwt.ValidationErrorExpired != 0 {
        log.Printf("Expired token: %v", err)
        return nil, status.Errorf(
          codes.Unauthenticated,
          "token is expired: %v",
        )
      } else {
        log.Printf("Invalid token: %v", err)
        return nil, status.Errorf(
          codes.Unauthenticated,
          "token is invalid",
        )
      }
    } else {
      log.Printf("Invalid token: %v", err)
      return nil, status.Errorf(
        codes.Unauthenticated,
        "token is invalid",
      )
    }
  }
  log.Printf("TokenString: %v", tokenString)
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
func GetTokenString(ctx context.Context, key string) *auth.UserClaims {
  return ctx.Value(key).(*auth.UserClaims)
}


