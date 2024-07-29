package task

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func init() {
	// 定时统计任务
	if config.Cfg.Statistics.Cron != "" {
		_, _ = gcron.AddSingleton(gctx.New(), config.Cfg.Statistics.Cron, func(ctx context.Context) {
			service.Statistics().StatisticsTask(ctx)
		})
	}
}
