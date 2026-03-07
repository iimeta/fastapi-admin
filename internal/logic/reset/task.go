package reset

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
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

	s.resetUser(ctx, now)
	s.resetReseller(ctx, now)
	s.resetApp(ctx, now)
	s.resetAppKey(ctx, now)
	s.resetGroup(ctx, now)

	if _, err := redis.Set(ctx, consts.TASK_RESET_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sReset) resetUser(ctx context.Context, now int64) {

	users, err := dao.User.Find(ctx, bson.M{"is_cycle_reset_quota": true, "status": 1, "next_reset_at": bson.M{"$gt": 0, "$lte": now}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	for _, user := range users {

		oldData := *user
		nextResetAt := common.CalcNextNaturalResetAt(time.UnixMilli(now).In(util.Location), user.CyclePeriod, user.PeriodUnit)

		newData, err := dao.User.FindOneAndUpdateById(ctx, user.Id, bson.M{
			"quota":         user.ResetQuota,
			"used_quota":    0,
			"reset_at":      now,
			"next_reset_at": nextResetAt,
		})
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		if _, err = redis.HSet(ctx, fmt.Sprintf(consts.API_USER_USAGE_KEY, user.UserId), g.Map{
			consts.USER_QUOTA_FIELD: user.ResetQuota,
		}); err != nil {
			logger.Error(ctx, err)
		}

		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
			Action:  consts.ACTION_UPDATE,
			OldData: &oldData,
			NewData: newData,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}
}

func (s *sReset) resetReseller(ctx context.Context, now int64) {

	resellers, err := dao.Reseller.Find(ctx, bson.M{"is_cycle_reset_quota": true, "status": 1, "next_reset_at": bson.M{"$gt": 0, "$lte": now}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	for _, reseller := range resellers {

		oldData := *reseller
		nextResetAt := common.CalcNextNaturalResetAt(time.UnixMilli(now).In(util.Location), reseller.CyclePeriod, reseller.PeriodUnit)

		newData, err := dao.Reseller.FindOneAndUpdateById(ctx, reseller.Id, bson.M{
			"quota":         reseller.ResetQuota,
			"used_quota":    0,
			"reset_at":      now,
			"next_reset_at": nextResetAt,
		})
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		if _, err = redis.HSet(ctx, fmt.Sprintf(consts.API_RESELLER_USAGE_KEY, reseller.UserId), g.Map{
			consts.RESELLER_QUOTA_FIELD: reseller.ResetQuota,
		}); err != nil {
			logger.Error(ctx, err)
		}

		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
			Action:  consts.ACTION_UPDATE,
			OldData: &oldData,
			NewData: newData,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}
}

func (s *sReset) resetApp(ctx context.Context, now int64) {

	apps, err := dao.App.Find(ctx, bson.M{"is_cycle_reset_quota": true, "status": 1, "next_reset_at": bson.M{"$gt": 0, "$lte": now}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	for _, app := range apps {

		oldData := *app
		nextResetAt := common.CalcNextNaturalResetAt(time.UnixMilli(now).In(util.Location), app.CyclePeriod, app.PeriodUnit)

		newData, err := dao.App.FindOneAndUpdateById(ctx, app.Id, bson.M{
			"quota":         app.ResetQuota,
			"used_quota":    0,
			"reset_at":      now,
			"next_reset_at": nextResetAt,
		})
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		if _, err = redis.HSet(ctx, fmt.Sprintf(consts.API_USER_USAGE_KEY, app.UserId), g.Map{
			fmt.Sprintf(consts.APP_QUOTA_FIELD, app.AppId):          app.ResetQuota,
			fmt.Sprintf(consts.APP_IS_LIMIT_QUOTA_FIELD, app.AppId): app.IsLimitQuota,
		}); err != nil {
			logger.Error(ctx, err)
		}

		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP, model.PubMessage{
			Action:  consts.ACTION_UPDATE,
			OldData: &oldData,
			NewData: newData,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}
}

func (s *sReset) resetAppKey(ctx context.Context, now int64) {

	keys, err := dao.AppKey.Find(ctx, bson.M{"is_cycle_reset_quota": true, "status": 1, "next_reset_at": bson.M{"$gt": 0, "$lte": now}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	for _, key := range keys {

		oldData := *key
		nextResetAt := common.CalcNextNaturalResetAt(time.UnixMilli(now).In(util.Location), key.CyclePeriod, key.PeriodUnit)

		newData, err := dao.AppKey.FindOneAndUpdateById(ctx, key.Id, bson.M{
			"quota":         key.ResetQuota,
			"used_quota":    0,
			"reset_at":      now,
			"next_reset_at": nextResetAt,
		})
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		if _, err = redis.HSet(ctx, fmt.Sprintf(consts.API_USER_USAGE_KEY, key.UserId), g.Map{
			fmt.Sprintf(consts.APP_KEY_QUOTA_FIELD, key.AppId, key.Key): key.ResetQuota,
		}); err != nil {
			logger.Error(ctx, err)
		}

		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP_KEY, model.PubMessage{
			Action:  consts.ACTION_UPDATE,
			OldData: &oldData,
			NewData: newData,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}
}

func (s *sReset) resetGroup(ctx context.Context, now int64) {

	groups, err := dao.Group.Find(ctx, bson.M{"is_cycle_reset_quota": true, "status": 1, "next_reset_at": bson.M{"$gt": 0, "$lte": now}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	for _, group := range groups {

		oldData := *group
		nextResetAt := common.CalcNextNaturalResetAt(time.UnixMilli(now).In(util.Location), group.CyclePeriod, group.PeriodUnit)

		newData, err := dao.Group.FindOneAndUpdateById(ctx, group.Id, bson.M{
			"quota":         group.ResetQuota,
			"used_quota":    0,
			"reset_at":      now,
			"next_reset_at": nextResetAt,
		})
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		if _, err = redis.HSetInt(ctx, consts.API_GROUP_USAGE_KEY, group.Id, group.ResetQuota); err != nil {
			logger.Error(ctx, err)
		}

		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
			Action:  consts.ACTION_UPDATE,
			OldData: &oldData,
			NewData: newData,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}
}
