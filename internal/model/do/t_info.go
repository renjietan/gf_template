// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TInfo is the golang structure of table t_info for DAO operations like Where/Data.
type TInfo struct {
	g.Meta   `orm:"table:t_info, do:true"`
	Id       any         //
	Info     any         //
	DeleteAt any         //
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
}
