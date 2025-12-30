package interceptor

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/frame/g"

	"gf_template/utility/ternary"
)

// LogProcessor 日志处理器
type LogProcessor struct {
	logChan         chan *LogInfo // 日志通道
	workerExit      chan bool     // 协程退出信号
	workerOnce      sync.Once     // 确保协程只启动一次
	isWorkerRunning bool          // 协程运行状态
	mutex           sync.RWMutex  // 读写锁
	context         chan *context.Context
}

// StartWorker 启动日志处理协程
func (lp *LogProcessor) StartWorker() {
	lp.mutex.Lock()
	defer lp.mutex.Unlock()

	if !lp.isWorkerRunning {
		lp.workerOnce.Do(func() {
			lp.isWorkerRunning = true
			go lp.workerRoutine()
		})
	}
}

// workerRoutine 日志处理协程（单协程处理所有日志）
func (lp *LogProcessor) workerRoutine() {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Error(context.Background(), "错误：", err)
			// 重新启动协程
			lp.mutex.Lock()
			lp.isWorkerRunning = false
			lp.workerOnce = sync.Once{}
			lp.mutex.Unlock()
			lp.StartWorker()
		}
	}()

	for {
		select {
		case logInfo := <-lp.logChan:
			// 这里处理日志，可以存储到数据库、发送到日志系统等
			// 注意：这里运行在同一个协程中
			lp.processLog(logInfo)
		case <-lp.workerExit:
			g.Log().Debug(context.Background(), "日志协程退出-----")
			return
		}
	}
}

// processLog 处理日志的实际逻辑
func (lp *LogProcessor) processLog(logInfo *LogInfo) {
	g.Log().Debugf(context.Background(), "【START】%v(%v): (耗时: %v)%v", logInfo.Method, logInfo.Status, logInfo.Duration, logInfo.URL)
	info, err := PrintAsJSON(logInfo)
	info = ternary.If(err == nil, info, err.Error())
	g.Log().Info(context.Background(), info)
	g.Log().Debugf(context.Background(), "【END】%v(%v): (耗时: %v)%v", logInfo.Method, logInfo.Status, logInfo.Duration, logInfo.URL)
}

// SendLog 发送日志到处理协程
func (lp *LogProcessor) SendLog(logInfo *LogInfo) {
	// 使用select避免阻塞，如果通道满则丢弃日志
	select {
	case lp.logChan <- logInfo:
		// 成功发送
	default:
		// 通道已满，记录警告
		g.Log().Warning(context.Background(), "Log channel is full, dropping log")
	}
}

// StopWorker 停止日志处理协程
func (lp *LogProcessor) StopWorker() {
	lp.mutex.Lock()
	defer lp.mutex.Unlock()

	if lp.isWorkerRunning {
		close(lp.workerExit)
		lp.isWorkerRunning = false
	}
}
