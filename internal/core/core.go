package core

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	USER_ID_AUTO_INCREMENT_CFG = "core.user_id_auto_increment"
	USER_ID_AUTO_INCREMENT_KEY = "CORE:USER_ID_AUTO_INCREMENT"
	APP_ID_AUTO_INCREMENT_CFG  = "core.app_id_auto_increment"
	APP_ID_AUTO_INCREMENT_KEY  = "CORE:APP_ID_AUTO_INCREMENT"
)

func init() {

	ctx := gctx.New()

	userId := config.GetInt(ctx, USER_ID_AUTO_INCREMENT_CFG, 10000)
	if userId == 10000 {
		if maxUserId := getMaxUserId(ctx); maxUserId != 0 {
			userId = maxUserId
		}
	}

	// 自增起始UserId
	_, _ = redis.SetNX(ctx, USER_ID_AUTO_INCREMENT_KEY, userId)

	appId := config.GetInt(ctx, APP_ID_AUTO_INCREMENT_CFG, 10000)
	if appId == 10000 {
		if maxAppId := getMaxAppId(ctx); maxAppId != 0 {
			appId = maxAppId
		}
	}

	// 自增起始AppId
	_, _ = redis.SetNX(ctx, APP_ID_AUTO_INCREMENT_KEY, appId)

}

func IncrUserId(ctx context.Context) int {

	reply, err := redis.Incr(ctx, USER_ID_AUTO_INCREMENT_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return 0
	}

	return int(reply)
}

func IncrAppId(ctx context.Context) int {

	reply, err := redis.Incr(ctx, APP_ID_AUTO_INCREMENT_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return 0
	}

	return int(reply)
}

// 获取最大用户ID
func getMaxUserId(ctx context.Context) int {

	if user, _ := dao.User.FindOne(ctx, bson.M{}, "-user_id"); user != nil {
		return user.UserId
	}

	return 0
}

// 获取最大应用ID
func getMaxAppId(ctx context.Context) int {

	if app, _ := dao.App.FindOne(ctx, bson.M{}, "-app_id"); app != nil {
		return app.AppId
	}

	return 0
}
