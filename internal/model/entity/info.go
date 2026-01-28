// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Info is the golang structure for table info.
type Info struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Info      string      `json:"info"      orm:"info"       description:""`
	DeleteAt  int         `json:"deleteAt"  orm:"delete_at"  description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
}
