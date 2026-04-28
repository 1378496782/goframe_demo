package main

import (
	_ "my-mvc-demo-chat/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"my-mvc-demo-chat/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
