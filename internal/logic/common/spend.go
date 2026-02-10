package common

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/common"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// 记录花费
func RecordSpend(ctx context.Context, userId, appId int, appKey string, rid int, key string, spend common.Spend) error {

	now := gtime.TimestampMilli()
	defer func() {
		logger.Debugf(ctx, "RecordSpend time: %d", gtime.TimestampMilli()-now)
	}()

	totalSpendTokens := int(spend.TotalSpendTokens)

	if totalSpendTokens == 0 {
		return nil
	}

	logger.Infof(ctx, "RecordSpend rid: %d, userId: %d, appId: %d, appKey: %s, totalSpendTokens: %d, key: %s", rid, userId, appId, appKey, totalSpendTokens, key)

	usageKey := getUserUsageKey(userId)

	currentQuota, err := redisSpendQuota(ctx, usageKey, consts.USER_QUOTA_FIELD, totalSpendTokens)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = mongoSpendQuota(ctx, func() error {

		if err := dao.User.UpdateOne(ctx, bson.M{"user_id": userId}, bson.M{
			"$inc": bson.M{
				"quota":      -totalSpendTokens,
				"used_quota": totalSpendTokens,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		return nil
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action: consts.ACTION_CACHE,
		NewData: &model.UserQuota{
			UserId:       userId,
			CurrentQuota: currentQuota,
		},
	}); err != nil {
		logger.Error(ctx, err)
	}

	if rid != 0 {

		currentQuota, err = redisSpendQuota(ctx, getResellerUsageKey(rid), consts.RESELLER_QUOTA_FIELD, totalSpendTokens)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = mongoSpendQuota(ctx, func() error {

			if err := dao.Reseller.UpdateOne(ctx, bson.M{"user_id": userId}, bson.M{
				"$inc": bson.M{
					"quota":      -totalSpendTokens,
					"used_quota": totalSpendTokens,
				},
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}

			return nil
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	app, err := dao.App.FindOne(ctx, bson.M{"app_id": appId})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if app.IsLimitQuota {

		currentQuota, err = redisSpendQuota(ctx, usageKey, getAppTotalTokensField(appId), totalSpendTokens)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = mongoSpendQuota(ctx, func() error {

			if err := dao.App.UpdateOne(ctx, bson.M{"app_id": appId}, bson.M{
				"$inc": bson.M{
					"quota":      -totalSpendTokens,
					"used_quota": totalSpendTokens,
				},
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}

			return nil
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

	} else {
		if err = mongoUsedQuota(ctx, func() error {

			if err := dao.App.UpdateOne(ctx, bson.M{"app_id": appId}, bson.M{
				"$inc": bson.M{
					"used_quota": totalSpendTokens,
				},
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}

			return nil
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	appkey, err := dao.AppKey.FindOne(ctx, bson.M{"key": appKey})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if appkey.IsLimitQuota {

		currentQuota, err = redisSpendQuota(ctx, usageKey, getAppKeyTotalTokensField(appId, appKey), totalSpendTokens)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = mongoSpendQuota(ctx, func() error {

			if err := dao.AppKey.UpdateOne(ctx, bson.M{"key": appKey}, bson.M{
				"$inc": bson.M{
					"quota":      -totalSpendTokens,
					"used_quota": totalSpendTokens,
				},
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}

			return nil
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

	} else {
		if err = mongoUsedQuota(ctx, func() error {

			if err := dao.AppKey.UpdateOne(ctx, bson.M{"key": appKey}, bson.M{
				"$inc": bson.M{
					"used_quota": totalSpendTokens,
				},
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}

			return nil
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if err = mongoUsedQuota(ctx, func() error {

		if err := dao.Key.UpdateOne(ctx, bson.M{"key": key}, bson.M{
			"$inc": bson.M{
				"used_quota": totalSpendTokens,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		return nil
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if spend.GroupId != "" {

		group, err := dao.Group.FindById(ctx, spend.GroupId)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if group.IsLimitQuota {

			currentQuota, err = redisSpendQuota(ctx, consts.API_GROUP_USAGE_KEY, group.Id, totalSpendTokens)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			if err = mongoSpendQuota(ctx, func() error {

				if err := dao.Group.UpdateById(ctx, group.Id, bson.M{
					"$inc": bson.M{
						"quota":      -totalSpendTokens,
						"used_quota": totalSpendTokens,
					},
				}); err != nil {
					logger.Error(ctx, err)
					return err
				}

				return nil
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}

		} else {
			if err = mongoUsedQuota(ctx, func() error {

				if err := dao.Group.UpdateById(ctx, group.Id, bson.M{
					"$inc": bson.M{
						"used_quota": totalSpendTokens,
					},
				}); err != nil {
					logger.Error(ctx, err)
					return err
				}

				return nil
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

func redisSpendQuota(ctx context.Context, usageKey, field string, totalTokens int, retry ...int) (int, error) {

	currentQuota, err := redis.HIncrBy(ctx, usageKey, field, int64(-totalTokens))
	if err != nil {
		logger.Errorf(ctx, "redisSpendQuota usageKey: %s, field: %s, totalTokens: %d, error: %v", usageKey, field, totalTokens, err)

		if len(retry) == 10 {
			return -1, err
		}

		retry = append(retry, 1)

		time.Sleep(time.Duration(len(retry)*5) * time.Second)

		logger.Errorf(ctx, "redisSpendQuota usageKey: %s, field: %s, totalTokens: %d, retry: %d", usageKey, field, totalTokens, len(retry))

		return redisSpendQuota(ctx, usageKey, field, totalTokens, retry...)
	}

	return int(currentQuota), nil
}

func mongoSpendQuota(ctx context.Context, f func() error, retry ...int) error {

	if err := f(); err != nil {
		logger.Errorf(ctx, "mongoSpendQuota error: %v", err)

		if len(retry) == 10 {
			return err
		}

		retry = append(retry, 1)

		time.Sleep(time.Duration(len(retry)*5) * time.Second)

		logger.Errorf(ctx, "mongoSpendQuota retry: %d", len(retry))

		return mongoSpendQuota(ctx, f, retry...)
	}

	return nil
}

func mongoUsedQuota(ctx context.Context, f func() error, retry ...int) error {

	if err := f(); err != nil {
		logger.Errorf(ctx, "mongoUsedQuota error: %v", err)

		if len(retry) == 10 {
			return err
		}

		retry = append(retry, 1)

		time.Sleep(time.Duration(len(retry)*5) * time.Second)

		logger.Errorf(ctx, "mongoUsedQuota retry: %d", len(retry))

		return mongoUsedQuota(ctx, f, retry...)
	}

	return nil
}

func getUserUsageKey(userId int) string {
	return fmt.Sprintf(consts.API_USER_USAGE_KEY, userId)
}

func getAppTotalTokensField(appId int) string {
	return fmt.Sprintf(consts.APP_QUOTA_FIELD, appId)
}

func getAppKeyTotalTokensField(appId int, appKey string) string {
	return fmt.Sprintf(consts.APP_KEY_QUOTA_FIELD, appId, appKey)
}

func getResellerUsageKey(rid int) string {
	return fmt.Sprintf(consts.API_RESELLER_USAGE_KEY, rid)
}
