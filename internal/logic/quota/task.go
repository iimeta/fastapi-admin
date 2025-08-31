package quota

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/logic/common"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/email"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

// 额度通知任务
func (s *sQuota) NoticeTask(ctx context.Context) {

	logger.Info(ctx, "sQuota NoticeTask start")

	now := gtime.TimestampMilli()

	mutex := s.noticeRedsync.NewMutex(consts.TASK_QUOTA_NOTICE_LOCK_KEY, redsync.WithExpiry(config.Cfg.QuotaTask.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sQuota NoticeTask", err)
		logger.Debugf(ctx, "sQuota NoticeTask end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sQuota NoticeTask lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sQuota NoticeTask unlock")
		}
		logger.Debugf(ctx, "sQuota NoticeTask end time: %d", gtime.TimestampMilli()-now)
	}()

	if users, err := dao.User.Find(ctx, bson.M{"status": 1}); err == nil {

		for _, user := range users {

			if err = email.Verify(user.Email); err != nil {
				logger.Infof(ctx, "sQuota NoticeTask user: %d, error: %v", user.UserId, err)
				continue
			}

			var (
				warningThreshold       = user.WarningThreshold
				expireWarningThreshold = user.ExpireWarningThreshold
				scene                  string
				title                  string
				content                string
				noticeTemplate         *model.NoticeTemplate
				siteConfig             *entity.SiteConfig
				dialer                 = email.NewDefaultDialer()
			)

			if warningThreshold == 0 {
				warningThreshold = config.Cfg.Quota.Threshold
			}

			if expireWarningThreshold == 0 {
				expireWarningThreshold = config.Cfg.Quota.ExpiredThreshold
			}

			if config.Cfg.Quota.Warning && !user.WarningNotice && !user.ExhaustionNotice && user.UsedQuota != 0 && user.Quota <= warningThreshold && (user.QuotaExpiresAt == 0 || user.QuotaExpiresAt > gtime.TimestampMilli()) {
				scene = consts.SCENE_QUOTA_WARNING
			} else if config.Cfg.Quota.ExhaustedNotice && !user.ExhaustionNotice && user.UsedQuota != 0 && user.Quota <= 0 && (user.QuotaExpiresAt == 0 || user.QuotaExpiresAt > gtime.TimestampMilli()) {
				scene = consts.SCENE_QUOTA_EXHAUSTION
			} else if config.Cfg.Quota.ExpiredWarning && !user.ExpireWarningNotice && user.Quota > 0 && user.QuotaExpiresAt > 0 && gtime.TimestampMilli() > user.QuotaExpiresAt-(expireWarningThreshold*gtime.D).Milliseconds() {
				scene = consts.SCENE_QUOTA_EXPIRE_WARNING
			} else if config.Cfg.Quota.ExpiredNotice && !user.ExpireNotice && user.Quota > 0 && user.QuotaExpiresAt > 0 && gtime.TimestampMilli() > user.QuotaExpiresAt {
				scene = consts.SCENE_QUOTA_EXPIRE
			}

			if scene == "" {
				continue
			}

			if (scene == consts.SCENE_QUOTA_WARNING || scene == consts.SCENE_QUOTA_EXPIRE_WARNING) && !user.QuotaWarning && user.WarningThreshold > 0 && user.ExpireWarningThreshold > 0 {
				continue
			}

			account, err := dao.Account.FindOne(ctx, bson.M{"user_id": user.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
			if err != nil {
				logger.Error(ctx, err)
				continue
			}

			if account == nil {
				logger.Infof(ctx, "sQuota NoticeTask user: %d, 因无可用账号, 不发送提醒邮件", user.UserId)
				continue
			}

			if user.Rid > 0 {

				isConfigEmail := false

				if siteConfig = service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain); siteConfig != nil && siteConfig.Rid == user.Rid && siteConfig.Host != "" {
					dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
					isConfigEmail = true
				}

				if !isConfigEmail {
					siteConfigs := service.SiteConfig().GetSiteConfigsByRid(ctx, user.Rid)
					for _, siteConfig = range siteConfigs {
						if siteConfig != nil && siteConfig.Host != "" {
							dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
							isConfigEmail = true
							break
						}
					}
				}

				if !isConfigEmail {
					logger.Infof(ctx, "sQuota NoticeTask 因代理商: %d, 所有站点未配置邮箱, 不发送提醒邮件", user.Rid)
					continue
				}

			} else {

				siteConfig = service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain)
				if siteConfig == nil {
					if siteConfig, err = dao.SiteConfig.FindOne(ctx, bson.M{"user_id": 1, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}}); err != nil {
						logger.Error(ctx, err)
					}
				}

				if siteConfig != nil && siteConfig.Host != "" {
					dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
				} else {
					logger.Infof(ctx, "sQuota NoticeTask 因站点 %s 未配置邮箱, 默认使用系统配置邮箱", account.LoginDomain)
				}
			}

			if noticeTemplate, err = service.NoticeTemplate().GetNoticeTemplateByScene(ctx, scene, []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL}); err != nil {
				logger.Error(ctx, err)
				continue
			}

			data := common.GetVariableData(ctx, user, nil, siteConfig, noticeTemplate.Variables)

			data["name"] = user.Name

			if user.Quota < 0 {
				data["quota"] = fmt.Sprintf("-$%f", util.Round(math.Abs(float64(user.Quota))/consts.QUOTA_USD_UNIT, 6))
			} else {
				data["quota"] = fmt.Sprintf("$%f", util.Round(float64(user.Quota)/consts.QUOTA_USD_UNIT, 6))
			}

			if scene == consts.SCENE_QUOTA_WARNING {
				data["warning_threshold"] = fmt.Sprintf("$%d", warningThreshold/consts.QUOTA_USD_UNIT)
			} else if scene == consts.SCENE_QUOTA_EXPIRE_WARNING || scene == consts.SCENE_QUOTA_EXPIRE {
				data["quota_expires_at"] = util.FormatDateTime(user.QuotaExpiresAt)
			}

			if title, content, err = util.RenderTemplate(noticeTemplate.Title, noticeTemplate.Content, data); err != nil {
				logger.Error(ctx, err)
				continue
			}

			if err = email.SendMailTask(ctx, email.NewMessage([]string{user.Email}, title, content), dialer); err != nil {
				logger.Errorf(ctx, "sQuota NoticeTask user: %d, email: %s, SendMailTask %s error: %v", user.UserId, user.Email, title, err)
				continue
			}

			logger.Infof(ctx, "sQuota NoticeTask user: %d, email: %s, SendMailTask %s success", user.UserId, user.Email, title)

			if err = dao.User.UpdateById(ctx, user.Id, bson.M{
				consts.QUOTA_NOTICE[scene]: true,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}

	} else {
		logger.Error(ctx, err)
	}

	if resellers, err := dao.Reseller.Find(ctx, bson.M{"status": 1}); err == nil {

		for _, reseller := range resellers {

			if err = email.Verify(reseller.Email); err != nil {
				logger.Infof(ctx, "sQuota NoticeTask reseller: %d, error: %v", reseller.UserId, err)
				continue
			}

			var (
				warningThreshold       = reseller.WarningThreshold
				expireWarningThreshold = reseller.ExpireWarningThreshold
				scene                  string
				title                  string
				content                string
				noticeTemplate         *model.NoticeTemplate
				siteConfig             *entity.SiteConfig
				dialer                 = email.NewDefaultDialer()
			)

			if warningThreshold == 0 {
				warningThreshold = config.Cfg.Quota.Threshold
			}

			if expireWarningThreshold == 0 {
				expireWarningThreshold = config.Cfg.Quota.ExpiredThreshold
			}

			if config.Cfg.Quota.Warning && !reseller.WarningNotice && !reseller.ExhaustionNotice && reseller.UsedQuota != 0 && reseller.Quota <= warningThreshold && (reseller.QuotaExpiresAt == 0 || reseller.QuotaExpiresAt > gtime.TimestampMilli()) {
				scene = consts.SCENE_QUOTA_WARNING
			} else if config.Cfg.Quota.ExhaustedNotice && !reseller.ExhaustionNotice && reseller.UsedQuota != 0 && reseller.Quota <= 0 && (reseller.QuotaExpiresAt == 0 || reseller.QuotaExpiresAt > gtime.TimestampMilli()) {
				scene = consts.SCENE_QUOTA_EXHAUSTION
			} else if config.Cfg.Quota.ExpiredWarning && !reseller.ExpireWarningNotice && reseller.Quota > 0 && reseller.QuotaExpiresAt > 0 && gtime.TimestampMilli() > reseller.QuotaExpiresAt-(expireWarningThreshold*gtime.D).Milliseconds() {
				scene = consts.SCENE_QUOTA_EXPIRE_WARNING
			} else if config.Cfg.Quota.ExpiredNotice && !reseller.ExpireNotice && reseller.Quota > 0 && reseller.QuotaExpiresAt > 0 && gtime.TimestampMilli() > reseller.QuotaExpiresAt {
				scene = consts.SCENE_QUOTA_EXPIRE
			}

			if scene == "" {
				continue
			}

			if (scene == consts.SCENE_QUOTA_WARNING || scene == consts.SCENE_QUOTA_EXPIRE_WARNING) && !reseller.QuotaWarning && reseller.WarningThreshold > 0 && reseller.ExpireWarningThreshold > 0 {
				continue
			}

			account, err := dao.ResellerAccount.FindOne(ctx, bson.M{"user_id": reseller.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
			if err != nil {
				logger.Error(ctx, err)
				continue
			}

			if account == nil {
				logger.Infof(ctx, "sQuota NoticeTask reseller: %d, 因无可用账号, 不发送提醒邮件", reseller.UserId)
				continue
			}

			siteConfig = service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain)
			if siteConfig == nil {
				if siteConfig, err = dao.SiteConfig.FindOne(ctx, bson.M{"user_id": 1, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}}); err != nil {
					logger.Error(ctx, err)
				}
			}

			if siteConfig != nil && siteConfig.Host != "" {
				dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
			} else {
				logger.Infof(ctx, "sQuota NoticeTask 因站点 %s 未配置邮箱, 默认使用系统配置邮箱", account.LoginDomain)
			}

			if noticeTemplate, err = service.NoticeTemplate().GetNoticeTemplateByScene(ctx, scene, []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL}); err != nil {
				logger.Error(ctx, err)
				continue
			}

			data := common.GetVariableData(ctx, nil, reseller, siteConfig, noticeTemplate.Variables)

			data["name"] = reseller.Name

			if reseller.Quota < 0 {
				data["quota"] = fmt.Sprintf("-$%f", util.Round(math.Abs(float64(reseller.Quota))/consts.QUOTA_USD_UNIT, 6))
			} else {
				data["quota"] = fmt.Sprintf("$%f", util.Round(float64(reseller.Quota)/consts.QUOTA_USD_UNIT, 6))
			}

			if scene == consts.SCENE_QUOTA_WARNING {
				data["warning_threshold"] = fmt.Sprintf("$%d", warningThreshold/consts.QUOTA_USD_UNIT)
			} else if scene == consts.SCENE_QUOTA_EXPIRE_WARNING || scene == consts.SCENE_QUOTA_EXPIRE {
				data["quota_expires_at"] = util.FormatDateTime(reseller.QuotaExpiresAt)
			}

			if title, content, err = util.RenderTemplate(noticeTemplate.Title, noticeTemplate.Content, data); err != nil {
				logger.Error(ctx, err)
				continue
			}

			if err = email.SendMailTask(ctx, email.NewMessage([]string{reseller.Email}, title, content), dialer); err != nil {
				logger.Errorf(ctx, "sQuota NoticeTask reseller: %d, email: %s, SendMailTask %s error: %v", reseller.UserId, reseller.Email, title, err)
				continue
			}

			logger.Infof(ctx, "sQuota NoticeTask reseller: %d, email: %s, SendMailTask %s success", reseller.UserId, reseller.Email, title)

			if err = dao.Reseller.UpdateById(ctx, reseller.Id, bson.M{
				consts.QUOTA_NOTICE[scene]: true,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}

	} else {
		logger.Error(ctx, err)
	}

	if _, err := redis.Set(ctx, consts.TASK_QUOTA_NOTICE_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}

// 额度清零任务
func (s *sQuota) ClearTask(ctx context.Context) {

	logger.Info(ctx, "sQuota ClearTask start")

	now := gtime.TimestampMilli()

	mutex := s.noticeRedsync.NewMutex(consts.TASK_QUOTA_CLEAR_LOCK_KEY, redsync.WithExpiry(config.Cfg.QuotaTask.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sQuota ClearTask", err)
		logger.Debugf(ctx, "sQuota ClearTask end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sQuota ClearTask lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sQuota ClearTask unlock")
		}
		logger.Debugf(ctx, "sQuota ClearTask end time: %d", gtime.TimestampMilli()-now)
	}()

	if users, err := dao.User.Find(ctx, bson.M{"quota": bson.M{"$gt": 0}, "quota_expires_at": bson.M{
		"$gt":  0,
		"$lte": gtime.TimestampMilli() - (config.Cfg.Quota.ExpiredClearDefer * time.Minute).Milliseconds(),
	}}); err == nil {

		for _, user := range users {

			newData, err := dao.User.FindOneAndUpdate(ctx, bson.M{"user_id": user.UserId}, bson.M{
				"$inc": bson.M{
					"quota": -user.Quota,
				},
			})
			if err != nil {
				logger.Error(ctx, err)
			}

			if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_USER_USAGE_KEY, user.UserId), consts.USER_QUOTA_FIELD, int64(-user.Quota)); err != nil {
				logger.Error(ctx, err)
			}

			// 交易记录
			if _, err = dao.DealRecord.Insert(ctx, &do.DealRecord{
				UserId: user.UserId,
				Quota:  -user.Quota,
				Type:   4,
				Status: 1,
				Rid:    user.Rid,
			}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: user,
				NewData: newData,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}

	} else {
		logger.Error(ctx, err)
	}

	if resellers, err := dao.Reseller.Find(ctx, bson.M{"quota": bson.M{"$gt": 0}, "quota_expires_at": bson.M{
		"$gt":  0,
		"$lte": gtime.TimestampMilli() - (config.Cfg.Quota.ExpiredClearDefer * time.Minute).Milliseconds(),
	}}); err == nil {

		for _, reseller := range resellers {

			newData, err := dao.Reseller.FindOneAndUpdate(ctx, bson.M{"user_id": reseller.UserId}, bson.M{
				"$inc": bson.M{
					"quota": -reseller.Quota,
				},
			})
			if err != nil {
				logger.Error(ctx, err)
			}

			if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_RESELLER_USAGE_KEY, reseller.UserId), consts.RESELLER_QUOTA_FIELD, int64(-reseller.Quota)); err != nil {
				logger.Error(ctx, err)
			}

			// 交易记录
			if _, err = dao.DealRecord.Insert(ctx, &do.DealRecord{
				UserId: reseller.UserId,
				Quota:  -reseller.Quota,
				Type:   4,
				Status: 1,
				Rid:    reseller.UserId,
			}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: reseller,
				NewData: newData,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}

	} else {
		logger.Error(ctx, err)
	}

	if _, err := redis.Set(ctx, consts.TASK_QUOTA_CLEAR_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
