package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/service"
)

var (
	fileCron  string
	fileEntry *gcron.Entry
)

// 文件任务
func fileTask(ctx context.Context) {

	if config.Cfg.FileTask.Open {

		if fileCron != config.Cfg.FileTask.Cron {
			fileCron = config.Cfg.FileTask.Cron
			if fileEntry != nil {
				fileEntry.Stop()
			}
		} else if fileEntry != nil {
			return
		}

		fileEntry, _ = gcron.AddSingleton(ctx, config.Cfg.FileTask.Cron, func(ctx context.Context) {

			service.TaskFile().Task(gctx.New())

		})

	} else {
		if fileEntry != nil {
			fileEntry.Stop()
			fileCron = ""
			fileEntry = nil
		}
	}
}
