package user

import (
	"context"
	v1 "gf_template/api/user/v1"
	"gf_template/internal/dao"
	"gf_template/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.PagerReq) (res *v1.PagerRes, err error) {
	res = &v1.PagerRes{}
	totalInt := int(res.Total)
	totalIntPtr := &totalInt
	var query *gdb.Model = dao.User.Ctx(ctx).With(entity.Famliy{})
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
	err = dao.User.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		info_id, err_info := dao.Info.Ctx(ctx).Data(entity.Info{
			Info: "张三--详情",
		}).OmitEmpty().InsertAndGetId()
		if err_info != nil {
			return err_info
		}
		_, err_user := dao.User.Ctx(ctx).Data(entity.User{
			Name:   "张三",
			FId:    1,
			InfoId: int(info_id),
		}).OmitEmpty().Insert()
		if err_user != nil {
			return err_user
		}
		return nil
	})
	return
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	_, err = dao.User.Ctx(ctx).Delete("id = ", req.Id)
	return
}

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	res = &v1.GetOneRes{}
	err = dao.User.Ctx(ctx).Where("id", req.Id).Scan(&res.User)
	return
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
