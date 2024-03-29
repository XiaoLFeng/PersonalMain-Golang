package cmd

import (
	authUser "PersonalMain/internal/controller/auth"
	"PersonalMain/internal/controller/auth/token"
	"PersonalMain/internal/controller/link/custom/friend"
	"PersonalMain/internal/controller/link/custom/sponsor"
	"PersonalMain/internal/controller/user"
	"PersonalMain/internal/middleware"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.JsonResponseMiddleware)
				group.Middleware(middleware.TimestampMiddleware)

				/*
				 * 用户管理
				 */
				group.Group("/auth", func(group *ghttp.RouterGroup) {
					group.Group("/user", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.VerifyTokenMiddleware)
						group.Bind(
							authUser.NewAuthV1(),
						)
					})
					group.Group("/token", func(group *ghttp.RouterGroup) {
						group.Bind(
							token.NewTokenV1(),
						)
					})
				})
				/*
				 * 博客连接管理
				 */
				group.Group("/link", func(group *ghttp.RouterGroup) {
					// 公共访问
					group.Group("/custom", func(group *ghttp.RouterGroup) {
						group.Group("/friend", func(group *ghttp.RouterGroup) {
							group.Bind(
								friend.NewLinkCustomFriendV1(),
							)
						})
						group.Group("/sponsor", func(group *ghttp.RouterGroup) {
							group.Bind(
								sponsor.NewSponsorV1().AddSponsor,
								sponsor.NewSponsorV1().GetSponsor,
								sponsor.NewSponsorV1().EditSponsor,
								sponsor.NewSponsorV1().DeleteSponsor,
							)
							group.Group("/check", func(group *ghttp.RouterGroup) {
								group.Middleware(middleware.VerifyTokenMiddleware)
								group.Bind(
									sponsor.NewSponsorV1().CheckSponsor,
									sponsor.NewSponsorV1().GetCheckSponsor,
								)
							})
						})
					})
					// 管理员访问
					//group.Group("/admin", func(group *ghttp.RouterGroup) {
					//
					//})
				})
				/*
				 * 用户管理
				 */
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.VerifyTokenMiddleware)
					group.Bind(
						user.NewUserV1(),
					)
				})
			})
			//s.SetServerRoot("resource/public")
			s.Run()
			return nil
		},
	}
)
