package task

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	logDelCron  string
	logDelEntry *gcron.Entry
)

// 日志删除任务
func logDelTask(ctx context.Context) {

	if config.Cfg.Log.Open {

		if logDelCron != config.Cfg.Log.Cron {
			logDelCron = config.Cfg.Log.Cron
			if logDelEntry != nil {
				logDelEntry.Stop()
			}
		} else if logDelEntry != nil {
			return
		}

		logDelEntry, _ = gcron.AddSingleton(ctx, config.Cfg.Log.Cron, func(ctx context.Context) {

			if config.Cfg.Log.ChatReserve > 0 {
				if deletedCount, err := dao.Chat.DeleteMany(ctx, bson.M{
					"req_time": bson.M{
						"$lte": gtime.Now().StartOfDay().TimestampMilli() - (config.Cfg.Log.ChatReserve * gtime.D).Milliseconds() - 1,
					},
				}); err == nil {
					logger.Infof(ctx, "聊天日志已删除 %d 条记录", deletedCount)
				} else {
					logger.Error(ctx, err)
				}
			}

			if config.Cfg.Log.ImageReserve > 0 {
				if deletedCount, err := dao.Image.DeleteMany(ctx, bson.M{
					"req_time": bson.M{
						"$lte": gtime.Now().StartOfDay().TimestampMilli() - (config.Cfg.Log.ImageReserve * gtime.D).Milliseconds() - 1,
					},
				}); err == nil {
					logger.Infof(ctx, "绘图日志已删除 %d 条记录", deletedCount)
				} else {
					logger.Error(ctx, err)
				}
			}

			if config.Cfg.Log.AudioReserve > 0 {
				if deletedCount, err := dao.Audio.DeleteMany(ctx, bson.M{
					"req_time": bson.M{
						"$lte": gtime.Now().StartOfDay().TimestampMilli() - (config.Cfg.Log.AudioReserve * gtime.D).Milliseconds() - 1,
					},
				}); err == nil {
					logger.Infof(ctx, "音频日志已删除 %d 条记录", deletedCount)
				} else {
					logger.Error(ctx, err)
				}
			}
		})

	} else {
		if logDelEntry != nil {
			logDelEntry.Stop()
			logDelCron = ""
		}
	}
}
