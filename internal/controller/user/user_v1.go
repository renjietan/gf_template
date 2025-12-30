package user

import (
	"context"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf_template/api/user/v1"
	"gf_template/internal/dao"
	"gf_template/internal/model/entity"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.PagerReq) (res *v1.PagerRes, err error) {
	res = &v1.PagerRes{}
	totalInt := int(res.Total)
	totalIntPtr := &totalInt
	var query *gdb.Model = dao.TUser.Ctx(ctx).With(entity.FFamliy{})
	// 指针类型 可判断 为是否 nil
	if req.Name != nil {
		err = query.WhereLike("name", "%"+*req.Name+"%").Limit(
			int(req.Page)-1,
			int(req.Size),
		).ScanAndCount(&res.List, totalIntPtr, false)
	} else {
		err = query.Limit(
			int(req.Page)-1,
			int(req.Size),
		).ScanAndCount(&res.List, &res.Total, false)
		res.Length = len(res.List)
	}
	return
}

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = dao.TUser.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		info_id, err_info := dao.TInfo.Ctx(ctx).Data(entity.TInfo{
			Info: "张三--详情",
		}).OmitEmpty().InsertAndGetId()
		if err_info != nil {
			return err_info
		}
		_, err_user := dao.TUser.Ctx(ctx).Data(entity.TUser{
			Name:   "张三",
			FId:    1,
			InfoId: info_id,
		}).OmitEmpty().Insert()
		if err_user != nil {
			return err_user
		}
		return nil
	})
	return
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	_, err = dao.TUser.Ctx(ctx).Delete("id = ", req.Id)
	return
}

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	res = &v1.GetOneRes{}
	err = dao.TUser.Ctx(ctx).Where("id", req.Id).Scan(&res.TUser)
	return
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
