package main

import (
	_ "PersonalMain/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"PersonalMain/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
