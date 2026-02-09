package model

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf_template/internal/model/entity"
)

// Context 请求上下文结构
type Context struct {
	User     *entity.User // 上下文用户信息
	Response *Response    // 请求响应
	Data     g.Map        // 自定kv变量 业务模块根据需要设置，不固定
}
