package notice

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
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

// 额度预警任务
func (s *sNotice) QuotaWarningTask(ctx context.Context) {

	logger.Info(ctx, "sNotice QuotaWarningTask start")

	now := gtime.TimestampMilli()

	mutex := s.noticeRedsync.NewMutex(consts.TASK_QUOTA_WARNING_LOCK_KEY, redsync.WithExpiry(config.Cfg.Notice.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sNotice QuotaWarningTask", err)
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

	if users, err := dao.User.Find(ctx, bson.M{"status": 1}); err == nil {

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

			if user.WarningThreshold == 0 {
				user.WarningThreshold = config.Cfg.QuotaWarning.Threshold
			}

			if user.ExpireWarningThreshold == 0 {
				user.ExpireWarningThreshold = config.Cfg.QuotaWarning.ExpireThreshold
			}

			if !user.WarningNotice && !user.ExhaustionNotice && user.UsedQuota != 0 && user.Quota <= user.WarningThreshold {
				action = consts.ACTION_WARNING_NOTICE
			} else if config.Cfg.QuotaWarning.ExhaustionNotice && !user.ExhaustionNotice && user.UsedQuota != 0 && user.Quota <= 0 {
				action = consts.ACTION_EXHAUSTION_NOTICE
			} else if config.Cfg.QuotaWarning.ExpireWarning && !user.ExpireWarningNotice && user.Quota > 0 && user.QuotaExpiresAt > 0 && gtime.TimestampMilli() > user.QuotaExpiresAt-(user.ExpireWarningThreshold*gtime.D).Milliseconds() {
				action = consts.ACTION_EXPIRE_WARNING_NOTICE
			} else if config.Cfg.QuotaWarning.ExpireNotice && !user.ExpireNotice && user.Quota > 0 && user.QuotaExpiresAt > 0 && gtime.TimestampMilli() > user.QuotaExpiresAt {
				action = consts.ACTION_EXPIRE_NOTICE
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
				data["warning_threshold"] = user.WarningThreshold / consts.QUOTA_USD_UNIT
			} else if action == consts.ACTION_EXPIRE_WARNING_NOTICE || action == consts.ACTION_EXPIRE_NOTICE {
				data["quota_expires_at"] = util.FormatDateTime(user.QuotaExpiresAt)
			}

			if template, err = util.RenderTemplate(data, action); err != nil {
				logger.Error(ctx, err)
				continue
			}

			dialer := email.NewDefaultDialer()

			account, err := dao.Account.FindOne(ctx, bson.M{"user_id": user.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
			if err != nil {
				logger.Error(ctx, err)
				continue
			}

			if account == nil {
				logger.Infof(ctx, "sNotice QuotaWarningTask user: %d, 因无可用账号, 不发送提醒邮件", user.UserId)
				continue
			}

			if user.Rid > 0 {

				isConfigEmail := false

				if account.LoginDomain != "" {
					siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain)
					if siteConfig != nil && siteConfig.Rid == user.Rid && siteConfig.Host != "" {
						dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password)
						isConfigEmail = true
					}
				}

				if !isConfigEmail {
					siteConfigs := service.SiteConfig().GetSiteConfigsByRid(ctx, user.Rid)
					for _, siteConfig := range siteConfigs {
						if siteConfig != nil && siteConfig.Host != "" {
							dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password)
							isConfigEmail = true
							break
						}
					}
				}

				if !isConfigEmail {
					logger.Infof(ctx, "sNotice QuotaWarningTask 因代理商: %d, 所有站点未配置邮箱, 不发送提醒邮件", user.Rid)
					continue
				}

			} else if account.LoginDomain != "" {
				siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain)
				if siteConfig != nil && siteConfig.Host != "" {
					dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password)
				} else {
					logger.Infof(ctx, "sNotice QuotaWarningTask 因站点 %s 未配置邮箱, 默认使用系统配置邮箱", account.LoginDomain)
				}
			}

			if err = email.SendMail(email.NewMessage([]string{user.Email}, consts.ACTION_MAP[action], template), dialer); err != nil {
				logger.Errorf(ctx, "sNotice QuotaWarningTask user: %d, email: %s, SendMail %s error: %v", user.UserId, user.Email, consts.ACTION_MAP[action], err)
				continue
			}

			logger.Infof(ctx, "sNotice QuotaWarningTask user: %d, email: %s, SendMail %s success", user.UserId, user.Email, consts.ACTION_MAP[action])

			if err = dao.User.UpdateById(ctx, user.Id, bson.M{
				action: true,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}

	} else {
		logger.Error(ctx, err)
	}

	if resellers, err := dao.Reseller.Find(ctx, bson.M{"status": 1}); err == nil {

		for _, reseller := range resellers {

			if !reseller.QuotaWarning && reseller.WarningThreshold > 0 && reseller.ExpireWarningThreshold > 0 {
				continue
			}

			if err = g.Validator().Data(reseller.Email).Rules("email").Run(ctx); err != nil {
				logger.Infof(ctx, "sNotice QuotaWarningTask reseller: %d, error: %v", reseller.UserId, err)
				continue
			}

			var (
				action   string
				template string
			)

			if reseller.WarningThreshold == 0 {
				reseller.WarningThreshold = config.Cfg.QuotaWarning.Threshold
			}

			if reseller.ExpireWarningThreshold == 0 {
				reseller.ExpireWarningThreshold = config.Cfg.QuotaWarning.ExpireThreshold
			}

			if !reseller.WarningNotice && !reseller.ExhaustionNotice && reseller.UsedQuota != 0 && reseller.Quota <= reseller.WarningThreshold {
				action = consts.ACTION_WARNING_NOTICE
			} else if config.Cfg.QuotaWarning.ExhaustionNotice && !reseller.ExhaustionNotice && reseller.UsedQuota != 0 && reseller.Quota <= 0 {
				action = consts.ACTION_EXHAUSTION_NOTICE
			} else if config.Cfg.QuotaWarning.ExpireWarning && !reseller.ExpireWarningNotice && reseller.Quota > 0 && reseller.QuotaExpiresAt > 0 && gtime.TimestampMilli() > reseller.QuotaExpiresAt-(reseller.ExpireWarningThreshold*gtime.D).Milliseconds() {
				action = consts.ACTION_EXPIRE_WARNING_NOTICE
			} else if config.Cfg.QuotaWarning.ExpireNotice && !reseller.ExpireNotice && reseller.Quota > 0 && reseller.QuotaExpiresAt > 0 && gtime.TimestampMilli() > reseller.QuotaExpiresAt {
				action = consts.ACTION_EXPIRE_NOTICE
			}

			if action == "" {
				continue
			}

			data := make(map[string]any)
			quota := util.Round(float64(reseller.Quota)/consts.QUOTA_USD_UNIT, 6)
			if quota < 0 {
				data["quota"] = fmt.Sprintf("-$%f", math.Abs(quota))
			} else {
				data["quota"] = fmt.Sprintf("$%f", quota)
			}

			if action == consts.ACTION_WARNING_NOTICE {
				data["warning_threshold"] = reseller.WarningThreshold / consts.QUOTA_USD_UNIT
			} else if action == consts.ACTION_EXPIRE_WARNING_NOTICE || action == consts.ACTION_EXPIRE_NOTICE {
				data["quota_expires_at"] = util.FormatDateTime(reseller.QuotaExpiresAt)
			}

			if template, err = util.RenderTemplate(data, action); err != nil {
				logger.Error(ctx, err)
				continue
			}

			dialer := email.NewDefaultDialer()

			account, err := dao.ResellerAccount.FindOne(ctx, bson.M{"user_id": reseller.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
			if err != nil {
				logger.Error(ctx, err)
				continue
			}

			if account == nil {
				logger.Infof(ctx, "sNotice QuotaWarningTask reseller: %d, 因无可用账号, 不发送提醒邮件", reseller.UserId)
				continue
			}

			if account.LoginDomain != "" {
				siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain)
				if siteConfig != nil && siteConfig.Host != "" {
					dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password)
				} else {
					logger.Infof(ctx, "sNotice QuotaWarningTask 因站点 %s 未配置邮箱, 默认使用系统配置邮箱", account.LoginDomain)
				}
			}

			if err = email.SendMail(email.NewMessage([]string{reseller.Email}, consts.ACTION_MAP[action], template), dialer); err != nil {
				logger.Errorf(ctx, "sNotice QuotaWarningTask reseller: %d, email: %s, SendMail %s error: %v", reseller.UserId, reseller.Email, consts.ACTION_MAP[action], err)
				continue
			}

			logger.Infof(ctx, "sNotice QuotaWarningTask reseller: %d, email: %s, SendMail %s success", reseller.UserId, reseller.Email, consts.ACTION_MAP[action])

			if err = dao.Reseller.UpdateById(ctx, reseller.Id, bson.M{
				action: true,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}

	} else {
		logger.Error(ctx, err)
	}

	if _, err := redis.Set(ctx, consts.TASK_QUOTA_WARNING_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
