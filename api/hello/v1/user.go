package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserListReq struct {
	g.Meta `path:"/user" tags:"用户管理" method:"get" summary:"用户列表"`
}
type UserListRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
