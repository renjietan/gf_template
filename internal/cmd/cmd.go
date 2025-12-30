package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf_template/internal/consts"
	"gf_template/internal/controller/user"
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
			// s.Use(interceptor.Init)
			interceptor.Init() // 先初始化日志管理器
			defer func() {
				interceptor.Shutdown()
			}() // 程序退出时，确保关闭日志管理器
			s.Use(interceptor.Logger)

			// 5、默认的中间件处理程序响应对象及其错误。
			s.Use(ghttp.MiddlewareHandlerResponse)
			//
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(
					user.NewV1(),
				)
			})
			s.Run()
			// 优雅关机
			setupGracefulShutdown(s)
			return nil
		},
	}
)

func setupGracefulShutdown(s *ghttp.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		g.Log().Info(context.Background(), "关闭日志处理器...")

		// 关闭日志处理器
		interceptor.Shutdown()
		s.Shutdown()
	}()
}
