package main

import (
	_ "zfw_proxima/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"zfw_proxima/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
