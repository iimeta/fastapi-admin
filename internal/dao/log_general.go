package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var LogGeneral = NewLogGeneralDao()

type LogGeneralDao struct {
	*MongoDB[entity.LogGeneral]
}

func NewLogGeneralDao(database ...string) *LogGeneralDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &LogGeneralDao{
		MongoDB: NewMongoDB[entity.LogGeneral](database[0], LOG_GENERAL),
	}
}
