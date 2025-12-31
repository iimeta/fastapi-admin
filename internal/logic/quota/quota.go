package quota

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
)

type sQuota struct {
	noticeRedsync *redsync.Redsync
}

func init() {
	service.RegisterQuota(New())
}

func New() service.IQuota {
	return &sQuota{
		noticeRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}
