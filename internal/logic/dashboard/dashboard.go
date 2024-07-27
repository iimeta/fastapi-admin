package dashboard

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

type sDashboard struct{}

func init() {
	service.RegisterDashboard(New())
}

func New() service.IDashboard {
	return &sDashboard{}
}

// 基础数据
func (s *sDashboard) BaseData(ctx context.Context) (dashboard *model.Dashboard, err error) {

	dashboard = new(model.Dashboard)

	if service.Session().IsUserRole(ctx) {
		if dashboard.App, err = dao.App.CountDocuments(ctx, bson.M{"user_id": service.Session().GetUserId(ctx)}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	} else {
		if dashboard.App, err = dao.App.EstimatedDocumentCount(ctx); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if service.Session().IsUserRole(ctx) {
		if models := service.Session().GetUser(ctx).Models; len(models) > 0 {
			if dashboard.Model, err = dao.Model.CountDocuments(ctx, bson.M{"_id": bson.M{"$in": models}}); err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
		}
	} else {
		if dashboard.Model, err = dao.Model.EstimatedDocumentCount(ctx); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if service.Session().IsUserRole(ctx) {
		if dashboard.AppKey, err = dao.Key.CountDocuments(ctx, bson.M{"user_id": service.Session().GetUserId(ctx), "type": 1}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	} else {
		if dashboard.AppKey, err = dao.Key.CountDocuments(ctx, bson.M{"type": 1}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if service.Session().IsAdminRole(ctx) {

		if dashboard.TodayApp, err = dao.App.CountDocuments(ctx, bson.M{
			"created_at": bson.M{
				"$gte": gtime.Now().StartOfDay().TimestampMilli(),
				"$lte": gtime.Now().EndOfDay(true).TimestampMilli(),
			},
		}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if dashboard.ModelKey, err = dao.Key.CountDocuments(ctx, bson.M{"type": 2, "status": 1}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if dashboard.User, err = dao.User.EstimatedDocumentCount(ctx); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if dashboard.TodayUser, err = dao.User.CountDocuments(ctx, bson.M{
			"created_at": bson.M{
				"$gte": gtime.Now().StartOfDay().TimestampMilli(),
				"$lte": gtime.Now().EndOfDay(true).TimestampMilli(),
			},
		}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
			},
		},
		{
			"$group": bson.M{
				"_id":  nil,
				"call": bson.M{"$sum": 1},
			},
		},
	}

	if service.Session().IsUserRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["user_id"] = service.Session().GetUserId(ctx)
	}

	result := make([]map[string]interface{}, 0)
	if err = dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if len(result) > 0 {
		dashboard.Call = gconv.Int(result[0]["call"])
	}

	return dashboard, nil
}

// 调用数据
func (s *sDashboard) CallData(ctx context.Context, params model.DashboardCallDataReq) ([]*model.CallData, error) {

	startTime := gtime.Now().AddDate(0, 0, -(params.Days - 1)).StartOfDay()
	endTime := gtime.Now().EndOfDay(true)

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"req_time": bson.M{
					"$gte": startTime.TimestampMilli(),
					"$lte": endTime.TimestampMilli(),
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
				"user":   bson.M{"$addToSet": "$user_id"},
				"app":    bson.M{"$addToSet": "$app_id"},
			},
		},
	}

	if service.Session().IsUserRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["user_id"] = service.Session().GetUserId(ctx)
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	resultMap := make(map[string]*model.CallData)
	for _, res := range result {
		resultMap[gconv.String(res["_id"])] = &model.CallData{
			Date:   gconv.String(res["_id"])[5:],
			Call:   gconv.Int(res["count"]),
			Tokens: gconv.Int(res["tokens"]),
			User:   len(gconv.SliceAny(res["user"])),
			App:    len(gconv.SliceAny(res["app"])),
		}
	}

	abnormalPipeline := []bson.M{
		{
			"$match": bson.M{
				"req_time": bson.M{
					"$gte": startTime.TimestampMilli(),
					"$lte": endTime.TimestampMilli(),
				},
				"status": bson.M{"$ne": 1},
			},
		},
		{
			"$group": bson.M{
				"_id":   "$req_date",
				"count": bson.M{"$sum": 1},
			},
		},
	}

	if service.Session().IsUserRole(ctx) {
		match := abnormalPipeline[0]["$match"].(bson.M)
		match["user_id"] = service.Session().GetUserId(ctx)
		match["is_smart_match"] = bson.M{"$ne": true}
		match["is_retry"] = bson.M{"$ne": true}
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
			resultMap[gconv.String(res["_id"])] = &model.CallData{
				Date:     gconv.String(res["_id"])[5:],
				Abnormal: len(gconv.SliceAny(res["count"])),
			}
		}
	}

	items := make([]*model.CallData, 0)
	days := util.Day(startTime.String(), endTime.String())

	for _, day := range days {
		callData := resultMap[day.StartDate]
		if callData == nil {
			callData = &model.CallData{Date: day.StartDate[5:]}
		}
		items = append(items, callData)
	}

	return items, nil
}

// 费用
func (s *sDashboard) Expense(ctx context.Context) (*model.Expense, error) {

	if service.Session().IsAdminRole(ctx) {
		return &model.Expense{}, nil
	}

	user, err := service.User().GetUserByUserId(ctx, service.Session().GetUserId(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.Expense{
		Quota:          user.Quota,
		QuotaUSD:       util.Round(float64(user.Quota)/consts.QUOTA_USD_UNIT, 4),
		UsedQuota:      user.UsedQuota,
		UsedQuotaUSD:   util.Round(float64(user.UsedQuota)/consts.QUOTA_USD_UNIT, 4),
		QuotaExpiresAt: user.QuotaExpiresAt,
	}, nil
}

// 数据TOP
func (s *sDashboard) DataTop(ctx context.Context, params model.DashboardDataTopReq) ([]*model.DataTop, error) {

	startTime := gtime.Now().AddDate(0, 0, -(params.Days - 1)).StartOfDay()
	endTime := gtime.Now().EndOfDay(true)

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"req_time": bson.M{
					"$gte": startTime.TimestampMilli(),
					"$lte": endTime.TimestampMilli(),
				},
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
			},
		},
	}

	switch params.DataType {
	case "user":
		pipeline = append(pipeline, bson.M{
			"$group": bson.M{
				"_id":    "$user_id",
				"count":  bson.M{"$sum": 1},
				"models": bson.M{"$addToSet": "$model"},
				"tokens": bson.M{"$sum": "$total_tokens"},
			},
		})
	case "app":
		pipeline = append(pipeline, bson.M{
			"$group": bson.M{
				"_id":     "$app_id",
				"user_id": bson.M{"$first": "$user_id"},
				"count":   bson.M{"$sum": 1},
				"models":  bson.M{"$addToSet": "$model"},
				"tokens":  bson.M{"$sum": "$total_tokens"},
			},
		})
	case "app_key":
		pipeline = append(pipeline, bson.M{
			"$group": bson.M{
				"_id":    "$creator",
				"app_id": bson.M{"$first": "$app_id"},
				"count":  bson.M{"$sum": 1},
				"tokens": bson.M{"$sum": "$total_tokens"},
			},
		})
	case "model":
		pipeline = append(pipeline, bson.M{
			"$group": bson.M{
				"_id":    "$model",
				"count":  bson.M{"$sum": 1},
				"users":  bson.M{"$addToSet": "$user_id"},
				"apps":   bson.M{"$addToSet": "$app_id"},
				"tokens": bson.M{"$sum": "$total_tokens"},
			},
		})
	}

	pipeline = append(pipeline, bson.M{"$sort": bson.M{"tokens": -1}}, bson.M{"$limit": 10})

	if service.Session().IsUserRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["user_id"] = service.Session().GetUserId(ctx)
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.DataTop, 0)
	switch params.DataType {
	case "user":
		for _, res := range result {
			items = append(items, &model.DataTop{
				UserId: gconv.Int(res["_id"]),
				Call:   gconv.Int(res["count"]),
				Models: len(gconv.SliceAny(res["models"])),
				Tokens: gconv.Int(res["tokens"]),
			})
		}
	case "app":
		for _, res := range result {
			items = append(items, &model.DataTop{
				AppId:  gconv.Int(res["_id"]),
				UserId: gconv.Int(res["user_id"]),
				Call:   gconv.Int(res["count"]),
				Models: len(gconv.SliceAny(res["models"])),
				Tokens: gconv.Int(res["tokens"]),
			})
		}
	case "app_key":
		for _, res := range result {
			items = append(items, &model.DataTop{
				AppKey: gconv.String(res["_id"]),
				AppId:  gconv.Int(res["app_id"]),
				Call:   gconv.Int(res["count"]),
				Tokens: gconv.Int(res["tokens"]),
			})
		}
	case "model":
		for _, res := range result {
			items = append(items, &model.DataTop{
				Model:  gconv.String(res["_id"]),
				Call:   gconv.Int(res["count"]),
				User:   len(gconv.SliceAny(res["users"])),
				App:    len(gconv.SliceAny(res["apps"])),
				Tokens: gconv.Int(res["tokens"]),
			})
		}
	}

	return items, nil
}

// 模型占比
func (s *sDashboard) ModelPercent(ctx context.Context, params model.DashboardModelPercentReq) ([]string, []*model.ModelPercent, error) {

	startTime := gtime.Now().AddDate(0, 0, -(params.Days - 1)).StartOfDay()
	endTime := gtime.Now().EndOfDay(true)

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"req_time": bson.M{
					"$gte": startTime.TimestampMilli(),
					"$lte": endTime.TimestampMilli(),
				},
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
			},
		},
		{
			"$group": bson.M{
				"_id":   "$model",
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$sort": bson.M{
				"count": -1,
			},
		},
		{
			"$limit": 10,
		},
	}

	if service.Session().IsUserRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["user_id"] = service.Session().GetUserId(ctx)
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, nil, err
	}

	models := make([]string, 0)
	items := make([]*model.ModelPercent, 0)
	for _, res := range result {
		models = append(models, gconv.String(res["_id"]))
		items = append(items, &model.ModelPercent{
			Name:  gconv.String(res["_id"]),
			Value: gconv.Int(res["count"]),
		})
	}

	return models, items, nil
}

// 每分钟数据
func (s *sDashboard) PerMinute(ctx context.Context, params model.DashboardPerMinuteReq) (int, int, error) {

	startTime := gtime.Now().TimestampMilli() - 60000
	endTime := gtime.Now().TimestampMilli()

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"req_time": bson.M{
					"$gte": startTime,
					"$lte": endTime,
				},
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
			},
		},
		{
			"$group": bson.M{
				"_id": nil,
				"rpm": bson.M{"$sum": 1},
				"tpm": bson.M{"$sum": "$total_tokens"},
			},
		},
	}

	match := pipeline[0]["$match"].(bson.M)

	if service.Session().IsUserRole(ctx) {
		match["user_id"] = service.Session().GetUserId(ctx)
	} else if params.UserId != 0 {
		match["user_id"] = params.UserId
	}

	if params.AppId != 0 {
		match["app_id"] = params.AppId
	}

	if params.Key != "" {
		match["creator"] = params.Key
	}

	if len(params.Models) > 0 {
		match["model_id"] = bson.M{
			"$in": params.Models,
		}
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return 0, 0, err
	}

	if len(result) > 0 {
		return gconv.Int(result[0]["rpm"]), gconv.Int(result[0]["tpm"]), nil
	}

	return 0, 0, nil
}
