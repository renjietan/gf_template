package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"

	"gf_template/internal/consts"
	"gf_template/internal/controller/user"
	"gf_template/middlewave/interceptor"
	"gf_template/middlewave/reponse"
)

var (
	Main = &gcmd.Command{
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
			interceptor.Init()           // 先初始化日志管理器
			defer interceptor.Shutdown() // 程序退出时，确保关闭日志管理器
			s.Use(interceptor.Logger)    // 日志中间件

			// 5、处理程序响应对象及其错误。
			s.Use(reponse.MiddlewareHandlerResponse)

			// 6、设置时区
			timeZone := g.Cfg().MustGet(ctx, "system.timeZone").String()
			if err := gtime.SetTimeZone(timeZone); err != nil {
				g.Log().Fatalf(ctx, "时区设置异常 err: %+v", err)
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(
					user.NewV1(),
				)
			})
			s.Run()
			// 优雅关闭日志管理器
			setupGracefulShutdown(s)
			return nil
		},
	}
	Help = &gcmd.Command{
		Name:  "help",
		Brief: "查看帮助",
		Description: `
		命令提示符
		---------------------------------------------------------------------------------
		启动服务
		>> 所有服务  [go run main.go]   热编译  [gf run main.go]
		>> HTTP服务  [go run main.go http]
		---------------------------------------------------------------------------------
    `,
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

func init() {
	if err := Main.AddCommand(Help); err != nil {
		panic(err)
	}
}
