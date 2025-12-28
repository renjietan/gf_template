package v1

import (
	"gf_template/internal/model/entity"
	"gf_template/utility"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/user" method:"get" tags:"User" summary:"用户列表"`
	// 可填 可不填 必须是指针类型
	Name *string `dc:"用户名称"`
	utility.CommonPaginationReq
}
type GetListRes struct {
	utility.CommonPaginationRes[*entity.TUser]
}

type GetOneReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"User" summary:"Get one user"`
	Id     int `v:"required" dc:"user id"`
}
type GetOneRes struct {
	*entity.TUser `dc:"user"`
}

type CreateReq struct {
	g.Meta `path:"/user" method:"post" tags:"User" summary:"Create user"`
	// Name   string `v:"required|length:3,10" dc:"user name"`
	// Age    uint   `v:"required|between:18,200" dc:"user age"`
}
type CreateRes struct {
	Id int64 `json:"id" dc:"user id"`
}

type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"User" summary:"删除用户"`
	utility.CommonIdReq
}
type DeleteRes struct{}

type UpdateReq struct {
	g.Meta `path:"/user/{id}" method:"put" tags:"User" summary:"Update user"`
	Id     uint   `v:"required" dc:"user id"`
	Name   string `v:"length:3,10" dc:"user name"`
}
type UpdateRes struct{}
