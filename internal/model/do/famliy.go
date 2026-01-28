// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Famliy is the golang structure of table gf_famliy for DAO operations like Where/Data.
type Famliy struct {
	g.Meta    `orm:"table:gf_famliy, do:true"`
	Id        any         //
	Name      any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
