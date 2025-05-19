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
	"math"
	"time"
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

	if service.Session().IsResellerRole(ctx) {
		if dashboard.App, err = dao.App.CountDocuments(ctx, bson.M{"rid": service.Session().GetRid(ctx)}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	} else if service.Session().IsUserRole(ctx) {
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

	if service.Session().IsResellerRole(ctx) {

		models := service.Session().GetReseller(ctx).Models
		if groupModels, err := service.Group().GetModelsByGroups(ctx, service.Session().GetReseller(ctx).Groups...); err != nil {
			logger.Error(ctx, err)
		} else {
			models = append(models, groupModels...)
		}

		if len(models) > 0 {
			if dashboard.Model, err = dao.Model.CountDocuments(ctx, bson.M{"_id": bson.M{"$in": models}}); err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
		}

	} else if service.Session().IsUserRole(ctx) {

		models := service.Session().GetUser(ctx).Models
		if groupModels, err := service.Group().GetModelsByGroups(ctx, service.Session().GetUser(ctx).Groups...); err != nil {
			logger.Error(ctx, err)
		} else {
			models = append(models, groupModels...)
		}

		if len(models) > 0 {
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

	if service.Session().IsResellerRole(ctx) {
		if dashboard.AppKey, err = dao.Key.CountDocuments(ctx, bson.M{"rid": service.Session().GetRid(ctx), "type": 1}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	} else if service.Session().IsUserRole(ctx) {
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

	if service.Session().IsResellerRole(ctx) {

		if dashboard.TodayApp, err = dao.App.CountDocuments(ctx, bson.M{
			"rid": service.Session().GetRid(ctx),
			"created_at": bson.M{
				"$gte": gtime.Now().StartOfDay().TimestampMilli(),
				"$lte": gtime.Now().EndOfDay(true).TimestampMilli(),
			},
		}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if groups := service.Session().GetUser(ctx).Groups; len(groups) > 0 {
			if dashboard.Group, err = dao.Group.CountDocuments(ctx, bson.M{"_id": bson.M{"$in": groups}}); err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
		}

		if dashboard.User, err = dao.User.CountDocuments(ctx, bson.M{"rid": service.Session().GetRid(ctx)}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if dashboard.TodayUser, err = dao.User.CountDocuments(ctx, bson.M{
			"rid": service.Session().GetRid(ctx),
			"created_at": bson.M{
				"$gte": gtime.Now().StartOfDay().TimestampMilli(),
				"$lte": gtime.Now().EndOfDay(true).TimestampMilli(),
			},
		}); err != nil {
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
			"$match": bson.M{},
		},
		{
			"$group": bson.M{
				"_id":  nil,
				"call": bson.M{"$sum": "$total"},
			},
		},
	}

	if service.Session().IsResellerRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["user_id"] = service.Session().GetUserId(ctx)
	}

	result := make([]map[string]interface{}, 0)
	if err = dao.StatisticsUser.Aggregate(ctx, pipeline, &result); err != nil {
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
				"stat_time": bson.M{
					"$gte": startTime.TimestampMilli(),
					"$lte": endTime.TimestampMilli(),
				},
			},
		},
		{
			"$group": bson.M{
				"_id":      "$stat_date",
				"count":    bson.M{"$sum": "$total"},
				"tokens":   bson.M{"$sum": "$tokens"},
				"abnormal": bson.M{"$sum": "$abnormal"},
				"user":     bson.M{"$addToSet": "$user_id"},
			},
		},
	}

	if service.Session().IsResellerRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {

		match := pipeline[0]["$match"].(bson.M)
		match["user_id"] = service.Session().GetUserId(ctx)

		group := pipeline[1]["$group"].(bson.M)
		group["app"] = bson.M{"$addToSet": "$app_id"}
		delete(group, "user")
	}

	result := make([]map[string]interface{}, 0)

	if service.Session().IsUserRole(ctx) {
		if err := dao.StatisticsApp.Aggregate(ctx, pipeline, &result); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	} else {
		if err := dao.StatisticsUser.Aggregate(ctx, pipeline, &result); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	resultMap := make(map[string]*model.CallData)
	for _, res := range result {
		resultMap[gconv.String(res["_id"])] = &model.CallData{
			Date:     gconv.String(res["_id"])[5:],
			Spend:    util.QuotaConv(gconv.Int(res["tokens"])),
			Call:     gconv.Int(res["count"]),
			Tokens:   gconv.Int(res["tokens"]),
			Abnormal: gconv.Int(res["abnormal"]),
			User:     len(gconv.SliceAny(res["user"])),
			App:      len(gconv.SliceAny(res["app"])),
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

	var (
		quota                  int
		usedQuota              int
		allocatedQuota         int
		toBeAllocated          int
		quotaExpiresAt         string
		quotaWarning           bool
		warningThreshold       int
		expireWarningThreshold time.Duration
	)

	if service.Session().IsResellerRole(ctx) {

		reseller, err := service.Reseller().GetResellerByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		quota = reseller.Quota
		usedQuota = reseller.UsedQuota
		quotaExpiresAt = reseller.QuotaExpiresAt
		quotaWarning = reseller.QuotaWarning
		warningThreshold = reseller.WarningThreshold
		expireWarningThreshold = reseller.ExpireWarningThreshold

		users, err := dao.User.Find(ctx, bson.M{"rid": reseller.UserId})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		for _, user := range users {

			if user.Quota > 0 {
				allocatedQuota += user.Quota
			}

			allocatedQuota += user.UsedQuota
		}

		toBeAllocated = quota + usedQuota - allocatedQuota
	}

	if service.Session().IsUserRole(ctx) {

		user, err := service.User().GetUserByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		quota = user.Quota
		usedQuota = user.UsedQuota
		quotaExpiresAt = user.QuotaExpiresAt
		quotaWarning = user.QuotaWarning
		warningThreshold = user.WarningThreshold
		expireWarningThreshold = user.ExpireWarningThreshold
	}

	return &model.Expense{
		Quota:                  quota,
		QuotaUSD:               util.Round(float64(quota)/consts.QUOTA_USD_UNIT, 4),
		UsedQuota:              usedQuota,
		UsedQuotaUSD:           util.Round(float64(usedQuota)/consts.QUOTA_USD_UNIT, 4),
		AllocatedQuota:         allocatedQuota,
		ToBeAllocated:          toBeAllocated,
		QuotaExpiresAt:         quotaExpiresAt,
		QuotaWarning:           quotaWarning,
		WarningThreshold:       warningThreshold / consts.QUOTA_USD_UNIT,
		ExpireWarningThreshold: expireWarningThreshold,
	}, nil
}

// 数据TOP
func (s *sDashboard) DataTop(ctx context.Context, params model.DashboardDataTopReq) ([]*model.DataTop, error) {

	startTime := gtime.Now().AddDate(0, 0, -(params.Days - 1)).StartOfDay()
	endTime := gtime.Now().EndOfDay(true)

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"stat_time": bson.M{
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
				"count":  bson.M{"$sum": "$total"},
				"tokens": bson.M{"$sum": "$tokens"},
			},
		})
	case "app":
		pipeline = append(pipeline, bson.M{
			"$group": bson.M{
				"_id":     "$app_id",
				"user_id": bson.M{"$first": "$user_id"},
				"count":   bson.M{"$sum": "$total"},
				"tokens":  bson.M{"$sum": "$tokens"},
			},
		})
	case "app_key":
		pipeline = append(pipeline, bson.M{
			"$group": bson.M{
				"_id":    "$app_key",
				"app_id": bson.M{"$first": "$app_id"},
				"count":  bson.M{"$sum": "$total"},
				"tokens": bson.M{"$sum": "$tokens"},
			},
		})
	case "model":
		pipeline = append(pipeline, bson.M{"$unwind": "$model_stats"})
		pipeline = append(pipeline, bson.M{
			"$group": bson.M{
				"_id":    "$model_stats.model",
				"count":  bson.M{"$sum": "$model_stats.total"},
				"tokens": bson.M{"$sum": "$model_stats.tokens"},
			},
		})
	}

	pipeline = append(pipeline, bson.M{"$sort": bson.M{"tokens": -1}}, bson.M{"$limit": 10})

	if service.Session().IsResellerRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["user_id"] = service.Session().GetUserId(ctx)
	}

	result := make([]map[string]interface{}, 0)
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
				Tokens: gconv.Int(res["tokens"]),
			})
		}

	case "app":

		for _, res := range result {
			items = append(items, &model.DataTop{
				AppId:  gconv.Int(res["_id"]),
				UserId: gconv.Int(res["user_id"]),
				Call:   gconv.Int(res["count"]),
				Tokens: gconv.Int(res["tokens"]),
			})
		}

	case "app_key":

		for _, res := range result {
			items = append(items, &model.DataTop{
				AppKey: util.Desensitize(gconv.String(res["_id"])),
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
				"stat_time": bson.M{
					"$gte": startTime.TimestampMilli(),
					"$lte": endTime.TimestampMilli(),
				},
			},
		},
		{
			"$unwind": "$model_stats",
		},
		{
			"$group": bson.M{
				"_id":   "$model_stats.model",
				"count": bson.M{"$sum": "$model_stats.total"},
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

	if service.Session().IsResellerRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		match := pipeline[0]["$match"].(bson.M)
		match["user_id"] = service.Session().GetUserId(ctx)
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.StatisticsUser.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, nil, err
	}

	items := make([]*model.ModelPercent, 0)
	for _, res := range result {
		items = append(items, &model.ModelPercent{
			Name:  gconv.String(res["_id"]),
			Value: gconv.Int(res["count"]),
		})
	}

	models := make([]string, 0)
	for _, item := range items {
		models = append(models, item.Name)
	}

	return models, items, nil
}

// 每秒钟数据
func (s *sDashboard) PerSecond(ctx context.Context, params model.DashboardPerSecondReq) (int, int, error) {

	startTime := gtime.TimestampMilli() - 5000
	endTime := gtime.TimestampMilli()

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"created_at": bson.M{
					"$gte": startTime,
					"$lte": endTime,
				},
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
			},
		},
		{
			"$group": bson.M{
				"_id":               nil,
				"rps":               bson.M{"$sum": 1},
				"prompt_tokens":     bson.M{"$sum": "$prompt_tokens"},
				"completion_tokens": bson.M{"$sum": "$completion_tokens"},
			},
		},
	}

	match := pipeline[0]["$match"].(bson.M)

	if service.Session().IsResellerRole(ctx) {
		match["rid"] = service.Session().GetRid(ctx)
	}

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

	if len(params.ModelAgents) > 0 && service.Session().IsAdminRole(ctx) {
		match["model_agent_id"] = bson.M{
			"$in": params.ModelAgents,
		}
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return 0, 0, err
	}

	if len(result) > 0 {
		return int(math.Ceil(gconv.Float64(result[0]["rps"]) / 5)), (gconv.Int(result[0]["prompt_tokens"]) + gconv.Int(result[0]["completion_tokens"])) / 5, nil
	}

	return 0, 0, nil
}

// 每分钟数据
func (s *sDashboard) PerMinute(ctx context.Context, params model.DashboardPerMinuteReq) (int, int, error) {

	startTime := gtime.TimestampMilli() - 60000
	endTime := gtime.TimestampMilli()

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"created_at": bson.M{
					"$gte": startTime,
					"$lte": endTime,
				},
				"is_smart_match": bson.M{"$ne": true},
				"is_retry":       bson.M{"$ne": true},
			},
		},
		{
			"$group": bson.M{
				"_id":               nil,
				"rpm":               bson.M{"$sum": 1},
				"prompt_tokens":     bson.M{"$sum": "$prompt_tokens"},
				"completion_tokens": bson.M{"$sum": "$completion_tokens"},
			},
		},
	}

	match := pipeline[0]["$match"].(bson.M)

	if service.Session().IsResellerRole(ctx) {
		match["rid"] = service.Session().GetRid(ctx)
	}

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

	if len(params.ModelAgents) > 0 && service.Session().IsAdminRole(ctx) {
		match["model_agent_id"] = bson.M{
			"$in": params.ModelAgents,
		}
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return 0, 0, err
	}

	if len(result) > 0 {
		return gconv.Int(result[0]["rpm"]), gconv.Int(result[0]["prompt_tokens"]) + gconv.Int(result[0]["completion_tokens"]), nil
	}

	return 0, 0, nil
}

// 额度预警
func (s *sDashboard) QuotaWarning(ctx context.Context, params model.DashboardQuotaWarningReq) error {

	update := bson.M{
		"quota_warning":         params.QuotaWarning,
		"warning_notice":        false,
		"exhaustion_notice":     false,
		"expire_warning_notice": false,
		"expire_notice":         false,
	}

	if params.WarningThreshold > 0 {
		update["warning_threshold"] = params.WarningThreshold * consts.QUOTA_USD_UNIT
	}

	if params.ExpireWarningThreshold > 0 {
		update["expire_warning_threshold"] = params.ExpireWarningThreshold
	}

	if service.Session().IsResellerRole(ctx) {
		if err := dao.Reseller.UpdateOne(ctx, bson.M{"user_id": service.Session().GetUserId(ctx)}, update); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if service.Session().IsUserRole(ctx) {
		if err := dao.User.UpdateOne(ctx, bson.M{"user_id": service.Session().GetUserId(ctx)}, update); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}
