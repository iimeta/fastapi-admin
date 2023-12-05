package main

import (
	_ "github.com/iimeta/fastapi-admin/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/iimeta/fastapi-admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
