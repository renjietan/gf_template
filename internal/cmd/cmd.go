package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf_template/internal/consts"
	"gf_template/internal/controller/user"
	logger "gf_template/middlewave/Logger"
	"gf_template/middlewave/interceptor"
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
			// 4、自定义 拦截器 中间件
			s.Use(interceptor.Init)
			// 5、自定义 日志 中间件
			// s.Use(logger.Init(nil)) // 使用默认配置
			s.Use(logger.Init(&logger.LogConfig{ // 自定义 日志中间件 配置
				LogPath:       "./logs",
				MaxAge:        30,               // 保留30天
				RotationSize:  50 * 1024 * 1024, // 50MB切割
				RotationCount: 30,
			}))
			// 5、默认的中间件处理程序响应对象及其错误。
			s.Use(ghttp.MiddlewareHandlerResponse)
			//
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(
					user.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
