package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var SysAdmin = NewSysAdminDao()

type SysAdminDao struct {
	*MongoDB[entity.SysAdmin]
}

func NewSysAdminDao(database ...string) *SysAdminDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &SysAdminDao{
		MongoDB: NewMongoDB[entity.SysAdmin](database[0], do.SYS_ADMIN_COLLECTION),
	}
}
