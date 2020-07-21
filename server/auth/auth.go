package auth

import (
  "time"
  "log"
  "os"
  "context"
  "errors"

  jwt "github.com/dgrijalva/jwt-go"
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
  password := req.Password
  username := req.Username
  userinfo, err := db.Userinfos(db.UserinfoWhere.Username.EQ(username)).One(ctx, con)
  if err != nil {
    log.Printf("Verify db error: %v", err.Error())
    return false, ""
  }
  log.Printf("UserInfo: %+v", userinfo)
  ok,err := ComparePasswordAndHash(password, userinfo.PasswordHash)
  if err != nil {
    log.Printf("ComparePasswordAndHash Error: %v", err)
    return false, ""
  }
  if ok != true {
    log.Printf("Password not match")
    return false, ""
  }
  return true, userinfo.UserType
}

// for register
// usertype, error
func Register(ctx context.Context, req *pbuser.Request, con boil.ContextExecutor) (string, error) {
  username := req.Username
  // already exists or not
  ok, err := db.Userinfos(db.UserinfoWhere.Username.EQ(username)).Exists(ctx, con)
  if err != nil {
    log.Printf("Register: DB error %v", err)
    return "", errors.New("Internal error")
  }
  if ok {
    log.Printf("Already Exists name")
    return "", errors.New("This Username is already exists.")
  }
  // make hash
  password_hash,err := GenerateFromPassword(req.Password)
  if err != nil {
    log.Printf("Register: Password_Hash generate error %v", err)
    return "", errors.New("Internal Error")
  }

  // register
  userinfo := db.Userinfo{ ID: 1, Username: req.Username, PasswordHash: password_hash, UserType: "player" }
  err = userinfo.Insert(ctx, con, boil.Blacklist("id"))
  if err != nil {
    log.Printf("Register error: %v", err.Error())
    return "", err
  }
  return "player", nil
}
