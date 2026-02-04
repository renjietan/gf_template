// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table gf_user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:gf_user, do:true"`
	Id        any         //
	Name      any         //
	FId       any         //
	DeleteAt  *gtime.Time //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	InfoId    any         //
	Delete    any         // 删除标识
}
