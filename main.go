package main

import (
	"github.com/gogf/gf/v2/os/gtime"

	_ "github.com/iimeta/fastapi-admin/internal/packed"

	_ "github.com/iimeta/fastapi-admin/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/iimeta/fastapi-admin/internal/cmd"
)

func main() {

	// 设置进程全局时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
