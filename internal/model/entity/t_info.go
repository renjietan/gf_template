// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TInfo is the golang structure for table t_info.
type TInfo struct {
	Id       int         `json:"id"       orm:"id"        description:""` //
	Info     string      `json:"info"     orm:"Info"      description:""` //
	DeleteAt int         `json:"deleteAt" orm:"delete_at" description:""` //
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:""` //
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:""` //
}
