package handler

import (
  "context"
  _ "fmt"
  pbuser "github.com/uta8a/suburi/server/proto/user"
)

type app struct {}

func NewApp() *app {
  return &app{}
}

func (s *app) GetToken(ctx context.Context, req *pbuser.Request) (*pbuser.Result, error) {
  return &pbuser.Result{ Token: "xxxyyyzzz" }, nil
}
