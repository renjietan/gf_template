package interceptor

import (
	"context"
	"gf_template/utility/ternary"
	"sync"
	"sync/atomic"

	"github.com/gogf/gf/v2/frame/g"
)

// LogProcessor 日志处理器
type LogProcessor struct {
	logChan         chan *LogInfo   // 日志通道
	workerExit      chan bool       // 协程退出信号
	workerOnce      sync.Once       // 确保协程只启动一次
	isWorkerRunning bool            // 协程运行状态
	mutex           sync.RWMutex    // 读写锁
	ctx             context.Context // 协程上下文（持久化）
	closing         int32           // 关闭标志，原子操作
	wg              sync.WaitGroup  // 等待协程退出
}

// StartWorker 启动日志处理协程
func (lp *LogProcessor) StartWorker() {
	lp.mutex.Lock()
	defer lp.mutex.Unlock()

	if !lp.isWorkerRunning {
		lp.workerOnce.Do(func() {
			lp.isWorkerRunning = true
			lp.wg.Add(1)
			go lp.workerRoutine()
		})
	}
}

// workerRoutine 日志处理协程（单协程处理所有日志）
func (lp *LogProcessor) workerRoutine() {
	defer func() {
		lp.wg.Done()
		if err := recover(); err != nil {
			// 使用协程的持久化上下文记录错误
			g.Log().Error(lp.ctx, "Log processor panic recovered:", err)
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
		case logInfo, ok := <-lp.logChan:
			if !ok {
				g.Log().Info(lp.ctx, "通道已关闭, 退出")
				return
			}
			lp.processLog(logInfo)
		case <-lp.workerExit:
			g.Log().Info(lp.ctx, "日志管理器已退出")
			return
		}
	}
}

// processLog 处理日志的实际逻辑
func (lp *LogProcessor) processLog(logInfo *LogInfo) {
	logCtx := context.WithValue(lp.ctx, "traceId", logInfo.TraceID)
	g.Log().Debugf(logCtx, "【START】%v(%v): (耗时: %v)%v", logInfo.Method, logInfo.Status, logInfo.Duration, logInfo.URL)
	info, err := PrintAsJSON(logInfo)
	info = ternary.If(err == "", info, err)
	g.Log().Info(logCtx, info)
	g.Log().Debugf(logCtx, "【END】%v(%v): (耗时: %v)%v", logInfo.Method, logInfo.Status, logInfo.Duration, logInfo.URL)
}

// 发送日志到处理协程
func (lp *LogProcessor) SendLog(logInfo *LogInfo) {
	// 检查是否正在关闭
	if atomic.LoadInt32(&lp.closing) == 1 {
		return
	}
	// 使用select避免阻塞，如果通道满则丢弃日志
	select {
	case lp.logChan <- logInfo:
		// 成功发送
	default:
		// 通道已满，使用协程的持久化上下文记录警告
		g.Log().Warning(lp.ctx, "Log channel is full, dropping log")
	}
}

// StopWorker 停止日志处理协程
func (lp *LogProcessor) StopWorker() {
	lp.mutex.Lock()
	defer lp.mutex.Unlock()

	if lp.isWorkerRunning {
		// 设置关闭标志，阻止新的日志发送
		atomic.StoreInt32(&lp.closing, 1)
		// 关闭通道，这样workerRoutine会在处理完通道中已有的日志后退出
		close(lp.logChan)
		// 等待协程退出
		lp.wg.Wait()
		// 关闭退出信号通道，尽管我们已经通过关闭logChan让协程退出了，但这里还是关闭workerExit，防止协程阻塞在workerExit上
		close(lp.workerExit)
		lp.isWorkerRunning = false
	}
}
