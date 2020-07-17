package handler

import (
	"context"
	"fmt"
	pb "suburi/server/helloworld"
)

type app struct{}

func NewApp() *app {
	return &app{}
}

func (s *app) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *app) TellMeSecret(
	ctx context.Context,
	req *pb.TellMeSecretRequest,
) (*pb.TellMeSecretResponse, error) {
	return &pb.TellMeSecretResponse{Answer: fmt.Sprintf("Secret: %s", req.Message)}, nil
}
