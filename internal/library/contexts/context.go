package contexts

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"gf_template/internal/consts"
	"gf_template/internal/model"
	"gf_template/internal/model/entity"
)

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改
func Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextHTTPKey, customCtx)
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func SetUser(ctx context.Context, user *entity.User) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetUser 中 c 是 nil ")
		return
	}
	c.User = user
}

// GetUser 获取用户信息
func GetUser(ctx context.Context) *entity.User {
	c := Get(ctx)
	if c == nil {
		return nil
	}
	return c.User
}

// SetResponse 设置组件响应 用于访问日志使用
func SetResponse(ctx context.Context, response *model.Response) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetResponse 中 c 是 nil ")
		return
	}
	c.Response = response
}

// GetResponse 设置组件响应 用于访问日志使用
func GetResponse(ctx context.Context) *model.Response {
	c := Get(ctx)
	if c == nil {
		return nil
	}
	return c.Response
}

// SetData 设置额外数据 只设置 单对 1 对 1
func SetData(ctx context.Context, k string, v interface{}) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetData 中 c == nil ")
		return
	}
	Get(ctx).Data[k] = v
}

// SetDataMap 设置额外数据
func SetDataMap(ctx context.Context, vs g.Map) {
	c := Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetDataMap 中 c == nil ")
		return
	}

	for k, v := range vs {
		Get(ctx).Data[k] = v
	}
}

// GetData 获取额外数据
func GetData(ctx context.Context) g.Map {
	c := Get(ctx)
	if c == nil {
		return nil
	}
	return c.Data
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextHTTPKey)
	if value == nil {
		return nil
	}
	if _ctx, ok := value.(*model.Context); ok {
		return _ctx
	}
	return nil
}
