package logger

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

// LogConfig 日志配置
type LogConfig struct {
	LogPath       string // 日志路径（相对路径）
	MaxAge        int    // 文件保留天数，默认7天
	RotationSize  int64  // 日志文件切割大小，默认100MB
	RotationCount uint   // 日志文件保留数量，默认7个
}

// DefaultLogConfig 默认日志配置
func DefaultLogConfig() *LogConfig {
	return &LogConfig{
		LogPath:       "./logs",
		MaxAge:        7,
		RotationSize:  100 * 1024 * 1024, // 100MB
		RotationCount: 7,
	}
}

// LogMiddleware 日志中间件
func Init(config *LogConfig) ghttp.HandlerFunc {
	if config == nil {
		config = DefaultLogConfig()
	}

	// 初始化日志配置
	initLogConfig(config)

	return func(r *ghttp.Request) {
		startTime := time.Now()

		// 处理请求
		r.Middleware.Next()

		// 记录访问日志
		logAccess(r, startTime)
	}
}

// initLogConfig 初始化日志配置
func initLogConfig(config *LogConfig) {
	// 确保日志目录存在
	logPath := config.LogPath

	// 配置不同级别的日志
	setupLogger("info", logPath, config, "info")
	setupLogger("debug", logPath, config, "debug")
	setupLogger("error", logPath, config, "error")
	setupLogger("warning", logPath, config, "warning")
	setupLogger("access", logPath, config, "access")
}

// setupLogger 设置单个日志实例
func setupLogger(name, basePath string, config *LogConfig, level string) {
	// 每个级别有自己的目录
	levelPath := filepath.Join(basePath, level)

	logger := glog.New()

	// 设置日志级别
	switch level {
	case "info":
		logger.SetLevelStr("INFO")
	case "debug":
		logger.SetLevelStr("DEBUG")
	case "error":
		logger.SetLevelStr("ERROR")
	case "warning":
		logger.SetLevelStr("WARN")
	default:
		logger.SetLevelStr("ALL")
	}

	// 配置日志输出
	cfg := fmt.Sprintf(
		`{
            "path": "%s",
            "file": "{Y-m-d}.log",
            "level": "%s",
            "StdoutPrint": false,
            "RotateSize": %d,
            "RotateExpire": %dh,
            "RotateBackupLimit": %d,
            "RotateBackupExpire": %dh,
            "RotateBackupCompress": 0
        }`,
		levelPath,
		level,
		config.RotationSize,
		config.MaxAge*24, // 转换为小时
		config.RotationCount,
		config.MaxAge*24,
	)

	if err := logger.SetConfigWithMap(map[string]any{
		"config": cfg,
	}); err != nil {
		panic(fmt.Sprintf("初始化日志失败: %v", err))
	}

	// 存储到全局变量或上下文中
	// 这里使用默认实例和自定义实例两种方式
	if name == "info" {
		g.Log(name).SetWriter(logger.GetWriter())
	}
}

// logAccess 记录访问日志
func logAccess(r *ghttp.Request, startTime time.Time) {
	endTime := time.Now()
	latency := endTime.Sub(startTime)

	// 获取请求信息
	status := r.Response.Status
	clientIP := r.GetClientIp()
	method := r.Method
	path := r.URL.Path
	userAgent := r.Header.Get("User-Agent")

	// 获取日志实例
	accessLogger := getLogger("access")

	// 记录访问日志
	accessLogger.Info(context.Background(),
		fmt.Sprintf("[%s] %s %s %d %v %s",
			clientIP,
			method,
			path,
			status,
			latency,
			userAgent,
		))
}

// 全局日志实例存储
var logInstances = make(map[string]*glog.Logger)

// getLogger 获取日志实例
func getLogger(name string) *glog.Logger {
	// 简单实现，实际使用时可以更完善
	return glog.New()
}

// LogDebug 记录调试日志
func LogDebug(ctx context.Context, v ...interface{}) {
	logger := getLogger("debug")
	logger.Debug(ctx, v...)
}

// LogInfo 记录信息日志
func LogInfo(ctx context.Context, v ...interface{}) {
	logger := getLogger("info")
	logger.Info(ctx, v...)
}

// LogWarning 记录警告日志
func LogWarning(ctx context.Context, v ...interface{}) {
	logger := getLogger("warning")
	logger.Warning(ctx, v...)
}

// LogError 记录错误日志
func LogError(ctx context.Context, v ...interface{}) {
	logger := getLogger("error")
	logger.Error(ctx, v...)
}

// LogErrorf 格式化记录错误日志
func LogErrorf(ctx context.Context, format string, v ...interface{}) {
	logger := getLogger("error")
	logger.Errorf(ctx, format, v...)
}

// LogPanic 记录Panic日志
func LogPanic(ctx context.Context, v ...interface{}) {
	logger := getLogger("error")
	logger.Panic(ctx, v...)
}

// Logger 封装日志记录器
type Logger struct {
	name string
}

// NewLogger 创建新的日志记录器
func NewLogger(name string) *Logger {
	return &Logger{name: name}
}

// WithCtx 使用上下文记录日志
func (l *Logger) WithCtx(ctx context.Context) *Logger {
	// 可以实现更复杂的上下文处理
	return l
}

// Debug 记录调试日志
func (l *Logger) Debug(v ...interface{}) {
	LogDebug(context.Background(), append([]interface{}{l.name + ":"}, v...)...)
}

// Debugf 格式化记录调试日志
func (l *Logger) Debugf(format string, v ...interface{}) {
	LogDebug(context.Background(), l.name+": "+fmt.Sprintf(format, v...))
}

// Info 记录信息日志
func (l *Logger) Info(v ...interface{}) {
	LogInfo(context.Background(), append([]interface{}{l.name + ":"}, v...)...)
}

// Infof 格式化记录信息日志
func (l *Logger) Infof(format string, v ...interface{}) {
	LogInfo(context.Background(), l.name+": "+fmt.Sprintf(format, v...))
}
