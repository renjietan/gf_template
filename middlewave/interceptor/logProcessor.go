package interceptor

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/gogf/gf/v2/frame/g"

	"gf_template/utility/ternary"
)

// 日志处理器
type LogManager struct {
	logChan         chan *LogInfo   // 日志通道
	workerExit      chan bool       // 协程退出信号
	workerOnce      sync.Once       // 确保协程只启动一次
	isWorkerRunning atomic.Bool     // 协程运行状态
	mutex           sync.RWMutex    // 读写锁
	ctx             context.Context // 协程上下文（持久化）
	closing         atomic.Bool     // 关闭标志，原子操作
	wg              sync.WaitGroup  // 等待协程退出
}

// 启动日志处理协程
func (logger *LogManager) StartWorker() {
	logger.isWorkerRunning.Store(true)
	if !logger.isWorkerRunning.Load() {
		logger.workerOnce.Do(func() {
			logger.isWorkerRunning.Store(true)
			logger.wg.Add(1)
			go logger.workerRoutine()
		})
	}
}

// 日志处理协程（单协程处理所有日志）
func (logger *LogManager) workerRoutine() {
	defer func() {
		logger.wg.Done()
		if err := recover(); err != nil {
			// 使用协程的持久化上下文记录错误
			g.Log().Error(logger.ctx, "Log processor panic recovered:", err)
			// 重新启动协程
			logger.isWorkerRunning.Store(false)
			logger.mutex.Lock()
			logger.workerOnce = sync.Once{}
			logger.mutex.Unlock()
			logger.StartWorker()
		}
	}()

	for {
		select {
		case logInfo, ok := <-logger.logChan:
			if !ok {
				g.Log().Info(logger.ctx, "通道已关闭, 退出")
				return
			}
			logger.processLog(logInfo)
		case <-logger.workerExit:
			g.Log().Info(logger.ctx, "日志管理器已退出")
			return
		}
	}
}

// processLog 处理日志的实际逻辑
func (logger *LogManager) processLog(logInfo *LogInfo) {
	logCtx := context.WithValue(logger.ctx, "traceId", logInfo.TraceID)
	g.Log().Debugf(logCtx, "【START】%v(%v): (耗时: %v)%v", logInfo.Method, logInfo.Status, logInfo.Duration, logInfo.URL)
	info, err := PrintAsJSON(logInfo)
	info = ternary.If(err == "", info, err)
	g.Log().Info(logCtx, info)
	g.Log().Debugf(logCtx, "【END】%v(%v): (耗时: %v)%v", logInfo.Method, logInfo.Status, logInfo.Duration, logInfo.URL)
}

// 发送日志到处理协程
func (logger *LogManager) SendLog(logInfo *LogInfo) {
	if logger.closing.Load() {
		return
	}
	select {
	case logger.logChan <- logInfo:
		// 防止堵塞
	default:
		g.Log().Warning(logger.ctx, "通道已满")
	}
}

// StopWorker 停止日志处理协程
func (logger *LogManager) StopWorker() {
	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	if logger.isWorkerRunning.Load() {
		logger.closing.Store(true)
		close(logger.logChan)
		logger.wg.Wait()
		close(logger.workerExit)
		logger.isWorkerRunning.Store(false)
	}
}
