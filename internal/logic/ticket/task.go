package ticket

import (
	"context"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *sTicket) AutoCloseTask(ctx context.Context) {

	logger.Info(ctx, "sTicket AutoCloseTask start")

	now := gtime.TimestampMilli()
	mutex := s.ticketRedsync.NewMutex(consts.TASK_TICKET_LOCK_KEY, redsync.WithExpiry(config.Cfg.Ticket.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sTicket AutoCloseTask", err)
		logger.Debugf(ctx, "sTicket AutoCloseTask end time: %d", gtime.TimestampMilli()-now)
		return
	}

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		}
		logger.Debugf(ctx, "sTicket AutoCloseTask end time: %d", gtime.TimestampMilli()-now)
	}()

	if config.Cfg.Ticket == nil || config.Cfg.Ticket.AutoCloseDays <= 0 {
		return
	}

	deadline := gtime.TimestampMilli() - (config.Cfg.Ticket.AutoCloseDays * 24 * time.Hour).Milliseconds()
	filter := bson.M{
		"status":     bson.M{"$in": []int{consts.STATUS_REPLIED, consts.STATUS_RESOLVED}},
		"updated_at": bson.M{"$lte": deadline},
	}

	tickets, err := dao.Ticket.Find(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	if len(tickets) == 0 {
		if _, err := redis.Set(ctx, consts.TASK_TICKET_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
			logger.Error(ctx, err)
		}
		return
	}

	if err := dao.Ticket.UpdateMany(ctx, filter, bson.M{
		"status":     consts.STATUS_CLOSED,
		"updated_at": gtime.TimestampMilli(),
	}); err != nil {
		logger.Error(ctx, err)
		return
	}

	for _, ticket := range tickets {
		ticket.Status = consts.STATUS_CLOSED
		ticket.UpdatedAt = gtime.TimestampMilli()
		s.noticeAutoClosedUser(ctx, ticket)
	}

	if _, err := redis.Set(ctx, consts.TASK_TICKET_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
