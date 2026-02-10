// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_template/internal/model"

	"github.com/gogf/gf/v2/encoding/gjson"
)

type (
	ISysLog interface {
		// AutoLog 根据配置自动记录请求日志
		AutoLog(ctx context.Context) error
		// AnalysisLog 解析日志数据
		AnalysisLog(ctx context.Context) model.SysLog
		SimplifyHeaderParams(data *gjson.Json) *gjson.Json
	}
)

var (
	localSysLog ISysLog
)

func SysLog() ISysLog {
	if localSysLog == nil {
		panic("implement not found for interface ISysLog, forgot register?")
	}
	return localSysLog
}

func RegisterSysLog(i ISysLog) {
	localSysLog = i
}
