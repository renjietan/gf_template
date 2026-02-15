package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gogf/gf/v2/container/gqueue"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
)

func TestContext(t *testing.T) {
	q := gqueue.NewTQueue[string]()
	ctx := gctx.New()
	// 数据生产者，每隔1秒往队列写数据
	gtimer.SetInterval(ctx, time.Second, func(ctx context.Context) {
		v := gtime.Now().String()
		q.Push(v)
		fmt.Println("Push:", v)
	})

	// 3秒后关闭队列
	gtimer.SetTimeout(ctx, 3*time.Second, func(ctx context.Context) {
		q.Close()
	})

	// 消费者，不停读取队列数据并输出到终端
	for {
		if v := q.Pop(); v != "" {
			fmt.Println("Pop:", v)
		} else {
			break
		}
	}
}
