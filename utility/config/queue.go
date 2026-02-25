package sysconfig

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

func Queue(ctx context.Context) *gvar.Var {
	res := g.Cfg().MustGet(ctx, "queue")
	return res
}
