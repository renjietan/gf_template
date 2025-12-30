package utility

type CommonPaginationReq struct {
	// in:"query": 从URL的查询参数中获取（即?后面的参数）
	Page int `json:"page" in:"query" d:"1"  v:"min:1#gf.gvalid.rule.min"     dc:"分页号码,默认1"`
	Size int `json:"size" in:"query" d:"10" v:"min:1#gf.gvalid.rule.min|max:100#gf.gvalid.rule.max" dc:"分页数量,最大100"`
}

type CommonPaginationRes[T any] struct {
	List   []T `json:"list" dc:"列表数据"`
	Total  int `json:"total" dc:"总数"`
	Page   int `json:"page" dc:"分页号码"`
	Size   int `json:"size" dc:"分页数量"`
	Length int `json:"length" dc:"当前页条数"`
}

type CommonIdReq struct {
	Id uint `json:"id" v:"required"`
}
