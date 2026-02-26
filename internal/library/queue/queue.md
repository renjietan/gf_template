- 使用方式
```aiignore
simple.SafeGo(ctx, func(ctx context.Context) {
    msg := queues.SysServeLog{
        Id:          0,
        TraceId:     "TraceId",
        LevelFormat: "WARN",
        Content:     "我是 AAAAAA",
        Stack:       nil,
        Line:        "",
        TriggerNs:   0,
        Status:      0,
        CreatedAt:   nil,
        UpdatedAt:   nil,
    }
    # 将消息 丢入 队列中
    queue.Push(consts.QueueServeLogTopic, msg)
})
```