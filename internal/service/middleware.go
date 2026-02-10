// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// 初始化 ctx上下文中间件
		Ctx(r *ghttp.Request)
		// 跨域
		CORS(r *ghttp.Request)
		// 将用户信息传递到上下文中
		// TODO: 目前只是一个示例，实际项目中需要根据自己的业务逻辑来实现
		DeliverUserContext(r *ghttp.Request) (err error)
		// IsExceptAuth 是否是不需要验证权限的路由地址
		IsExceptAuth(ctx context.Context, appName string, path string) bool
		// IsExceptLogin 是否是不需要登录的路由地址
		IsExceptLogin(ctx context.Context, appName string, path string) bool
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
