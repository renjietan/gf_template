package user

import (
	"context"

	v1 "gf_template/api/user/v1"
	"gf_template/internal/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.PagerReq) (res *v1.PagerRes, err error) {
	res, err = service.User().GetList(ctx, req)
	return
}

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return service.User().Create(ctx, req)
}

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	// _, err = dao.User.Ctx(ctx).Data(g.Map{
	// 	"'delete'": 1,
	// }).Where("id = ?", req.Id).Update()
	// _, err = dao.User.Ctx(ctx).Where("id = ?", req.Id).Delete()
	res, err = service.User().Delete(ctx, req)
	return
}

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	return service.User().GetOne(ctx, req)
}

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return service.User().Update(ctx, req)
}
