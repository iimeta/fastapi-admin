package core

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
)

const (
	USER_ID_AUTO_INCREMENT_KEY   = "CORE:USER_ID_AUTO_INCREMENT"
	GROUP_ID_AUTO_INCREMENT_KEY  = "CORE:GROUP_ID_AUTO_INCREMENT"
	RECORD_ID_AUTO_INCREMENT_KEY = "CORE:RECORD_ID_AUTO_INCREMENT"
)

const (
	USER_ID_AUTO_INCREMENT_CFG   = "CORE.USER_ID_AUTO_INCREMENT"
	GROUP_ID_AUTO_INCREMENT_CFG  = "CORE.GROUP_ID_AUTO_INCREMENT"
	RECORD_ID_AUTO_INCREMENT_CFG = "CORE.RECORD_ID_AUTO_INCREMENT"
)

func init() {

	ctx := gctx.New()

	// 默认自增起始UserId
	_, _ = redis.SetNX(ctx, USER_ID_AUTO_INCREMENT_KEY, config.GetInt(ctx, USER_ID_AUTO_INCREMENT_CFG, 10000))

	// 默认自增起始GroupId
	_, _ = redis.SetNX(ctx, GROUP_ID_AUTO_INCREMENT_KEY, config.GetInt(ctx, GROUP_ID_AUTO_INCREMENT_CFG, 10000))

	// 默认自增起始RecordId
	_, _ = redis.SetNX(ctx, RECORD_ID_AUTO_INCREMENT_KEY, config.GetInt(ctx, RECORD_ID_AUTO_INCREMENT_CFG))
}

func IncrUserId(ctx context.Context) int {

	reply, err := redis.Incr(ctx, USER_ID_AUTO_INCREMENT_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return 0
	}

	return int(reply)
}

func IncrGroupId(ctx context.Context) int {

	reply, err := redis.Incr(ctx, GROUP_ID_AUTO_INCREMENT_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return 0
	}

	return int(reply)
}

func IncrRecordId(ctx context.Context) int {

	reply, err := redis.Incr(ctx, RECORD_ID_AUTO_INCREMENT_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return 0
	}

	return int(reply)
}
