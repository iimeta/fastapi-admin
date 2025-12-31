package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

var (
	errorCheckEntryList []*gcron.Entry
)

// 错误检查任务
func errorCheckTask(ctx context.Context) {

	if config.Cfg.AutoEnableError.Open {

		if errorCheckEntryList != nil {
			for _, errorCheckEntry := range errorCheckEntryList {
				errorCheckEntry.Stop()
			}
		}

		if len(config.Cfg.AutoEnableError.EnableErrors) > 0 {

			errorCheckEntryList = make([]*gcron.Entry, 0)
			for _, enableError := range config.Cfg.AutoEnableError.EnableErrors {

				entry, _ := gcron.AddSingleton(ctx, enableError.Cron, func(ctx context.Context) {
					service.Key().CheckTask(gctx.New(), enableError)
				})

				errorCheckEntryList = append(errorCheckEntryList, entry)
			}
		}

	} else {
		if errorCheckEntryList != nil {
			for _, errorCheckEntry := range errorCheckEntryList {
				errorCheckEntry.Stop()
			}
			errorCheckEntryList = nil
		}
	}
}
