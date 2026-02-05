package cmd

import (
	"context"
	"fmt"

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
	Brief: "HTTP",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		s := g.Server()
		// 自定义 swagger 模板
		s.SetSwaggerUITemplate(consts.SwaggerTpl)

		// 处理程序响应对象及其错误。
		s.Use(reponse.MiddlewareHandlerResponse)

		// TODO: 需要单独管理
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Bind(
				user.NewV1(),
			)
		})
		serverWg.Add(1)
		fmt.Println("======================= 开启HTTP ==========================")
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
