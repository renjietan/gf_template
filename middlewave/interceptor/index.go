package interceptor

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	// 全局日志处理器单例
	globalProcessor *LogManager
	processorOnce   sync.Once
)

// 首次调用将初始化，第 N 次调用获取日志管理器指针
func GetLogManager() *LogManager {
	processorOnce.Do(func() {
		// 创建协程的持久化上下文
		workerCtx := context.Background()
		workerCtx = context.WithValue(workerCtx, "worker", "log_processor")

		globalProcessor = &LogManager{
			logChan:    make(chan *LogInfo, 1000),
			workerExit: make(chan bool),
			ctx:        workerCtx,
		}
		// 启动协程
		globalProcessor.StartWorker()
	})
	return globalProcessor
}

// 全局日志中间件
func Logger(r *ghttp.Request) {
	startTime := time.Now()

	// 生成请求追踪ID
	traceID := generateTraceID()

	// 创建日志信息结构
	logInfo := &LogInfo{
		TraceID:     traceID,
		Method:      r.Method,
		Path:        r.URL.Path,
		URL:         r.Request.URL.String(),
		QueryParams: r.URL.RawQuery,
		RequestTime: startTime.Format("2006-05-04 15:02:01"),
		Status:      r.Response.Status,
	}

	// 将traceID存入上下文，方便后续使用
	ctx := context.WithValue(r.Context(), "traceId", traceID)
	r.SetCtx(ctx)

	// 根据请求方法获取请求参数
	switch r.Method {
	case "GET", "DELETE":
		// GET 和 DELETE 请求通常使用 Query 参数
		if params := r.GetMap(); len(params) > 0 {
			logInfo.BodyParams = params
		}
	case "POST", "PUT", "PATCH":
		// POST、PUT、PATCH 请求获取 Body 参数
		contentType := r.Header.Get("Content-Type")
		switch {
		case contentType == "application/json":
			// JSON 请求体
			var jsonBody interface{}
			if err := r.Parse(&jsonBody); err == nil && jsonBody != nil {
				logInfo.BodyParams = jsonBody
			}
		case contentType == "application/x-www-form-urlencoded",
			contentType == "multipart/form-data":
			// Form 表单数据
			if formData := r.GetFormMap(); len(formData) > 0 {
				logInfo.BodyParams = formData
			}
		default:
			// 其他类型，尝试获取原始 Body
			if bodyBytes := r.GetBody(); len(bodyBytes) > 0 {
				// 如果是批量操作，可能包含数组
				if r.URL.Path == "/batch" || len(bodyBytes) < 1024 { // 小数据直接转换
					logInfo.BodyParams = gconv.String(bodyBytes)
				} else {
					// 大数据只记录长度
					logInfo.BodyParams = map[string]interface{}{
						"body_length":  len(bodyBytes),
						"content_type": contentType,
					}
				}
			}
		}
	}

	// 保存请求日志信息到上下文
	ctx = context.WithValue(r.Context(), "requestLogInfo", logInfo)
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

	// 发送日志到全局协程处理器
	processor := GetLogManager()
	processor.SendLog(logInfo)

	// 将完整日志信息存入上下文
	ctx = context.WithValue(r.Context(), "responseLogInfo", logInfo)
	r.SetCtx(ctx)
}

// 生成请求追踪ID
func generateTraceID() string {
	// 使用时间戳和随机数生成traceID
	return gconv.String(time.Now().UnixNano()) + "-" + gconv.String(rand.Intn(10000))
}

// 从上下文中获取traceID
func GetTraceIDFromCtx(ctx context.Context) string {
	if value := ctx.Value("traceId"); value != nil {
		if traceID, ok := value.(string); ok {
			return traceID
		}
	}
	return ""
}

// 从上下文中获取日志信息
func GetLogInfoFromCtx(ctx context.Context) *LogInfo {
	if value := ctx.Value("responseLogInfo"); value != nil {
		if info, ok := value.(*LogInfo); ok {
			return info
		}
	}
	return nil
}

// 从上下文中获取请求信息
func GetRequestInfoFromCtx(ctx context.Context) *LogInfo {
	if value := ctx.Value("requestLogInfo"); value != nil {
		if info, ok := value.(*LogInfo); ok {
			return info
		}
	}
	return nil
}

// 初始化日志处理器
func Init() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 启动时初始化全局处理器
	_ = GetLogManager()
	g.Log().Info(context.Background(), "初始化日志管理器")
}

// 关闭日志处理器
func Shutdown() {
	if globalProcessor != nil {
		globalProcessor.StopWorker()
		g.Log().Info(context.Background(), "关闭日志管理器")
	}
}
