package sysconfig

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func GetCacheAdapter(ctx context.Context) string {
	res := g.Cfg().MustGet(ctx, "cache.adapter").String()
	return res
}

func GetCacheFileDir(ctx context.Context) string {
	res := g.Cfg().MustGet(ctx, "cache.fileDir").String()
	return res
}
