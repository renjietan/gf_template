package interceptor

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type LogInfo struct {
	Method      string      `json:"method"`      // 请求方法
	Path        string      `json:"path"`        // 请求路径
	URL         string      `json:"url"`         // 完整URL
	QueryParams interface{} `json:"queryParams"` // 查询参数
	BodyParams  interface{} `json:"bodyParams"`  // 请求体参数
	RequestTime string      `json:"requestTime"` // 请求时间
	Response    interface{} `json:"response"`    // 响应结果
	Status      int         `json:"status"`      // 响应状态码
	Duration    int64       `json:"duration"`    // 处理时长(毫秒)
	Error       string      `json:"error"`       // 错误信息
}

func (log *LogInfo) cPrint(ctx *context.Context) {
	g.Log().Debugf(*ctx, "%v(%v): %v", log.RequestTime, log.Method, log.URL)
	PrintAsJSON(log)
	g.Log().Debugf(*ctx, "%v(%v): %v", log.RequestTime, log.Method, log.URL)
}
