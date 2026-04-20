package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

var (
	modelAgentSessionKeepCron  string
	modelAgentSessionKeepEntry *gcron.Entry
)

// 会话保持清理任务
func sessionKeepCleanupTask(ctx context.Context) {

	if config.Cfg.ModelAgentSessionKeepTask.Open {

		if modelAgentSessionKeepCron != config.Cfg.ModelAgentSessionKeepTask.Cron {
			modelAgentSessionKeepCron = config.Cfg.ModelAgentSessionKeepTask.Cron
			if modelAgentSessionKeepEntry != nil {
				modelAgentSessionKeepEntry.Stop()
			}
		} else if modelAgentSessionKeepEntry != nil {
			return
		}

		modelAgentSessionKeepEntry, _ = gcron.AddSingleton(ctx, config.Cfg.ModelAgentSessionKeepTask.Cron, func(ctx context.Context) {
			service.ModelAgent().SessionKeepCleanupTask(gctx.New())
		})

	} else {
		if modelAgentSessionKeepEntry != nil {
			modelAgentSessionKeepEntry.Stop()
			modelAgentSessionKeepCron = ""
			modelAgentSessionKeepEntry = nil
		}
	}
}
