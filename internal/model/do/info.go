// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Info is the golang structure of table gf_info for DAO operations like Where/Data.
type Info struct {
	g.Meta   `orm:"table:gf_info, do:true"`
	Id       any         //
	Info     any         //
	DeleteAt any         //
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
}
