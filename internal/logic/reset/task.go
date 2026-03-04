package reset

import (
	"context"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
)

// 重置任务
func (s *sReset) Task(ctx context.Context) {

	logger.Info(ctx, "sReset Task start")

	now := gtime.TimestampMilli()

	mutex := s.resetRedsync.NewMutex(consts.TASK_RESET_LOCK_KEY, redsync.WithExpiry(config.Cfg.ResetTask.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sReset Task", err)
		logger.Debugf(ctx, "sReset Task end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sReset Task lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sReset Task unlock")
		}
		logger.Debugf(ctx, "sReset Task end time: %d", gtime.TimestampMilli()-now)
	}()

	/////

	if _, err := redis.Set(ctx, consts.TASK_RESET_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
