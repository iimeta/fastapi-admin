package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var SysMenu = NewSysMenuDao()

type SysMenuDao struct {
	*MongoDB[entity.SysMenu]
}

func NewSysMenuDao(database ...string) *SysMenuDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &SysMenuDao{
		MongoDB: NewMongoDB[entity.SysMenu](database[0], do.SYS_MENU_COLLECTION),
	}
}
