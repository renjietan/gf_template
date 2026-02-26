package global

import (
	"context"
	"fmt"
	"gf_template/internal/consts"
	"gf_template/internal/library/queue"
	"gf_template/internal/queues"
	charset "gf_template/utility/chatset"
	"runtime"
	"strings"

	"gf_template/internal/library/cache"
	sysconfig "gf_template/utility/config"
	"gf_template/utility/validate"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"
)

func Init(ctx context.Context) {
	// 默认使用开发环境配置文件
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("development")

	// 设置数据库调试模式
	var db_conf = g.Cfg().MustGet(ctx, "database.default").Map()
	g.DB().SetDebug(db_conf["debug"].(bool))

	// 设置默认语言
	g.I18n().SetLanguage(sysconfig.GetLanguage(ctx))

	// 根据运行模式，重置默认的开发环境配置文件
	mode := sysconfig.GetMode(ctx)
	fmt.Printf("当前运行环境: %v, 当前运行模式: %v 运行根路径为: %v  gf版本: %v \n", runtime.GOOS, mode, gfile.Pwd(), gf.VERSION)
	if mode != gmode.DEVELOP && mode != gmode.NOT_SET {
		g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("production")
	}

	// 设置gf运行模式
	SetGFMode(ctx)

	// 设置服务日志处理
	glog.SetDefaultHandler(LoggingServeLogHandler)

	// 默认上海时区
	timezone := sysconfig.GetTimeZone(ctx)
	if err := gtime.SetTimeZone(timezone); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常 err: %+v", err)
		return
	}

	// 初始化，并设置缓存适配器
	cache.SetAdapter(ctx)
}

// LoggingServeLogHandler 服务日志处理
// 需要将异常日志保存到服务日志时可以通过SetHandlers设置此方法
func LoggingServeLogHandler(ctx context.Context, in *glog.HandlerInput) {
	in.Next(ctx)

	//err := g.Try(ctx, func(ctx context.Context) {
	//	var err error
	//	defer func() {
	//		if err != nil {
	//			panic(err)
	//		}
	//	}()
	//})

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

		conf, err := sysconfig.GetServeLog(ctx)
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

		var data queues.SysServeLog
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
			// TODO: 如果未配置消息队列， 这里需要做些什么？
		}
	})

	if err != nil {
		g.Dump("日志管理器报错(LoggingServeLogHandler):", err)
	}
}

func SetGFMode(ctx context.Context) {
	mode := sysconfig.GetMode(ctx)

	var modes = []string{gmode.DEVELOP, gmode.TESTING, gmode.STAGING, gmode.PRODUCT}

	// 如果是有效的运行模式，就进行设置
	if validate.InSlice(modes, mode) {
		gmode.Set(mode)
	}
}
