package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var AppKey = NewAppKeyDao()

type AppKeyDao struct {
	*MongoDB[entity.AppKey]
}

func NewAppKeyDao(database ...string) *AppKeyDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &AppKeyDao{
		MongoDB: NewMongoDB[entity.AppKey](database[0], do.APP_KEY_COLLECTION),
	}
}
