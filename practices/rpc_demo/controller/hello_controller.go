package controller

import (
	"context"
	"rpc_demo/protobuf"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type HelloController struct {
	protobuf.UnimplementedGreeterServer
}

func Register(s *grpcx.GrpcServer) {
	protobuf.RegisterGreeterServer(s.Server, &HelloController{})
}

func (c *HelloController) SayHello(ctx context.Context, req *protobuf.HelloRequest) (*protobuf.HelloReply, error) {
	return &protobuf.HelloReply{
		Message: "Hello, " + req.GetName(),
	}, nil
}
