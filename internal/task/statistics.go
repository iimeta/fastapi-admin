package task

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/service"
)

var (
	statisticsCron  string
	statisticsEntry *gcron.Entry
)

// 统计任务
func statisticsTask(ctx context.Context) {

	if config.Cfg.Statistics.Open {

		if statisticsCron != config.Cfg.Statistics.Cron {
			statisticsCron = config.Cfg.Statistics.Cron
			if statisticsEntry != nil {
				statisticsEntry.Stop()
			}
		} else if statisticsEntry != nil {
			return
		}

		statisticsEntry, _ = gcron.AddSingleton(ctx, config.Cfg.Statistics.Cron, func(ctx context.Context) {
			service.Statistics().StatisticsTask(gctx.New())
		})

	} else {
		if statisticsEntry != nil {
			statisticsEntry.Stop()
			statisticsCron = ""
			statisticsEntry = nil
		}
	}
}
