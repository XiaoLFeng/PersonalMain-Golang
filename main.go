package main

import (
	_ "PersonalMain/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"PersonalMain/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
