package dao

import (
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
)

var StatisticsAppKey = NewStatisticsAppKeyDao()

type StatisticsAppKeyDao struct {
	*MongoDB[entity.StatisticsAppKey]
}

func NewStatisticsAppKeyDao(database ...string) *StatisticsAppKeyDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &StatisticsAppKeyDao{
		MongoDB: NewMongoDB[entity.StatisticsAppKey](database[0], STATISTICS_APP_KEY),
	}
}
