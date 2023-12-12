package core

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
)

const (
	USER_ID_AUTO_INCREMENT_KEY = "CORE:USER_ID_AUTO_INCREMENT"
	APP_ID_AUTO_INCREMENT_KEY  = "CORE:APP_ID_AUTO_INCREMENT"
)

const (
	USER_ID_AUTO_INCREMENT_CFG = "CORE.USER_ID_AUTO_INCREMENT"
	APP_ID_AUTO_INCREMENT_CFG  = "CORE.APP_ID_AUTO_INCREMENT"
)

func init() {

	ctx := gctx.New()

	// 默认自增起始UserId
	_, _ = redis.SetNX(ctx, USER_ID_AUTO_INCREMENT_KEY, config.GetInt(ctx, USER_ID_AUTO_INCREMENT_CFG, 10000))

	// 默认自增起始AppId
	_, _ = redis.SetNX(ctx, APP_ID_AUTO_INCREMENT_KEY, config.GetInt(ctx, APP_ID_AUTO_INCREMENT_CFG, 10000))
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
