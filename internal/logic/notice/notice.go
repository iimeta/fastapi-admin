package notice

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/email"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"math"
	"time"
)

type sNotice struct {
	noticeRedsync *redsync.Redsync
}

func init() {
	service.RegisterNotice(New())
}

func New() service.INotice {
	return &sNotice{
		noticeRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}

// 额度预警任务
func (s *sNotice) QuotaWarningTask(ctx context.Context) {

	logger.Info(ctx, "sNotice QuotaWarningTask start")

	now := gtime.TimestampMilli()

	mutex := s.noticeRedsync.NewMutex(consts.TASK_QUOTA_WARNING_LOCK_KEY, redsync.WithExpiry(config.Cfg.Notice.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, err)
		logger.Debugf(ctx, "sNotice QuotaWarningTask end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sNotice QuotaWarningTask lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sNotice QuotaWarningTask unlock")
		}
		logger.Debugf(ctx, "sNotice QuotaWarningTask end time: %d", gtime.TimestampMilli()-now)
	}()

	users, err := dao.User.Find(ctx, bson.M{"status": 1})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	dialer := email.NewDefaultDialer()

	for _, user := range users {

		if !user.QuotaWarning && user.WarningThreshold > 0 && user.ExpireWarningThreshold > 0 {
			continue
		}

		if err = g.Validator().Data(user.Email).Rules("email").Run(ctx); err != nil {
			logger.Infof(ctx, "sNotice QuotaWarningTask user: %d, error: %v", user.UserId, err)
			continue
		}

		var (
			action   string
			template string
		)

		if !user.WarningNotice && !user.ExhaustionNotice && user.UsedQuota != 0 && user.Quota <= config.Cfg.QuotaWarning.Threshold {
			action = consts.ACTION_WARNING_NOTICE
		} else if config.Cfg.QuotaWarning.ExhaustionNotice && !user.ExhaustionNotice && user.UsedQuota != 0 && user.Quota <= 0 {
			action = consts.ACTION_EXHAUSTION_NOTICE
		}

		if action == "" {
			continue
		}

		data := make(map[string]any)
		quota := util.Round(float64(user.Quota)/consts.QUOTA_USD_UNIT, 6)
		if quota < 0 {
			data["quota"] = fmt.Sprintf("-$%f", math.Abs(quota))
		} else {
			data["quota"] = fmt.Sprintf("$%f", quota)
		}

		if action == consts.ACTION_WARNING_NOTICE {
			data["warning_threshold"] = config.Cfg.QuotaWarning.Threshold / consts.QUOTA_USD_UNIT
		}

		if action == consts.ACTION_WARNING_NOTICE {
			if template, err = util.RenderQuotaWarningTemplate(data); err != nil {
				logger.Error(ctx, err)
				continue
			}
		} else {
			if template, err = util.RenderExhaustionNoticeTemplate(data); err != nil {
				logger.Error(ctx, err)
				continue
			}
		}

		// 发送邮件验证码
		if err = email.SendMail(email.NewMessage([]string{user.Email}, consts.ACTION_MAP[action], template), dialer); err != nil {
			logger.Errorf(ctx, "sNotice QuotaWarningTask user: %d, email: %s, SendMail %s error: %v", user.UserId, user.Email, consts.ACTION_MAP[action], err)
			continue
		}

		logger.Infof(ctx, "sNotice QuotaWarningTask user: %d, email: %s, SendMail %s success", user.UserId, user.Email, consts.ACTION_MAP[action])

		if action == consts.ACTION_WARNING_NOTICE {
			if err = dao.User.UpdateById(ctx, user.Id, bson.M{
				"warning_notice": true,
			}); err != nil {
				logger.Error(ctx, err)
			}
		} else {
			if err = dao.User.UpdateById(ctx, user.Id, bson.M{
				"exhaustion_notice": true,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}
	}

	if _, err = redis.Set(ctx, consts.TASK_QUOTA_WARNING_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
