package handler

import (
  "time"
  "os"
	_ "context"
	"github.com/dgrijalva/jwt-go"
  pbuser "github.com/uta8a/suburi/server/proto/user"
)

const Tokenkey = "token"

func CreateToken(username string) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": username,
    "iat": time.Now().Unix(),
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  })
  tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  if err != nil {
    return "", err
  }
  return tokenString, nil
}

func Verify(req *pbuser.Request) (bool, error) {
  
  return true, nil
}

