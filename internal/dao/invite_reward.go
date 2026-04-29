package dao

import (
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
)

var InviteReward = NewInviteRewardDao()

type InviteRewardDao struct {
	*MongoDB[entity.InviteReward]
}

func NewInviteRewardDao(database ...string) *InviteRewardDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &InviteRewardDao{
		MongoDB: NewMongoDB[entity.InviteReward](database[0], INVITE_REWARD),
	}
}
