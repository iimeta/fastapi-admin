package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var SysSettings = NewSysSettingsDao()

type SysSettingsDao struct {
	*MongoDB[entity.SysSettings]
}

func NewSysSettingsDao(database ...string) *SysSettingsDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &SysSettingsDao{
		MongoDB: NewMongoDB[entity.SysSettings](database[0], do.SYS_SETTINGS_COLLECTION),
	}
}
