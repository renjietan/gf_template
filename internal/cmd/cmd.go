package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf_template/internal/consts"
	"gf_template/internal/controller/user"
	"gf_template/utility"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 1、自定义 swagger 模板
			s.SetSwaggerUITemplate(consts.SwaggerTpl)
			// 2、执行SQL时，打印SQL语句
			g.DB().SetDebug(true)
			// 3、开启 国际化 & 设置 默认语言
			g.I18n().SetLanguage("zh-CN")
			// 4、拦截器
			s.Use(utility.MiddlewareRequest)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					user.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
