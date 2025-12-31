package dao

import (
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
)

var NoticeTemplate = NewNoticeTemplateDao()

type NoticeTemplateDao struct {
	*MongoDB[entity.NoticeTemplate]
}

func NewNoticeTemplateDao(database ...string) *NoticeTemplateDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &NoticeTemplateDao{
		MongoDB: NewMongoDB[entity.NoticeTemplate](database[0], NOTICE_TEMPLATE),
	}
}
