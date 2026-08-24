package admin_user

import (
	"context"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// 用户过期任务
func (s *sAdminUser) ExpireTask(ctx context.Context) {

	logger.Info(ctx, "sAdminUser ExpireTask start")

	now := gtime.TimestampMilli()

	if config.Cfg.UserExpireTask == nil {
		logger.Debugf(ctx, "sAdminUser ExpireTask end time: %d", gtime.TimestampMilli()-now)
		return
	}

	mutex := s.expireRedsync.NewMutex(consts.TASK_USER_EXPIRE_LOCK_KEY, redsync.WithExpiry(config.Cfg.UserExpireTask.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sAdminUser ExpireTask", err)
		logger.Debugf(ctx, "sAdminUser ExpireTask end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sAdminUser ExpireTask lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sAdminUser ExpireTask unlock")
		}
		logger.Debugf(ctx, "sAdminUser ExpireTask end time: %d", gtime.TimestampMilli()-now)
	}()

	users, err := dao.User.Find(ctx, bson.M{
		"status": 1,
		"expires_at": bson.M{
			"$gt":  0,
			"$lte": now,
		},
	})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	for _, user := range users {
		if err := s.ChangeStatus(ctx, model.UserChangeStatusReq{
			Id:     user.Id,
			Status: 2,
		}); err != nil {
			logger.Error(ctx, err)
			continue
		}
		logger.Infof(ctx, "sAdminUser ExpireTask disable expired user: %d", user.UserId)
	}

	if _, err := redis.Set(ctx, consts.TASK_USER_EXPIRE_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
