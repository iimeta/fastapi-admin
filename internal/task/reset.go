package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

var (
	resetCron  string
	resetEntry *gcron.Entry
)

// 重置任务
func resetTask(ctx context.Context) {

	if config.Cfg.ResetTask.Open {

		if resetCron != config.Cfg.ResetTask.Cron {
			resetCron = config.Cfg.ResetTask.Cron
			if resetEntry != nil {
				resetEntry.Stop()
			}
		} else if resetEntry != nil {
			return
		}

		resetEntry, _ = gcron.AddSingleton(ctx, config.Cfg.ResetTask.Cron, func(ctx context.Context) {
			service.Reset().Task(gctx.New())
		})

	} else {
		if resetEntry != nil {
			resetEntry.Stop()
			resetCron = ""
			resetEntry = nil
		}
	}
}
