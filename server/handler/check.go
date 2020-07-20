package handler

import (
  "context"
  
  pbcheck "github.com/uta8a/suburi/server/proto/check"
)


func (s *App) HealthCheck(ctx context.Context, req *pbcheck.Request) (*pbcheck.Response, error) {
  return &pbcheck.Response{ DisPlayMessage: "health check is ok!" }, nil
}

func (s *App) SecretCheck(ctx context.Context, req *pbcheck.Request) (*pbcheck.Response, error) {
  return &pbcheck.Response{ DisplayMessage: "This is secret only for Admin." }, nil
}

