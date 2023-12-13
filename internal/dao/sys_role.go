package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var SysRole = NewSysRoleDao()

type SysRoleDao struct {
	*MongoDB[entity.SysRole]
}

func NewSysRoleDao(database ...string) *SysRoleDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &SysRoleDao{
		MongoDB: NewMongoDB[entity.SysRole](database[0], do.SYS_ROLE_COLLECTION),
	}
}
