package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf_template/utility/simple"
)

var (
	Main = &gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "开启主服务，与 ALL类似",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, "参数:", gcmd.GetArgAll())
			return All.Func(ctx, parser)
		},
	}
	All = &gcmd.Command{
		Name:        "全部服务",
		Brief:       "开启全部服务",
		Description: "用于启动所有服务器的命令输入项",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, "\n======================= 所有服务 ==========================")

			// 需要启动的服务
			var allServers = []*gcmd.Command{Http, Cron}

			for _, server := range allServers {
				var cmd = server
				simple.SafeGo(ctx, func(ctx context.Context) {
					if err := cmd.Func(ctx, parser); err != nil {
						g.Log().Fatalf(ctx, "%v 启动失败:%v", cmd.Name, err)
					}
				})
			}

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			<-serverCloseSignal
			serverWg.Wait()
			return
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
		>> 消息队列  [go run main.go queue]
		>> 定时任务  [go run main.go cron]
		>> 查看帮助  [go run main.go help]
		---------------------------------------------------------------------------------
    `,
	}
)

func init() {
	if err := Main.AddCommand(Help, Http, Cron); err != nil {
		panic(err)
	}
}
