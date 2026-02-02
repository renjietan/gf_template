package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // 导入 MySQL 驱动
	_ "github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"gf_template/internal/cmd"
	"gf_template/internal/global"
	_ "gf_template/internal/packed"
)

func main() {
	var ctx = gctx.GetInitCtx()
	g.Log().Error(ctx, "11111111111111")
	global.Init(ctx)
	cmd.Main.Run(ctx)
}
