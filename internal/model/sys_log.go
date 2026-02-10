package model

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// LogConfig 日志配置
type LogConfig struct {
	Switch   bool     `json:"switch"`
	Queue    bool     `json:"queue"`
	Module   []string `json:"module"`
	SkipCode []string `json:"skipCode"`
}

type SysLog struct {
	Id         int64       `json:"id"         orm:"id"           description:"日志ID"`
	ReqId      string      `json:"reqId"      orm:"req_id"       description:"对外ID"`
	UserId     int64       `json:"userId"   orm:"user_id"    description:"用户ID"`
	Method     string      `json:"method"     orm:"method"       description:"提交类型"`
	Url        string      `json:"url"        orm:"url"          description:"提交url"`
	GetData    *gjson.Json `json:"getData"    orm:"get_data"     description:"get数据"`
	PostData   *gjson.Json `json:"postData"   orm:"post_data"    description:"post数据"`
	HeaderData *gjson.Json `json:"headerData" orm:"header_data"  description:"header数据"`
	ErrorCode  int         `json:"errorCode"  orm:"error_code"   description:"报错code"`
	ErrorMsg   string      `json:"errorMsg"   orm:"error_msg"    description:"对外错误提示"`
	ErrorData  *gjson.Json `json:"errorData"  orm:"error_data"   description:"报错日志"`
	UserAgent  string      `json:"userAgent"  orm:"user_agent"   description:"UA信息"`
	TakeUpTime int64       `json:"takeUpTime" orm:"take_up_time" description:"请求耗时"`
	Timestamp  int64       `json:"timestamp"  orm:"timestamp"    description:"响应时间"`
	Status     int         `json:"status"     orm:"status"       description:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"   description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"   description:"修改时间"`
}
