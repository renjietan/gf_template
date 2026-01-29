package global

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

// LoggingServeLogHandler 服务日志处理
// 需要将异常日志保存到服务日志时可以通过SetHandlers设置此方法
func LoggingServeLogHandler(ctx context.Context, in *glog.HandlerInput) {
	in.Next(ctx)

	err := g.Try(ctx, func(ctx context.Context) {
		var err error
		defer func() {
			if err != nil {
				panic(err)
			}
		}()

		// web服务日志不做记录，因为会导致重复记录
		r := g.RequestFromCtx(ctx)
		if r != nil && r.Server != nil && in.Logger.GetConfig().Path == r.Server.Logger().GetConfig().Path {
			return
		}

		conf, err := service.SysConfig().GetLoadServeLog(ctx)
		if err != nil {
			return
		}

		if conf == nil {
			return
		}

		if !conf.Switch {
			return
		}

		if in.LevelFormat == "" || !gstr.InArray(conf.LevelFormat, in.LevelFormat) {
			return
		}

		if in.Stack == "" {
			in.Stack = in.Logger.GetStack()
		}

		if len(in.Content) == 0 {
			in.Content = gstr.StrLimit(gvar.New(in.Values).String(), consts.MaxServeLogContentLen)
		}

		var data entity.SysServeLog
		data.TraceId = gctx.CtxId(ctx)
		data.LevelFormat = in.LevelFormat
		data.Content = in.Content
		data.Stack = gjson.New(charset.ParseStack(in.Stack))
		data.Line = strings.TrimRight(in.CallerPath, ":")
		data.TriggerNs = in.Time.UnixNano()
		data.Status = consts.StatusEnabled

		if gstr.Contains(in.Content, `exception recovered`) {
			data.LevelFormat = "PANI"
		}

		if data.Stack.IsNil() {
			data.Stack = gjson.New(consts.NilJsonToString)
		}

		if conf.Queue {
			err = queue.Push(consts.QueueServeLogTopic, data)
		} else {
			err = service.SysServeLog().RealWrite(ctx, data)
		}
	})

	if err != nil {
		g.Dump("日志管理器报错(LoggingServeLogHandler):", err)
	}
}
