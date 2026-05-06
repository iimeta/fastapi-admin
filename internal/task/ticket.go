package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

var (
	ticketCron  string
	ticketEntry *gcron.Entry
)

// 工单任务
func ticketTask(ctx context.Context) {

	if config.Cfg.Ticket != nil && config.Cfg.Ticket.Open {

		if ticketCron != config.Cfg.Ticket.Cron {
			ticketCron = config.Cfg.Ticket.Cron
			if ticketEntry != nil {
				ticketEntry.Stop()
			}
		} else if ticketEntry != nil {
			return
		}

		ticketEntry, _ = gcron.AddSingleton(ctx, config.Cfg.Ticket.Cron, func(ctx context.Context) {
			service.Ticket().AutoCloseTask(gctx.New())
		})

	} else {
		if ticketEntry != nil {
			ticketEntry.Stop()
			ticketCron = ""
			ticketEntry = nil
		}
	}
}
