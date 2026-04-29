package dao

import (
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
)

var InviteRewardApply = NewInviteRewardApplyDao()

type InviteRewardApplyDao struct {
	*MongoDB[entity.InviteRewardApply]
}

func NewInviteRewardApplyDao(database ...string) *InviteRewardApplyDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &InviteRewardApplyDao{
		MongoDB: NewMongoDB[entity.InviteRewardApply](database[0], INVITE_REWARD_APPLY),
	}
}
