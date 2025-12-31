package dao

import (
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
)

var DealRecord = NewDealRecordDao()

type DealRecordDao struct {
	*MongoDB[entity.DealRecord]
}

func NewDealRecordDao(database ...string) *DealRecordDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &DealRecordDao{
		MongoDB: NewMongoDB[entity.DealRecord](database[0], DEAL_RECORD),
	}
}
