package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"gf_template/internal/model/entity"
	"gf_template/utility"
)

type PagerReq struct {
	g.Meta `path:"/user" method:"get" tags:"User" summary:"用户列表"`
	// 可填 可不填 必须是指针类型
	Name *string `dc:"用户名称"`
	utility.CommonPaginationReq
}
type PagerRes struct {
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
	Name   string `json:"name" v:"required" dc:"用户名称"`
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
