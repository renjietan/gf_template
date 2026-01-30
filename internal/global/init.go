package global

import (
	"context"
	"fmt"
	"runtime"

	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmode"

	"gf_template/internal/consts"
	"gf_template/internal/library/cache"
	"gf_template/utility/simple"
	"gf_template/utility/validate"
)

func Init(ctx context.Context) {
	// 设置gf运行模式
	SetGFMode(ctx)

	// 设置服务日志处理
	glog.SetDefaultHandler(LoggingServeLogHandler)
	g.Cfg()
	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常 err: %+v", err)
		return
	}

	fmt.Printf("当前运行环境：%v, 运行根路径为：%v  gf版本：%v \n", runtime.GOOS, gfile.Pwd(), gf.VERSION)

	// 初始化链路追踪
	InitTrace(ctx)

	// 设置缓存适配器
	cache.SetAdapter(ctx)

	// // 初始化功能库配置
	// service.SysConfig().InitConfig(ctx)

	// // 加载超管数据
	// service.AdminMember().LoadSuperAdmin(ctx)

	// // 订阅集群同步
	// SubscribeClusterSync(ctx)
}

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
	})

	if err != nil {
		g.Dump("日志管理器报错(LoggingServeLogHandler):", err)
	}
}

// InitTrace 初始化链路追踪
func InitTrace(ctx context.Context) {
	if !g.Cfg().MustGet(ctx, "jaeger.switch").Bool() {
		return
	}

	tp, err := jaeger.Init(simple.AppName(ctx), g.Cfg().MustGet(ctx, "jaeger.endpoint").String())
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	simple.Event().Register(consts.EventServerClose, func(ctx context.Context, args ...interface{}) {
		_ = tp.Shutdown(ctx)
		g.Log().Debug(ctx, "jaeger closed ..")
	})
}

func SetGFMode(ctx context.Context) {
	mode := g.Cfg().MustGet(ctx, "system.mode").String()
	if len(mode) == 0 {
		mode = gmode.NOT_SET
	}

	var modes = []string{gmode.DEVELOP, gmode.TESTING, gmode.STAGING, gmode.PRODUCT}

	// 如果是有效的运行模式，就进行设置
	if validate.InSlice(modes, mode) {
		gmode.Set(mode)
	}
}
