package queues

import (
	"context"
	"encoding/json"
	"gf_template/internal/consts"
	"gf_template/internal/library/queue"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type SysServeLog struct {
	Id          int64       `json:"id"          orm:"id"           description:"日志ID"`
	TraceId     string      `json:"traceId"     orm:"trace_id"     description:"链路ID"`
	LevelFormat string      `json:"levelFormat" orm:"level_format" description:"日志级别"`
	Content     string      `json:"content"     orm:"content"      description:"日志内容"`
	Stack       *gjson.Json `json:"stack"       orm:"stack"        description:"打印堆栈"`
	Line        string      `json:"line"        orm:"line"         description:"调用行"`
	TriggerNs   int64       `json:"triggerNs"   orm:"trigger_ns"   description:"触发时间(ns)"`
	Status      int         `json:"status"      orm:"status"       description:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"修改时间"`
}

func init() {
	queue.RegisterConsumer(ServeLog)
}

// ServeLog 登录日志
var ServeLog = &qServeLog{}

type qServeLog struct{}

// GetTopic 主题
func (q *qServeLog) GetTopic() string {
	return consts.QueueServeLogTopic
}

// Handle 处理消息
func (q *qServeLog) Handle(ctx context.Context, mqMsg queue.MqMsg) error {
	var data SysServeLog
	if err := json.Unmarshal(mqMsg.Body, &data); err != nil {
		g.Dump("ServeLog Handle Unmarshal err:", err)
		return nil
	}

	return nil
}
