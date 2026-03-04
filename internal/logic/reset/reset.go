package reset

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
)

type sReset struct {
	resetRedsync *redsync.Redsync
}

func init() {
	service.RegisterReset(New())
}

func New() service.IReset {
	return &sReset{
		resetRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}
