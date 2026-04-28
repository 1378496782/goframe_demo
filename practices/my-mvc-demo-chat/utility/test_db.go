package main

import (
	"context"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	ctx := context.Background()

	g.Log().Info(ctx, "Testing database connection...")

	// 测试连接
	err := g.DB().PingMaster()
	if err != nil {
		g.Log().Fatalf(ctx, "Database connection failed: %v", err)
	}
	g.Log().Info(ctx, "✅ Database connected successfully!")

	// 用简单的方式查询
	result, err := g.DB().GetAll(ctx, "SHOW TABLES")
	if err != nil {
		g.Log().Warningf(ctx, "Failed to query tables: %v", err)
	} else {
		g.Log().Infof(ctx, "Query result: %v", result)
		if len(result) == 0 {
			g.Log().Warning(ctx, "⚠️ No tables found. You need to create the user table first!")
		}
	}
}
