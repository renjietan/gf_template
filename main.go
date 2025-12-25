package main

import (
	_ "gf_template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf_template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
