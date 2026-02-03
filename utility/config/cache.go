package sysconfig

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func GetCacheAdapter(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.cache.adapter").String()
}

func GetCacheFileDir(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.cache.fileDir").String()
}
