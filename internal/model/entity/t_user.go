// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUser is the golang structure for table t_user.
type TUser struct {
	Id       int64       `json:"id"       orm:"id"        description:""` //
	Name     string      `json:"name"     orm:"name"      description:""` //
	FId      int         `json:"fId"      orm:"f_id"      description:""` //
	InfoId   int64       `json:"infoId"   orm:"infoId"    description:""` //
	DeleteAt int         `json:"deleteAt" orm:"delete_at" description:""` //
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:""` //
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:""` //
}
