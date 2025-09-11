package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var Notice = NewNoticeDao()

type NoticeDao struct {
	*MongoDB[entity.Notice]
}

func NewNoticeDao(database ...string) *NoticeDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &NoticeDao{
		MongoDB: NewMongoDB[entity.Notice](database[0], NOTICE),
	}
}
