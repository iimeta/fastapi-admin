package task

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

var (
	imageStorageCron  string
	imageStorageEntry *gcron.Entry
)

// 绘图转储过期清理任务
func imageStorageTask(ctx context.Context) {

	if config.Cfg.ImageStorage.Open {

		if imageStorageCron != config.Cfg.ImageStorage.Cron {
			imageStorageCron = config.Cfg.ImageStorage.Cron
			if imageStorageEntry != nil {
				imageStorageEntry.Stop()
			}
		} else if imageStorageEntry != nil {
			return
		}

		imageStorageEntry, _ = gcron.AddSingleton(ctx, config.Cfg.ImageStorage.Cron, func(ctx context.Context) {
			service.LogImage().StorageCleanTask(gctx.New())
		})

	} else {
		if imageStorageEntry != nil {
			imageStorageEntry.Stop()
			imageStorageCron = ""
			imageStorageEntry = nil
		}
	}
}
