package main

import (
	"rpc_demo/controller"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

func main() {
	s := grpcx.Server.New()
	controller.Register(s)
	s.Run()
}
