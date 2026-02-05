package main

import (
	_ "gf_template/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // 导入 MySQL 驱动
	_ "github.com/gogf/gf/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gf_template/internal/cmd"
	"gf_template/internal/global"
	_ "gf_template/internal/packed"
)

func main() {
	var ctx = gctx.GetInitCtx()
	global.Init(ctx)
	cmd.Main.Run(ctx)
}
