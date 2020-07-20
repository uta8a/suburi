package handler

import (
  "time"
  "log"
  "os"
  "context"

  "github.com/dgrijalva/jwt-go"
  _ "golang.org/x/crypto/argon2"
  "github.com/volatiletech/sqlboiler/v4/boil"

  pbuser "github.com/uta8a/suburi/server/proto/user"
  "github.com/uta8a/suburi/server/db"
)

func CreateToken(username string, usertype string) (string, error) {
  
  TokenSecret := os.Getenv("TOKEN_SECRET")
  if TokenSecret == "" {
    log.Fatal("Env is not set")
  }
  //log.Printf("ENV: %s", TokenSecret)
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": username,
    "usertype": usertype,
    "iat": time.Now().Unix(),
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  })
  tokenString, err := token.SignedString([]byte(TokenSecret))
  if err != nil {
    return "", err
  }
  return tokenString, nil
}

func Verify(ctx context.Context, req *pbuser.Request, con boil.ContextExecutor) (bool, string) {
  // DB username password_hash

  // for dev, raw password
  // password_hash := hash(req.password)
  password_hash := req.Password
  username := req.Username
  
  userinfo, err := db.Userinfos(db.UserinfoWhere.Username.EQ(username)).One(ctx, con)
  if err != nil {
    log.Printf("Verify db error: %v", err.Error())
    return false, ""
  }
  log.Printf("UserInfo: %+v", userinfo)
  if password_hash != userinfo.PasswordHash {
    return false, ""
  }
  
  return true, userinfo.UserType
}

