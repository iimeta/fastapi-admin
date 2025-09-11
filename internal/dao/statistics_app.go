package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var StatisticsApp = NewStatisticsAppDao()

type StatisticsAppDao struct {
	*MongoDB[entity.StatisticsApp]
}

func NewStatisticsAppDao(database ...string) *StatisticsAppDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &StatisticsAppDao{
		MongoDB: NewMongoDB[entity.StatisticsApp](database[0], STATISTICS_APP),
	}
}
