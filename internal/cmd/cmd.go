package cmd

import (
	"PersonalMain/internal/controller/auth"
	"PersonalMain/internal/middleware"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"PersonalMain/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.TimestampMiddleware)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(ghttp.MiddlewareHandlerResponse)
					group.Bind(
						hello.NewV1(),
					)
				})
				group.Group("/auth", func(group *ghttp.RouterGroup) {
					group.Middleware(ghttp.MiddlewareHandlerResponse)
					group.Bind(
						auth.NewAuthV1(),
					)
				})
			})
			s.Run()
			s.SetServerRoot("public")
			return nil
		},
	}
)
