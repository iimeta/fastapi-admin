package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
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
		MongoDB: NewMongoDB[entity.DealRecord](database[0], do.DEAL_RECORD_COLLECTION),
	}
}
