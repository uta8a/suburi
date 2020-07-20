package auth

import (
  "time"
  "log"
  "os"
  "context"

  jwt "github.com/dgrijalva/jwt-go"
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
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
    username,
    usertype,
    jwt.StandardClaims {
      IssuedAt: time.Now().Unix(),
      ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
    },
  })
  tokenString, err := token.SignedString([]byte(TokenSecret))
  if err != nil {
    return "", err
  }
  return tokenString, nil
}

func Verify(ctx context.Context, req *pbuser.Request, con boil.ContextExecutor) (bool, string) {
  // DB username password_hash

  // TODO argon2 verify
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

