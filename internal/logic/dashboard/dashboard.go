package dashboard

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
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
		if dashboard.Model, err = dao.Model.CountDocuments(ctx, bson.M{"is_public": true}); err != nil {
			logger.Error(ctx, err)
			return nil, err
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

		if dashboard.TodayUser, err = dao.App.CountDocuments(ctx, bson.M{
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
			"$match": bson.M{},
		},
		{
			"$group": bson.M{
				"_id": "$trace_id",
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
	dashboard.Call = len(result)

	return dashboard, nil
}

// 调用数据
func (s *sDashboard) CallData(ctx context.Context, params model.DashboardCallDataReq) ([]*model.CallData, error) {

	startTime := gtime.Now().AddDate(0, 0, -params.Days-1).StartOfDay()
	endTime := gtime.Now().EndOfDay(true)

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"req_time": bson.M{
					"$gte": startTime.TimestampMilli(),
					"$lte": endTime.TimestampMilli(),
				},
			},
		},
		{
			"$group": bson.M{
				"_id":    "$req_date",
				"count":  bson.M{"$addToSet": "$trace_id"},
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
			Call:   len(gconv.SliceAny(res["count"])),
			Tokens: gconv.Int(res["tokens"]),
			User:   len(gconv.SliceAny(res["user"])),
			App:    len(gconv.SliceAny(res["app"])),
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

	quota, err := redis.HGetInt(ctx, fmt.Sprintf(consts.API_USAGE_KEY, service.Session().GetUserId(ctx)), consts.USER_TOTAL_TOKENS_FIELD)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.Expense{
		Quota: quota,
	}, nil
}

// 数据TOP5
func (s *sDashboard) DataTop5(ctx context.Context, params model.DashboardDataTop5Req) ([]*model.DataTop5, error) {

	startTime := gtime.Now().AddDate(0, 0, -params.Days-1).StartOfDay()
	endTime := gtime.Now().EndOfDay(true)

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"req_time": bson.M{
					"$gte": startTime.TimestampMilli(),
					"$lte": endTime.TimestampMilli(),
				},
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
				"_id":    "$app_id",
				"count":  bson.M{"$sum": 1},
				"models": bson.M{"$addToSet": "$model"},
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

	pipeline = append(pipeline, bson.M{"$sort": bson.M{"count": -1}}, bson.M{"$limit": 5})

	if service.Session().IsUserRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["user_id"] = service.Session().GetUserId(ctx)
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.DataTop5, 0)
	switch params.DataType {
	case "user":
		for _, res := range result {
			items = append(items, &model.DataTop5{
				UserId: gconv.Int(res["_id"]),
				Call:   gconv.Int(res["count"]),
				Models: len(gconv.SliceAny(res["models"])),
				Tokens: gconv.Int(res["tokens"]),
			})
		}
	case "app":
		for _, res := range result {
			items = append(items, &model.DataTop5{
				AppId:  gconv.Int(res["_id"]),
				Call:   gconv.Int(res["count"]),
				Models: len(gconv.SliceAny(res["models"])),
				Tokens: gconv.Int(res["tokens"]),
			})
		}
	case "model":
		for _, res := range result {
			items = append(items, &model.DataTop5{
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

	startTime := gtime.Now().AddDate(0, 0, -params.Days-1).StartOfDay()
	endTime := gtime.Now().EndOfDay(true)

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"req_time": bson.M{
					"$gte": startTime.TimestampMilli(),
					"$lte": endTime.TimestampMilli(),
				},
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
