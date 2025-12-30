package interceptor

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// Logger 全局日志中间件
func Init(r *ghttp.Request) {
	startTime := time.Now()

	// 创建日志信息结构
	logInfo := &LogInfo{
		Method:      r.Method,
		Path:        r.URL.Path,
		URL:         r.Request.URL.String(),
		QueryParams: r.URL.RawQuery,
		RequestTime: startTime.Format("2006-01-02 15:04:05"),
		Status:      r.Response.Status,
	}
	// 根据请求方法获取请求参数
	switch r.Method {
	case "GET", "DELETE":
		if len(r.GetMap()) > 0 {
			logInfo.QueryParams = r.GetMap()
		}
	case "POST", "PUT", "PATCH":
		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/x-www-form-urlencoded", "multipart/form-data":
			// Form 表单数据
			if formData := r.GetFormMap(); len(formData) > 0 {
				logInfo.BodyParams = formData
			}
		default:
			if bodyBytes := r.GetBody(); len(bodyBytes) > 0 {
				var bodyJson = map[string]any{}
				err := json.Unmarshal(bodyBytes, &bodyJson)
				if err == nil {
					logInfo.BodyParams = bodyJson
				} else {
					logInfo.BodyParams = bodyBytes
				}
			}
		}
	}

	// 保存日志信息到上下文，方便后续使用
	ctx := context.WithValue(r.Context(), "requestLogInfo", logInfo)
	r.SetCtx(ctx)

	r.Middleware.Next()

	// 获取响应结果
	logInfo.Status = r.Response.Status
	logInfo.Duration = time.Since(startTime).Milliseconds()

	// 获取响应数据
	if r.Response.BufferLength() > 0 {
		// 如果有缓冲区数据
		responseBytes := r.Response.Buffer()
		var responseData interface{}

		// 尝试解析 JSON 响应
		if err := gconv.Scan(responseBytes, &responseData); err == nil {
			logInfo.Response = responseData
		} else {
			// 如果不是 JSON，存储为字符串
			logInfo.Response = gconv.String(responseBytes)
		}
	} else if r.Response.Writer != nil {
		// 对于流式响应，可以记录状态信息
		logInfo.Response = map[string]interface{}{
			"type":   "stream_response",
			"status": r.Response.Status,
		}
	}

	// 如果发生错误，记录错误信息
	if err := r.GetError(); err != nil {
		logInfo.Error = err.Error()
	}

	// 根据需要将最终日志信息存入上下文
	ctx = context.WithValue(r.Context(), "responseLogInfo", logInfo)
	r.SetCtx(ctx)
	logInfo.cPrint(&ctx)
}

// GetLogInfoFromCtx 从上下文中获取日志信息
func GetLogInfoFromCtx(ctx context.Context) *LogInfo {
	if value := ctx.Value("responseLogInfo"); value != nil {
		if info, ok := value.(*LogInfo); ok {
			return info
		}
	}
	return nil
}

// GetRequestInfoFromCtx 从上下文中获取请求信息（在 Next 之前调用）
func GetRequestInfoFromCtx(ctx context.Context) *LogInfo {
	if value := ctx.Value("requestLogInfo"); value != nil {
		if info, ok := value.(*LogInfo); ok {
			return info
		}
	}
	return nil
}
