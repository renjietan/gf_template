// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUser is the golang structure of table t_user for DAO operations like Where/Data.
type TUser struct {
	g.Meta   `orm:"table:t_user, do:true"`
	Id       any         //
	Name     any         //
	FId      any         //
	InfoId   any         //
	DeleteAt any         //
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
}
