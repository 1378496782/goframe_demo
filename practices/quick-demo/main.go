package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Write("Hello, World! zfw")
	})
	s.SetPort(8000)
	s.Run()
}
