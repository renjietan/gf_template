package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf_template/internal/consts"
	"gf_template/internal/controller/user"
	"gf_template/middlewave/reponse"
)

var Http = &gcmd.Command{
	Name:  "http",
	Usage: "http",
	Brief: "HTTP服务，也可以称为主服务，包含http、websocket、tcpserver多个可对外服务",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		s := g.Server()
		// 1、自定义 swagger 模板
		s.SetSwaggerUITemplate(consts.SwaggerTpl)
		var db_conf = g.Cfg().MustGet(ctx, "database.default").Map()
		// 2、执行SQL时，打印SQL语句
		g.DB().SetDebug(db_conf["debug"].(bool))

		// 3、开启 国际化 & 设置 默认语言
		g.I18n().SetLanguage("zh-CN")

		// 5、处理程序响应对象及其错误。
		s.Use(reponse.MiddlewareHandlerResponse)

		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Bind(
				user.NewV1(),
			)
		})
		serverWg.Add(1)

		// 信号监听
		signalListen(ctx, signalHandlerForOverall)

		go func() {
			<-serverCloseSignal
			_ = s.Shutdown() // 关闭http服务，主服务建议放在最后一个关闭
			g.Log().Debug(ctx, "\n收到关闭信号, 关闭HTTP服务")
			serverWg.Done()
		}()
		s.Run()
		return nil
	},
}
