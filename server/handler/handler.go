package handler

import (
  "context"
  "fmt"
  "log"
  "os"
  "time"
  "database/sql"
  _ "github.com/lib/pq"

  "github.com/volatiletech/sqlboiler/v4/boil"

  _ "github.com/uta8a/suburi/server/db"
  pbuser "github.com/uta8a/suburi/server/proto/user"
  auth "github.com/uta8a/suburi/server/auth"
)

const (
  GetTokenSuccessMessage = "This is your token."
  GetTokenErrorMessage = ""
  LoginSuccessMessage = "Login Success"
  LoginErrorMessage = "Login Failed"
  RegisterSuccessMessage = "Register Success"
  LogoutSuccessMessage = "Logout Success"
)

func NewApp(con boil.ContextExecutor) *App {
  return &App{
    con: con,
  }
}

func (s *App) GetToken(ctx context.Context, req *pbuser.Request) (*pbuser.Result, error) {
  // return sample token for debug
  token, err := auth.CreateToken(req.Username)
  if err != nil {
    log.Printf("GetTokenError: cannot get token %v", err)
    return &pbuser.Result{ Token: "", DisplayMessage: GetTokenErrorMessage }, err
  }
  return &pbuser.Result{ Token: token, DisplayMessage: GetTokenSuccessMessage }, nil
}
func (s *App) Login(ctx context.Context, req *pbuser.Request) (*pbuser.Result, error) {
  // DB? or Auth Server
  con := s.con
  if auth.Verify(ctx, req, con) != true  {
    return &pbuser.Result{ Token: "", DisplayMessage: LoginErrorMessage }, nil
  }
  token, err := auth.CreateToken(req.Username)
  if err != nil {
    log.Printf("LoginError: cannot create token %v", err)
    return &pbuser.Result{ Token: "", DisplayMessage: LoginErrorMessage }, nil
  }
  return &pbuser.Result{ Token: token, DisplayMessage: LoginSuccessMessage }, nil
}
func (s *App) Register(ctx context.Context, req *pbuser.Request) (*pbuser.Result, error) {
  // nearly equal Login
  return &pbuser.Result{ Token: "xxx", DisplayMessage: RegisterSuccessMessage }, nil
}
func (s *App) Logout(ctx context.Context, req *pbuser.Request) (*pbuser.Result, error) {
  // delete token
  return &pbuser.Result{ Token: "xxx", DisplayMessage: LogoutSuccessMessage }, nil
}

func InitDB() (boil.ContextExecutor, error) {
  dbname := os.Getenv("POSTGRES_DB")
  dbuser := os.Getenv("POSTGRES_USER")
  dbpass := os.Getenv("POSTGRES_PASSWORD")
  dbport := os.Getenv("POSTGRES_PORT")
  dbhost := os.Getenv("POSTGRES_HOST")
  // TODO sslmode ON in production
  con,err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s host=%s password=%s port=%s sslmode=disable", dbuser, dbname, dbhost, dbpass, dbport))
  if err != nil {
    return nil, err
  }
  con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(10)
	con.SetConnMaxLifetime(300 * time.Second)

  boil.SetDB(con)
  boil.DebugMode = true
  return con, nil
}
