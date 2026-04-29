package dao

import (
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
)

var InviteRelation = NewInviteRelationDao()

type InviteRelationDao struct {
	*MongoDB[entity.InviteRelation]
}

func NewInviteRelationDao(database ...string) *InviteRelationDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &InviteRelationDao{
		MongoDB: NewMongoDB[entity.InviteRelation](database[0], INVITE_RELATION),
	}
}
