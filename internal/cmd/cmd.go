package cmd

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/controller/admin_user"
	"github.com/iimeta/fastapi-admin/internal/controller/app"
	"github.com/iimeta/fastapi-admin/internal/controller/auth"
	"github.com/iimeta/fastapi-admin/internal/controller/chat"
	"github.com/iimeta/fastapi-admin/internal/controller/common"
	"github.com/iimeta/fastapi-admin/internal/controller/corp"
	"github.com/iimeta/fastapi-admin/internal/controller/dashboard"
	"github.com/iimeta/fastapi-admin/internal/controller/key"
	"github.com/iimeta/fastapi-admin/internal/controller/model"
	"github.com/iimeta/fastapi-admin/internal/controller/model_agent"
	"github.com/iimeta/fastapi-admin/internal/controller/sys_admin"
	"github.com/iimeta/fastapi-admin/internal/controller/user"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"net/http"
	"strings"

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

			s.SetServerRoot("./resource/fastapi-web/")

			s.AddStaticPath("/login", "./resource/fastapi-web/")
			s.AddStaticPath("/admin", "./resource/fastapi-web/")
			s.AddStaticPath("/dashboard/workplace", "./resource/fastapi-web/")
			s.AddStaticPath("/my/model", "./resource/fastapi-web/")
			s.AddStaticPath("/app/list", "./resource/fastapi-web/")
			s.AddStaticPath("/app/create", "./resource/fastapi-web/")
			s.AddStaticPath("/app/update", "./resource/fastapi-web/")
			s.AddStaticPath("/app/detail", "./resource/fastapi-web/")
			s.AddStaticPath("/app/key", "./resource/fastapi-web/")
			s.AddStaticPath("/model/list", "./resource/fastapi-web/")
			s.AddStaticPath("/model/create", "./resource/fastapi-web/")
			s.AddStaticPath("/model/update", "./resource/fastapi-web/")
			s.AddStaticPath("/model/detail", "./resource/fastapi-web/")
			s.AddStaticPath("/key/list", "./resource/fastapi-web/")
			s.AddStaticPath("/key/create", "./resource/fastapi-web/")
			s.AddStaticPath("/key/update", "./resource/fastapi-web/")
			s.AddStaticPath("/key/detail", "./resource/fastapi-web/")
			s.AddStaticPath("/agent/list", "./resource/fastapi-web/")
			s.AddStaticPath("/agent/create", "./resource/fastapi-web/")
			s.AddStaticPath("/agent/update", "./resource/fastapi-web/")
			s.AddStaticPath("/agent/detail", "./resource/fastapi-web/")
			s.AddStaticPath("/user/list", "./resource/fastapi-web/")
			s.AddStaticPath("/user/create", "./resource/fastapi-web/")
			s.AddStaticPath("/user/update", "./resource/fastapi-web/")
			s.AddStaticPath("/user/detail", "./resource/fastapi-web/")
			s.AddStaticPath("/user/center", "./resource/fastapi-web/")
			s.AddStaticPath("/chat/list", "./resource/fastapi-web/")
			s.AddStaticPath("/corp/list", "./resource/fastapi-web/")
			s.AddStaticPath("/corp/create", "./resource/fastapi-web/")
			s.AddStaticPath("/corp/update", "./resource/fastapi-web/")
			s.AddStaticPath("/corp/detail", "./resource/fastapi-web/")

			s.AddStaticPath("/public", "./resource/public")

			s.Group("/", func(g *ghttp.RouterGroup) {
				g.Middleware(ghttp.MiddlewareHandlerResponse)
				g.Middleware(middleware)
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
					g.Middleware(middleware)
					g.Bind(
						user.NewV1(),
					)
				})

				v1.Group("/admin/user", func(g *ghttp.RouterGroup) {
					g.Middleware(sysMiddleware)
					g.Bind(
						admin_user.NewV1(),
					)
				})

				v1.Group("/app", func(g *ghttp.RouterGroup) {
					g.Middleware(middleware)
					g.Bind(
						app.NewV1(),
					)
				})

				v1.Group("/model", func(g *ghttp.RouterGroup) {
					g.Middleware(middleware)
					g.Bind(
						model.NewV1(),
					)
				})

				v1.Group("/model/agent", func(g *ghttp.RouterGroup) {
					g.Middleware(middleware)
					g.Bind(
						model_agent.NewV1(),
					)
				})

				v1.Group("/key", func(g *ghttp.RouterGroup) {
					g.Middleware(middleware)
					g.Bind(
						key.NewV1(),
					)
				})

				v1.Group("/dashboard", func(g *ghttp.RouterGroup) {
					g.Middleware(middleware)
					g.Bind(
						dashboard.NewV1(),
					)
				})

				v1.Group("/chat", func(g *ghttp.RouterGroup) {
					g.Middleware(middleware)
					g.Bind(
						chat.NewV1(),
					)
				})

				v1.Group("/corp", func(g *ghttp.RouterGroup) {
					g.Middleware(middleware)
					g.Bind(
						corp.NewV1(),
					)
				})
			})

			s.Group("/api/v1/sys", func(v1 *ghttp.RouterGroup) {

				v1.Middleware(ghttp.MiddlewareHandlerResponse)
				v1.Middleware(sysMiddleware)

				v1.Group("/admin", func(g *ghttp.RouterGroup) {
					g.Bind(
						sys_admin.NewV1(),
					)
				})
			})

			if config.Cfg.AdminServerAddress != "" {
				s.SetAddr(config.Cfg.AdminServerAddress)
			}

			s.Run()
			return nil
		},
	}
)

func beforeServeHook(r *ghttp.Request) {
	logger.Debugf(r.GetCtx(), "beforeServeHook [isFile: %t] URI: %s", r.IsFileRequest(), r.RequestURI)
	r.Response.CORSDefault()
}

func middleware(r *ghttp.Request) {

	token := r.GetHeader("Authorization")
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

	if token == "" {
		token = r.Get("token").String()
	}

	if token == "" {
		r.Response.Header().Set("Content-Type", "application/json")
		r.Response.WriteStatus(http.StatusUnauthorized, g.Map{"code": 401, "message": "Unauthorized"})
		r.Exit()
		return
	}

	if gstr.HasPrefix(token, consts.USER_TOKEN_PREFIX) {

		user, err := service.Auth().GetUserByToken(r.GetCtx(), token)
		if err != nil {
			r.Response.Header().Set("Content-Type", "application/json")
			r.Response.WriteStatus(http.StatusUnauthorized, g.Map{"code": 401, "message": "Unauthorized"})
			r.Exit()
			return
		}

		if err = service.Session().SaveUser(r.GetCtx(), token, user); err != nil {
			r.Response.Header().Set("Content-Type", "application/json")
			r.Response.WriteStatus(http.StatusUnauthorized, g.Map{"code": 401, "message": "Unauthorized"})
			r.Exit()
			return
		}

	} else {

		admin, err := service.Auth().GetAdminByToken(r.GetCtx(), token)
		if err != nil {
			r.Response.Header().Set("Content-Type", "application/json")
			r.Response.WriteStatus(http.StatusUnauthorized, g.Map{"code": 401, "message": "Unauthorized"})
			r.Exit()
			return
		}

		if err = service.Session().SaveAdmin(r.GetCtx(), token, admin); err != nil {
			r.Response.Header().Set("Content-Type", "application/json")
			r.Response.WriteStatus(http.StatusUnauthorized, g.Map{"code": 401, "message": "Unauthorized"})
			r.Exit()
			return
		}
	}

	if config.Cfg.Debug {
		if gstr.HasPrefix(r.GetHeader("Content-Type"), "application/json") {
			logger.Debugf(r.GetCtx(), "url: %s, request body: %s", r.GetUrl(), r.GetBodyString())
		} else {
			logger.Debugf(r.GetCtx(), "url: %s, Content-Type: %s", r.GetUrl(), r.GetHeader("Content-Type"))
		}
	}

	r.Middleware.Next()
}

func sysMiddleware(r *ghttp.Request) {

	token := r.GetHeader("Authorization")
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

	if token == "" {
		token = r.Get("token").String()
	}

	if token == "" {
		r.Response.Header().Set("Content-Type", "application/json")
		r.Response.WriteStatus(http.StatusUnauthorized, g.Map{"code": 401, "message": "Unauthorized"})
		r.Exit()
		return
	}

	admin, err := service.Auth().GetAdminByToken(r.GetCtx(), token)
	if err != nil {
		r.Response.Header().Set("Content-Type", "application/json")
		r.Response.WriteStatus(http.StatusUnauthorized, g.Map{"code": 401, "message": "Unauthorized"})
		r.Exit()
		return
	}

	if err = service.Session().SaveAdmin(r.GetCtx(), token, admin); err != nil {
		r.Response.Header().Set("Content-Type", "application/json")
		r.Response.WriteStatus(http.StatusUnauthorized, g.Map{"code": 401, "message": "Unauthorized"})
		r.Exit()
		return
	}

	if gstr.HasPrefix(r.GetHeader("Content-Type"), "application/json") {
		logger.Debugf(r.GetCtx(), "url: %s, request body: %s", r.GetUrl(), r.GetBodyString())
	} else {
		logger.Debugf(r.GetCtx(), "url: %s, Content-Type: %s", r.GetUrl(), r.GetHeader("Content-Type"))
	}

	r.Middleware.Next()
}
