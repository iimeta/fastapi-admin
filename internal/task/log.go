package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

var (
	logDelCron  string
	logDelEntry *gcron.Entry
)

// 日志删除任务
func logDelTask(ctx context.Context) {

	if config.Cfg.Log.Open {

		if logDelCron != config.Cfg.Log.Cron {
			logDelCron = config.Cfg.Log.Cron
			if logDelEntry != nil {
				logDelEntry.Stop()
			}
		} else if logDelEntry != nil {
			return
		}

		logDelEntry, _ = gcron.AddSingleton(ctx, config.Cfg.Log.Cron, func(ctx context.Context) {
			service.Log().DelTask(gctx.New())
		})

	} else {
		if logDelEntry != nil {
			logDelEntry.Stop()
			logDelCron = ""
			logDelEntry = nil
		}
	}
}
