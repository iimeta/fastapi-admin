package cmd

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/controller/auth"
	"github.com/iimeta/fastapi-admin/internal/controller/common"
	"github.com/iimeta/fastapi-admin/internal/controller/user"
	"github.com/iimeta/fastapi-admin/utility/logger"

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

			s.BindHookHandler("/*", ghttp.HookBeforeServe, beforeServeHook)

			s.Group("/", func(g *ghttp.RouterGroup) {
				g.Middleware(ghttp.MiddlewareHandlerResponse)
				g.Middleware(MiddlewareAuth)
				g.Bind()
			})

			s.Group("/api/v1", func(v1 *ghttp.RouterGroup) {

				v1.Middleware(ghttp.MiddlewareHandlerResponse)

				v1.Group("/common", func(g *ghttp.RouterGroup) {
					g.Bind(
						common.NewV1(),
					)
				})

				v1.Group("/auth", func(g *ghttp.RouterGroup) {
					g.Bind(
						auth.NewV1(),
					)
				})

				v1.Group("/user", func(g *ghttp.RouterGroup) {
					g.Middleware(MiddlewareAuth)
					g.Bind(
						user.NewV1(),
					)
				})
			})

			s.Run()
			return nil
		},
	}
)

func beforeServeHook(r *ghttp.Request) {
	logger.Debugf(r.GetCtx(), "beforeServeHook [isFile: %v] URI: %s", r.IsFileRequest(), r.RequestURI)
	r.Response.CORSDefault()
}

func MiddlewareAuth(r *ghttp.Request) {
	//middleware.Auth(r, config.Cfg.Jwt.Secret, "api", cache.NewTokenSessionStorage(redis.Client))
}
