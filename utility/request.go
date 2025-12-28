package utility

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareRequest(r *ghttp.Request) {
	Header := r.Response.Request.Header
	g.Log().Debug(r.Context(), "token", Header.Get("Token"))
	r.Middleware.Next()
}
