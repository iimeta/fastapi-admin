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
	RESELLER_ID_AUTO_INCREMENT_CFG = "core.reseller_id_auto_increment"
	RESELLER_ID_AUTO_INCREMENT_KEY = "CORE:RESELLER_ID_AUTO_INCREMENT"
	USER_ID_AUTO_INCREMENT_CFG     = "core.user_id_auto_increment"
	USER_ID_AUTO_INCREMENT_KEY     = "CORE:USER_ID_AUTO_INCREMENT"
	APP_ID_AUTO_INCREMENT_CFG      = "core.app_id_auto_increment"
	APP_ID_AUTO_INCREMENT_KEY      = "CORE:APP_ID_AUTO_INCREMENT"
)

func init() {

	ctx := gctx.New()

	resellerId := config.GetInt(ctx, RESELLER_ID_AUTO_INCREMENT_CFG, 80000)
	if resellerId == 80000 {
		if maxResellerId := getMaxResellerId(ctx); maxResellerId != 0 {
			resellerId = maxResellerId
		}
	}

	// 自增起始resellerId
	_, _ = redis.SetNX(ctx, RESELLER_ID_AUTO_INCREMENT_KEY, resellerId)

	userId := config.GetInt(ctx, USER_ID_AUTO_INCREMENT_CFG, 10000)
	if userId == 10000 {
		if maxUserId := getMaxUserId(ctx); maxUserId != 0 {
			userId = maxUserId
		}
	}

	// 自增起始userId
	_, _ = redis.SetNX(ctx, USER_ID_AUTO_INCREMENT_KEY, userId)

	appId := config.GetInt(ctx, APP_ID_AUTO_INCREMENT_CFG, 10000)
	if appId == 10000 {
		if maxAppId := getMaxAppId(ctx); maxAppId != 0 {
			appId = maxAppId
		}
	}

	// 自增起始appId
	_, _ = redis.SetNX(ctx, APP_ID_AUTO_INCREMENT_KEY, appId)
}

func IncrResellerId(ctx context.Context) int {

	reply, err := redis.Incr(ctx, RESELLER_ID_AUTO_INCREMENT_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return 0
	}

	return int(reply)
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

// 获取最大代理商ID
func getMaxResellerId(ctx context.Context) int {

	if user, _ := dao.Reseller.FindOne(ctx, bson.M{}, &dao.FindOptions{SortFields: []string{"-user_id"}}); user != nil {
		return user.UserId
	}

	return 0
}

// 获取最大用户ID
func getMaxUserId(ctx context.Context) int {

	if user, _ := dao.User.FindOne(ctx, bson.M{}, &dao.FindOptions{SortFields: []string{"-user_id"}}); user != nil {
		return user.UserId
	}

	return 0
}

// 获取最大应用ID
func getMaxAppId(ctx context.Context) int {

	if app, _ := dao.App.FindOne(ctx, bson.M{}, &dao.FindOptions{SortFields: []string{"-app_id"}}); app != nil {
		return app.AppId
	}

	return 0
}
