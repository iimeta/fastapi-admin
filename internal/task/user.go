package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

var (
	userExpireCron  string
	userExpireEntry *gcron.Entry
)

// 用户过期任务
func userExpireTask(ctx context.Context) {

	if config.Cfg.UserExpireTask != nil && config.Cfg.UserExpireTask.Open {

		if userExpireCron != config.Cfg.UserExpireTask.Cron {
			userExpireCron = config.Cfg.UserExpireTask.Cron
			if userExpireEntry != nil {
				userExpireEntry.Stop()
			}
		} else if userExpireEntry != nil {
			return
		}

		userExpireEntry, _ = gcron.AddSingleton(ctx, config.Cfg.UserExpireTask.Cron, func(ctx context.Context) {
			service.AdminUser().ExpireTask(gctx.New())
		})

	} else {
		if userExpireEntry != nil {
			userExpireEntry.Stop()
			userExpireCron = ""
			userExpireEntry = nil
		}
	}
}
