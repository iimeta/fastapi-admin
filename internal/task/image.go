package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

var (
	imageCron  string
	imageEntry *gcron.Entry
)

// 绘图任务
func imageTask(ctx context.Context) {

	if config.Cfg.ImageTask.Open {

		if imageCron != config.Cfg.ImageTask.Cron {
			imageCron = config.Cfg.ImageTask.Cron
			if imageEntry != nil {
				imageEntry.Stop()
			}
		} else if imageEntry != nil {
			return
		}

		imageEntry, _ = gcron.AddSingleton(ctx, config.Cfg.ImageTask.Cron, func(ctx context.Context) {
			service.TaskImage().Task(gctx.New())
		})

	} else {
		if imageEntry != nil {
			imageEntry.Stop()
			imageCron = ""
			imageEntry = nil
		}
	}
}
