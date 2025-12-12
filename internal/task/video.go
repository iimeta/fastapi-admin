package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/service"
)

var (
	videoCron  string
	videoEntry *gcron.Entry
)

// 视频任务
func videoTask(ctx context.Context) {

	if config.Cfg.VideoTask.Open {

		if videoCron != config.Cfg.VideoTask.Cron {
			videoCron = config.Cfg.VideoTask.Cron
			if videoEntry != nil {
				videoEntry.Stop()
			}
		} else if videoEntry != nil {
			return
		}

		videoEntry, _ = gcron.AddSingleton(ctx, config.Cfg.VideoTask.Cron, func(ctx context.Context) {

			service.TaskVideo().Task(gctx.New())

		})

	} else {
		if videoEntry != nil {
			videoEntry.Stop()
			videoCron = ""
			videoEntry = nil
		}
	}
}
