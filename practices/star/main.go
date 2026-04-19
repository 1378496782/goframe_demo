package main

import (
	"errors"
	_ "star/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"star/internal/cmd"
)

func main() {
	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")
	err := connDb()
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(gctx.GetInitCtx())
}

// connDb 检查数据库连接是否正常
func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("连接到数据库失败")
	}
	return nil
}
