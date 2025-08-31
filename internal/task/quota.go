package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/service"
)

var (
	quotaCron  string
	quotaEntry *gcron.Entry
)

// 额度任务
func quotaTask(ctx context.Context) {

	if config.Cfg.QuotaTask.Open {

		if quotaCron != config.Cfg.QuotaTask.Cron {
			quotaCron = config.Cfg.QuotaTask.Cron
			if quotaEntry != nil {
				quotaEntry.Stop()
			}
		} else if quotaEntry != nil {
			return
		}

		quotaEntry, _ = gcron.AddSingleton(ctx, config.Cfg.QuotaTask.Cron, func(ctx context.Context) {

			if config.Cfg.Quota.Warning || config.Cfg.Quota.ExpiredWarning || config.Cfg.Quota.ExhaustedNotice || config.Cfg.Quota.ExpiredNotice {
				service.Quota().NoticeTask(gctx.New())
			}

			if config.Cfg.Quota.ExpiredClear {
				service.Quota().ClearTask(gctx.New())
			}
		})

	} else {
		if quotaEntry != nil {
			quotaEntry.Stop()
			quotaCron = ""
			quotaEntry = nil
		}
	}
}
