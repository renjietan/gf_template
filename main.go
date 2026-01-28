package main

import (
	_ "github.com/gogf/gf/v2" // 导入 MySQL 驱动
	"github.com/gogf/gf/v2/os/gctx"

	"gf_template/internal/cmd"
	_ "gf_template/internal/packed"
)

func main() {
	var ctx = gctx.GetInitCtx()
	cmd.Main.Run(ctx)
}
