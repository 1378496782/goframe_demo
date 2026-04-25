package main

import (
	"rpc_demo/protobuf"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx    = gctx.New()
		conn   = grpcx.Client.MustNewGrpcClientConn("rpc_demo")
		client = protobuf.NewGreeterClient(conn)
	)

	res, err := client.SayHello(ctx, &protobuf.HelloRequest{
		Name: "zfw111",
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Debug(ctx, "Response:", res.Message)
}
