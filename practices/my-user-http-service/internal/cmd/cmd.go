package cmd

import (
	"context"

	"my-user-http-service/internal/controller/user"
	"my-user-http-service/internal/service/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func:  mainFunc,
	}
)

func mainFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	var (
		s             = g.Server()
		middlewareSvc = middleware.New()
	)
	s.Use(ghttp.MiddlewareHandlerResponse)
	s.Group("/", func(group *ghttp.RouterGroup) {
		// Group middlewares.
		group.Middleware(
			middlewareSvc.Ctx,
			ghttp.MiddlewareCORS,
		)
		// Register route handlers.
		var (
			userCtrl = user.NewV1()
		)
		group.Bind(
			userCtrl,
		)
	})
	s.Run()
	return nil
}
