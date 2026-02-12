package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf_template/internal/consts"
	"gf_template/internal/controller/user"
	"gf_template/internal/service"
)

var Http = &gcmd.Command{
	Name:  "http",
	Usage: "http",
	Brief: "HTTP",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		s := g.Server()
		// 自定义 swagger 模板
		s.SetSwaggerUITemplate(consts.SwaggerTpl)

		// 处理程序响应对象及其错误。
		// s.Use(reponse.MiddlewareHandlerResponse)
		// 初始化请求前回调
		s.BindHookHandler("/*any", ghttp.HookBeforeServe, service.Hook().BeforeServe)

		// 请求响应结束后回调
		s.BindHookHandler("/*any", ghttp.HookAfterOutput, service.Hook().AfterOutput)

		// 注册全局中间件，按照注册顺序执行
		s.BindMiddleware("/*any", []ghttp.HandlerFunc{
			service.Middleware().Ctx,  // 初始化请求上下文，需要第一个进行加载，后续中间件存在依赖关系, 否则后续中间价无法找到上下文数据
			service.Middleware().CORS, // 跨域中间件，自动处理跨域问题
		}...)
		// TODO: 需要单独管理
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Bind(
				user.NewV1(),
			)
		})
		// 服务性能工具：http://localhost:8199/debug/pprof
		// 安装 Graphviz 图形化工具，
		// 		(1) 执行命令：go tool pprof -http :8080 "http://127.0.0.1:8199/debug/pprof/profile"（等待30秒）
		// 		(2-1) curl http://127.0.0.1:8199/debug/pprof/profile > pprof.profile
		// 		(2-2) go tool pprof -http :8080 pprof.profile
		// heap: 报告内存分配样本；用于监视当前和历史内存使用情况，并检查内存泄漏。
		// threadcreate: 报告了导致创建新OS线程的程序部分。
		// goroutine: 报告所有当前 goroutine 的堆栈跟踪。
		// block: 显示 goroutine 在哪里阻塞同步原语（包括计时器通道）的等待。默认情况下未启用，需要手动调用 runtime.SetBlockProfileRate 启用。
		// mutex: 报告锁竞争。默认情况下未启用，需要手动调用 runtime.SetMutexProfileFraction 启用。
		go ghttp.StartPProfServer(consts.PprofPath)
		// 开启服务性能工具
		s.EnablePProf()
		// 设置优雅模式 涉及到 shutdown restart
		s.SetGraceful(true)
		// 通过web网页控制服务 重启 与 关闭，默认地址：http://localhost:8000/debug/admin
		s.EnableAdmin()
		serverWg.Add(1)
		fmt.Println("======================= 开启HTTP ==========================")
		// 信号监听
		signalListen(ctx, signalHandlerForOverall)

		go func() {
			<-serverCloseSignal
			_ = s.Shutdown() // 关闭http服务，主服务建议放在最后一个关闭
			g.Log().Debug(ctx, "\n收到关闭信号, 关闭HTTP服务")
			serverWg.Done()
		}()
		s.Run()
		return nil
	},
}
