package statistics

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/common"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type sStatistics struct{}

func init() {
	service.RegisterStatistics(New())
}

func New() service.IStatistics {
	return &sStatistics{}
}

// 用户数据
func (s *sStatistics) DataUser(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error) {

	logger.Debugf(ctx, "sStatistics DataUser userId: %d statStartTime: %d statEndTime: %d start", params.UserId, params.StatStartTime, params.StatEndTime)

	now := gtime.TimestampMilli()
	defer func() {
		logger.Debugf(ctx, "sStatistics DataUser userId: %d statStartTime: %d statEndTime: %d end time: %d", params.UserId, params.StatStartTime, params.StatEndTime, gtime.TimestampMilli()-now)
	}()

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"user_id": params.UserId,
				"req_time": bson.M{
					"$gte": params.StatStartTime,
					"$lte": params.StatEndTime,
				},
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
			},
		},
		{
			"$group": bson.M{
				"_id":    bson.M{"req_date": "$req_date", "model": "$model"},
				"count":  bson.M{"$sum": 1},
				"tokens": bson.M{"$sum": "$total_tokens"},
			},
		},
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	resultMap := make(map[string]*do.StatisticsUser)
	resultModelMap := make(map[string]map[string]*common.ModelStat) // [日期][模型]模型数据

	for _, res := range result {

		_id := res["_id"].(map[string]interface{})
		reqDate := gconv.String(_id["req_date"])
		model := gconv.String(_id["model"])
		count := gconv.Int(res["count"])
		tokens := gconv.Int(res["tokens"])

		statisticsUser := resultMap[reqDate]

		if statisticsUser != nil {

			statisticsUser.Total += count
			statisticsUser.Tokens += tokens

			modelStat := &common.ModelStat{
				Model:  model,
				Total:  count,
				Tokens: tokens,
			}

			resultMap[reqDate].ModelStats = append(resultMap[reqDate].ModelStats, modelStat)
			resultModelMap[reqDate][model] = modelStat

		} else {

			modelStat := &common.ModelStat{
				Model:  model,
				Total:  count,
				Tokens: tokens,
			}

			resultMap[reqDate] = &do.StatisticsUser{
				UserId:     params.UserId,
				StatDate:   reqDate,
				StatTime:   gtime.NewFromStrFormat(reqDate, time.DateOnly).TimestampMilli(),
				Total:      count,
				Tokens:     tokens,
				ModelStats: []*common.ModelStat{modelStat},
			}

			resultModelMap[reqDate] = make(map[string]*common.ModelStat)
			resultModelMap[reqDate][model] = modelStat
		}
	}

	match := pipeline[0]["$match"].(bson.M)
	match["status"] = bson.M{"$ne": 1} // 异常状态数据

	abnormalResult := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &abnormalResult); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, res := range abnormalResult {

		_id := res["_id"].(map[string]interface{})
		reqDate := gconv.String(_id["req_date"])
		model := gconv.String(_id["model"])
		count := gconv.Int(res["count"])
		tokens := gconv.Int(res["tokens"])

		resultMap[reqDate].Abnormal += count
		resultMap[reqDate].AbnormalTokens += tokens
		resultModelMap[reqDate][model].Abnormal += count
		resultModelMap[reqDate][model].AbnormalTokens += tokens
	}

	for _, statisticsUser := range resultMap {
		if _, err := dao.StatisticsUser.Insert(ctx, statisticsUser); err != nil {
			logger.Error(ctx, err)
		}
	}

	return nil, nil
}

// 应用数据
func (s *sStatistics) DataApp(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error) {

	logger.Debugf(ctx, "sStatistics DataApp userId: %d appId: %d statStartTime: %d statEndTime: %d start", params.UserId, params.AppId, params.StatStartTime, params.StatEndTime)

	now := gtime.TimestampMilli()
	defer func() {
		logger.Debugf(ctx, "sStatistics DataApp userId: %d appId: %d statStartTime: %d statEndTime: %d end time: %d", params.UserId, params.AppId, params.StatStartTime, params.StatEndTime, gtime.TimestampMilli()-now)
	}()

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"app_id": params.AppId,
				"req_time": bson.M{
					"$gte": params.StatStartTime,
					"$lte": params.StatEndTime,
				},
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
			},
		},
		{
			"$group": bson.M{
				"_id":    bson.M{"req_date": "$req_date", "model": "$model"},
				"count":  bson.M{"$sum": 1},
				"tokens": bson.M{"$sum": "$total_tokens"},
			},
		},
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	resultMap := make(map[string]*do.StatisticsApp)
	resultModelMap := make(map[string]map[string]*common.ModelStat) // [日期][模型]模型数据

	for _, res := range result {

		_id := res["_id"].(map[string]interface{})
		reqDate := gconv.String(_id["req_date"])
		model := gconv.String(_id["model"])
		count := gconv.Int(res["count"])
		tokens := gconv.Int(res["tokens"])

		statisticsApp := resultMap[reqDate]

		if statisticsApp != nil {

			statisticsApp.Total += count
			statisticsApp.Tokens += tokens

			modelStat := &common.ModelStat{
				Model:  model,
				Total:  count,
				Tokens: tokens,
			}

			resultMap[reqDate].Models = append(resultMap[reqDate].Models, modelStat)
			resultModelMap[reqDate][model] = modelStat

		} else {

			modelStat := &common.ModelStat{
				Model:  model,
				Total:  count,
				Tokens: tokens,
			}

			resultMap[reqDate] = &do.StatisticsApp{
				UserId:   params.UserId,
				AppId:    params.AppId,
				StatDate: reqDate,
				StatTime: gtime.NewFromStrFormat(reqDate, time.DateOnly).TimestampMilli(),
				Total:    count,
				Tokens:   tokens,
				Models:   []*common.ModelStat{modelStat},
			}

			resultModelMap[reqDate] = make(map[string]*common.ModelStat)
			resultModelMap[reqDate][model] = modelStat
		}
	}

	match := pipeline[0]["$match"].(bson.M)
	match["status"] = bson.M{"$ne": 1} // 异常状态数据

	abnormalResult := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &abnormalResult); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, res := range abnormalResult {

		_id := res["_id"].(map[string]interface{})
		reqDate := gconv.String(_id["req_date"])
		model := gconv.String(_id["model"])
		count := gconv.Int(res["count"])
		tokens := gconv.Int(res["tokens"])

		resultMap[reqDate].Abnormal += count
		resultMap[reqDate].AbnormalTokens += tokens
		resultModelMap[reqDate][model].Abnormal += count
		resultModelMap[reqDate][model].AbnormalTokens += tokens
	}

	for _, statisticsApp := range resultMap {
		if _, err := dao.StatisticsApp.Insert(ctx, statisticsApp); err != nil {
			logger.Error(ctx, err)
		}
	}

	return nil, nil
}

// 应用密钥数据
func (s *sStatistics) DataAppKey(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error) {

	logger.Debugf(ctx, "sStatistics DataAppKey userId: %d appId: %d appKey: %s statStartTime: %d statEndTime: %d start", params.UserId, params.AppId, params.AppKey, params.StatStartTime, params.StatEndTime)

	now := gtime.TimestampMilli()
	defer func() {
		logger.Debugf(ctx, "sStatistics DataAppKey userId: %d appId: %d appKey: %s statStartTime: %d statEndTime: %d end time: %d", params.UserId, params.AppId, params.AppKey, params.StatStartTime, params.StatEndTime, gtime.TimestampMilli()-now)
	}()

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"creator": params.AppKey,
				"req_time": bson.M{
					"$gte": params.StatStartTime,
					"$lte": params.StatEndTime,
				},
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
			},
		},
		{
			"$group": bson.M{
				"_id":    bson.M{"req_date": "$req_date", "model": "$model"},
				"count":  bson.M{"$sum": 1},
				"tokens": bson.M{"$sum": "$total_tokens"},
			},
		},
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	resultMap := make(map[string]*do.StatisticsAppKey)
	resultModelMap := make(map[string]map[string]*common.ModelStat) // [日期][模型]模型数据

	for _, res := range result {

		_id := res["_id"].(map[string]interface{})
		reqDate := gconv.String(_id["req_date"])
		model := gconv.String(_id["model"])
		count := gconv.Int(res["count"])
		tokens := gconv.Int(res["tokens"])

		statisticsAppKey := resultMap[reqDate]

		if statisticsAppKey != nil {

			statisticsAppKey.Total += count
			statisticsAppKey.Tokens += tokens

			modelStat := &common.ModelStat{
				Model:  model,
				Total:  count,
				Tokens: tokens,
			}

			resultMap[reqDate].Models = append(resultMap[reqDate].Models, modelStat)
			resultModelMap[reqDate][model] = modelStat

		} else {

			modelStat := &common.ModelStat{
				Model:  model,
				Total:  count,
				Tokens: tokens,
			}

			resultMap[reqDate] = &do.StatisticsAppKey{
				UserId:   params.UserId,
				AppId:    params.AppId,
				AppKey:   params.AppKey,
				StatDate: reqDate,
				StatTime: gtime.NewFromStrFormat(reqDate, time.DateOnly).TimestampMilli(),
				Total:    count,
				Tokens:   tokens,
				Models:   []*common.ModelStat{modelStat},
			}

			resultModelMap[reqDate] = make(map[string]*common.ModelStat)
			resultModelMap[reqDate][model] = modelStat
		}
	}

	match := pipeline[0]["$match"].(bson.M)
	match["status"] = bson.M{"$ne": 1} // 异常状态数据

	abnormalResult := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &abnormalResult); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, res := range abnormalResult {

		_id := res["_id"].(map[string]interface{})
		reqDate := gconv.String(_id["req_date"])
		model := gconv.String(_id["model"])
		count := gconv.Int(res["count"])
		tokens := gconv.Int(res["tokens"])

		resultMap[reqDate].Abnormal += count
		resultMap[reqDate].AbnormalTokens += tokens
		resultModelMap[reqDate][model].Abnormal += count
		resultModelMap[reqDate][model].AbnormalTokens += tokens
	}

	for _, statisticsAppKey := range resultMap {
		if _, err := dao.StatisticsAppKey.Insert(ctx, statisticsAppKey); err != nil {
			logger.Error(ctx, err)
		}
	}

	return nil, nil
}

// 统计任务
func (s *sStatistics) StatisticsTask(ctx context.Context) {

	logger.Info(ctx, "sStatistics StatisticsTask start")

	now := gtime.TimestampMilli()
	defer func() {
		logger.Infof(ctx, "sStatistics StatisticsTask end time: %d", gtime.TimestampMilli()-now)
	}()

	startDate := gtime.Now().AddDate(0, 0, -config.Cfg.Statistics.Days)
	statStartTime := startDate.StartOfDay().TimestampMilli()
	statEndTime := startDate.EndOfDay(true).TimestampMilli()

	if config.Cfg.Statistics.Days == 0 {
		statStartTime = 0
		statEndTime = gtime.Now().AddDate(0, 0, -config.Cfg.Statistics.Days).EndOfDay(true).TimestampMilli()
	}

	if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

		logger.Info(ctx, "sStatistics StatisticsTask User start")

		now := gtime.TimestampMilli()
		defer func() {
			logger.Infof(ctx, "sStatistics StatisticsTask User end time: %d", gtime.TimestampMilli()-now)
		}()

		users, err := dao.User.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
		}

		for _, user := range users {
			if _, err = s.DataUser(ctx, model.StatisticsDataReq{
				UserId:        user.UserId,
				StatStartTime: statStartTime,
				StatEndTime:   statEndTime,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}

	}, nil); err != nil {
		logger.Error(ctx, err)
	}

	if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

		logger.Info(ctx, "sStatistics StatisticsTask App start")

		now := gtime.TimestampMilli()
		defer func() {
			logger.Infof(ctx, "sStatistics StatisticsTask App end time: %d", gtime.TimestampMilli()-now)
		}()

		apps, err := dao.App.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
		}

		for _, app := range apps {
			if _, err = s.DataApp(ctx, model.StatisticsDataReq{
				UserId:        app.UserId,
				AppId:         app.AppId,
				StatStartTime: statStartTime,
				StatEndTime:   statEndTime,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}

	}, nil); err != nil {
		logger.Error(ctx, err)
	}

	if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

		logger.Info(ctx, "sStatistics StatisticsTask AppKey start")

		now := gtime.TimestampMilli()
		defer func() {
			logger.Infof(ctx, "sStatistics StatisticsTask AppKey end time: %d", gtime.TimestampMilli()-now)
		}()

		keys, err := dao.Key.Find(ctx, bson.M{"type": 1})
		if err != nil {
			logger.Error(ctx, err)
		}

		for _, key := range keys {
			if _, err = s.DataAppKey(ctx, model.StatisticsDataReq{
				UserId:        key.UserId,
				AppId:         key.AppId,
				AppKey:        key.Key,
				StatStartTime: statStartTime,
				StatEndTime:   statEndTime,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}

	}, nil); err != nil {
		logger.Error(ctx, err)
	}
}

// 统计聊天数据
func (s *sStatistics) StatisticsChat(ctx context.Context, paging *db.Paging) {

	if paging == nil {
		paging = &db.Paging{
			Page:     1,
			PageSize: 100,
		}
	}

	filter := bson.M{
		"created_at": bson.M{
			"$gte": 0,
		},
		"is_smart_match": bson.M{"$ne": true},
		"is_retry":       bson.M{"$ne": true},
	}

	results, err := dao.Chat.FindByPage(ctx, paging, filter, "", "created_at")
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	userMap := make(map[string]map[int]*entity.StatisticsUser)                // map[req_date][user_id]entity.StatisticsUser
	userModelStatMap := make(map[string]map[int]map[string]*common.ModelStat) // map[req_date][user_id][model_id]common.ModelStat
	//appMap := make(map[string]map[int]*entity.StatisticsApp)          // map[req_date][app_id]entity.StatisticsApp
	//appKeyMap := make(map[string]map[string]*entity.StatisticsAppKey) // map[req_date][app_key]entity.StatisticsAppKey

	for _, result := range results {

		reqDate := userMap[result.ReqDate]
		if reqDate == nil {
			userMap[result.ReqDate] = make(map[int]*entity.StatisticsUser)
			userModelStatMap[result.ReqDate] = make(map[int]map[string]*common.ModelStat)
		}

		user := reqDate[result.UserId]
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

		modelStat := userModelStatMap[result.ReqDate][result.UserId][result.ModelId]
		if modelStat == nil {
			modelStat = &common.ModelStat{
				ModelId: result.ModelId,
				Model:   result.Model,
			}
			userModelStatMap[result.ReqDate][result.UserId][result.ModelId] = modelStat
		}

		user.Total += 1
		user.Tokens += result.TotalTokens
		modelStat.Total += 1
		modelStat.Tokens += result.TotalTokens

		if result.Status != 1 {
			user.Abnormal += 1
			user.AbnormalTokens += result.TotalTokens
			modelStat.Abnormal += 1
			modelStat.AbnormalTokens += result.TotalTokens
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

	if paging.Page < paging.PageCount {
		paging.Page += 1
		s.StatisticsChat(ctx, paging)
	}
}
