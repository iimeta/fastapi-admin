package dao

import (
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
)

var StatisticsCursor = NewStatisticsCursorDao()

type StatisticsCursorDao struct {
	*MongoDB[entity.StatisticsCursor]
}

func NewStatisticsCursorDao(database ...string) *StatisticsCursorDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &StatisticsCursorDao{
		MongoDB: NewMongoDB[entity.StatisticsCursor](database[0], STATISTICS_CURSOR),
	}
}
