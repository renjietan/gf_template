// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	Name      string      `json:"name"      orm:"name"       description:""`
	FId       int         `json:"fId"       orm:"f_id"       description:""`
	DeleteAt  bool        `json:"deleteAt"  orm:"delete_at"  description:""`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	InfoId    int         `json:"infoId"    orm:"infoId"     description:""`
	Delete    int         `json:"delete"    orm:"delete"     description:"删除标识"`
}
