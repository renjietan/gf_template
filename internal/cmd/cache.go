package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var Cache = &gcmd.Command{
	Name:  "cache",
	Usage: "cache",
	Brief: "缓存管理器",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

		return nil
	},
}
