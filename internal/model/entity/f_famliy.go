// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FFamliy is the golang structure for table f_famliy.
type FFamliy struct {
	Id       int         `json:"id"       orm:"id"        description:""` //
	Name     string      `json:"name"     orm:"name"      description:""` //
	DeleteAt int         `json:"deleteAt" orm:"delete_at" description:""` //
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:""` //
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:""` //
}
