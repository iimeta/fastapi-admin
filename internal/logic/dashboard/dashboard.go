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
func (s *sDashboard) BaseData(ctx context.Context) (*model.Dashboard, error) {

	appCount, err := dao.App.EstimatedDocumentCount(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	todayAppCount, err := dao.App.CountDocuments(ctx, bson.M{
		"created_at": bson.M{
			"$gte": gtime.Now().StartOfDay().Unix(),
			"$lte": gtime.Now().EndOfDay().Unix(),
		},
	})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelCount, err := dao.Model.EstimatedDocumentCount(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	appKeyCount, err := dao.Key.CountDocuments(ctx, bson.M{"type": 1})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelKeyCount, err := dao.Key.CountDocuments(ctx, bson.M{"type": 2})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	userCount, err := dao.User.EstimatedDocumentCount(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	todayUserCount, err := dao.App.CountDocuments(ctx, bson.M{
		"created_at": bson.M{
			"$gte": gtime.Now().StartOfDay().Unix(),
			"$lte": gtime.Now().EndOfDay().Unix(),
		},
	})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.Dashboard{
		App:       appCount,
		TodayApp:  todayAppCount,
		Model:     modelCount,
		AppKey:    appKeyCount,
		ModelKey:  modelKeyCount,
		User:      userCount,
		TodayUser: todayUserCount,
	}, nil
}

// 调用数据
func (s *sDashboard) CallData(ctx context.Context) ([]*model.CallData, error) {

	startTime := gtime.Now().AddDate(0, 0, -9).StartOfDay()
	endTime := gtime.Now().EndOfDay()

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
				"_id":   "$req_date",
				"count": bson.M{"$sum": 1},
			},
		},
	}

	result := make([]map[string]interface{}, 0)
	if err := dao.Chat.Aggregate(ctx, pipeline, &result); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	resultMap := make(map[string]int)
	for _, res := range result {
		resultMap[gconv.String(res["_id"])] = gconv.Int(res["count"])
	}

	items := make([]*model.CallData, 0)
	days := util.Day(startTime.String(), endTime.String())

	for _, day := range days {
		items = append(items, &model.CallData{
			Date:  day.StartDate,
			Count: resultMap[day.StartDate],
		})
	}

	return items, nil
}

// 费用
func (s *sDashboard) Expense(ctx context.Context) (*model.Expense, error) {

	if service.Session().GetRole(ctx) == consts.SESSION_ADMIN {
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
