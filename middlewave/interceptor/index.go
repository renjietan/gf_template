package interceptor

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

var (
	// globalProcessor 全局日志处理器单例
	globalProcessor *LogProcessor
	processorOnce   sync.Once
)

// GetLogProcessor 获取全局日志处理器（单例模式）
func GetLogProcessor() *LogProcessor {
	processorOnce.Do(func() {
		globalProcessor = &LogProcessor{
			logChan:    make(chan *LogInfo, 1000), // 缓冲通道，防止阻塞
			workerExit: make(chan bool),
		}
		// 启动协程
		globalProcessor.StartWorker()
	})
	return globalProcessor
}

// Logger 全局日志中间件
func Logger(r *ghttp.Request) {
	startTime := time.Now()

	// 生成请求追踪ID
	traceID := generateTraceID()

	// 创建日志信息结构
	logInfo := &LogInfo{
		Method:      r.Method,
		Path:        r.URL.Path,
		URL:         r.Request.URL.String(),
		QueryParams: r.URL.RawQuery,
		RequestTime: startTime.Format("2006-01-02 15:04:05"),
		Status:      r.Response.Status,
	}

	// 将traceID存入上下文，方便后续使用
	ctx := context.WithValue(r.Context(), "traceId", traceID)
	r.SetCtx(ctx)

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
		// 对于流式响应，记录状态信息
		logInfo.Response = map[string]interface{}{
			"type":        "stream_response",
			"status_code": r.Response.Status,
			"size":        r.Response.BufferLength(),
		}
	}

	// 如果发生错误，记录错误信息
	if err := r.GetError(); err != nil {
		logInfo.Error = err.Error()
	}

	// 将完整日志信息存入上下文
	ctx = context.WithValue(r.Context(), "gServer_info", logInfo)
	r.SetCtx(ctx)

	// 发送日志到全局协程处理器
	processor := GetLogProcessor()
	processor.SendLog(logInfo)
}

// generateTraceID 生成请求追踪ID
func generateTraceID() string {
	// 使用时间戳和随机数生成traceID
	// 实际项目中可以使用UUID或雪花算法
	return gconv.String(time.Now().UnixNano()) + "-" + gconv.String(grand.N(0, 10000))
}

// GetTraceIDFromCtx 从上下文中获取traceID
func GetTraceIDFromCtx(ctx context.Context) string {
	if value := ctx.Value("traceId"); value != nil {
		if traceID, ok := value.(string); ok {
			return traceID
		}
	}
	return ""
}

// GetLogInfoFromCtx 从上下文中获取日志信息
func GetLogInfoFromCtx(ctx context.Context) *LogInfo {
	if value := ctx.Value("gServer_info"); value != nil {
		if info, ok := value.(*LogInfo); ok {
			return info
		}
	}
	return nil
}

// Init 初始化日志处理器
func Init() {
	// 启动时初始化全局处理器
	_ = GetLogProcessor()
	g.Log().Info(context.Background(), "日志管理器初始化！")
}

// Shutdown 关闭日志处理器
func Shutdown() {
	if globalProcessor != nil {
		globalProcessor.StopWorker()
		g.Log().Info(context.Background(), "日志关闭")
	}
}
