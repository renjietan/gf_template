package middleware

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// LogInfo 日志信息结构体
type LogInfo struct {
	Method      string      `json:"method"`      // 请求方法
	Path        string      `json:"path"`        // 请求路径
	URL         string      `json:"url"`         // 完整URL
	QueryParams string      `json:"queryParams"` // 查询参数
	BodyParams  interface{} `json:"bodyParams"`  // 请求体参数
	RequestTime time.Time   `json:"requestTime"` // 请求时间
	Response    interface{} `json:"response"`    // 响应结果
	Status      int         `json:"status"`      // 响应状态码
	Duration    int64       `json:"duration"`    // 处理时长(毫秒)
	Error       string      `json:"error"`       // 错误信息
}

func (log *LogInfo) cPrint(ctx *context.Context) {
	g.Log().Debugf(*ctx, "==================== %v ========================", log.URL)
	fmt.Printf("Method: %v\n", log.Method)
	fmt.Printf("Path: %v\n", log.Path)
	fmt.Printf("Uri: %v\n", log.URL)
	fmt.Printf("Status: %v\n", log.Status)
	fmt.Printf("Duration: %vms\n", log.Duration)
	fmt.Printf("QueryParams: %v\n", log.QueryParams)
	fmt.Printf("BodyParams: %v\n", log.BodyParams)
	fmt.Printf("RequestTime: %v\n", log.RequestTime)
	print_reflect_map(log.Response)
	// fmt.Printf("Response: %v\n", log.Response)
	fmt.Printf("Error: %v\n", log.Error)
	g.Log().Debugf(*ctx, "==================== %v ========================", log.URL)
}

func print_reflect_map(x interface{}) {
	t := reflect.TypeOf(x)
	kind := t.Kind()
	switch kind {
	case reflect.Map:
		for key, value := range x.(g.Map) {
			t2 := reflect.TypeOf(value)
			kind2 := t2.Kind()
			fmt.Println("key:", key)
			fmt.Println("kind2:", kind2)
			fmt.Println("value:", value)
			switch kind2 {
			case reflect.Map:
				print_reflect_map(value)
			case reflect.Slice:
				print_reflect_map(value)
			default:
				fmt.Printf("Response(%v): %v\n", key, value)
			}
		}
	case reflect.Slice:

	default:
		fmt.Printf("Response: %v\n", x)
	}
}

// Logger 全局日志中间件
func Logger(r *ghttp.Request) {
	startTime := time.Now()

	// 创建日志信息结构
	logInfo := &LogInfo{
		Method:      r.Method,
		Path:        r.URL.Path,
		URL:         r.Request.URL.String(),
		QueryParams: r.URL.RawQuery,
		RequestTime: startTime,
		Status:      r.Response.Status,
	}

	// 根据请求方法获取请求参数
	switch r.Method {
	case "GET", "DELETE":
		// GET 和 DELETE 请求通常使用 Query 参数
		if len(r.GetMap()) > 0 {
			logInfo.BodyParams = r.GetMap()
		}
	case "POST", "PUT", "PATCH":
		// POST、PUT、PATCH 请求获取 Body 参数
		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			var jsonBody interface{}
			if err := r.Parse(&jsonBody); err == nil && jsonBody != nil {
				logInfo.BodyParams = jsonBody
			}
		case "application/x-www-form-urlencoded", "multipart/form-data":
			// Form 表单数据
			if formData := r.GetFormMap(); len(formData) > 0 {
				logInfo.BodyParams = formData
			}
		default:
			if bodyBytes := r.GetBody(); len(bodyBytes) > 0 {
				logInfo.BodyParams = gconv.String(bodyBytes)
			}
		}
	}

	// 保存日志信息到上下文，方便后续使用
	ctx := context.WithValue(r.Context(), "requestLogInfo", logInfo)
	r.SetCtx(ctx)

	// 执行后续中间件和业务逻辑
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

	// 可以根据需要将最终日志信息存入上下文
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
