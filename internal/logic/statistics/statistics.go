package statistics

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model/common"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type sStatistics struct{}

var statisticsRedsync *redsync.Redsync

func init() {
	service.RegisterStatistics(New())
	statisticsRedsync = redsync.New(goredis.NewPool(redis.UniversalClient))
}

func New() service.IStatistics {
	return &sStatistics{}
}

// 统计任务
func (s *sStatistics) StatisticsTask(ctx context.Context) {

	logger.Info(ctx, "sStatistics StatisticsTask start")

	now := gtime.TimestampMilli()

	mutex := statisticsRedsync.NewMutex(consts.STATISTICS_LOCK_KEY, redsync.WithExpiry(time.Minute*config.Cfg.Statistics.LockMinutes))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, err)
		logger.Debugf(ctx, "sStatistics StatisticsTask end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sStatistics StatisticsTask lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sStatistics StatisticsTask unlock")
		}
		logger.Debugf(ctx, "sStatistics StatisticsTask end time: %d", gtime.TimestampMilli()-now)
	}()

	// 统计聊天数据
	s.StatisticsChat(ctx)
	// 统计绘图数据
	s.StatisticsImage(ctx)
	// 统计音频数据
	s.StatisticsAudio(ctx)

	if _, err := redis.Set(ctx, consts.STATISTICS_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}

// 统计聊天数据
func (s *sStatistics) StatisticsChat(ctx context.Context) {

	logger.Debugf(ctx, "sStatistics StatisticsChat start")

	now := gtime.TimestampMilli()
	defer func() {
		logger.Debugf(ctx, "sStatistics StatisticsChat end time: %d", gtime.TimestampMilli()-now)
	}()

	lastTime, err := redis.GetInt64(ctx, consts.STATISTICS_CHAT_LAST_TIME_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	lastId, err := redis.GetStr(ctx, consts.STATISTICS_CHAT_LAST_ID_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	filter := bson.M{
		"created_at": bson.M{
			"$gte": lastTime,
		},
		"is_smart_match": bson.M{"$ne": true},
		"is_retry":       bson.M{"$ne": true},
	}

	results, err := dao.Chat.FindByPage(ctx, &db.Paging{Page: 1, PageSize: config.Cfg.Statistics.Limit}, filter, "created_at_1_is_smart_match_-1_is_retry_-1", "created_at")
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	userMap := make(map[string]map[int]*entity.StatisticsUser)                     // map[req_date][user_id]entity.StatisticsUser
	userModelStatMap := make(map[string]map[int]map[string]*common.ModelStat)      // map[req_date][user_id][model_id]common.ModelStat
	appMap := make(map[string]map[int]*entity.StatisticsApp)                       // map[req_date][app_id]entity.StatisticsApp
	appModelStatMap := make(map[string]map[int]map[string]*common.ModelStat)       // map[req_date][app_id][model_id]common.ModelStat
	appKeyMap := make(map[string]map[string]*entity.StatisticsAppKey)              // map[req_date][app_key]entity.StatisticsAppKey
	appKeyModelStatMap := make(map[string]map[string]map[string]*common.ModelStat) // map[req_date][app_key][model_id]common.ModelStat

	for _, result := range results {

		if result.Id == lastId {
			continue
		}

		lastTime = result.CreatedAt
		lastId = result.Id

		if userMap[result.ReqDate] == nil {
			userMap[result.ReqDate] = make(map[int]*entity.StatisticsUser)
			userModelStatMap[result.ReqDate] = make(map[int]map[string]*common.ModelStat)
		}

		user := userMap[result.ReqDate][result.UserId]
		if user == nil {

			if user, err = dao.StatisticsUser.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "user_id": result.UserId}); err != nil {
				user = &entity.StatisticsUser{
					UserId:   result.UserId,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			userMap[result.ReqDate][result.UserId] = user

			if userModelStatMap[result.ReqDate][result.UserId] == nil {
				userModelStatMap[result.ReqDate][result.UserId] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range user.ModelStats {
				userModelStatMap[result.ReqDate][result.UserId][modelStat.ModelId] = modelStat
			}
		}

		userModelStat := userModelStatMap[result.ReqDate][result.UserId][result.ModelId]
		if userModelStat == nil {
			userModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			userModelStatMap[result.ReqDate][result.UserId][result.ModelId] = userModelStat
		}

		user.Total += 1
		user.Tokens += result.TotalTokens
		userModelStat.Total += 1
		userModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			user.Abnormal += 1
			user.AbnormalTokens += result.TotalTokens
			userModelStat.Abnormal += 1
			userModelStat.AbnormalTokens += result.TotalTokens
		}

		if appMap[result.ReqDate] == nil {
			appMap[result.ReqDate] = make(map[int]*entity.StatisticsApp)
			appModelStatMap[result.ReqDate] = make(map[int]map[string]*common.ModelStat)
		}

		app := appMap[result.ReqDate][result.AppId]
		if app == nil {

			if app, err = dao.StatisticsApp.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "app_id": result.AppId}); err != nil {
				app = &entity.StatisticsApp{
					UserId:   result.UserId,
					AppId:    result.AppId,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			appMap[result.ReqDate][result.AppId] = app

			if appModelStatMap[result.ReqDate][result.AppId] == nil {
				appModelStatMap[result.ReqDate][result.AppId] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range app.ModelStats {
				appModelStatMap[result.ReqDate][result.AppId][modelStat.ModelId] = modelStat
			}
		}

		appModelStat := appModelStatMap[result.ReqDate][result.AppId][result.ModelId]
		if appModelStat == nil {
			appModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			appModelStatMap[result.ReqDate][result.AppId][result.ModelId] = appModelStat
		}

		app.Total += 1
		app.Tokens += result.TotalTokens
		appModelStat.Total += 1
		appModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			app.Abnormal += 1
			app.AbnormalTokens += result.TotalTokens
			appModelStat.Abnormal += 1
			appModelStat.AbnormalTokens += result.TotalTokens
		}

		if appKeyMap[result.ReqDate] == nil {
			appKeyMap[result.ReqDate] = make(map[string]*entity.StatisticsAppKey)
			appKeyModelStatMap[result.ReqDate] = make(map[string]map[string]*common.ModelStat)
		}

		appKey := appKeyMap[result.ReqDate][result.Creator]
		if appKey == nil {

			if appKey, err = dao.StatisticsAppKey.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "app_key": result.Creator}); err != nil {
				appKey = &entity.StatisticsAppKey{
					UserId:   result.UserId,
					AppId:    result.AppId,
					AppKey:   result.Creator,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			appKeyMap[result.ReqDate][result.Creator] = appKey

			if appKeyModelStatMap[result.ReqDate][result.Creator] == nil {
				appKeyModelStatMap[result.ReqDate][result.Creator] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range appKey.ModelStats {
				appKeyModelStatMap[result.ReqDate][result.Creator][modelStat.ModelId] = modelStat
			}
		}

		appKeyModelStat := appKeyModelStatMap[result.ReqDate][result.Creator][result.ModelId]
		if appKeyModelStat == nil {
			appKeyModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			appKeyModelStatMap[result.ReqDate][result.Creator][result.ModelId] = appKeyModelStat
		}

		appKey.Total += 1
		appKey.Tokens += result.TotalTokens
		appKeyModelStat.Total += 1
		appKeyModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			appKey.Abnormal += 1
			appKey.AbnormalTokens += result.TotalTokens
			appKeyModelStat.Abnormal += 1
			appKeyModelStat.AbnormalTokens += result.TotalTokens
		}
	}

	for reqDate, data := range userMap {
		for userId, user := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range userModelStatMap[reqDate][userId] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsUser := &do.StatisticsUser{
				UserId:         user.UserId,
				StatDate:       user.StatDate,
				StatTime:       user.StatTime,
				Total:          user.Total,
				Tokens:         user.Tokens,
				Abnormal:       user.Abnormal,
				AbnormalTokens: user.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        user.Creator,
				CreatedAt:      user.CreatedAt,
			}

			if user.Id != "" {
				if err = dao.StatisticsUser.UpdateById(ctx, user.Id, statisticsUser); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsUser.Insert(ctx, statisticsUser); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	for reqDate, data := range appMap {
		for appId, app := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range appModelStatMap[reqDate][appId] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsApp := &do.StatisticsApp{
				UserId:         app.UserId,
				AppId:          app.AppId,
				StatDate:       app.StatDate,
				StatTime:       app.StatTime,
				Total:          app.Total,
				Tokens:         app.Tokens,
				Abnormal:       app.Abnormal,
				AbnormalTokens: app.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        app.Creator,
				CreatedAt:      app.CreatedAt,
			}

			if app.Id != "" {
				if err = dao.StatisticsApp.UpdateById(ctx, app.Id, statisticsApp); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsApp.Insert(ctx, statisticsApp); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	for reqDate, data := range appKeyMap {
		for app_key, appKey := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range appKeyModelStatMap[reqDate][app_key] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsAppKey := &do.StatisticsAppKey{
				UserId:         appKey.UserId,
				AppId:          appKey.AppId,
				AppKey:         appKey.AppKey,
				StatDate:       appKey.StatDate,
				StatTime:       appKey.StatTime,
				Total:          appKey.Total,
				Tokens:         appKey.Tokens,
				Abnormal:       appKey.Abnormal,
				AbnormalTokens: appKey.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        appKey.Creator,
				CreatedAt:      appKey.CreatedAt,
			}

			if appKey.Id != "" {
				if err = dao.StatisticsAppKey.UpdateById(ctx, appKey.Id, statisticsAppKey); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsAppKey.Insert(ctx, statisticsAppKey); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	if _, err = redis.Set(ctx, consts.STATISTICS_CHAT_LAST_TIME_KEY, lastTime); err != nil {
		logger.Error(ctx, err)
		return
	}

	if _, err = redis.Set(ctx, consts.STATISTICS_CHAT_LAST_ID_KEY, lastId); err != nil {
		logger.Error(ctx, err)
		return
	}

	if int64(len(results)) == config.Cfg.Statistics.Limit {
		s.StatisticsChat(ctx)
	}
}

// 统计绘图数据
func (s *sStatistics) StatisticsImage(ctx context.Context) {

	logger.Debugf(ctx, "sStatistics StatisticsImage start")

	now := gtime.TimestampMilli()
	defer func() {
		logger.Debugf(ctx, "sStatistics StatisticsImage end time: %d", gtime.TimestampMilli()-now)
	}()

	lastTime, err := redis.GetInt64(ctx, consts.STATISTICS_IMAGE_LAST_TIME_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	lastId, err := redis.GetStr(ctx, consts.STATISTICS_IMAGE_LAST_ID_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	filter := bson.M{
		"created_at": bson.M{
			"$gte": lastTime,
		},
		"is_smart_match": bson.M{"$ne": true},
		"is_retry":       bson.M{"$ne": true},
	}

	results, err := dao.Image.FindByPage(ctx, &db.Paging{Page: 1, PageSize: config.Cfg.Statistics.Limit}, filter, "", "created_at")
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	userMap := make(map[string]map[int]*entity.StatisticsUser)                     // map[req_date][user_id]entity.StatisticsUser
	userModelStatMap := make(map[string]map[int]map[string]*common.ModelStat)      // map[req_date][user_id][model_id]common.ModelStat
	appMap := make(map[string]map[int]*entity.StatisticsApp)                       // map[req_date][app_id]entity.StatisticsApp
	appModelStatMap := make(map[string]map[int]map[string]*common.ModelStat)       // map[req_date][app_id][model_id]common.ModelStat
	appKeyMap := make(map[string]map[string]*entity.StatisticsAppKey)              // map[req_date][app_key]entity.StatisticsAppKey
	appKeyModelStatMap := make(map[string]map[string]map[string]*common.ModelStat) // map[req_date][app_key][model_id]common.ModelStat

	for _, result := range results {

		if result.Id == lastId {
			continue
		}

		lastTime = result.CreatedAt
		lastId = result.Id

		if userMap[result.ReqDate] == nil {
			userMap[result.ReqDate] = make(map[int]*entity.StatisticsUser)
			userModelStatMap[result.ReqDate] = make(map[int]map[string]*common.ModelStat)
		}

		user := userMap[result.ReqDate][result.UserId]
		if user == nil {

			if user, err = dao.StatisticsUser.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "user_id": result.UserId}); err != nil {
				user = &entity.StatisticsUser{
					UserId:   result.UserId,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			userMap[result.ReqDate][result.UserId] = user

			if userModelStatMap[result.ReqDate][result.UserId] == nil {
				userModelStatMap[result.ReqDate][result.UserId] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range user.ModelStats {
				userModelStatMap[result.ReqDate][result.UserId][modelStat.ModelId] = modelStat
			}
		}

		userModelStat := userModelStatMap[result.ReqDate][result.UserId][result.ModelId]
		if userModelStat == nil {
			userModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			userModelStatMap[result.ReqDate][result.UserId][result.ModelId] = userModelStat
		}

		user.Total += 1
		user.Tokens += result.TotalTokens
		userModelStat.Total += 1
		userModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			user.Abnormal += 1
			user.AbnormalTokens += result.TotalTokens
			userModelStat.Abnormal += 1
			userModelStat.AbnormalTokens += result.TotalTokens
		}

		if appMap[result.ReqDate] == nil {
			appMap[result.ReqDate] = make(map[int]*entity.StatisticsApp)
			appModelStatMap[result.ReqDate] = make(map[int]map[string]*common.ModelStat)
		}

		app := appMap[result.ReqDate][result.AppId]
		if app == nil {

			if app, err = dao.StatisticsApp.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "app_id": result.AppId}); err != nil {
				app = &entity.StatisticsApp{
					UserId:   result.UserId,
					AppId:    result.AppId,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			appMap[result.ReqDate][result.AppId] = app

			if appModelStatMap[result.ReqDate][result.AppId] == nil {
				appModelStatMap[result.ReqDate][result.AppId] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range app.ModelStats {
				appModelStatMap[result.ReqDate][result.AppId][modelStat.ModelId] = modelStat
			}
		}

		appModelStat := appModelStatMap[result.ReqDate][result.AppId][result.ModelId]
		if appModelStat == nil {
			appModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			appModelStatMap[result.ReqDate][result.AppId][result.ModelId] = appModelStat
		}

		app.Total += 1
		app.Tokens += result.TotalTokens
		appModelStat.Total += 1
		appModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			app.Abnormal += 1
			app.AbnormalTokens += result.TotalTokens
			appModelStat.Abnormal += 1
			appModelStat.AbnormalTokens += result.TotalTokens
		}

		if appKeyMap[result.ReqDate] == nil {
			appKeyMap[result.ReqDate] = make(map[string]*entity.StatisticsAppKey)
			appKeyModelStatMap[result.ReqDate] = make(map[string]map[string]*common.ModelStat)
		}

		appKey := appKeyMap[result.ReqDate][result.Creator]
		if appKey == nil {

			if appKey, err = dao.StatisticsAppKey.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "app_key": result.Creator}); err != nil {
				appKey = &entity.StatisticsAppKey{
					UserId:   result.UserId,
					AppId:    result.AppId,
					AppKey:   result.Creator,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			appKeyMap[result.ReqDate][result.Creator] = appKey

			if appKeyModelStatMap[result.ReqDate][result.Creator] == nil {
				appKeyModelStatMap[result.ReqDate][result.Creator] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range appKey.ModelStats {
				appKeyModelStatMap[result.ReqDate][result.Creator][modelStat.ModelId] = modelStat
			}
		}

		appKeyModelStat := appKeyModelStatMap[result.ReqDate][result.Creator][result.ModelId]
		if appKeyModelStat == nil {
			appKeyModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			appKeyModelStatMap[result.ReqDate][result.Creator][result.ModelId] = appKeyModelStat
		}

		appKey.Total += 1
		appKey.Tokens += result.TotalTokens
		appKeyModelStat.Total += 1
		appKeyModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			appKey.Abnormal += 1
			appKey.AbnormalTokens += result.TotalTokens
			appKeyModelStat.Abnormal += 1
			appKeyModelStat.AbnormalTokens += result.TotalTokens
		}
	}

	for reqDate, data := range userMap {
		for userId, user := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range userModelStatMap[reqDate][userId] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsUser := &do.StatisticsUser{
				UserId:         user.UserId,
				StatDate:       user.StatDate,
				StatTime:       user.StatTime,
				Total:          user.Total,
				Tokens:         user.Tokens,
				Abnormal:       user.Abnormal,
				AbnormalTokens: user.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        user.Creator,
				CreatedAt:      user.CreatedAt,
			}

			if user.Id != "" {
				if err = dao.StatisticsUser.UpdateById(ctx, user.Id, statisticsUser); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsUser.Insert(ctx, statisticsUser); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	for reqDate, data := range appMap {
		for appId, app := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range appModelStatMap[reqDate][appId] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsApp := &do.StatisticsApp{
				UserId:         app.UserId,
				AppId:          app.AppId,
				StatDate:       app.StatDate,
				StatTime:       app.StatTime,
				Total:          app.Total,
				Tokens:         app.Tokens,
				Abnormal:       app.Abnormal,
				AbnormalTokens: app.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        app.Creator,
				CreatedAt:      app.CreatedAt,
			}

			if app.Id != "" {
				if err = dao.StatisticsApp.UpdateById(ctx, app.Id, statisticsApp); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsApp.Insert(ctx, statisticsApp); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	for reqDate, data := range appKeyMap {
		for app_key, appKey := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range appKeyModelStatMap[reqDate][app_key] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsAppKey := &do.StatisticsAppKey{
				UserId:         appKey.UserId,
				AppId:          appKey.AppId,
				AppKey:         appKey.AppKey,
				StatDate:       appKey.StatDate,
				StatTime:       appKey.StatTime,
				Total:          appKey.Total,
				Tokens:         appKey.Tokens,
				Abnormal:       appKey.Abnormal,
				AbnormalTokens: appKey.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        appKey.Creator,
				CreatedAt:      appKey.CreatedAt,
			}

			if appKey.Id != "" {
				if err = dao.StatisticsAppKey.UpdateById(ctx, appKey.Id, statisticsAppKey); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsAppKey.Insert(ctx, statisticsAppKey); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	if _, err = redis.Set(ctx, consts.STATISTICS_IMAGE_LAST_TIME_KEY, lastTime); err != nil {
		logger.Error(ctx, err)
		return
	}

	if _, err = redis.Set(ctx, consts.STATISTICS_IMAGE_LAST_ID_KEY, lastId); err != nil {
		logger.Error(ctx, err)
		return
	}

	if int64(len(results)) == config.Cfg.Statistics.Limit {
		s.StatisticsImage(ctx)
	}
}

// 统计音频数据
func (s *sStatistics) StatisticsAudio(ctx context.Context) {

	logger.Debugf(ctx, "sStatistics StatisticsAudio start")

	now := gtime.TimestampMilli()
	defer func() {
		logger.Debugf(ctx, "sStatistics StatisticsAudio end time: %d", gtime.TimestampMilli()-now)
	}()

	lastTime, err := redis.GetInt64(ctx, consts.STATISTICS_AUDIO_LAST_TIME_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	lastId, err := redis.GetStr(ctx, consts.STATISTICS_AUDIO_LAST_ID_KEY)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	filter := bson.M{
		"created_at": bson.M{
			"$gte": lastTime,
		},
		"is_smart_match": bson.M{"$ne": true},
		"is_retry":       bson.M{"$ne": true},
	}

	results, err := dao.Audio.FindByPage(ctx, &db.Paging{Page: 1, PageSize: config.Cfg.Statistics.Limit}, filter, "", "created_at")
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	userMap := make(map[string]map[int]*entity.StatisticsUser)                     // map[req_date][user_id]entity.StatisticsUser
	userModelStatMap := make(map[string]map[int]map[string]*common.ModelStat)      // map[req_date][user_id][model_id]common.ModelStat
	appMap := make(map[string]map[int]*entity.StatisticsApp)                       // map[req_date][app_id]entity.StatisticsApp
	appModelStatMap := make(map[string]map[int]map[string]*common.ModelStat)       // map[req_date][app_id][model_id]common.ModelStat
	appKeyMap := make(map[string]map[string]*entity.StatisticsAppKey)              // map[req_date][app_key]entity.StatisticsAppKey
	appKeyModelStatMap := make(map[string]map[string]map[string]*common.ModelStat) // map[req_date][app_key][model_id]common.ModelStat

	for _, result := range results {

		if result.Id == lastId {
			continue
		}

		lastTime = result.CreatedAt
		lastId = result.Id

		if userMap[result.ReqDate] == nil {
			userMap[result.ReqDate] = make(map[int]*entity.StatisticsUser)
			userModelStatMap[result.ReqDate] = make(map[int]map[string]*common.ModelStat)
		}

		user := userMap[result.ReqDate][result.UserId]
		if user == nil {

			if user, err = dao.StatisticsUser.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "user_id": result.UserId}); err != nil {
				user = &entity.StatisticsUser{
					UserId:   result.UserId,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			userMap[result.ReqDate][result.UserId] = user

			if userModelStatMap[result.ReqDate][result.UserId] == nil {
				userModelStatMap[result.ReqDate][result.UserId] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range user.ModelStats {
				userModelStatMap[result.ReqDate][result.UserId][modelStat.ModelId] = modelStat
			}
		}

		userModelStat := userModelStatMap[result.ReqDate][result.UserId][result.ModelId]
		if userModelStat == nil {
			userModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			userModelStatMap[result.ReqDate][result.UserId][result.ModelId] = userModelStat
		}

		user.Total += 1
		user.Tokens += result.TotalTokens
		userModelStat.Total += 1
		userModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			user.Abnormal += 1
			user.AbnormalTokens += result.TotalTokens
			userModelStat.Abnormal += 1
			userModelStat.AbnormalTokens += result.TotalTokens
		}

		if appMap[result.ReqDate] == nil {
			appMap[result.ReqDate] = make(map[int]*entity.StatisticsApp)
			appModelStatMap[result.ReqDate] = make(map[int]map[string]*common.ModelStat)
		}

		app := appMap[result.ReqDate][result.AppId]
		if app == nil {

			if app, err = dao.StatisticsApp.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "app_id": result.AppId}); err != nil {
				app = &entity.StatisticsApp{
					UserId:   result.UserId,
					AppId:    result.AppId,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			appMap[result.ReqDate][result.AppId] = app

			if appModelStatMap[result.ReqDate][result.AppId] == nil {
				appModelStatMap[result.ReqDate][result.AppId] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range app.ModelStats {
				appModelStatMap[result.ReqDate][result.AppId][modelStat.ModelId] = modelStat
			}
		}

		appModelStat := appModelStatMap[result.ReqDate][result.AppId][result.ModelId]
		if appModelStat == nil {
			appModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			appModelStatMap[result.ReqDate][result.AppId][result.ModelId] = appModelStat
		}

		app.Total += 1
		app.Tokens += result.TotalTokens
		appModelStat.Total += 1
		appModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			app.Abnormal += 1
			app.AbnormalTokens += result.TotalTokens
			appModelStat.Abnormal += 1
			appModelStat.AbnormalTokens += result.TotalTokens
		}

		if appKeyMap[result.ReqDate] == nil {
			appKeyMap[result.ReqDate] = make(map[string]*entity.StatisticsAppKey)
			appKeyModelStatMap[result.ReqDate] = make(map[string]map[string]*common.ModelStat)
		}

		appKey := appKeyMap[result.ReqDate][result.Creator]
		if appKey == nil {

			if appKey, err = dao.StatisticsAppKey.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "app_key": result.Creator}); err != nil {
				appKey = &entity.StatisticsAppKey{
					UserId:   result.UserId,
					AppId:    result.AppId,
					AppKey:   result.Creator,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			appKeyMap[result.ReqDate][result.Creator] = appKey

			if appKeyModelStatMap[result.ReqDate][result.Creator] == nil {
				appKeyModelStatMap[result.ReqDate][result.Creator] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range appKey.ModelStats {
				appKeyModelStatMap[result.ReqDate][result.Creator][modelStat.ModelId] = modelStat
			}
		}

		appKeyModelStat := appKeyModelStatMap[result.ReqDate][result.Creator][result.ModelId]
		if appKeyModelStat == nil {
			appKeyModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			appKeyModelStatMap[result.ReqDate][result.Creator][result.ModelId] = appKeyModelStat
		}

		appKey.Total += 1
		appKey.Tokens += result.TotalTokens
		appKeyModelStat.Total += 1
		appKeyModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			appKey.Abnormal += 1
			appKey.AbnormalTokens += result.TotalTokens
			appKeyModelStat.Abnormal += 1
			appKeyModelStat.AbnormalTokens += result.TotalTokens
		}
	}

	for reqDate, data := range userMap {
		for userId, user := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range userModelStatMap[reqDate][userId] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsUser := &do.StatisticsUser{
				UserId:         user.UserId,
				StatDate:       user.StatDate,
				StatTime:       user.StatTime,
				Total:          user.Total,
				Tokens:         user.Tokens,
				Abnormal:       user.Abnormal,
				AbnormalTokens: user.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        user.Creator,
				CreatedAt:      user.CreatedAt,
			}

			if user.Id != "" {
				if err = dao.StatisticsUser.UpdateById(ctx, user.Id, statisticsUser); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsUser.Insert(ctx, statisticsUser); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	for reqDate, data := range appMap {
		for appId, app := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range appModelStatMap[reqDate][appId] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsApp := &do.StatisticsApp{
				UserId:         app.UserId,
				AppId:          app.AppId,
				StatDate:       app.StatDate,
				StatTime:       app.StatTime,
				Total:          app.Total,
				Tokens:         app.Tokens,
				Abnormal:       app.Abnormal,
				AbnormalTokens: app.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        app.Creator,
				CreatedAt:      app.CreatedAt,
			}

			if app.Id != "" {
				if err = dao.StatisticsApp.UpdateById(ctx, app.Id, statisticsApp); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsApp.Insert(ctx, statisticsApp); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	for reqDate, data := range appKeyMap {
		for app_key, appKey := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range appKeyModelStatMap[reqDate][app_key] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsAppKey := &do.StatisticsAppKey{
				UserId:         appKey.UserId,
				AppId:          appKey.AppId,
				AppKey:         appKey.AppKey,
				StatDate:       appKey.StatDate,
				StatTime:       appKey.StatTime,
				Total:          appKey.Total,
				Tokens:         appKey.Tokens,
				Abnormal:       appKey.Abnormal,
				AbnormalTokens: appKey.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        appKey.Creator,
				CreatedAt:      appKey.CreatedAt,
			}

			if appKey.Id != "" {
				if err = dao.StatisticsAppKey.UpdateById(ctx, appKey.Id, statisticsAppKey); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsAppKey.Insert(ctx, statisticsAppKey); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	if _, err = redis.Set(ctx, consts.STATISTICS_AUDIO_LAST_TIME_KEY, lastTime); err != nil {
		logger.Error(ctx, err)
		return
	}

	if _, err = redis.Set(ctx, consts.STATISTICS_AUDIO_LAST_ID_KEY, lastId); err != nil {
		logger.Error(ctx, err)
		return
	}

	if int64(len(results)) == config.Cfg.Statistics.Limit {
		s.StatisticsAudio(ctx)
	}
}

// 统计数据
func (s *sStatistics) StatisticsData(ctx context.Context, collection, index, lastTimeKey, lastIdKey string) {

	logger.Debugf(ctx, "sStatistics StatisticsData start")

	now := gtime.TimestampMilli()
	defer func() {
		logger.Debugf(ctx, "sStatistics StatisticsData end time: %d", gtime.TimestampMilli()-now)
	}()

	lastTime, err := redis.GetInt64(ctx, lastTimeKey)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	lastId, err := redis.GetStr(ctx, lastIdKey)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	filter := bson.M{
		"updated_at": bson.M{
			"$gte": lastTime,
		},
		"is_smart_match": bson.M{"$ne": true},
		"is_retry":       bson.M{"$ne": true},
	}

	results, err := dao.NewMongoDB[entity.StatisticsData](db.DefaultDatabase, collection).FindByPage(ctx, &db.Paging{Page: 1, PageSize: config.Cfg.Statistics.Limit}, filter, index, "updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	userMap := make(map[string]map[int]*entity.StatisticsUser)                     // map[req_date][user_id]entity.StatisticsUser
	userModelStatMap := make(map[string]map[int]map[string]*common.ModelStat)      // map[req_date][user_id][model_id]common.ModelStat
	appMap := make(map[string]map[int]*entity.StatisticsApp)                       // map[req_date][app_id]entity.StatisticsApp
	appModelStatMap := make(map[string]map[int]map[string]*common.ModelStat)       // map[req_date][app_id][model_id]common.ModelStat
	appKeyMap := make(map[string]map[string]*entity.StatisticsAppKey)              // map[req_date][app_key]entity.StatisticsAppKey
	appKeyModelStatMap := make(map[string]map[string]map[string]*common.ModelStat) // map[req_date][app_key][model_id]common.ModelStat

	for _, result := range results {

		if result.Id == lastId {
			continue
		}

		lastTime = result.UpdatedAt
		lastId = result.Id

		if userMap[result.ReqDate] == nil {
			userMap[result.ReqDate] = make(map[int]*entity.StatisticsUser)
			userModelStatMap[result.ReqDate] = make(map[int]map[string]*common.ModelStat)
		}

		user := userMap[result.ReqDate][result.UserId]
		if user == nil {

			if user, err = dao.StatisticsUser.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "user_id": result.UserId}); err != nil {
				user = &entity.StatisticsUser{
					UserId:   result.UserId,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			userMap[result.ReqDate][result.UserId] = user

			if userModelStatMap[result.ReqDate][result.UserId] == nil {
				userModelStatMap[result.ReqDate][result.UserId] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range user.ModelStats {
				userModelStatMap[result.ReqDate][result.UserId][modelStat.ModelId] = modelStat
			}
		}

		userModelStat := userModelStatMap[result.ReqDate][result.UserId][result.ModelId]
		if userModelStat == nil {
			userModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			userModelStatMap[result.ReqDate][result.UserId][result.ModelId] = userModelStat
		}

		user.Total += 1
		user.Tokens += result.TotalTokens
		userModelStat.Total += 1
		userModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			user.Abnormal += 1
			user.AbnormalTokens += result.TotalTokens
			userModelStat.Abnormal += 1
			userModelStat.AbnormalTokens += result.TotalTokens
		}

		if appMap[result.ReqDate] == nil {
			appMap[result.ReqDate] = make(map[int]*entity.StatisticsApp)
			appModelStatMap[result.ReqDate] = make(map[int]map[string]*common.ModelStat)
		}

		app := appMap[result.ReqDate][result.AppId]
		if app == nil {

			if app, err = dao.StatisticsApp.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "app_id": result.AppId}); err != nil {
				app = &entity.StatisticsApp{
					UserId:   result.UserId,
					AppId:    result.AppId,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			appMap[result.ReqDate][result.AppId] = app

			if appModelStatMap[result.ReqDate][result.AppId] == nil {
				appModelStatMap[result.ReqDate][result.AppId] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range app.ModelStats {
				appModelStatMap[result.ReqDate][result.AppId][modelStat.ModelId] = modelStat
			}
		}

		appModelStat := appModelStatMap[result.ReqDate][result.AppId][result.ModelId]
		if appModelStat == nil {
			appModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			appModelStatMap[result.ReqDate][result.AppId][result.ModelId] = appModelStat
		}

		app.Total += 1
		app.Tokens += result.TotalTokens
		appModelStat.Total += 1
		appModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			app.Abnormal += 1
			app.AbnormalTokens += result.TotalTokens
			appModelStat.Abnormal += 1
			appModelStat.AbnormalTokens += result.TotalTokens
		}

		if appKeyMap[result.ReqDate] == nil {
			appKeyMap[result.ReqDate] = make(map[string]*entity.StatisticsAppKey)
			appKeyModelStatMap[result.ReqDate] = make(map[string]map[string]*common.ModelStat)
		}

		appKey := appKeyMap[result.ReqDate][result.Creator]
		if appKey == nil {

			if appKey, err = dao.StatisticsAppKey.FindOne(ctx, bson.M{"stat_date": result.ReqDate, "app_key": result.Creator}); err != nil {
				appKey = &entity.StatisticsAppKey{
					UserId:   result.UserId,
					AppId:    result.AppId,
					AppKey:   result.Creator,
					StatDate: result.ReqDate,
					StatTime: gtime.NewFromStrFormat(result.ReqDate, time.DateOnly).TimestampMilli(),
				}
			}

			appKeyMap[result.ReqDate][result.Creator] = appKey

			if appKeyModelStatMap[result.ReqDate][result.Creator] == nil {
				appKeyModelStatMap[result.ReqDate][result.Creator] = make(map[string]*common.ModelStat)
			}

			for _, modelStat := range appKey.ModelStats {
				appKeyModelStatMap[result.ReqDate][result.Creator][modelStat.ModelId] = modelStat
			}
		}

		appKeyModelStat := appKeyModelStatMap[result.ReqDate][result.Creator][result.ModelId]
		if appKeyModelStat == nil {
			appKeyModelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			appKeyModelStatMap[result.ReqDate][result.Creator][result.ModelId] = appKeyModelStat
		}

		appKey.Total += 1
		appKey.Tokens += result.TotalTokens
		appKeyModelStat.Total += 1
		appKeyModelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			appKey.Abnormal += 1
			appKey.AbnormalTokens += result.TotalTokens
			appKeyModelStat.Abnormal += 1
			appKeyModelStat.AbnormalTokens += result.TotalTokens
		}
	}

	for reqDate, data := range userMap {
		for userId, user := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range userModelStatMap[reqDate][userId] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsUser := &do.StatisticsUser{
				UserId:         user.UserId,
				StatDate:       user.StatDate,
				StatTime:       user.StatTime,
				Total:          user.Total,
				Tokens:         user.Tokens,
				Abnormal:       user.Abnormal,
				AbnormalTokens: user.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        user.Creator,
				CreatedAt:      user.CreatedAt,
			}

			if user.Id != "" {
				if err = dao.StatisticsUser.UpdateById(ctx, user.Id, statisticsUser); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsUser.Insert(ctx, statisticsUser); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	for reqDate, data := range appMap {
		for appId, app := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range appModelStatMap[reqDate][appId] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsApp := &do.StatisticsApp{
				UserId:         app.UserId,
				AppId:          app.AppId,
				StatDate:       app.StatDate,
				StatTime:       app.StatTime,
				Total:          app.Total,
				Tokens:         app.Tokens,
				Abnormal:       app.Abnormal,
				AbnormalTokens: app.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        app.Creator,
				CreatedAt:      app.CreatedAt,
			}

			if app.Id != "" {
				if err = dao.StatisticsApp.UpdateById(ctx, app.Id, statisticsApp); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsApp.Insert(ctx, statisticsApp); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	for reqDate, data := range appKeyMap {
		for app_key, appKey := range data {

			modelStats := make([]*common.ModelStat, 0)
			for _, modelStat := range appKeyModelStatMap[reqDate][app_key] {
				modelStats = append(modelStats, modelStat)
			}

			statisticsAppKey := &do.StatisticsAppKey{
				UserId:         appKey.UserId,
				AppId:          appKey.AppId,
				AppKey:         appKey.AppKey,
				StatDate:       appKey.StatDate,
				StatTime:       appKey.StatTime,
				Total:          appKey.Total,
				Tokens:         appKey.Tokens,
				Abnormal:       appKey.Abnormal,
				AbnormalTokens: appKey.AbnormalTokens,
				ModelStats:     modelStats,
				Creator:        appKey.Creator,
				CreatedAt:      appKey.CreatedAt,
			}

			if appKey.Id != "" {
				if err = dao.StatisticsAppKey.UpdateById(ctx, appKey.Id, statisticsAppKey); err != nil {
					logger.Error(ctx, err)
				}
			} else {
				if _, err = dao.StatisticsAppKey.Insert(ctx, statisticsAppKey); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	if _, err = redis.Set(ctx, lastTimeKey, lastTime); err != nil {
		logger.Error(ctx, err)
		return
	}

	if _, err = redis.Set(ctx, lastIdKey, lastId); err != nil {
		logger.Error(ctx, err)
		return
	}

	if int64(len(results)) == config.Cfg.Statistics.Limit {
		s.StatisticsData(ctx, collection, index, lastTimeKey, lastIdKey)
	}
}
