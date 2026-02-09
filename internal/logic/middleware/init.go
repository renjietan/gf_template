package middleware

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"

	"gf_template/internal/library/contexts"
	"gf_template/internal/model"
	"gf_template/internal/service"
	"gf_template/utility/simple"
	"gf_template/utility/validate"
)

type sMiddleware struct {
	LoginUrl         string // 登录路由地址
	NotRecordRequest g.Map  // 不记录请求数据的路由（当前请求数据过大时会影响响应效率，可以将路径放到该选项中改善）
}

func init() {
	service.RegisterMiddleware(NewMiddleware())
}

func NewMiddleware() *sMiddleware {
	return &sMiddleware{}
}

// 初始化 ctx上下文中间件
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	r.SetCtx(gi18n.WithLanguage(r.Context(), simple.GetHeaderLocale(r.Context())))
	data := make(g.Map)
	if _, ok := s.NotRecordRequest[r.URL.Path]; ok {
		data["request.body"] = gjson.New(nil)
	} else {
		data["request.body"] = gjson.New(r.GetBodyString())
	}
	contexts.Init(r, &model.Context{
		Data: data,
	})
	if len(r.Cookie.GetSessionId()) == 0 {
		r.Cookie.SetSessionId(gctx.CtxId(r.Context()))
	}
	r.SetCtx(r.GetNeverDoneCtx())
	r.Middleware.Next()
}

// 跨域
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 将用户信息传递到上下文中
// TODO: 目前只是一个示例，实际项目中需要根据自己的业务逻辑来实现
func (s *sMiddleware) DeliverUserContext(r *ghttp.Request) (err error) {
	return
}

// IsExceptAuth 是否是不需要验证权限的路由地址
func (s *sMiddleware) IsExceptAuth(ctx context.Context, appName, path string) bool {
	pathList := g.Cfg().MustGet(ctx, fmt.Sprintf("router.%v.exceptAuth", appName)).Strings()

	for i := 0; i < len(pathList); i++ {
		if validate.InSliceExistStr(pathList[i], path) {
			return true
		}
	}
	return false
}

// IsExceptLogin 是否是不需要登录的路由地址
func (s *sMiddleware) IsExceptLogin(ctx context.Context, appName, path string) bool {
	pathList := g.Cfg().MustGet(ctx, fmt.Sprintf("router.%v.exceptLogin", appName)).Strings()

	for i := 0; i < len(pathList); i++ {
		if validate.InSliceExistStr(pathList[i], path) {
			return true
		}
	}
	return false
}
