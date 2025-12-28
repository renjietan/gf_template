// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 用于 DAO 操作（如查询/数据操作）的表 t_user 的 Go 语言结构体。
type TUser struct {
	g.Meta   `orm:"table:t_user, do:true"`
	Id       any         //
	Name     any         //
	FId      any         //
	DeleteAt any         //
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
	InfoId   any         //
}
