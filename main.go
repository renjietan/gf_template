package main

import (
	"github.com/gogf/gf/v2/os/gctx"

	"gf_template/internal/cmd"
	_ "gf_template/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
