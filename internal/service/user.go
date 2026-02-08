// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gf_template/api/user/v1"
)

type (
	IUser interface {
		GetList(ctx context.Context, req *v1.PagerReq) (res *v1.PagerRes, err error)
		Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
		Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
		GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
		Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
