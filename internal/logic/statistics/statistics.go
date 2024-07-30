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
	"github.com/iimeta/fastapi-admin/internal/service"
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
				"_id":    "$req_date",
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
	for _, res := range result {
		resultMap[gconv.String(res["_id"])] = &do.StatisticsUser{
			UserId:   params.UserId,
			StatDate: gconv.String(res["_id"]),
			StatTime: gtime.NewFromStrFormat(gconv.String(res["_id"]), time.DateOnly).TimestampMilli(),
			Total:    gconv.Int(res["count"]),
			Tokens:   gconv.Int(res["tokens"]),
		}
	}

	abnormalPipeline := []bson.M{
		{
			"$match": bson.M{
				"user_id": params.UserId,
				"req_time": bson.M{
					"$gte": params.StatStartTime,
					"$lte": params.StatEndTime,
				},
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
				"status":         bson.M{"$ne": 1},
			},
		},
		{
			"$group": bson.M{
				"_id":   "$req_date",
				"count": bson.M{"$sum": 1},
			},
		},
	}

	abnormalResult := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, abnormalPipeline, &abnormalResult); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, res := range abnormalResult {
		if resultMap[gconv.String(res["_id"])] != nil {
			resultMap[gconv.String(res["_id"])].Abnormal = gconv.Int(res["count"])
		} else {
			resultMap[gconv.String(res["_id"])] = &do.StatisticsUser{
				UserId:   params.UserId,
				StatDate: gconv.String(res["_id"]),
				StatTime: gtime.NewFromStrFormat(gconv.String(res["_id"]), time.DateOnly).TimestampMilli(),
				Abnormal: gconv.Int(res["count"]),
			}
		}
	}

	modelsPipeline := []bson.M{
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
				"_id":   bson.M{"req_date": "$req_date", "model": "$model"},
				"count": bson.M{"$sum": 1},
			},
		},
	}

	modelsResult := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, modelsPipeline, &modelsResult); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, res := range modelsResult {

		_id := res["_id"].(map[string]interface{})

		if resultMap[gconv.String(_id["req_date"])] != nil {
			resultMap[gconv.String(_id["req_date"])].Models = append(resultMap[gconv.String(_id["req_date"])].Models, common.ModelStat{
				Model: gconv.String(_id["model"]),
				Total: gconv.Int(res["count"]),
			})
		} else {
			resultMap[gconv.String(_id["req_date"])] = &do.StatisticsUser{
				UserId:   params.UserId,
				StatDate: gconv.String(_id["req_date"]),
				StatTime: gtime.NewFromStrFormat(gconv.String(_id["req_date"]), time.DateOnly).TimestampMilli(),
				Models: []common.ModelStat{{
					Model: gconv.String(_id["model"]),
					Total: gconv.Int(res["count"]),
				}},
			}
		}
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
				"_id":    "$req_date",
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
	for _, res := range result {
		resultMap[gconv.String(res["_id"])] = &do.StatisticsApp{
			UserId:   params.UserId,
			AppId:    params.AppId,
			StatDate: gconv.String(res["_id"]),
			StatTime: gtime.NewFromStrFormat(gconv.String(res["_id"]), time.DateOnly).TimestampMilli(),
			Total:    gconv.Int(res["count"]),
			Tokens:   gconv.Int(res["tokens"]),
		}
	}

	abnormalPipeline := []bson.M{
		{
			"$match": bson.M{
				"app_id": params.AppId,
				"req_time": bson.M{
					"$gte": params.StatStartTime,
					"$lte": params.StatEndTime,
				},
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
				"status":         bson.M{"$ne": 1},
			},
		},
		{
			"$group": bson.M{
				"_id":   "$req_date",
				"count": bson.M{"$sum": 1},
			},
		},
	}

	abnormalResult := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, abnormalPipeline, &abnormalResult); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, res := range abnormalResult {
		if resultMap[gconv.String(res["_id"])] != nil {
			resultMap[gconv.String(res["_id"])].Abnormal = gconv.Int(res["count"])
		} else {
			resultMap[gconv.String(res["_id"])] = &do.StatisticsApp{
				UserId:   params.UserId,
				AppId:    params.AppId,
				StatDate: gconv.String(res["_id"]),
				StatTime: gtime.NewFromStrFormat(gconv.String(res["_id"]), time.DateOnly).TimestampMilli(),
				Abnormal: gconv.Int(res["count"]),
			}
		}
	}

	modelsPipeline := []bson.M{
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
				"_id":   bson.M{"req_date": "$req_date", "model": "$model"},
				"count": bson.M{"$sum": 1},
			},
		},
	}

	modelsResult := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, modelsPipeline, &modelsResult); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, res := range modelsResult {

		_id := res["_id"].(map[string]interface{})

		if resultMap[gconv.String(_id["req_date"])] != nil {
			resultMap[gconv.String(_id["req_date"])].Models = append(resultMap[gconv.String(_id["req_date"])].Models, common.ModelStat{
				Model: gconv.String(_id["model"]),
				Total: gconv.Int(res["count"]),
			})
		} else {
			resultMap[gconv.String(_id["req_date"])] = &do.StatisticsApp{
				UserId:   params.UserId,
				AppId:    params.AppId,
				StatDate: gconv.String(_id["req_date"]),
				StatTime: gtime.NewFromStrFormat(gconv.String(_id["req_date"]), time.DateOnly).TimestampMilli(),
				Models: []common.ModelStat{{
					Model: gconv.String(_id["model"]),
					Total: gconv.Int(res["count"]),
				}},
			}
		}
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
				"_id":    "$req_date",
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
	for _, res := range result {
		resultMap[gconv.String(res["_id"])] = &do.StatisticsAppKey{
			UserId:   params.UserId,
			AppId:    params.AppId,
			AppKey:   params.AppKey,
			StatDate: gconv.String(res["_id"]),
			StatTime: gtime.NewFromStrFormat(gconv.String(res["_id"]), time.DateOnly).TimestampMilli(),
			Total:    gconv.Int(res["count"]),
			Tokens:   gconv.Int(res["tokens"]),
		}
	}

	abnormalPipeline := []bson.M{
		{
			"$match": bson.M{
				"creator": params.AppKey,
				"req_time": bson.M{
					"$gte": params.StatStartTime,
					"$lte": params.StatEndTime,
				},
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
				"status":         bson.M{"$ne": 1},
			},
		},
		{
			"$group": bson.M{
				"_id":   "$req_date",
				"count": bson.M{"$sum": 1},
			},
		},
	}

	abnormalResult := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, abnormalPipeline, &abnormalResult); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, res := range abnormalResult {
		if resultMap[gconv.String(res["_id"])] != nil {
			resultMap[gconv.String(res["_id"])].Abnormal = gconv.Int(res["count"])
		} else {
			resultMap[gconv.String(res["_id"])] = &do.StatisticsAppKey{
				UserId:   params.UserId,
				AppId:    params.AppId,
				AppKey:   params.AppKey,
				StatDate: gconv.String(res["_id"]),
				StatTime: gtime.NewFromStrFormat(gconv.String(res["_id"]), time.DateOnly).TimestampMilli(),
				Abnormal: gconv.Int(res["count"]),
			}
		}
	}

	modelsPipeline := []bson.M{
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
				"_id":   bson.M{"req_date": "$req_date", "model": "$model"},
				"count": bson.M{"$sum": 1},
			},
		},
	}

	modelsResult := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, modelsPipeline, &modelsResult); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	for _, res := range modelsResult {

		_id := res["_id"].(map[string]interface{})

		if resultMap[gconv.String(_id["req_date"])] != nil {
			resultMap[gconv.String(_id["req_date"])].Models = append(resultMap[gconv.String(_id["req_date"])].Models, common.ModelStat{
				Model: gconv.String(_id["model"]),
				Total: gconv.Int(res["count"]),
			})
		} else {
			resultMap[gconv.String(_id["req_date"])] = &do.StatisticsAppKey{
				UserId:   params.UserId,
				AppId:    params.AppId,
				AppKey:   params.AppKey,
				StatDate: gconv.String(_id["req_date"]),
				StatTime: gtime.NewFromStrFormat(gconv.String(_id["req_date"]), time.DateOnly).TimestampMilli(),
				Models: []common.ModelStat{{
					Model: gconv.String(_id["model"]),
					Total: gconv.Int(res["count"]),
				}},
			}
		}
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
