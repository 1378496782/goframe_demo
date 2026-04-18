package main

import (
	_ "quick-demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"quick-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
