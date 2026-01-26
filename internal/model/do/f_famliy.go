// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FFamliy is the golang structure of table f_famliy for DAO operations like Where/Data.
type FFamliy struct {
	g.Meta   `orm:"table:f_famliy, do:true"`
	Id       any         //
	Name     any         //
	DeleteAt any         //
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
}
