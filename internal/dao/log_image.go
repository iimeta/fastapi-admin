package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var LogImage = NewLogImageDao()

type LogImageDao struct {
	*MongoDB[entity.LogImage]
}

func NewLogImageDao(database ...string) *LogImageDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &LogImageDao{
		MongoDB: NewMongoDB[entity.LogImage](database[0], LOG_IMAGE),
	}
}
