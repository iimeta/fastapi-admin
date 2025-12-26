package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/service"
)

var (
	batchCron  string
	batchEntry *gcron.Entry
)

// 批处理任务
func batchTask(ctx context.Context) {

	if config.Cfg.BatchTask.Open {

		if batchCron != config.Cfg.BatchTask.Cron {
			batchCron = config.Cfg.BatchTask.Cron
			if batchEntry != nil {
				batchEntry.Stop()
			}
		} else if batchEntry != nil {
			return
		}

		batchEntry, _ = gcron.AddSingleton(ctx, config.Cfg.BatchTask.Cron, func(ctx context.Context) {

			service.TaskBatch().Task(gctx.New())

		})

	} else {
		if batchEntry != nil {
			batchEntry.Stop()
			batchCron = ""
			batchEntry = nil
		}
	}
}
