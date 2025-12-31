package dao

import (
	"context"

	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/crypto"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
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
		MongoDB: NewMongoDB[entity.SysAdmin](database[0], SYS_ADMIN),
	}
}

func (d *SysAdminDao) ChangePassword(ctx context.Context, uid string, password string) error {

	salt := grand.Letters(8)
	if err := SysAdmin.UpdateById(ctx, uid, bson.M{
		"password": crypto.EncryptPassword(password + salt),
		"salt":     salt,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
