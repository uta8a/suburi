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

// for login
func Verify(ctx context.Context, req *pbuser.Request, con boil.ContextExecutor) (bool, string) {
  // DB username password_hash

  // TODO argon2 verify
  // for dev, raw password
  // password_hash := hash(req.password)
  password_hash := req.Password
  username := req.Username
  log.Printf("user:pass = %s:%s", username, password_hash)
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

// for register
// usertype, error
func Register(ctx context.Context, req *pbuser.Request, con boil.ContextExecutor) (string, error) {
  userinfo := db.Userinfo{ ID: 1, Username: req.Username, PasswordHash: req.Password, UserType: "player" }
  err := userinfo.Insert(ctx, con, boil.Blacklist("id"))
  if err != nil {
    log.Printf("Register error: %v", err.Error())
    return "", err
  }
  return "player", nil
}
