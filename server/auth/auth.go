package handler

import (
  "time"
  "log"
  "os"
  _ "context"
  "github.com/dgrijalva/jwt-go"
  pbuser "github.com/uta8a/suburi/server/proto/user"
)

func CreateToken(username string) (string, error) {
  
  TokenSecret := os.Getenv("TOKEN_SECRET")
  if TokenSecret == "" {
    log.Fatal("Env is not set")
  }
  //log.Printf("ENV: %s", TokenSecret)
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": username,
    "iat": time.Now().Unix(),
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  })
  tokenString, err := token.SignedString([]byte(TokenSecret))
  if err != nil {
    return "", err
  }
  return tokenString, nil
}

func Verify(req *pbuser.Request) (bool, error) {
  
  return true, nil
}

