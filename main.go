package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"gf_template/internal/cmd"
	_ "gf_template/internal/packed"

	_ "github.com/gogf/gf/v2" // 导入 MySQL 驱动
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
