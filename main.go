package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/v2/internal/cmd"
	_ "github.com/iimeta/fastapi-admin/v2/internal/logic"
	_ "github.com/iimeta/fastapi-admin/v2/internal/task"
)

func main() {

	// 设置进程全局时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
