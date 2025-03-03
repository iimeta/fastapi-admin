package task

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/service"
)

var (
	noticeCron  string
	noticeEntry *gcron.Entry
)

// 通知任务
func noticeTask(ctx context.Context) {

	if config.Cfg.Notice.Open {

		if noticeCron != config.Cfg.Notice.Cron {
			noticeCron = config.Cfg.Notice.Cron
			if noticeEntry != nil {
				noticeEntry.Stop()
			}
		} else if noticeEntry != nil {
			return
		}

		noticeEntry, _ = gcron.AddSingleton(ctx, config.Cfg.Notice.Cron, func(ctx context.Context) {
			service.Notice().QuotaWarningTask(gctx.New())
		})

	} else {
		if noticeEntry != nil {
			noticeEntry.Stop()
			noticeCron = ""
			noticeEntry = nil
		}
	}
}
