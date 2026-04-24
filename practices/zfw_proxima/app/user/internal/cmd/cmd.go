package cmd

import (
	"context"

	"zfw_proxima/app/user/internal/controller/account"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "user grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			account.Register(s)
			// 注册反射服务，便于 grpcurl 和 Postman 调用
			reflection.Register(s.Server)
			s.Run()
			return nil
		},
	}
)
