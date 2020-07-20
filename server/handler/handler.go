package handler

import (
  "context"
  _ "fmt"
  pbuser "github.com/uta8a/suburi/server/proto/user"
)

type app struct {}

const (
  GetTokenSuccessMessage = "This is your token."
  GetTokenErrorMessage = ""
  LoginSuccessMessage = "Login Success"
  RegisterSuccessMessage = "Register Success"
  LogoutSuccessMessage = "Logout Success"
)

func NewApp() *app {
  return &app{}
}

func (s *app) GetToken(ctx context.Context, req *pbuser.Request) (*pbuser.Result, error) {
  // return sample token for debug
  return &pbuser.Result{ Token: "xxx", DisplayMessage: GetTokenSuccessMessage }, nil
}
func (s *app) Login(ctx context.Context, req *pbuser.Request) (*pbuser.Result, error) {
  // DB? or Auth Server
  return &pbuser.Result{ Token: "xxx", DisplayMessage: LoginSuccessMessage }, nil
}
func (s *app) Register(ctx context.Context, req *pbuser.Request) (*pbuser.Result, error) {
  // nearly equal Login
  return &pbuser.Result{ Token: "xxx", DisplayMessage: RegisterSuccessMessage }, nil
}
func (s *app) Logout(ctx context.Context, req *pbuser.Request) (*pbuser.Result, error) {
  // delete token
  return &pbuser.Result{ Token: "xxx", DisplayMessage: LogoutSuccessMessage }, nil
}


