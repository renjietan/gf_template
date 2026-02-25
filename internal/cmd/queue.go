package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"

	"gf_template/internal/global"
	"gf_template/internal/library/queue"
	_ "gf_template/internal/queues"
	"gf_template/utility/simple"
)

var (
	Queue = &gcmd.Command{
		Name:        "queue",
		Brief:       "消息队列",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			queue.Init()
			// 服务日志处理
			queue.Logger().SetHandlers(global.LoggingServeLogHandler)

			simple.SafeGo(ctx, func(ctx context.Context) {
				queue.Logger().Debug(ctx, "start queue consumer..")
				queue.StartConsumersListener(ctx)
				queue.Logger().Debug(ctx, "start queue consumer success..")
			})

			serverWg.Add(1)

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			<-serverCloseSignal
			queue.Logger().Debug(ctx, "queue successfully closed ..")
			serverWg.Done()
			return
		},
	}
)
