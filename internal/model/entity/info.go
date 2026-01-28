// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Info is the golang structure for table info.
type Info struct {
	Id       int         `json:"id"       orm:"id"        description:""`
	Info     string      `json:"info"     orm:"info"      description:""`
	DeleteAt int         `json:"deleteAt" orm:"delete_at" description:""`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:""`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:""`
}
