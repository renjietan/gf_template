package cache

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"

	"gf_template/internal/library/cache/file"
	sysconfig "gf_template/utility/config"
)

// cache 缓存驱动
var cache *gcache.Cache

// Instance 缓存实例
func Instance() *gcache.Cache {
	if cache == nil {
		panic("缓存未初始化")
	}
	return cache
}

// SetAdapter 设置缓存适配器
func SetAdapter(ctx context.Context) {
	var adapter gcache.Adapter
	switch sysconfig.GetCacheAdapter(ctx) {
	case "redis":
		adapter = gcache.NewAdapterRedis(g.Redis())
	case "file":
		fileDir := sysconfig.GetCacheFileDir(ctx)
		if fileDir == "" {
			g.Log().Fatal(ctx, "缓存文件夹路径未配置")
			return
		}

		if !gfile.Exists(fileDir) {
			if err := gfile.Mkdir(fileDir); err != nil {
				g.Log().Fatalf(ctx, "缓存文件夹创建失败, err:%+v", err)
				return
			}
		}
		adapter = file.NewAdapterFile(fileDir)
	default:
		adapter = gcache.NewAdapterMemory()
	}

	// 数据库缓存，默认和通用缓冲驱动一致，后续有需要再自定义
	g.DB().GetCache().SetAdapter(adapter)

	// 通用缓存
	cache = gcache.New()
	cache.SetAdapter(adapter)
}
