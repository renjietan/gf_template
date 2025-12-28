### 第一步: 
执行 gf gen dao
```
- dao: 生成的dao文件有两个
    - internal/dao/internal/user.go
      用于封装对数据表user的访问。该文件自动生成了一些数据结构和方法，简化对数据表的CRUD操作。该文件每次生成都会覆盖，由开发工具自动维护
    - internal/dao/user.go
      是对internal/dao/internal/user.go的进一步封装，用于供其他模块直接调用访问。该文件开发者可以随意修改，或者扩展dao的能力。
- do: internal/model/do/user.go
  数据转换模型的使用，我们将在后续代码逻辑中做展示
- entity: internal/model/entity/user.go
  实体类，与数据库表一一对应
```

### 第二步
新增 api/user/v1/xxx.go 文件
```
// 新增接口
type CreateReq struct {
    g.Meta `path:"/user" method:"post" tags:"User" summary:"Create user"`
    Name   string `v:"required|length:3,10" dc:"user name"`
    Age    uint   `v:"required|between:18,200" dc:"user age"`
}
type CreateRes struct {
    Id int64 `json:"id" dc:"user id"`
}

// 删除接口
type DeleteReq struct {
    g.Meta `path:"/user/{id}" method:"delete" tags:"User" summary:"Delete user"`
    Id     int64 `v:"required" dc:"user id"`
}
type DeleteRes struct{}

// 更新接口：对Status参数的校验使用了in:0,1校验规则，该规则将会校验传递的Status的值必需是我们定义的常量的两个值StatusOK/StatusDisabled，即0/1
type Status int
const (
    StatusOK       Status = 0 // User is OK.
    StatusDisabled Status = 1 // User is disabled.
)

type UpdateReq struct {
    g.Meta `path:"/user/{id}" method:"put" tags:"User" summary:"Update user"`
    Id     int64   `v:"required" dc:"user id"`
    Name   *string `v:"length:3,10" dc:"user name"`
    Age    *uint   `v:"between:18,200" dc:"user age"`
    Status *Status `v:"in:0,1" dc:"user status"`
}
type UpdateRes struct{}

// 查询单个接口
type GetOneReq struct {
    g.Meta `path:"/user/{id}" method:"get" tags:"User" summary:"Get one user"`
    Id     int64 `v:"required" dc:"user id"`
}
type GetOneRes struct {
    *entity.User `dc:"user"`
}

// 查询接口（列表）
type GetListReq struct {
    g.Meta `path:"/user" method:"get" tags:"User" summary:"Get users"`
    Age    *uint   `v:"between:18,200" dc:"user age"`
    Status *Status `v:"in:0,1" dc:"user status"`
}
type GetListRes struct {
    List []*entity.User `json:"list" dc:"user list"`
}
```

### 第三步
执行 gf gen ctrl
- internal/controller/user/user.go
  空的源码文件，可用于定义一些控制器内部使用的数据结构、常量等内容
- internal/controller/user/user_new.go
  路由对象创建文件
- internal/controller/user/user_v1_get_list.go
  接口的实现模板，我们完善这个路由函数的具体业务逻辑即可

### 第四步（关联查询）
- 修改: internal/model/entity/t_user.go
```
type TUser struct {
	Id      int      `json:"id"   orm:"id"   description:""` // ID
	Name    string   `json:"name" orm:"name" description:""` // 姓名
	FId     int      `json:"-"  orm:"f_id" description:""`   // 添加 json:"-" 隐藏这个字段
	FFamliy *FFamliy `orm:"with:id=f_id"`
}
```
- 修改: internal\controller\user\user_v1_get_list.go
```
func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	err = dao.TUser.Ctx(ctx).With(entity.FFamliy{}).WhereLike("name", "%"+*req.Name+"%").Limit(
		req.Page-1,
		req.Size,
	).ScanAndCount(&res.List, &res.Total, false)
	return
}
```

### 第五步
