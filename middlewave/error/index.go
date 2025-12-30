package errorHandle

import "github.com/gogf/gf/v2/net/ghttp"

func Init(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.ClearBuffer()
		r.Response.Writef("%+v", err)
	}
}
