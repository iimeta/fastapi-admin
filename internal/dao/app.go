package dao

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
	"go.mongodb.org/mongo-driver/bson"
)

var App = NewAppDao()

type AppDao struct {
	*MongoDB[entity.App]
}

func NewAppDao(database ...string) *AppDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &AppDao{
		MongoDB: NewMongoDB[entity.App](database[0], do.APP_COLLECTION),
	}
}

// 根据appId查询应用
func (d *AppDao) FindByAppId(ctx context.Context, appId int) (*entity.App, error) {
	return d.FindOne(ctx, bson.M{"app_id": appId})
}
