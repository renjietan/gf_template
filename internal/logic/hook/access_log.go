package hook

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"gf_template/internal/library/contexts"
)

// 忽略的请求方式
var ignoredRequestMethods = []string{"HEAD", "PRI"}

// accessLog 访问日志
func (s *sHook) accessLog(r *ghttp.Request) {
	fmt.Println("Method:", r.Method)
	if s.isIgnoredRequest(r) {
		return
	}
	var ctx = r.Context()
	var ctx1 = contexts.Get(ctx)
	if ctx1 == nil {
		return
	}
	// 添加 额外数据 到上下文中
	contexts.SetDataMap(ctx, g.Map{
		"request.takeUpTime": gtime.Now().Sub(gtime.New(r.EnterTime)).Milliseconds(),
	})
	var ctx11 = contexts.Get(ctx)
	fmt.Println("accessLog ctx=========================", ctx11)
	// simple.SafeGo(ctx, func(ctx context.Context) {
	// 	if err := service.SysLog().AutoLog(ctx); err != nil {
	// 		g.Log().Infof(ctx, "hook accessLog err:%+v", err)
	// 	}
	// })
}

// isIgnoredRequest 是否忽略请求
func (s *sHook) isIgnoredRequest(r *ghttp.Request) bool {
	if r.IsFileRequest() {
		return true
	}

	if gstr.InArray(ignoredRequestMethods, strings.ToUpper(r.Method)) {
		return true
	}
	return false
}
