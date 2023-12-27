package dashboard

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/logger"
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
