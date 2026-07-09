package statistics

import (
	"context"
	"sync"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	mcommon "github.com/iimeta/fastapi-admin/v2/internal/model/common"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type sStatistics struct {
	logCollections    []string
	statisticsRedsync *redsync.Redsync
}

func init() {
	service.RegisterStatistics(New())
}

func New() service.IStatistics {
	return &sStatistics{
		logCollections: []string{
			dao.LOG_TEXT,
			dao.LOG_IMAGE,
			dao.LOG_AUDIO,
			dao.LOG_VIDEO,
			dao.LOG_FILE,
			dao.LOG_BATCH,
			dao.LOG_GENERAL,
		},
		statisticsRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}

// 用户数据
func (s *sStatistics) DataUser(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error) {

	match := s.buildMatchFilter(ctx, params.StatStartTime, params.StatEndTime, params.UserId, params.AppId)

	pipeline := []bson.M{
		{"$match": match},
		{"$group": bson.M{
			"_id":      nil,
			"total":    bson.M{"$sum": "$total"},
			"tokens":   bson.M{"$sum": "$tokens"},
			"abnormal": bson.M{"$sum": "$abnormal"},
		}},
	}

	result := make([]map[string]any, 0)
	if err := dao.StatisticsUser.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	res := &model.StatisticsDataRes{}
	if len(result) > 0 {
		res.Total = gconv.Int(result[0]["total"])
		res.Tokens = common.ConvQuotaUnitReverse(gconv.Int(result[0]["tokens"]))
		res.Abnormal = gconv.Int(result[0]["abnormal"])
	}

	return res, nil
}

// 应用数据
func (s *sStatistics) DataApp(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error) {

	match := s.buildMatchFilter(ctx, params.StatStartTime, params.StatEndTime, params.UserId, params.AppId)

	pipeline := []bson.M{
		{"$match": match},
		{"$group": bson.M{
			"_id":      nil,
			"total":    bson.M{"$sum": "$total"},
			"tokens":   bson.M{"$sum": "$tokens"},
			"abnormal": bson.M{"$sum": "$abnormal"},
		}},
	}

	result := make([]map[string]any, 0)
	if err := dao.StatisticsApp.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	res := &model.StatisticsDataRes{}
	if len(result) > 0 {
		res.Total = gconv.Int(result[0]["total"])
		res.Tokens = common.ConvQuotaUnitReverse(gconv.Int(result[0]["tokens"]))
		res.Abnormal = gconv.Int(result[0]["abnormal"])
	}

	return res, nil
}

// 应用密钥数据
func (s *sStatistics) DataAppKey(ctx context.Context, params model.StatisticsDataReq) (*model.StatisticsDataRes, error) {

	match := s.buildMatchFilter(ctx, params.StatStartTime, params.StatEndTime, params.UserId, params.AppId)

	if params.AppKey != "" {
		match["app_key"] = params.AppKey
	}

	pipeline := []bson.M{
		{"$match": match},
		{"$group": bson.M{
			"_id":      nil,
			"total":    bson.M{"$sum": "$total"},
			"tokens":   bson.M{"$sum": "$tokens"},
			"abnormal": bson.M{"$sum": "$abnormal"},
		}},
	}

	result := make([]map[string]any, 0)
	if err := dao.StatisticsAppKey.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	res := &model.StatisticsDataRes{}
	if len(result) > 0 {
		res.Total = gconv.Int(result[0]["total"])
		res.Tokens = common.ConvQuotaUnitReverse(gconv.Int(result[0]["tokens"]))
		res.Abnormal = gconv.Int(result[0]["abnormal"])
	}

	return res, nil
}

// 数据看板汇总
func (s *sStatistics) DataSummary(ctx context.Context, params model.StatisticsSummaryReq) (*model.StatisticsSummaryRes, error) {

	match := s.buildMatchFilter(ctx, params.StatStartTime, params.StatEndTime, params.UserId, params.AppId, params.AppKey, gconv.String(params.Rid), params.Key)

	resolvedModels, useModels := s.resolveModels(ctx, params.Models, params.Provider)

	// 用户维度汇总
	pipeline := []bson.M{
		{"$match": match},
		{"$group": bson.M{
			"_id":                nil,
			"total":              bson.M{"$sum": "$total"},
			"tokens":             bson.M{"$sum": "$tokens"},
			"abnormal":           bson.M{"$sum": "$abnormal"},
			"input_tokens":       bson.M{"$sum": "$input_tokens"},
			"output_tokens":      bson.M{"$sum": "$output_tokens"},
			"reasoning_tokens":   bson.M{"$sum": "$reasoning_tokens"},
			"cache_read_tokens":  bson.M{"$sum": "$cache_read_tokens"},
			"cache_write_tokens": bson.M{"$sum": "$cache_write_tokens"},
			"users":              bson.M{"$addToSet": "$user_id"},
		}},
	}

	// 如果指定了模型筛选, 使用 model_stats 展开
	if useModels {
		pipeline = []bson.M{
			{"$match": match},
			{"$unwind": "$model_stats"},
			{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}},
			{"$group": bson.M{
				"_id":                nil,
				"total":              bson.M{"$sum": "$model_stats.total"},
				"tokens":             bson.M{"$sum": "$model_stats.tokens"},
				"abnormal":           bson.M{"$sum": "$model_stats.abnormal"},
				"input_tokens":       bson.M{"$sum": "$model_stats.input_tokens"},
				"output_tokens":      bson.M{"$sum": "$model_stats.output_tokens"},
				"reasoning_tokens":   bson.M{"$sum": "$model_stats.reasoning_tokens"},
				"cache_read_tokens":  bson.M{"$sum": "$model_stats.cache_read_tokens"},
				"cache_write_tokens": bson.M{"$sum": "$model_stats.cache_write_tokens"},
				"users":              bson.M{"$addToSet": "$user_id"},
			}},
		}
	}

	result := make([]map[string]any, 0)
	var err error
	// 根据 appKey 选择查询维度
	if params.AppKey != "" {
		err = dao.StatisticsAppKey.Aggregate(ctx, pipeline, &result)
	} else {
		err = dao.StatisticsUser.Aggregate(ctx, pipeline, &result)
	}
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	res := &model.StatisticsSummaryRes{}
	if len(result) > 0 {
		res.Total = gconv.Int(result[0]["total"])
		res.Tokens = common.ConvQuotaUnitReverse(gconv.Int(result[0]["tokens"]))
		res.Abnormal = gconv.Int(result[0]["abnormal"])
		res.ActiveUsers = len(gconv.SliceAny(result[0]["users"]))
		res.InputTokens = gconv.Int64(result[0]["input_tokens"])
		res.OutputTokens = gconv.Int64(result[0]["output_tokens"])
		res.ReasoningTokens = gconv.Int64(result[0]["reasoning_tokens"])
		res.CacheReadTokens = gconv.Int64(result[0]["cache_read_tokens"])
		res.CacheWriteTokens = gconv.Int64(result[0]["cache_write_tokens"])
		if res.Total > 0 {
			res.AbnormalRate = float64(res.Abnormal) / float64(res.Total) * 100
		}
	}

	// 应用维度获取活跃应用数
	appMatch := s.buildMatchFilter(ctx, params.StatStartTime, params.StatEndTime, params.UserId, params.AppId, params.AppKey, gconv.String(params.Rid), params.Key)
	appPipeline := []bson.M{
		{"$match": appMatch},
		{"$group": bson.M{
			"_id":  nil,
			"apps": bson.M{"$addToSet": "$app_id"},
		}},
	}

	// 如果指定了模型筛选,需要先展开model_stats过滤
	if useModels {
		appPipeline = []bson.M{
			{"$match": appMatch},
			{"$unwind": "$model_stats"},
			{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}},
			{"$group": bson.M{
				"_id":  nil,
				"apps": bson.M{"$addToSet": "$app_id"},
			}},
		}
	}

	appResult := make([]map[string]any, 0)
	if err := dao.StatisticsApp.Aggregate(ctx, appPipeline, &appResult); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if len(appResult) > 0 {
		res.ActiveApps = len(gconv.SliceAny(appResult[0]["apps"]))
	}

	// 环比: 查询上一周期数据
	if params.StatStartTime > 0 && params.StatEndTime > 0 {
		duration := params.StatEndTime - params.StatStartTime
		prevStart := params.StatStartTime - duration
		prevEnd := params.StatStartTime - 1

		prevMatch := s.buildMatchFilter(ctx, prevStart, prevEnd, params.UserId, params.AppId, params.AppKey, gconv.String(params.Rid), params.Key)

		prevPipeline := []bson.M{
			{"$match": prevMatch},
			{"$group": bson.M{
				"_id":      nil,
				"total":    bson.M{"$sum": "$total"},
				"tokens":   bson.M{"$sum": "$tokens"},
				"abnormal": bson.M{"$sum": "$abnormal"},
			}},
		}

		if useModels {
			prevPipeline = []bson.M{
				{"$match": prevMatch},
				{"$unwind": "$model_stats"},
				{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}},
				{"$group": bson.M{
					"_id":      nil,
					"total":    bson.M{"$sum": "$model_stats.total"},
					"tokens":   bson.M{"$sum": "$model_stats.tokens"},
					"abnormal": bson.M{"$sum": "$model_stats.abnormal"},
				}},
			}
		}

		prevResult := make([]map[string]any, 0)
		var prevErr error
		if params.AppKey != "" {
			prevErr = dao.StatisticsAppKey.Aggregate(ctx, prevPipeline, &prevResult)
		} else {
			prevErr = dao.StatisticsUser.Aggregate(ctx, prevPipeline, &prevResult)
		}
		if prevErr != nil {
			logger.Error(ctx, prevErr)
		} else if len(prevResult) > 0 {
			res.PrevTotal = gconv.Int(prevResult[0]["total"])
			res.PrevTokens = common.ConvQuotaUnitReverse(gconv.Int(prevResult[0]["tokens"]))
			res.PrevAbnormal = gconv.Int(prevResult[0]["abnormal"])
			if res.PrevTotal > 0 {
				res.PrevAbnormalRate = float64(res.PrevAbnormal) / float64(res.PrevTotal) * 100
			}
		}
	}

	return res, nil
}

// 数据看板趋势
func (s *sStatistics) DataTrend(ctx context.Context, params model.StatisticsTrendReq) (*model.StatisticsTrendRes, error) {

	match := s.buildMatchFilter(ctx, params.StatStartTime, params.StatEndTime, params.UserId, params.AppId, params.AppKey, gconv.String(params.Rid), params.Key)

	resolvedModels, useModels := s.resolveModels(ctx, params.Models, params.Provider)

	pipeline := []bson.M{
		{"$match": match},
		{"$group": bson.M{
			"_id":      "$stat_date",
			"total":    bson.M{"$sum": "$total"},
			"tokens":   bson.M{"$sum": "$tokens"},
			"abnormal": bson.M{"$sum": "$abnormal"},
			"users":    bson.M{"$addToSet": "$user_id"},
		}},
	}

	if useModels {
		pipeline = []bson.M{
			{"$match": match},
			{"$unwind": "$model_stats"},
			{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}},
			{"$group": bson.M{
				"_id":      "$stat_date",
				"total":    bson.M{"$sum": "$model_stats.total"},
				"tokens":   bson.M{"$sum": "$model_stats.tokens"},
				"abnormal": bson.M{"$sum": "$model_stats.abnormal"},
				"users":    bson.M{"$addToSet": "$user_id"},
			}},
		}
	}

	result := make([]map[string]any, 0)
	var err error
	if params.AppKey != "" {
		err = dao.StatisticsAppKey.Aggregate(ctx, pipeline, &result)
	} else {
		err = dao.StatisticsUser.Aggregate(ctx, pipeline, &result)
	}
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	// 应用维度获取每日活跃应用数
	appMatch := s.buildMatchFilter(ctx, params.StatStartTime, params.StatEndTime, params.UserId, params.AppId, params.AppKey, gconv.String(params.Rid), params.Key)
	appPipeline := []bson.M{
		{"$match": appMatch},
		{"$group": bson.M{
			"_id":  "$stat_date",
			"apps": bson.M{"$addToSet": "$app_id"},
		}},
	}

	if useModels {
		appPipeline = []bson.M{
			{"$match": appMatch},
			{"$unwind": "$model_stats"},
			{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}},
			{"$group": bson.M{
				"_id":  "$stat_date",
				"apps": bson.M{"$addToSet": "$app_id"},
			}},
		}
	}

	appResult := make([]map[string]any, 0)
	if err := dao.StatisticsApp.Aggregate(ctx, appPipeline, &appResult); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	appMap := make(map[string]int)
	for _, res := range appResult {
		appMap[gconv.String(res["_id"])] = len(gconv.SliceAny(res["apps"]))
	}

	resultMap := make(map[string]*model.StatisticsTrendItem)
	for _, res := range result {
		date := gconv.String(res["_id"])
		total := gconv.Int(res["total"])
		abnormal := gconv.Int(res["abnormal"])
		abnormalRate := float64(0)
		if total > 0 {
			abnormalRate = float64(abnormal) / float64(total) * 100
		}
		resultMap[date] = &model.StatisticsTrendItem{
			Date:         date[5:],
			Total:        total,
			Tokens:       common.ConvQuotaUnitReverse(gconv.Int(res["tokens"])),
			Abnormal:     abnormal,
			AbnormalRate: abnormalRate,
			ActiveUsers:  len(gconv.SliceAny(res["users"])),
			ActiveApps:   appMap[date],
		}
	}

	// 填充空日期
	var days []*util.DateTime
	if params.StatStartTime > 0 && params.StatEndTime > 0 {
		startTime := gtime.NewFromTimeStamp(params.StatStartTime / 1000)
		endTime := gtime.NewFromTimeStamp(params.StatEndTime / 1000)
		days = util.Day(startTime.String(), endTime.String())
	} else {
		// 全部历史模式: 从结果中取最早和最晚日期
		var minDate, maxDate string
		for date := range resultMap {
			if minDate == "" || date < minDate {
				minDate = date
			}
			if maxDate == "" || date > maxDate {
				maxDate = date
			}
		}
		if minDate != "" && maxDate != "" {
			days = util.Day(minDate+" 00:00:00", maxDate+" 23:59:59")
		}
	}

	items := make([]*model.StatisticsTrendItem, 0)
	for _, day := range days {
		item := resultMap[day.StartDate]
		if item == nil {
			item = &model.StatisticsTrendItem{Date: day.StartDate[5:]}
		}
		items = append(items, item)
	}

	return &model.StatisticsTrendRes{Items: items}, nil
}

// 数据看板模型分布
func (s *sStatistics) DataModelPercent(ctx context.Context, params model.StatisticsModelPercentReq) (*model.StatisticsModelPercentRes, error) {

	match := s.buildMatchFilter(ctx, params.StatStartTime, params.StatEndTime, params.UserId, params.AppId, params.AppKey, gconv.String(params.Rid), params.Key)

	resolvedModels, useModels := s.resolveModels(ctx, params.Models, params.Provider)

	sumField := "$model_stats.total"
	if params.DataType == "tokens" {
		sumField = "$model_stats.tokens"
	}

	pipeline := []bson.M{
		{"$match": match},
		{"$unwind": "$model_stats"},
	}

	// 模型筛选
	if useModels {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}})
	}

	pipeline = append(pipeline,
		bson.M{"$group": bson.M{
			"_id":   "$model_stats.model",
			"count": bson.M{"$sum": sumField},
		}},
		bson.M{"$sort": bson.M{"count": -1}},
		bson.M{"$limit": 10},
	)

	result := make([]map[string]any, 0)
	var err error
	if params.AppKey != "" {
		err = dao.StatisticsAppKey.Aggregate(ctx, pipeline, &result)
	} else {
		err = dao.StatisticsUser.Aggregate(ctx, pipeline, &result)
	}
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.ModelPercent, 0)
	models := make([]string, 0)
	for _, res := range result {
		name := gconv.String(res["_id"])
		value := gconv.Float64(res["count"])
		if params.DataType == "tokens" {
			value = common.ConvQuotaUnitReverse(gconv.Int(res["count"]))
		}
		items = append(items, &model.ModelPercent{
			Name:  name,
			Value: value,
		})
		models = append(models, name)
	}

	return &model.StatisticsModelPercentRes{
		Models: models,
		Items:  items,
	}, nil
}

// 数据看板排行
func (s *sStatistics) DataTop(ctx context.Context, params model.StatisticsTopReq) (*model.StatisticsTopRes, error) {

	match := s.buildMatchFilter(ctx, params.StatStartTime, params.StatEndTime, params.UserId, params.AppId, params.AppKey, gconv.String(params.Rid), params.Key)

	resolvedModels, useModels := s.resolveModels(ctx, params.Models, params.Provider)

	limit := params.Limit
	if limit <= 0 {
		limit = 10
	}

	pipeline := []bson.M{{"$match": match}}

	// 如果指定了models筛选,需要先unwind model_stats并过滤
	needUnwindForModels := useModels && params.DataType != "model"

	switch params.DataType {
	case "user":
		if needUnwindForModels {
			pipeline = append(pipeline, bson.M{"$unwind": "$model_stats"})
			pipeline = append(pipeline, bson.M{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}})
			pipeline = append(pipeline, bson.M{
				"$group": bson.M{
					"_id":    "$user_id",
					"count":  bson.M{"$sum": "$model_stats.total"},
					"tokens": bson.M{"$sum": "$model_stats.tokens"},
				},
			})
		} else {
			pipeline = append(pipeline, bson.M{
				"$group": bson.M{
					"_id":    "$user_id",
					"count":  bson.M{"$sum": "$total"},
					"tokens": bson.M{"$sum": "$tokens"},
				},
			})
		}
	case "app":
		if needUnwindForModels {
			pipeline = append(pipeline, bson.M{"$unwind": "$model_stats"})
			pipeline = append(pipeline, bson.M{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}})
			pipeline = append(pipeline, bson.M{
				"$group": bson.M{
					"_id":     "$app_id",
					"user_id": bson.M{"$first": "$user_id"},
					"count":   bson.M{"$sum": "$model_stats.total"},
					"tokens":  bson.M{"$sum": "$model_stats.tokens"},
				},
			})
		} else {
			pipeline = append(pipeline, bson.M{
				"$group": bson.M{
					"_id":     "$app_id",
					"user_id": bson.M{"$first": "$user_id"},
					"count":   bson.M{"$sum": "$total"},
					"tokens":  bson.M{"$sum": "$tokens"},
				},
			})
		}
	case "app_key":
		if needUnwindForModels {
			pipeline = append(pipeline, bson.M{"$unwind": "$model_stats"})
			pipeline = append(pipeline, bson.M{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}})
			pipeline = append(pipeline, bson.M{
				"$group": bson.M{
					"_id":    "$app_key",
					"app_id": bson.M{"$first": "$app_id"},
					"count":  bson.M{"$sum": "$model_stats.total"},
					"tokens": bson.M{"$sum": "$model_stats.tokens"},
				},
			})
		} else {
			pipeline = append(pipeline, bson.M{
				"$group": bson.M{
					"_id":    "$app_key",
					"app_id": bson.M{"$first": "$app_id"},
					"count":  bson.M{"$sum": "$total"},
					"tokens": bson.M{"$sum": "$tokens"},
				},
			})
		}
	case "model":
		pipeline = append(pipeline, bson.M{"$unwind": "$model_stats"})
		if useModels {
			pipeline = append(pipeline, bson.M{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}})
		}
		pipeline = append(pipeline, bson.M{
			"$group": bson.M{
				"_id":    "$model_stats.model",
				"count":  bson.M{"$sum": "$model_stats.total"},
				"tokens": bson.M{"$sum": "$model_stats.tokens"},
			},
		})
	}

	// provider 维度需要查原始日志
	if params.DataType == "provider" {
		logMatch := bson.M{
			"is_smart_match": bson.M{"$ne": true},
			"is_retry":       bson.M{"$ne": true},
		}
		if params.StatStartTime > 0 {
			logMatch["created_at"] = bson.M{"$gte": params.StatStartTime}
		}
		if params.StatEndTime > 0 {
			if existing, ok := logMatch["created_at"]; ok {
				existing.(bson.M)["$lte"] = params.StatEndTime
			} else {
				logMatch["created_at"] = bson.M{"$lte": params.StatEndTime}
			}
		}
		if service.Session().IsResellerRole(ctx) {
			logMatch["rid"] = service.Session().GetRid(ctx)
		}
		if service.Session().IsUserRole(ctx) {
			logMatch["user_id"] = service.Session().GetUserId(ctx)
		}
		if service.Session().IsAdminRole(ctx) && params.Rid > 0 {
			logMatch["rid"] = params.Rid
		}
		if params.UserId > 0 {
			logMatch["user_id"] = params.UserId
		}
		if params.AppId > 0 {
			logMatch["app_id"] = params.AppId
		}
		if params.AppKey != "" {
			logMatch["app_key"] = params.AppKey
		}
		// 管理员可通过key筛选
		if params.Key != "" && service.Session().IsAdminRole(ctx) {
			logMatch["app_key"] = params.Key
		}
		// 模型筛选(models 优先级高于 provider)
		if len(params.Models) > 0 {
			logMatch["model"] = bson.M{"$in": params.Models}
		} else if params.Provider != "" {
			// 提供商筛选(按 provider_id)
			logMatch["provider_id"] = params.Provider
		}

		provPipeline := []bson.M{
			{"$match": logMatch},
			{"$group": bson.M{
				"_id":           "$provider_name",
				"provider_id":   bson.M{"$first": "$provider_id"},
				"provider_name": bson.M{"$first": "$provider_name"},
				"count":         bson.M{"$sum": 1},
				"tokens":        bson.M{"$sum": "$spend.total_spend_tokens"},
			}},
			{"$sort": bson.M{"count": -1}},
			{"$limit": limit},
		}

		provResult := make([]map[string]any, 0)
		if err := dao.LogText.Aggregate(ctx, provPipeline, &provResult); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		items := make([]*model.DataTop, 0)
		for _, res := range provResult {
			items = append(items, &model.DataTop{
				Provider:   gconv.String(res["provider_name"]),
				ProviderId: gconv.String(res["provider_id"]),
				Call:       gconv.Int(res["count"]),
				Tokens:     common.ConvQuotaUnitReverse(gconv.Int(res["tokens"])),
			})
		}
		return &model.StatisticsTopRes{Items: items}, nil
	}

	pipeline = append(pipeline, bson.M{"$sort": bson.M{"tokens": -1}}, bson.M{"$limit": limit})

	result := make([]map[string]any, 0)
	switch params.DataType {
	case "user":
		if err := dao.StatisticsUser.Aggregate(ctx, pipeline, &result); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	case "app":
		if err := dao.StatisticsApp.Aggregate(ctx, pipeline, &result); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	case "app_key":
		if err := dao.StatisticsAppKey.Aggregate(ctx, pipeline, &result); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	case "model":
		if err := dao.StatisticsUser.Aggregate(ctx, pipeline, &result); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	items := make([]*model.DataTop, 0)
	switch params.DataType {
	case "user":
		for _, res := range result {
			items = append(items, &model.DataTop{
				UserId: gconv.Int(res["_id"]),
				Call:   gconv.Int(res["count"]),
				Tokens: common.ConvQuotaUnitReverse(gconv.Int(res["tokens"])),
			})
		}
	case "app":
		for _, res := range result {
			items = append(items, &model.DataTop{
				AppId:  gconv.Int(res["_id"]),
				UserId: gconv.Int(res["user_id"]),
				Call:   gconv.Int(res["count"]),
				Tokens: common.ConvQuotaUnitReverse(gconv.Int(res["tokens"])),
			})
		}
	case "app_key":
		for _, res := range result {
			raw := gconv.String(res["_id"])
			items = append(items, &model.DataTop{
				AppKey:    util.Desensitize(raw),
				AppKeyRaw: raw,
				AppId:     gconv.Int(res["app_id"]),
				Call:      gconv.Int(res["count"]),
				Tokens:    common.ConvQuotaUnitReverse(gconv.Int(res["tokens"])),
			})
		}
	case "model":
		for _, res := range result {
			items = append(items, &model.DataTop{
				Model:  gconv.String(res["_id"]),
				Call:   gconv.Int(res["count"]),
				Tokens: common.ConvQuotaUnitReverse(gconv.Int(res["tokens"])),
			})
		}
	}

	return &model.StatisticsTopRes{Items: items}, nil
}

// 数据看板明细
func (s *sStatistics) DataDetail(ctx context.Context, params model.StatisticsDetailReq) (*model.StatisticsDetailRes, error) {

	filter := bson.M{}

	resolvedModels, useModels := s.resolveModels(ctx, params.Models, params.Provider)

	if params.StatStartTime > 0 && params.StatEndTime > 0 {
		filter["stat_time"] = bson.M{
			"$gte": params.StatStartTime,
			"$lte": params.StatEndTime,
		}
	} else if params.StatStartTime > 0 {
		filter["stat_time"] = bson.M{"$gte": params.StatStartTime}
	} else if params.StatEndTime > 0 {
		filter["stat_time"] = bson.M{"$lte": params.StatEndTime}
	}

	// 角色过滤
	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
	}

	// 管理员可按代理商筛选
	if service.Session().IsAdminRole(ctx) && params.Rid != 0 {
		filter["rid"] = params.Rid
	}

	if !service.Session().IsUserRole(ctx) && params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	if params.AppKey != "" {
		filter["app_key"] = params.AppKey
	}

	// 管理员可通过key筛选(匹配app_key字段)
	if params.Key != "" && service.Session().IsAdminRole(ctx) {
		filter["app_key"] = params.Key
	}

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	if paging.Page <= 0 {
		paging.Page = 1
	}
	if paging.PageSize <= 0 {
		paging.PageSize = 10
	}

	findOptions := &dao.FindOptions{
		SortFields: []string{"-stat_time", "-total"},
	}

	items := make([]*model.StatisticsDetailItem, 0)

	// 如果指定了models筛选,需要使用聚合管道
	needAggregateForModels := useModels && params.DataType != "model"

	switch params.DataType {
	case "user":
		if needAggregateForModels {
			// 使用聚合管道过滤models
			pipeline := []bson.M{
				{"$match": filter},
				{"$unwind": "$model_stats"},
				{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}},
				{"$group": bson.M{
					"_id":         bson.M{"user_id": "$user_id", "stat_date": "$stat_date"},
					"user_id":     bson.M{"$first": "$user_id"},
					"stat_date":   bson.M{"$first": "$stat_date"},
					"total":       bson.M{"$sum": "$model_stats.total"},
					"tokens":      bson.M{"$sum": "$model_stats.tokens"},
					"abnormal":    bson.M{"$sum": "$model_stats.abnormal"},
					"model_stats": bson.M{"$push": "$model_stats"},
				}},
				{"$sort": bson.M{"stat_date": -1, "total": -1}},
				{"$skip": (paging.Page - 1) * paging.PageSize},
				{"$limit": paging.PageSize},
			}
			result := make([]map[string]any, 0)
			if err := dao.StatisticsUser.Aggregate(ctx, pipeline, &result); err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			// 计算总数
			countPipeline := []bson.M{
				{"$match": filter},
				{"$unwind": "$model_stats"},
				{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}},
				{"$group": bson.M{"_id": bson.M{"user_id": "$user_id", "stat_date": "$stat_date"}}},
				{"$count": "total"},
			}
			countResult := make([]map[string]any, 0)
			if err := dao.StatisticsUser.Aggregate(ctx, countPipeline, &countResult); err == nil && len(countResult) > 0 {
				paging.Total = gconv.Int64(countResult[0]["total"])
			}
			for _, r := range result {
				total := gconv.Int(r["total"])
				abnormal := gconv.Int(r["abnormal"])
				abnormalRate := float64(0)
				if total > 0 {
					abnormalRate = float64(abnormal) / float64(total) * 100
				}
				items = append(items, &model.StatisticsDetailItem{
					UserId:       gconv.Int(r["user_id"]),
					StatDate:     gconv.String(r["stat_date"]),
					Total:        total,
					Tokens:       common.ConvQuotaUnitReverse(gconv.Int(r["tokens"])),
					Abnormal:     abnormal,
					AbnormalRate: abnormalRate,
					ModelStats:   convModelStats(r["model_stats"]),
				})
			}
		} else {
			results, err := dao.StatisticsUser.FindByPage(ctx, paging, filter, findOptions)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			for _, r := range results {
				total := r.Total
				abnormal := r.Abnormal
				abnormalRate := float64(0)
				if total > 0 {
					abnormalRate = float64(abnormal) / float64(total) * 100
				}
				items = append(items, &model.StatisticsDetailItem{
					UserId:       r.UserId,
					StatDate:     r.StatDate,
					Total:        total,
					Tokens:       common.ConvQuotaUnitReverse(r.Tokens),
					Abnormal:     abnormal,
					AbnormalRate: abnormalRate,
					ModelStats:   convModelStats(r.ModelStats),
				})
			}
		}

	case "app":
		if needAggregateForModels {
			pipeline := []bson.M{
				{"$match": filter},
				{"$unwind": "$model_stats"},
				{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}},
				{"$group": bson.M{
					"_id":         bson.M{"user_id": "$user_id", "app_id": "$app_id", "stat_date": "$stat_date"},
					"user_id":     bson.M{"$first": "$user_id"},
					"app_id":      bson.M{"$first": "$app_id"},
					"stat_date":   bson.M{"$first": "$stat_date"},
					"total":       bson.M{"$sum": "$model_stats.total"},
					"tokens":      bson.M{"$sum": "$model_stats.tokens"},
					"abnormal":    bson.M{"$sum": "$model_stats.abnormal"},
					"model_stats": bson.M{"$push": "$model_stats"},
				}},
				{"$sort": bson.M{"stat_date": -1, "total": -1}},
				{"$skip": (paging.Page - 1) * paging.PageSize},
				{"$limit": paging.PageSize},
			}
			result := make([]map[string]any, 0)
			if err := dao.StatisticsApp.Aggregate(ctx, pipeline, &result); err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			countPipeline := []bson.M{
				{"$match": filter},
				{"$unwind": "$model_stats"},
				{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}},
				{"$group": bson.M{"_id": bson.M{"user_id": "$user_id", "app_id": "$app_id", "stat_date": "$stat_date"}}},
				{"$count": "total"},
			}
			countResult := make([]map[string]any, 0)
			if err := dao.StatisticsApp.Aggregate(ctx, countPipeline, &countResult); err == nil && len(countResult) > 0 {
				paging.Total = gconv.Int64(countResult[0]["total"])
			}
			for _, r := range result {
				total := gconv.Int(r["total"])
				abnormal := gconv.Int(r["abnormal"])
				abnormalRate := float64(0)
				if total > 0 {
					abnormalRate = float64(abnormal) / float64(total) * 100
				}
				items = append(items, &model.StatisticsDetailItem{
					UserId:       gconv.Int(r["user_id"]),
					AppId:        gconv.Int(r["app_id"]),
					StatDate:     gconv.String(r["stat_date"]),
					Total:        total,
					Tokens:       common.ConvQuotaUnitReverse(gconv.Int(r["tokens"])),
					Abnormal:     abnormal,
					AbnormalRate: abnormalRate,
					ModelStats:   convModelStats(r["model_stats"]),
				})
			}
		} else {
			results, err := dao.StatisticsApp.FindByPage(ctx, paging, filter, findOptions)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			for _, r := range results {
				total := r.Total
				abnormal := r.Abnormal
				abnormalRate := float64(0)
				if total > 0 {
					abnormalRate = float64(abnormal) / float64(total) * 100
				}
				items = append(items, &model.StatisticsDetailItem{
					UserId:       r.UserId,
					AppId:        r.AppId,
					StatDate:     r.StatDate,
					Total:        total,
					Tokens:       common.ConvQuotaUnitReverse(r.Tokens),
					Abnormal:     abnormal,
					AbnormalRate: abnormalRate,
					ModelStats:   convModelStats(r.ModelStats),
				})
			}
		}

	case "app_key":
		if needAggregateForModels {
			pipeline := []bson.M{
				{"$match": filter},
				{"$unwind": "$model_stats"},
				{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}},
				{"$group": bson.M{
					"_id":         bson.M{"user_id": "$user_id", "app_id": "$app_id", "app_key": "$app_key", "stat_date": "$stat_date"},
					"user_id":     bson.M{"$first": "$user_id"},
					"app_id":      bson.M{"$first": "$app_id"},
					"app_key":     bson.M{"$first": "$app_key"},
					"stat_date":   bson.M{"$first": "$stat_date"},
					"total":       bson.M{"$sum": "$model_stats.total"},
					"tokens":      bson.M{"$sum": "$model_stats.tokens"},
					"abnormal":    bson.M{"$sum": "$model_stats.abnormal"},
					"model_stats": bson.M{"$push": "$model_stats"},
				}},
				{"$sort": bson.M{"stat_date": -1, "total": -1}},
				{"$skip": (paging.Page - 1) * paging.PageSize},
				{"$limit": paging.PageSize},
			}
			result := make([]map[string]any, 0)
			if err := dao.StatisticsAppKey.Aggregate(ctx, pipeline, &result); err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			countPipeline := []bson.M{
				{"$match": filter},
				{"$unwind": "$model_stats"},
				{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}},
				{"$group": bson.M{"_id": bson.M{"user_id": "$user_id", "app_id": "$app_id", "app_key": "$app_key", "stat_date": "$stat_date"}}},
				{"$count": "total"},
			}
			countResult := make([]map[string]any, 0)
			if err := dao.StatisticsAppKey.Aggregate(ctx, countPipeline, &countResult); err == nil && len(countResult) > 0 {
				paging.Total = gconv.Int64(countResult[0]["total"])
			}
			for _, r := range result {
				total := gconv.Int(r["total"])
				abnormal := gconv.Int(r["abnormal"])
				abnormalRate := float64(0)
				if total > 0 {
					abnormalRate = float64(abnormal) / float64(total) * 100
				}
				items = append(items, &model.StatisticsDetailItem{
					UserId:       gconv.Int(r["user_id"]),
					AppId:        gconv.Int(r["app_id"]),
					AppKey:       util.Desensitize(gconv.String(r["app_key"])),
					StatDate:     gconv.String(r["stat_date"]),
					Total:        total,
					Tokens:       common.ConvQuotaUnitReverse(gconv.Int(r["tokens"])),
					Abnormal:     abnormal,
					AbnormalRate: abnormalRate,
					ModelStats:   convModelStats(r["model_stats"]),
				})
			}
		} else {
			results, err := dao.StatisticsAppKey.FindByPage(ctx, paging, filter, findOptions)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			for _, r := range results {
				total := r.Total
				abnormal := r.Abnormal
				abnormalRate := float64(0)
				if total > 0 {
					abnormalRate = float64(abnormal) / float64(total) * 100
				}
				items = append(items, &model.StatisticsDetailItem{
					UserId:       r.UserId,
					AppId:        r.AppId,
					AppKey:       util.Desensitize(r.AppKey),
					StatDate:     r.StatDate,
					Total:        total,
					Tokens:       common.ConvQuotaUnitReverse(r.Tokens),
					Abnormal:     abnormal,
					AbnormalRate: abnormalRate,
					ModelStats:   convModelStats(r.ModelStats),
				})
			}
		}

	case "model":
		// 模型维度: 按用户聚合该模型的调用数据
		if params.ModelId == "" {
			break
		}
		match := bson.M{}
		if params.StatStartTime > 0 && params.StatEndTime > 0 {
			match["stat_time"] = bson.M{"$gte": params.StatStartTime, "$lte": params.StatEndTime}
		} else if params.StatStartTime > 0 {
			match["stat_time"] = bson.M{"$gte": params.StatStartTime}
		} else if params.StatEndTime > 0 {
			match["stat_time"] = bson.M{"$lte": params.StatEndTime}
		}

		if service.Session().IsResellerRole(ctx) {
			match["rid"] = service.Session().GetRid(ctx)
		}
		if service.Session().IsUserRole(ctx) {
			match["user_id"] = service.Session().GetUserId(ctx)
		}
		if service.Session().IsAdminRole(ctx) && params.Rid != 0 {
			match["rid"] = params.Rid
		}
		if !service.Session().IsUserRole(ctx) && params.UserId != 0 {
			match["user_id"] = params.UserId
		}
		if params.AppId != 0 {
			match["app_id"] = params.AppId
		}
		if params.AppKey != "" {
			match["app_key"] = params.AppKey
		}
		// 管理员可通过key筛选
		if params.Key != "" && service.Session().IsAdminRole(ctx) {
			match["app_key"] = params.Key
		}

		pipeline := []bson.M{
			{"$match": match},
			{"$unwind": "$model_stats"},
			{"$match": bson.M{"model_stats.model": params.ModelId}},
			{"$group": bson.M{
				"_id":      "$user_id",
				"total":    bson.M{"$sum": "$model_stats.total"},
				"tokens":   bson.M{"$sum": "$model_stats.tokens"},
				"abnormal": bson.M{"$sum": "$model_stats.abnormal"},
			}},
			{"$sort": bson.M{"total": -1}},
		}

		result := make([]map[string]any, 0)
		if err := dao.StatisticsUser.Aggregate(ctx, pipeline, &result); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		for _, r := range result {
			total := gconv.Int(r["total"])
			abnormal := gconv.Int(r["abnormal"])
			abnormalRate := float64(0)
			if total > 0 {
				abnormalRate = float64(abnormal) / float64(total) * 100
			}
			items = append(items, &model.StatisticsDetailItem{
				UserId:       gconv.Int(r["_id"]),
				Model:        params.ModelId,
				Total:        total,
				Tokens:       common.ConvQuotaUnitReverse(gconv.Int(r["tokens"])),
				Abnormal:     abnormal,
				AbnormalRate: abnormalRate,
			})
		}

		paging.Total = int64(len(items))
		// 手动分页
		start := int((paging.Page - 1) * paging.PageSize)
		end := start + int(paging.PageSize)
		if start > len(items) {
			items = items[:0]
		} else {
			if end > len(items) {
				end = len(items)
			}
			items = items[start:end]
		}

	case "provider":
		// 提供商维度: 按日期聚合该提供商下的调用数据, 从日志表查询
		if params.Provider == "" {
			break
		}
		logMatch := bson.M{
			"is_smart_match": bson.M{"$ne": true},
			"is_retry":       bson.M{"$ne": true},
			"provider_id":    params.Provider,
		}
		if params.StatStartTime > 0 {
			logMatch["created_at"] = bson.M{"$gte": params.StatStartTime}
		}
		if params.StatEndTime > 0 {
			if existing, ok := logMatch["created_at"]; ok {
				existing.(bson.M)["$lte"] = params.StatEndTime
			} else {
				logMatch["created_at"] = bson.M{"$lte": params.StatEndTime}
			}
		}
		if service.Session().IsResellerRole(ctx) {
			logMatch["rid"] = service.Session().GetRid(ctx)
		}
		if service.Session().IsUserRole(ctx) {
			logMatch["user_id"] = service.Session().GetUserId(ctx)
		}
		if service.Session().IsAdminRole(ctx) && params.Rid != 0 {
			logMatch["rid"] = params.Rid
		}
		if !service.Session().IsUserRole(ctx) && params.UserId != 0 {
			logMatch["user_id"] = params.UserId
		}
		if params.AppId != 0 {
			logMatch["app_id"] = params.AppId
		}
		if params.AppKey != "" {
			logMatch["app_key"] = params.AppKey
		}
		if params.Key != "" && service.Session().IsAdminRole(ctx) {
			logMatch["app_key"] = params.Key
		}
		if len(params.Models) > 0 {
			logMatch["model"] = bson.M{"$in": params.Models}
		}

		// 按 日期+模型 分组
		logPipeline := []bson.M{
			{"$match": logMatch},
			{"$group": bson.M{
				"_id": bson.M{
					"date": bson.M{
						"$dateToString": bson.M{
							"format": "%Y-%m-%d",
							"date":   bson.M{"$toDate": "$created_at"},
						},
					},
					"model": "$model",
				},
				"total":    bson.M{"$sum": 1},
				"tokens":   bson.M{"$sum": "$spend.total_spend_tokens"},
				"abnormal": bson.M{"$sum": bson.M{"$cond": bson.A{bson.M{"$eq": bson.A{"$status", -1}}, 1, 0}}},
			}},
		}

		logResults := s.aggregateAllLogs(ctx, logPipeline)

		// 按 date 聚合; model_stats收集该天各模型的数据
		type dayAgg struct {
			date       string
			total      int
			tokens     float64
			abnormal   int
			modelStats map[string]*mcommon.ModelStat
		}
		dayMap := make(map[string]*dayAgg)
		for _, r := range logResults {
			var date, modelName string
			switch id := r["_id"].(type) {
			case bson.D:
				for _, elem := range id {
					switch elem.Key {
					case "date":
						date = gconv.String(elem.Value)
					case "model":
						modelName = gconv.String(elem.Value)
					}
				}
			case bson.M:
				date = gconv.String(id["date"])
				modelName = gconv.String(id["model"])
			case map[string]any:
				date = gconv.String(id["date"])
				modelName = gconv.String(id["model"])
			}
			if date == "" {
				continue
			}
			total := gconv.Int(r["total"])
			tokens := gconv.Float64(r["tokens"])
			abnormal := gconv.Int(r["abnormal"])

			agg, ok := dayMap[date]
			if !ok {
				agg = &dayAgg{date: date, modelStats: make(map[string]*mcommon.ModelStat)}
				dayMap[date] = agg
			}
			agg.total += total
			agg.tokens += tokens
			agg.abnormal += abnormal
			if modelName != "" {
				ms, exist := agg.modelStats[modelName]
				if !exist {
					ms = &mcommon.ModelStat{Model: modelName}
					agg.modelStats[modelName] = ms
				}
				ms.Total += total
				ms.Tokens += tokens
				ms.Abnormal += abnormal
			}
		}

		// 按日期倒序
		dates := make([]string, 0, len(dayMap))
		for d := range dayMap {
			dates = append(dates, d)
		}
		for i := 0; i < len(dates); i++ {
			for j := i + 1; j < len(dates); j++ {
				if dates[j] > dates[i] {
					dates[i], dates[j] = dates[j], dates[i]
				}
			}
		}

		allItems := make([]*model.StatisticsDetailItem, 0, len(dates))
		for _, d := range dates {
			agg := dayMap[d]
			abnormalRate := float64(0)
			if agg.total > 0 {
				abnormalRate = float64(agg.abnormal) / float64(agg.total) * 100
			}
			stats := make([]*mcommon.ModelStat, 0, len(agg.modelStats))
			for _, ms := range agg.modelStats {
				ms.Tokens = common.ConvQuotaUnitReverse(int(ms.Tokens))
				stats = append(stats, ms)
			}
			allItems = append(allItems, &model.StatisticsDetailItem{
				StatDate:     d,
				Total:        agg.total,
				Tokens:       common.ConvQuotaUnitReverse(int(agg.tokens)),
				Abnormal:     agg.abnormal,
				AbnormalRate: abnormalRate,
				ModelStats:   stats,
			})
		}

		paging.Total = int64(len(allItems))
		// 手动分页
		startIdx := int((paging.Page - 1) * paging.PageSize)
		endIdx := startIdx + int(paging.PageSize)
		if startIdx > len(allItems) {
			items = items[:0]
		} else {
			if endIdx > len(allItems) {
				endIdx = len(allItems)
			}
			items = allItems[startIdx:endIdx]
		}

	default:
		// 默认用户维度
		results, err := dao.StatisticsUser.FindByPage(ctx, paging, filter, findOptions)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
		for _, r := range results {
			total := r.Total
			abnormal := r.Abnormal
			abnormalRate := float64(0)
			if total > 0 {
				abnormalRate = float64(abnormal) / float64(total) * 100
			}
			items = append(items, &model.StatisticsDetailItem{
				UserId:       r.UserId,
				StatDate:     r.StatDate,
				Total:        total,
				Tokens:       common.ConvQuotaUnitReverse(r.Tokens),
				Abnormal:     abnormal,
				AbnormalRate: abnormalRate,
				ModelStats:   convModelStats(r.ModelStats),
			})
		}
	}

	return &model.StatisticsDetailRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 数据看板全局总览
func (s *sStatistics) DataOverview(ctx context.Context, params model.StatisticsOverviewReq) (*model.StatisticsOverviewRes, error) {

	res := &model.StatisticsOverviewRes{}

	// 历史总调用/花费/异常
	match := s.buildMatchFilter(ctx, 0, 0, params.UserId, params.AppId, params.AppKey, gconv.String(params.Rid))
	pipeline := []bson.M{
		{"$match": match},
		{"$group": bson.M{
			"_id":      nil,
			"total":    bson.M{"$sum": "$total"},
			"tokens":   bson.M{"$sum": "$tokens"},
			"abnormal": bson.M{"$sum": "$abnormal"},
		}},
	}

	result := make([]map[string]any, 0)
	var err error
	if params.AppKey != "" {
		err = dao.StatisticsAppKey.Aggregate(ctx, pipeline, &result)
	} else {
		err = dao.StatisticsUser.Aggregate(ctx, pipeline, &result)
	}
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}
	if len(result) > 0 {
		res.TotalCalls = gconv.Int(result[0]["total"])
		res.TotalTokens = common.ConvQuotaUnitReverse(gconv.Int(result[0]["tokens"]))
		res.TotalAbnormal = gconv.Int(result[0]["abnormal"])
		if res.TotalCalls > 0 {
			res.AbnormalRate = float64(res.TotalAbnormal) / float64(res.TotalCalls) * 100
		}
	}

	// 今日数据
	todayMatch := s.buildMatchFilter(ctx, gtime.Now().StartOfDay().TimestampMilli(), gtime.Now().EndOfDay(true).TimestampMilli(), params.UserId, params.AppId, params.AppKey, gconv.String(params.Rid))
	todayPipeline := []bson.M{
		{"$match": todayMatch},
		{"$group": bson.M{
			"_id":      nil,
			"total":    bson.M{"$sum": "$total"},
			"tokens":   bson.M{"$sum": "$tokens"},
			"abnormal": bson.M{"$sum": "$abnormal"},
		}},
	}

	todayResult := make([]map[string]any, 0)
	var todayErr error
	if params.AppKey != "" {
		todayErr = dao.StatisticsAppKey.Aggregate(ctx, todayPipeline, &todayResult)
	} else {
		todayErr = dao.StatisticsUser.Aggregate(ctx, todayPipeline, &todayResult)
	}
	if todayErr != nil {
		logger.Error(ctx, todayErr)
	}
	if len(todayResult) > 0 {
		res.TodayCalls = gconv.Int(todayResult[0]["total"])
		res.TodayTokens = common.ConvQuotaUnitReverse(gconv.Int(todayResult[0]["tokens"]))
		res.TodayAbnormal = gconv.Int(todayResult[0]["abnormal"])
	}

	// 实体计数 - 根据角色
	if service.Session().IsAdminRole(ctx) {
		res.TotalUsers, err = dao.User.EstimatedDocumentCount(ctx)
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalApps, err = dao.App.EstimatedDocumentCount(ctx)
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalAppKeys, err = dao.AppKey.EstimatedDocumentCount(ctx)
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalModels, err = dao.Model.EstimatedDocumentCount(ctx)
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalModelKeys, err = dao.Key.CountDocuments(ctx, bson.M{"status": 1})
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalAgents, err = dao.ModelAgent.CountDocuments(ctx, bson.M{"status": 1})
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalProviders, err = dao.Provider.CountDocuments(ctx, bson.M{"status": 1})
		if err != nil {
			logger.Error(ctx, err)
		}
		todayFilter := bson.M{
			"created_at": bson.M{
				"$gte": gtime.Now().StartOfDay().TimestampMilli(),
				"$lte": gtime.Now().EndOfDay(true).TimestampMilli(),
			},
		}
		res.TodayUsers, err = dao.User.CountDocuments(ctx, todayFilter)
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TodayApps, err = dao.App.CountDocuments(ctx, todayFilter)
		if err != nil {
			logger.Error(ctx, err)
		}
		groupCount, err := dao.Group.EstimatedDocumentCount(ctx)
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalGroups = int(groupCount)
		res.TotalBatchTasks, err = dao.TaskBatch.EstimatedDocumentCount(ctx)
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalFileTasks, err = dao.TaskFile.EstimatedDocumentCount(ctx)
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalImageTasks, err = dao.TaskImage.EstimatedDocumentCount(ctx)
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalVideoTasks, err = dao.TaskVideo.EstimatedDocumentCount(ctx)
		if err != nil {
			logger.Error(ctx, err)
		}
	} else if service.Session().IsResellerRole(ctx) {
		rid := service.Session().GetRid(ctx)
		res.TotalUsers, err = dao.User.CountDocuments(ctx, bson.M{"rid": rid})
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalApps, err = dao.App.CountDocuments(ctx, bson.M{"rid": rid})
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalAppKeys, err = dao.AppKey.CountDocuments(ctx, bson.M{"rid": rid})
		if err != nil {
			logger.Error(ctx, err)
		}
		todayFilter := bson.M{
			"rid": rid,
			"created_at": bson.M{
				"$gte": gtime.Now().StartOfDay().TimestampMilli(),
				"$lte": gtime.Now().EndOfDay(true).TimestampMilli(),
			},
		}
		res.TodayUsers, err = dao.User.CountDocuments(ctx, todayFilter)
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TodayApps, err = dao.App.CountDocuments(ctx, todayFilter)
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalGroups = len(service.Session().GetReseller(ctx).Groups)
	} else {
		userId := service.Session().GetUserId(ctx)
		res.TotalApps, err = dao.App.CountDocuments(ctx, bson.M{"user_id": userId})
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalAppKeys, err = dao.AppKey.CountDocuments(ctx, bson.M{"user_id": userId})
		if err != nil {
			logger.Error(ctx, err)
		}
		res.TotalGroups = len(service.Session().GetUser(ctx).Groups)
	}

	return res, nil
}

// 数据看板模型趋势
func (s *sStatistics) DataModelTrend(ctx context.Context, params model.StatisticsModelTrendReq) (*model.StatisticsModelTrendRes, error) {

	match := s.buildMatchFilter(ctx, params.StatStartTime, params.StatEndTime, params.UserId, params.AppId, params.AppKey, gconv.String(params.Rid), params.Key)

	// 解析筛选 (models 优先级高于 provider): resolvedModels 仅作 $match 过滤条件
	resolvedModels, useFilter := s.resolveModels(ctx, params.Models, params.Provider)

	// 最终显示的模型列表:
	// - 用户明确传了 models: 按用户指定展示(即使某些模型当前区间无数据也按用户意图展示空线)
	// - 其他情况(仅 provider 或无筛选): 按实际数据取 top5, 确保只展示有数据的模型
	var models []string
	if len(params.Models) > 0 {
		models = params.Models
	} else {
		topPipeline := []bson.M{
			{"$match": match},
			{"$unwind": "$model_stats"},
		}
		if useFilter {
			topPipeline = append(topPipeline, bson.M{"$match": bson.M{"model_stats.model": bson.M{"$in": resolvedModels}}})
		}
		topPipeline = append(topPipeline,
			bson.M{"$group": bson.M{
				"_id":   "$model_stats.model",
				"count": bson.M{"$sum": "$model_stats.total"},
			}},
			bson.M{"$sort": bson.M{"count": -1}},
			bson.M{"$limit": 5},
		)

		topResult := make([]map[string]any, 0)
		var topErr error
		if params.AppKey != "" {
			topErr = dao.StatisticsAppKey.Aggregate(ctx, topPipeline, &topResult)
		} else {
			topErr = dao.StatisticsUser.Aggregate(ctx, topPipeline, &topResult)
		}
		if topErr != nil {
			logger.Error(ctx, topErr)
			return nil, topErr
		}

		for _, r := range topResult {
			if name := gconv.String(r["_id"]); name != "" {
				models = append(models, name)
			}
		}
	}

	if len(models) == 0 {
		return &model.StatisticsModelTrendRes{Models: []string{}, Dates: []string{}, Series: map[string]*model.ModelTrendSeries{}}, nil
	}

	// 按日期+模型分组
	pipeline := []bson.M{
		{"$match": match},
		{"$unwind": "$model_stats"},
		{"$match": bson.M{"model_stats.model": bson.M{"$in": models}}},
		{"$group": bson.M{
			"_id": bson.M{
				"date":  "$stat_date",
				"model": "$model_stats.model",
			},
			"total":    bson.M{"$sum": "$model_stats.total"},
			"tokens":   bson.M{"$sum": "$model_stats.tokens"},
			"abnormal": bson.M{"$sum": "$model_stats.abnormal"},
		}},
	}

	result := make([]map[string]any, 0)
	var err error
	if params.AppKey != "" {
		err = dao.StatisticsAppKey.Aggregate(ctx, pipeline, &result)
	} else {
		err = dao.StatisticsUser.Aggregate(ctx, pipeline, &result)
	}
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	// 构建日期列表
	var days []*util.DateTime
	if params.StatStartTime > 0 && params.StatEndTime > 0 {
		startTime := gtime.NewFromTimeStamp(params.StatStartTime / 1000)
		endTime := gtime.NewFromTimeStamp(params.StatEndTime / 1000)
		days = util.Day(startTime.String(), endTime.String())
	} else {
		// 全部历史模式: 从结果中取最早和最晚日期
		var minDate, maxDate string
		for _, r := range result {
			var date string
			switch id := r["_id"].(type) {
			case bson.D:
				for _, elem := range id {
					if elem.Key == "date" {
						date = gconv.String(elem.Value)
					}
				}
			case bson.M:
				date = gconv.String(id["date"])
			}
			if date == "" {
				continue
			}
			if minDate == "" || date < minDate {
				minDate = date
			}
			if maxDate == "" || date > maxDate {
				maxDate = date
			}
		}
		if minDate != "" && maxDate != "" {
			days = util.Day(minDate+" 00:00:00", maxDate+" 23:59:59")
		}
	}

	dates := make([]string, 0)
	for _, day := range days {
		dates = append(dates, day.StartDate[5:])
	}

	// 构建 series
	dataMap := make(map[string]map[string][3]float64) // model -> date -> [calls, tokens, abnormal]
	for _, r := range result {

		var date, modelName string
		switch id := r["_id"].(type) {
		case bson.D:
			for _, elem := range id {
				switch elem.Key {
				case "date":
					date = gconv.String(elem.Value)
				case "model":
					modelName = gconv.String(elem.Value)
				}
			}
		case bson.M:
			date = gconv.String(id["date"])
			modelName = gconv.String(id["model"])
		}
		if dataMap[modelName] == nil {
			dataMap[modelName] = make(map[string][3]float64)
		}
		dataMap[modelName][date] = [3]float64{gconv.Float64(r["total"]), gconv.Float64(r["tokens"]), gconv.Float64(r["abnormal"])}
	}

	series := make(map[string]*model.ModelTrendSeries)
	for _, m := range models {
		s := &model.ModelTrendSeries{
			Calls:    make([]int, len(days)),
			Tokens:   make([]float64, len(days)),
			Abnormal: make([]int, len(days)),
		}
		for i, day := range days {
			if v, ok := dataMap[m][day.StartDate]; ok {
				s.Calls[i] = int(v[0])
				s.Tokens[i] = common.ConvQuotaUnitReverse(int(v[1]))
				s.Abnormal[i] = int(v[2])
			}
		}
		series[m] = s
	}

	return &model.StatisticsModelTrendRes{
		Models: models,
		Dates:  dates,
		Series: series,
	}, nil
}

// 数据看板响应耗时趋势
func (s *sStatistics) DataLatencyTrend(ctx context.Context, params model.StatisticsLatencyTrendReq) (*model.StatisticsLatencyTrendRes, error) {

	if params.StatStartTime <= 0 || params.StatEndTime <= 0 {
		return &model.StatisticsLatencyTrendRes{
			Models: []string{},
			Dates:  []string{},
			Series: map[string]*model.LatencyTrendSeries{},
		}, nil
	}

	baseMatch := bson.M{
		"is_smart_match": bson.M{"$ne": true},
		"is_retry":       bson.M{"$ne": true},
		"created_at":     bson.M{"$gte": params.StatStartTime, "$lte": params.StatEndTime},
	}

	if service.Session().IsResellerRole(ctx) {
		baseMatch["rid"] = service.Session().GetRid(ctx)
	}
	if service.Session().IsUserRole(ctx) {
		baseMatch["user_id"] = service.Session().GetUserId(ctx)
	}
	if service.Session().IsAdminRole(ctx) && params.Rid != 0 {
		baseMatch["rid"] = params.Rid
	}
	if !service.Session().IsUserRole(ctx) && params.UserId != 0 {
		baseMatch["user_id"] = params.UserId
	}
	if params.AppId != 0 {
		baseMatch["app_id"] = params.AppId
	}
	if params.AppKey != "" {
		baseMatch["app_key"] = params.AppKey
	}

	// 管理员可通过key筛选(匹配app_key字段)
	if params.Key != "" && service.Session().IsAdminRole(ctx) {
		baseMatch["app_key"] = params.Key
	}

	// 模型筛选 (models 优先级高于 provider)
	if len(params.Models) > 0 {
		baseMatch["model"] = bson.M{"$in": params.Models}
	} else if params.Provider != "" {
		// 按 provider_id 过滤(params.Provider 是 provider 的 id)
		baseMatch["provider_id"] = params.Provider
	}

	// 第一步: 找出 top5 模型
	topPipeline := []bson.M{
		{"$match": baseMatch},
		{"$group": bson.M{
			"_id":   "$model",
			"count": bson.M{"$sum": 1},
		}},
		{"$sort": bson.M{"count": -1}},
		{"$limit": 5},
	}

	topResults := s.aggregateAllLogs(ctx, topPipeline)

	// 合并多集合 top 结果
	topMap := make(map[string]float64)
	for _, r := range topResults {
		name := gconv.String(r["_id"])
		if name == "" {
			continue
		}
		topMap[name] += gconv.Float64(r["count"])
	}

	type modelCount struct {
		name  string
		count float64
	}
	ranked := make([]modelCount, 0, len(topMap))
	for name, count := range topMap {
		ranked = append(ranked, modelCount{name, count})
	}
	for i := 0; i < len(ranked); i++ {
		for j := i + 1; j < len(ranked); j++ {
			if ranked[j].count > ranked[i].count {
				ranked[i], ranked[j] = ranked[j], ranked[i]
			}
		}
	}
	if len(ranked) > 5 {
		ranked = ranked[:5]
	}

	models := make([]string, 0, len(ranked))
	modelSet := make([]any, 0, len(ranked))
	for _, mc := range ranked {
		models = append(models, mc.name)
		modelSet = append(modelSet, mc.name)
	}

	if len(models) == 0 {
		return &model.StatisticsLatencyTrendRes{
			Models: []string{},
			Dates:  []string{},
			Series: map[string]*model.LatencyTrendSeries{},
		}, nil
	}

	// 第二步: 按日期+模型分组查询平均耗时
	dateExpr := bson.M{
		"$dateToString": bson.M{
			"format": "%Y-%m-%d",
			"date":   bson.M{"$toDate": "$created_at"},
		},
	}

	histMatch := bson.M{}
	for k, v := range baseMatch {
		histMatch[k] = v
	}
	histMatch["model"] = bson.M{"$in": modelSet}

	histPipeline := []bson.M{
		{"$match": histMatch},
		{"$group": bson.M{
			"_id": bson.M{
				"date":  dateExpr,
				"model": "$model",
			},
			"avg_total_time": bson.M{"$avg": "$total_time"},
			"count":          bson.M{"$sum": 1},
		}},
	}

	aggResults := s.aggregateAllLogs(ctx, histPipeline)

	// 合并多集合结果 (加权平均)
	type compositeKey struct {
		date  string
		model string
	}
	type avgEntry struct {
		totalTime float64
		count     float64
	}
	mergedMap := make(map[compositeKey]*avgEntry)
	for _, r := range aggResults {
		var date, modelName string
		switch id := r["_id"].(type) {
		case bson.D:
			for _, elem := range id {
				switch elem.Key {
				case "date":
					date = gconv.String(elem.Value)
				case "model":
					modelName = gconv.String(elem.Value)
				}
			}
		case bson.M:
			date = gconv.String(id["date"])
			modelName = gconv.String(id["model"])
		case map[string]any:
			date = gconv.String(id["date"])
			modelName = gconv.String(id["model"])
		}

		count := gconv.Float64(r["count"])
		avgTime := gconv.Float64(r["avg_total_time"])

		key := compositeKey{date, modelName}
		if existing, ok := mergedMap[key]; ok {
			totalCount := existing.count + count
			if totalCount > 0 {
				existing.totalTime = (existing.totalTime*existing.count + avgTime*count) / totalCount
			}
			existing.count = totalCount
		} else {
			mergedMap[key] = &avgEntry{totalTime: avgTime, count: count}
		}
	}

	// 构建日期列表
	startTime := gtime.NewFromTimeStamp(params.StatStartTime / 1000)
	endTime := gtime.NewFromTimeStamp(params.StatEndTime / 1000)
	days := util.Day(startTime.String(), endTime.String())

	dates := make([]string, 0, len(days))
	for _, day := range days {
		dates = append(dates, day.StartDate[5:])
	}

	// 构建 series
	series := make(map[string]*model.LatencyTrendSeries)
	for _, m := range models {
		ls := &model.LatencyTrendSeries{
			AvgTotalTime: make([]int64, len(days)),
		}
		for i, day := range days {
			key := compositeKey{day.StartDate, m}
			if entry, ok := mergedMap[key]; ok {
				ls.AvgTotalTime[i] = int64(entry.totalTime)
			}
		}
		series[m] = ls
	}

	return &model.StatisticsLatencyTrendRes{
		Models: models,
		Dates:  dates,
		Series: series,
	}, nil
}

// 数据看板任务状态分布
func (s *sStatistics) DataTaskStatus(ctx context.Context, params model.StatisticsTaskStatusReq) (*model.StatisticsTaskStatusRes, error) {

	res := &model.StatisticsTaskStatusRes{
		Batch: make([]*model.TaskStatusItem, 0),
		File:  make([]*model.TaskStatusItem, 0),
		Image: make([]*model.TaskStatusItem, 0),
		Video: make([]*model.TaskStatusItem, 0),
	}

	// 时间过滤条件
	match := bson.M{}
	if params.StatStartTime > 0 && params.StatEndTime > 0 {
		match["created_at"] = bson.M{
			"$gte": gtime.NewFromTimeStamp(params.StatStartTime),
			"$lte": gtime.NewFromTimeStamp(params.StatEndTime),
		}
	} else if params.StatStartTime > 0 {
		match["created_at"] = bson.M{"$gte": gtime.NewFromTimeStamp(params.StatStartTime)}
	} else if params.StatEndTime > 0 {
		match["created_at"] = bson.M{"$lte": gtime.NewFromTimeStamp(params.StatEndTime)}
	}

	// 角色过滤
	if service.Session().IsResellerRole(ctx) {
		match["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		match["user_id"] = service.Session().GetUserId(ctx)
	}

	// 管理员可按代理商筛选
	if service.Session().IsAdminRole(ctx) && params.Rid != 0 {
		match["rid"] = params.Rid
	}

	// 管理员/代理商可按用户筛选
	if !service.Session().IsUserRole(ctx) && params.UserId != 0 {
		match["user_id"] = params.UserId
	}

	if params.AppId != 0 {
		match["app_id"] = params.AppId
	}

	// 应用密钥筛选(用户和代理商使用)
	if params.AppKey != "" {
		match["creator"] = params.AppKey
	}

	// 管理员可通过key筛选(匹配creator字段)
	if params.Key != "" && service.Session().IsAdminRole(ctx) {
		match["creator"] = params.Key
	}

	// 模型筛选 (models 优先级高于 provider)
	taskResolvedModels, taskUseModels := s.resolveModels(ctx, params.Models, params.Provider)
	if taskUseModels {
		match["model"] = bson.M{"$in": taskResolvedModels}
	}

	buildPipeline := func() []bson.M {
		pipeline := make([]bson.M, 0)
		if len(match) > 0 {
			pipeline = append(pipeline, bson.M{"$match": match})
		}
		pipeline = append(pipeline,
			bson.M{"$group": bson.M{"_id": "$status", "count": bson.M{"$sum": 1}}},
			bson.M{"$sort": bson.M{"count": -1}},
		)
		return pipeline
	}

	// 批处理任务按状态分组
	batchResult := make([]map[string]any, 0)
	if err := dao.TaskBatch.Aggregate(ctx, buildPipeline(), &batchResult); err != nil {
		logger.Error(ctx, err)
	}
	for _, r := range batchResult {
		status := gconv.String(r["_id"])
		count := gconv.Int64(r["count"])
		res.Batch = append(res.Batch, &model.TaskStatusItem{Status: status, Count: count})
		if status == "in_progress" || status == "validating" || status == "finalizing" {
			res.ActiveBatch += count
		}
		if status == "queued" {
			res.QueuedBatch += count
		}
	}

	// 文件任务按状态分组
	fileResult := make([]map[string]any, 0)
	if err := dao.TaskFile.Aggregate(ctx, buildPipeline(), &fileResult); err != nil {
		logger.Error(ctx, err)
	}
	for _, r := range fileResult {
		res.File = append(res.File, &model.TaskStatusItem{
			Status: gconv.String(r["_id"]),
			Count:  gconv.Int64(r["count"]),
		})
	}

	// 绘图任务按状态分组
	imageResult := make([]map[string]any, 0)
	if err := dao.TaskImage.Aggregate(ctx, buildPipeline(), &imageResult); err != nil {
		logger.Error(ctx, err)
	}
	for _, r := range imageResult {
		status := gconv.String(r["_id"])
		count := gconv.Int64(r["count"])
		res.Image = append(res.Image, &model.TaskStatusItem{Status: status, Count: count})
		if status == "in_progress" {
			res.ActiveImage += count
		}
		if status == "queued" {
			res.QueuedImage += count
		}
	}

	// 视频任务按状态分组
	videoResult := make([]map[string]any, 0)
	if err := dao.TaskVideo.Aggregate(ctx, buildPipeline(), &videoResult); err != nil {
		logger.Error(ctx, err)
	}
	for _, r := range videoResult {
		status := gconv.String(r["_id"])
		count := gconv.Int64(r["count"])
		res.Video = append(res.Video, &model.TaskStatusItem{Status: status, Count: count})
		if status == "in_progress" {
			res.ActiveVideo += count
		}
		if status == "queued" {
			res.QueuedVideo += count
		}
	}

	return res, nil
}

// 数据看板代理状态
func (s *sStatistics) DataAgentStatus(ctx context.Context, params model.StatisticsAgentStatusReq) (*model.StatisticsAgentStatusRes, error) {

	res := &model.StatisticsAgentStatusRes{}

	// 时间过滤条件
	timeFilter := bson.M{}
	if params.StatStartTime > 0 && params.StatEndTime > 0 {
		timeFilter["updated_at"] = bson.M{
			"$gte": gtime.NewFromTimeStamp(params.StatStartTime),
			"$lte": gtime.NewFromTimeStamp(params.StatEndTime),
		}
	} else if params.StatStartTime > 0 {
		timeFilter["updated_at"] = bson.M{"$gte": gtime.NewFromTimeStamp(params.StatStartTime)}
	} else if params.StatEndTime > 0 {
		timeFilter["updated_at"] = bson.M{"$lte": gtime.NewFromTimeStamp(params.StatEndTime)}
	}

	activeFilter := bson.M{"status": 1}
	disabledFilter := bson.M{"status": 2}
	autoDisabledFilter := bson.M{"is_auto_disabled": true}
	for k, v := range timeFilter {
		activeFilter[k] = v
		disabledFilter[k] = v
		autoDisabledFilter[k] = v
	}

	var err error
	if len(timeFilter) > 0 {
		res.Total, err = dao.ModelAgent.CountDocuments(ctx, timeFilter)
	} else {
		res.Total, err = dao.ModelAgent.EstimatedDocumentCount(ctx)
	}
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	res.Active, err = dao.ModelAgent.CountDocuments(ctx, activeFilter)
	if err != nil {
		logger.Error(ctx, err)
	}

	res.Disabled, err = dao.ModelAgent.CountDocuments(ctx, disabledFilter)
	if err != nil {
		logger.Error(ctx, err)
	}

	res.AutoDisabled, err = dao.ModelAgent.CountDocuments(ctx, autoDisabledFilter)
	if err != nil {
		logger.Error(ctx, err)
	}

	// 按模型代理统计已用额度(汇总其下所有key的已用额度)
	agentMatchFilter := bson.M{"status": bson.M{"$ne": -1}}
	for k, v := range timeFilter {
		agentMatchFilter[k] = v
	}

	// 模型筛选(models 优先, 其次 provider → 解析为模型ID列表, ModelAgent.models 存的是模型ID)
	agentModelIds, agentUseModelIds := s.resolveModelIds(ctx, params.Models, params.Provider)
	if agentUseModelIds {
		agentMatchFilter["models"] = bson.M{"$in": agentModelIds}
	}

	agentPipeline := []bson.M{
		{"$match": agentMatchFilter},
		{"$sort": bson.M{"used_quota": -1}},
		{"$limit": 20},
		{"$project": bson.M{"_id": 1, "name": 1, "status": 1}},
	}

	agentResult := make([]map[string]any, 0)
	if err = dao.ModelAgent.Aggregate(ctx, agentPipeline, &agentResult); err != nil {
		logger.Error(ctx, err)
	}

	if len(agentResult) > 0 {
		// 收集代理ID
		agentIds := make([]string, 0, len(agentResult))
		agentNameMap := make(map[string]string)
		agentStatusMap := make(map[string]int)
		for _, r := range agentResult {
			id := gconv.String(r["_id"])
			agentIds = append(agentIds, id)
			agentNameMap[id] = gconv.String(r["name"])
			agentStatusMap[id] = gconv.Int(r["status"])
		}

		// 聚合key表, 按model_agents分组统计used_quota
		keyMatchFilter := bson.M{"model_agents": bson.M{"$in": agentIds}}

		// 应用密钥筛选(用户和代理商使用)
		if params.AppKey != "" {
			keyMatchFilter["creator"] = params.AppKey
		}

		// 管理员可通过key筛选(匹配key或creator字段)
		if params.Key != "" && service.Session().IsAdminRole(ctx) {
			keyMatchFilter["$or"] = bson.A{
				bson.M{"key": params.Key},
				bson.M{"creator": params.Key},
			}
		}

		// 模型筛选(Key.models 存的也是模型ID)
		if agentUseModelIds {
			keyMatchFilter["models"] = bson.M{"$in": agentModelIds}
		}

		keyPipeline := []bson.M{
			{"$match": keyMatchFilter},
			{"$unwind": "$model_agents"},
			{"$match": bson.M{"model_agents": bson.M{"$in": agentIds}}},
			{"$group": bson.M{
				"_id":        "$model_agents",
				"used_quota": bson.M{"$sum": "$used_quota"},
			}},
		}

		keyResult := make([]map[string]any, 0)
		if err = dao.Key.Aggregate(ctx, keyPipeline, &keyResult); err != nil {
			logger.Error(ctx, err)
		}

		quotaMap := make(map[string]int)
		for _, r := range keyResult {
			quotaMap[gconv.String(r["_id"])] = gconv.Int(r["used_quota"])
		}

		res.ByAgent = make([]*model.AgentStat, 0, len(agentIds))
		for _, id := range agentIds {
			res.ByAgent = append(res.ByAgent, &model.AgentStat{
				Name:      agentNameMap[id],
				Status:    agentStatusMap[id],
				UsedQuota: common.ConvQuotaUnitReverse(quotaMap[id]),
			})
		}
	}

	return res, nil
}

// 数据看板密钥状态
func (s *sStatistics) DataKeyStatus(ctx context.Context, params model.StatisticsKeyStatusReq) (*model.StatisticsKeyStatusRes, error) {

	res := &model.StatisticsKeyStatusRes{}

	// 时间过滤条件
	timeFilter := bson.M{}
	if params.StatStartTime > 0 && params.StatEndTime > 0 {
		timeFilter["updated_at"] = bson.M{
			"$gte": gtime.NewFromTimeStamp(params.StatStartTime),
			"$lte": gtime.NewFromTimeStamp(params.StatEndTime),
		}
	} else if params.StatStartTime > 0 {
		timeFilter["updated_at"] = bson.M{"$gte": gtime.NewFromTimeStamp(params.StatStartTime)}
	} else if params.StatEndTime > 0 {
		timeFilter["updated_at"] = bson.M{"$lte": gtime.NewFromTimeStamp(params.StatEndTime)}
	}

	// 应用密钥筛选(用户和代理商使用)
	if params.AppKey != "" {
		timeFilter["creator"] = params.AppKey
	}

	// 管理员可通过key筛选(匹配key或creator字段)
	if params.Key != "" && service.Session().IsAdminRole(ctx) {
		timeFilter["$or"] = bson.A{
			bson.M{"key": params.Key},
			bson.M{"creator": params.Key},
		}
	}

	activeFilter := bson.M{"status": 1}
	disabledFilter := bson.M{"status": 2}
	autoDisabledFilter := bson.M{"is_auto_disabled": true}
	for k, v := range timeFilter {
		activeFilter[k] = v
		disabledFilter[k] = v
		autoDisabledFilter[k] = v
	}

	var err error
	if len(timeFilter) > 0 {
		res.Total, err = dao.Key.CountDocuments(ctx, timeFilter)
	} else {
		res.Total, err = dao.Key.EstimatedDocumentCount(ctx)
	}
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	res.Active, err = dao.Key.CountDocuments(ctx, activeFilter)
	if err != nil {
		logger.Error(ctx, err)
	}

	res.Disabled, err = dao.Key.CountDocuments(ctx, disabledFilter)
	if err != nil {
		logger.Error(ctx, err)
	}

	res.AutoDisabled, err = dao.Key.CountDocuments(ctx, autoDisabledFilter)
	if err != nil {
		logger.Error(ctx, err)
	}

	// 按密钥列出状态(取最近的密钥)
	matchFilter := bson.M{"status": bson.M{"$ne": -1}}
	for k, v := range timeFilter {
		matchFilter[k] = v
	}

	// 模型筛选(models 优先, 其次 provider → 解析为模型ID列表, Key.models 存的是模型ID)
	keyModelIds, keyUseModelIds := s.resolveModelIds(ctx, params.Models, params.Provider)
	if keyUseModelIds {
		matchFilter["models"] = bson.M{"$in": keyModelIds}
	}

	pipeline := []bson.M{
		{"$match": matchFilter},
		{"$sort": bson.M{"used_quota": -1}},
		{"$limit": 20},
		{"$project": bson.M{"key": 1, "status": 1, "used_quota": 1}},
	}

	keyResult := make([]map[string]any, 0)
	if err = dao.Key.Aggregate(ctx, pipeline, &keyResult); err != nil {
		logger.Error(ctx, err)
	}

	res.ByKey = make([]*model.KeyStat, 0)
	for _, r := range keyResult {
		key := gconv.String(r["key"])
		if len(key) > 5 {
			key = key[len(key)-5:]
		}
		res.ByKey = append(res.ByKey, &model.KeyStat{
			Key:       key,
			Status:    gconv.Int(r["status"]),
			UsedQuota: common.ConvQuotaUnitReverse(gconv.Int(r["used_quota"])),
		})
	}

	return res, nil
}

// 构建基础匹配过滤器
func (s *sStatistics) buildMatchFilter(ctx context.Context, statStartTime, statEndTime int64, userId, appId int, extras ...string) bson.M {

	// extras: [0]=appKey, [1]=rid (as string), [2]=key (admin only)
	appKey := ""
	rid := 0
	key := ""

	if len(extras) > 0 {
		appKey = extras[0]
	}

	if len(extras) > 1 {
		rid = gconv.Int(extras[1])
	}

	if len(extras) > 2 {
		key = extras[2]
	}

	match := bson.M{}

	if statStartTime > 0 && statEndTime > 0 {
		match["stat_time"] = bson.M{
			"$gte": statStartTime,
			"$lte": statEndTime,
		}
	} else if statStartTime > 0 {
		match["stat_time"] = bson.M{"$gte": statStartTime}
	} else if statEndTime > 0 {
		match["stat_time"] = bson.M{"$lte": statEndTime}
	}

	// 角色过滤
	if service.Session().IsResellerRole(ctx) {
		match["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		match["user_id"] = service.Session().GetUserId(ctx)
	}

	// 管理员可按代理商筛选
	if service.Session().IsAdminRole(ctx) && rid != 0 {
		match["rid"] = rid
	}

	// 管理员/代理商可按用户筛选
	if !service.Session().IsUserRole(ctx) && userId != 0 {
		match["user_id"] = userId
	}

	if appId != 0 {
		match["app_id"] = appId
	}

	if appKey != "" {
		match["app_key"] = appKey
	}

	// 管理员可通过key筛选(匹配app_key字段)
	if key != "" && service.Session().IsAdminRole(ctx) {
		match["app_key"] = key
	}

	return match
}

// 并发查询所有日志集合并合并结果
func (s *sStatistics) aggregateAllLogs(ctx context.Context, pipeline []bson.M) []map[string]any {
	var mu sync.Mutex
	var wg sync.WaitGroup
	merged := make([]map[string]any, 0)

	for _, coll := range s.logCollections {
		wg.Add(1)
		go func(collection string) {
			defer wg.Done()
			result := make([]map[string]any, 0)
			if err := dao.Aggregate(ctx, db.DefaultDatabase, collection, pipeline, &result); err != nil {
				logger.Error(ctx, err)
				return
			}
			mu.Lock()
			merged = append(merged, result...)
			mu.Unlock()
		}(coll)
	}
	wg.Wait()
	return merged
}

// 将聚合管道$push出的model_stats转为[]*common.ModelStat
func convModelStats(v any) []*mcommon.ModelStat {
	if v == nil {
		return nil
	}

	toMap := func(item any) map[string]any {
		switch x := item.(type) {
		case map[string]any:
			return x
		case bson.M:
			return x
		case bson.D:
			m := make(map[string]any, len(x))
			for _, e := range x {
				m[e.Key] = e.Value
			}
			return m
		default:
			m := make(map[string]any)
			_ = gconv.Struct(item, &m)
			return m
		}
	}

	items := make([]any, 0)
	switch arr := v.(type) {
	case []any:
		items = arr
	case []*mcommon.ModelStat:
		for _, x := range arr {
			items = append(items, x)
		}
	case bson.A:
		for _, x := range arr {
			items = append(items, x)
		}
	default:
		return nil
	}

	if len(items) == 0 {
		return nil
	}

	stats := make([]*mcommon.ModelStat, 0, len(items))
	for _, item := range items {
		m := toMap(item)
		if len(m) == 0 {
			continue
		}
		stats = append(stats, &mcommon.ModelStat{
			ModelId:        gconv.String(m["model_id"]),
			Model:          gconv.String(m["model"]),
			Total:          gconv.Int(m["total"]),
			Tokens:         common.ConvQuotaUnitReverse(gconv.Int(m["tokens"])),
			Abnormal:       gconv.Int(m["abnormal"]),
			AbnormalTokens: common.ConvQuotaUnitReverse(gconv.Int(m["abnormal_tokens"])),
		})
	}
	return stats
}

// 按筛选条件解析出最终需要过滤的模型名列表
// 1. 若传入了 models(优先级最高), 直接返回
// 2. 否则若传入了 provider(provider_id), 通过 provider 查出对应模型名列表
// 3. 返回: models 列表, 是否应使用 models 过滤
// 说明: 当仅选了 provider 但该 provider 下无模型时, 返回 ([""], true), 配合 $in 查询保证查不到数据
func (s *sStatistics) resolveModels(ctx context.Context, models []string, provider string) ([]string, bool) {
	if len(models) > 0 {
		return models, true
	}
	if provider == "" {
		return nil, false
	}
	list, err := service.Model().List(ctx, model.ModelListReq{ProviderId: provider})
	if err != nil {
		logger.Error(ctx, err)
		return []string{""}, true
	}
	if len(list) == 0 {
		return []string{""}, true
	}
	names := make([]string, 0, len(list))
	for _, m := range list {
		if m.Model != "" {
			names = append(names, m.Model)
		}
	}
	if len(names) == 0 {
		return []string{""}, true
	}
	return names, true
}

// 按筛选条件解析出最终需要过滤的模型ID列表(用于 ModelAgent/Key 等使用模型ID绑定的表)
// 1. 若传入了 models(按名称), 通过名称查出 modelId 列表
// 2. 否则若传入了 provider(provider_id), 通过 provider 查出该 provider 下的所有模型ID
// 3. 返回: modelId 列表, 是否应使用 models 过滤
func (s *sStatistics) resolveModelIds(ctx context.Context, models []string, provider string) ([]string, bool) {
	if len(models) == 0 && provider == "" {
		return nil, false
	}
	req := model.ModelListReq{}
	if provider != "" && len(models) == 0 {
		req.ProviderId = provider
	}
	list, err := service.Model().List(ctx, req)
	if err != nil {
		logger.Error(ctx, err)
		return []string{""}, true
	}
	if len(list) == 0 {
		return []string{""}, true
	}
	ids := make([]string, 0, len(list))
	if len(models) > 0 {
		// 按名称匹配
		nameSet := make(map[string]struct{}, len(models))
		for _, n := range models {
			nameSet[n] = struct{}{}
		}
		for _, m := range list {
			if _, ok := nameSet[m.Model]; ok {
				ids = append(ids, m.Id)
			}
		}
	} else {
		for _, m := range list {
			ids = append(ids, m.Id)
		}
	}
	if len(ids) == 0 {
		return []string{""}, true
	}
	return ids, true
}
