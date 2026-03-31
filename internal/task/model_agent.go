package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

var (
	modelAgentHealthCheckCron  string
	modelAgentHealthCheckEntry *gcron.Entry
)

// 模型代理健康检查任务
func modelAgentHealthCheckTask(ctx context.Context) {

	if config.Cfg.ModelAgentHealthCheckTask.Open {

		if modelAgentHealthCheckCron != config.Cfg.ModelAgentHealthCheckTask.Cron {
			modelAgentHealthCheckCron = config.Cfg.ModelAgentHealthCheckTask.Cron
			if modelAgentHealthCheckEntry != nil {
				modelAgentHealthCheckEntry.Stop()
			}
		} else if modelAgentHealthCheckEntry != nil {
			return
		}

		modelAgentHealthCheckEntry, _ = gcron.AddSingleton(ctx, config.Cfg.ModelAgentHealthCheckTask.Cron, func(ctx context.Context) {
			service.ModelAgent().HealthCheckTask(gctx.New())
		})

	} else {
		if modelAgentHealthCheckEntry != nil {
			modelAgentHealthCheckEntry.Stop()
			modelAgentHealthCheckCron = ""
			modelAgentHealthCheckEntry = nil
		}
	}
}
