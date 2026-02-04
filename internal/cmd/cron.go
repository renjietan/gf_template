package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"

	"gf_template/internal/global"
	"gf_template/internal/library/cron"
)

var Cron = &gcmd.Command{
	Name:        "cron",
	Brief:       "定时任务",
	Description: ``,
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		// 服务日志处理
		cron.Logger().SetHandlers(global.LoggingServeLogHandler)

		// 启动定时任务
		// service.SysCron().StartCron(ctx)
		serverWg.Add(1)

		// 信号监听
		signalListen(ctx, signalHandlerForOverall)

		// 收到 关闭信号，停止定时任务
		<-serverCloseSignal
		cron.StopALL()
		cron.Logger().Debug(ctx, "\n收到 关闭信号, 定时任务已关闭")
		serverWg.Done()
		return
	},
}
