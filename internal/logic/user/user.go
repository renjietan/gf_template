package user

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"

	v1 "gf_template/api/user/v1"
)

type sUser struct{} // 结构体名以 's' 开头

func (s *sUser) GetList(ctx context.Context, req *v1.PagerReq) (res *v1.PagerRes, err error) {
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
