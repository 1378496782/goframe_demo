package main

import (
	_ "zfw_proxima/app/gateway/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"zfw_proxima/app/gateway/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
