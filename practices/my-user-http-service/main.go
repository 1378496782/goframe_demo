package main

import (
	_ "my-user-http-service/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "my-user-http-service/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"my-user-http-service/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
