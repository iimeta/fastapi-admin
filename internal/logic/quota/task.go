package quota

import (
	"context"
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/email"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type quotaNoticeTarget struct {
	Id                string
	UserId            int
	Rid               int
	Name              string
	Email             string
	Quota             int
	UsedQuota         int
	QuotaExpiresAt    int64
	QuotaWarning      bool
	WarningThresholds []int
	ExpireThresholds  []int
	WarningNotices    []int
	ExpireNotices     []int
	ExhaustionNotice  bool
	ExpireNotice      bool
	NoticeQuota       int
	NoticeQuotaExpiry int64
	IsReseller        bool
}

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
			s.handleQuotaNotice(ctx, quotaNoticeTarget{
				Id:                user.Id,
				UserId:            user.UserId,
				Rid:               user.Rid,
				Name:              user.Name,
				Email:             user.Email,
				Quota:             user.Quota,
				UsedQuota:         user.UsedQuota,
				QuotaExpiresAt:    user.QuotaExpiresAt,
				QuotaWarning:      user.QuotaWarning,
				WarningThresholds: user.WarningThresholds,
				ExpireThresholds:  user.ExpireWarningThresholds,
				WarningNotices:    user.WarningNotices,
				ExpireNotices:     user.ExpireWarningNotices,
				ExhaustionNotice:  user.ExhaustionNotice,
				ExpireNotice:      user.ExpireNotice,
				NoticeQuota:       user.NoticeQuota,
				NoticeQuotaExpiry: user.NoticeQuotaExpiry,
			})
		}
	} else {
		logger.Error(ctx, err)
	}

	if resellers, err := dao.Reseller.Find(ctx, bson.M{"status": 1}); err == nil {

		for _, reseller := range resellers {
			s.handleQuotaNotice(ctx, quotaNoticeTarget{
				Id:                reseller.Id,
				UserId:            reseller.UserId,
				Name:              reseller.Name,
				Email:             reseller.Email,
				Quota:             reseller.Quota,
				UsedQuota:         reseller.UsedQuota,
				QuotaExpiresAt:    reseller.QuotaExpiresAt,
				QuotaWarning:      reseller.QuotaWarning,
				WarningThresholds: reseller.WarningThresholds,
				ExpireThresholds:  reseller.ExpireWarningThresholds,
				WarningNotices:    reseller.WarningNotices,
				ExpireNotices:     reseller.ExpireWarningNotices,
				ExhaustionNotice:  reseller.ExhaustionNotice,
				ExpireNotice:      reseller.ExpireNotice,
				NoticeQuota:       reseller.NoticeQuota,
				NoticeQuotaExpiry: reseller.NoticeQuotaExpiry,
				IsReseller:        true,
			})
		}
	} else {
		logger.Error(ctx, err)
	}

	if _, err := redis.Set(ctx, consts.TASK_QUOTA_NOTICE_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sQuota) handleQuotaNotice(ctx context.Context, target quotaNoticeTarget) {

	if err := email.Verify(target.Email); err != nil {
		logger.Infof(ctx, "sQuota NoticeTask user: %d, error: %v", target.UserId, err)
		return
	}

	warningThresholds := common.QuotaWarningThresholds(target.WarningThresholds)
	expireThresholds := common.QuotaExpireWarningThresholds(target.ExpireThresholds)
	hasCustomWarning := len(target.WarningThresholds) > 0
	hasCustomExpire := len(target.ExpireThresholds) > 0
	now := gtime.TimestampMilli()
	quotaValid := target.QuotaExpiresAt == 0 || target.QuotaExpiresAt > now

	update := bson.M{}

	if target.NoticeQuota != 0 && target.Quota > target.NoticeQuota {
		target.WarningNotices = []int{}
		update["warning_notices"] = []int{}
	}

	if target.NoticeQuotaExpiry != 0 && target.NoticeQuotaExpiry != target.QuotaExpiresAt {
		target.ExpireNotices = []int{}
		update["expire_warning_notices"] = []int{}
	}

	if config.Cfg.Quota.Warning && !target.ExhaustionNotice && target.UsedQuota != 0 && quotaValid && (target.QuotaWarning || !hasCustomWarning) {
		if stages := unsentStages(quotaWarningStages(warningThresholds, target.Quota), target.WarningNotices); len(stages) > 0 {
			currentStage := slices.Max(stages)
			if err := s.sendQuotaNotice(ctx, target, consts.SCENE_QUOTA_WARNING, currentStage); err != nil {
				logger.Errorf(ctx, "sQuota NoticeTask user: %d, email: %s, error: %v", target.UserId, target.Email, err)
				return
			}
			target.WarningNotices = append(target.WarningNotices, stages...)
			update["warning_notices"] = target.WarningNotices
			update["notice_quota"] = target.Quota
		}
	}

	if config.Cfg.Quota.ExhaustedNotice && !target.ExhaustionNotice && target.UsedQuota != 0 && target.Quota <= 0 && quotaValid {
		if err := s.sendQuotaNotice(ctx, target, consts.SCENE_QUOTA_EXHAUSTION, 0); err != nil {
			logger.Errorf(ctx, "sQuota NoticeTask user: %d, email: %s, error: %v", target.UserId, target.Email, err)
			return
		}
		update["exhaustion_notice"] = true
	}

	if config.Cfg.Quota.ExpiredWarning && target.Quota > 0 && target.QuotaExpiresAt > 0 && (target.QuotaWarning || !hasCustomExpire) {
		days := int(math.Ceil(float64(target.QuotaExpiresAt-now) / float64(gtime.D.Milliseconds())))
		if days > 0 {
			if stages := unsentStages(quotaExpireWarningStages(expireThresholds, days), target.ExpireNotices); len(stages) > 0 {
				if err := s.sendQuotaNotice(ctx, target, consts.SCENE_QUOTA_EXPIRE_WARNING, slices.Max(stages)); err != nil {
					logger.Errorf(ctx, "sQuota NoticeTask user: %d, email: %s, error: %v", target.UserId, target.Email, err)
					return
				}
				target.ExpireNotices = append(target.ExpireNotices, stages...)
				update["expire_warning_notices"] = target.ExpireNotices
				update["notice_quota_expiry"] = target.QuotaExpiresAt
			}
		}
	}

	if config.Cfg.Quota.ExpiredNotice && !target.ExpireNotice && target.Quota > 0 && target.QuotaExpiresAt > 0 && now > target.QuotaExpiresAt {
		if err := s.sendQuotaNotice(ctx, target, consts.SCENE_QUOTA_EXPIRE, 0); err != nil {
			logger.Errorf(ctx, "sQuota NoticeTask user: %d, email: %s, error: %v", target.UserId, target.Email, err)
			return
		}
		update["expire_notice"] = true
	}

	if len(update) == 0 {
		return
	}

	if target.IsReseller {
		if err := dao.Reseller.UpdateById(ctx, target.Id, update); err != nil {
			logger.Error(ctx, err)
		}
		return
	}

	if err := dao.User.UpdateById(ctx, target.Id, update); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sQuota) sendQuotaNotice(ctx context.Context, target quotaNoticeTarget, scene string, stage int) error {

	var (
		title          string
		content        string
		noticeTemplate *model.NoticeTemplate
		siteConfig     *entity.SiteConfig
		dialer         = email.NewDefaultDialer()
		currencySymbol = "$"
		err            error
	)

	if target.IsReseller {
		account, err := dao.ResellerAccount.FindOne(ctx, bson.M{"user_id": target.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
		if err != nil {
			return err
		}
		if account == nil {
			logger.Infof(ctx, "sQuota NoticeTask reseller: %d, 因无可用账号, 不发送提醒邮件", target.UserId)
			return nil
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
	} else {
		account, err := dao.Account.FindOne(ctx, bson.M{"user_id": target.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
		if err != nil {
			return err
		}
		if account == nil {
			logger.Infof(ctx, "sQuota NoticeTask user: %d, 因无可用账号, 不发送提醒邮件", target.UserId)
			return nil
		}

		if target.Rid > 0 {
			isConfigEmail := false
			if siteConfig = service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain); siteConfig != nil && siteConfig.Rid == target.Rid && siteConfig.Host != "" {
				dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
				isConfigEmail = true
			}
			if !isConfigEmail {
				for _, item := range service.SiteConfig().GetSiteConfigsByRid(ctx, target.Rid) {
					if item != nil && item.Host != "" {
						siteConfig = item
						dialer = email.NewDialer(item.Host, item.Port, item.UserName, item.Password, item.FromName)
						isConfigEmail = true
						break
					}
				}
			}
			if !isConfigEmail {
				logger.Infof(ctx, "sQuota NoticeTask 因代理商: %d, 所有站点未配置邮箱, 不发送提醒邮件", target.Rid)
				return nil
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
	}

	if noticeTemplate, err = service.NoticeTemplate().GetNoticeTemplateByScene(ctx, scene, []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL}); err != nil {
		return err
	}

	if siteConfig != nil && siteConfig.CurrencySymbol != "" {
		currencySymbol = siteConfig.CurrencySymbol
	}

	var (
		user     *entity.User
		reseller *entity.Reseller
	)
	if target.IsReseller {
		reseller = &entity.Reseller{Name: target.Name, Quota: target.Quota, UsedQuota: target.UsedQuota, QuotaExpiresAt: target.QuotaExpiresAt}
	} else {
		user = &entity.User{Name: target.Name, Quota: target.Quota, UsedQuota: target.UsedQuota, QuotaExpiresAt: target.QuotaExpiresAt}
	}

	data := common.GetVariableData(ctx, user, reseller, siteConfig, noticeTemplate.Variables)
	data["name"] = target.Name
	if target.Quota < 0 {
		data["quota"] = fmt.Sprintf("-%s%f", currencySymbol, common.ConvQuotaUnitReverse(int(math.Abs(float64(target.Quota)))))
	} else {
		data["quota"] = fmt.Sprintf("%s%f", currencySymbol, common.ConvQuotaUnitReverse(target.Quota))
	}

	if scene == consts.SCENE_QUOTA_WARNING {
		data["warning_threshold"] = fmt.Sprintf("%s%d", currencySymbol, stage/consts.QUOTA_DEFAULT_UNIT)
	} else if scene == consts.SCENE_QUOTA_EXPIRE_WARNING || scene == consts.SCENE_QUOTA_EXPIRE {
		data["quota_expires_at"] = util.FormatDateTime(target.QuotaExpiresAt)
	}

	if title, content, err = util.RenderTemplate(noticeTemplate.Title, noticeTemplate.Content, data); err != nil {
		return err
	}

	if err = email.SendMailTask(ctx, email.NewMessage([]string{target.Email}, title, content), dialer); err != nil {
		return err
	}

	logger.Infof(ctx, "sQuota NoticeTask user: %d, email: %s, SendMailTask %s success", target.UserId, target.Email, title)
	return nil
}

func quotaWarningStages(thresholds []int, quota int) []int {
	stages := make([]int, 0)
	for _, threshold := range thresholds {
		if threshold > 0 && quota <= threshold {
			stages = append(stages, threshold)
		}
	}
	return stages
}

func quotaExpireWarningStages(thresholds []int, days int) []int {
	stages := make([]int, 0)
	for _, threshold := range thresholds {
		if threshold > 0 && days <= threshold {
			stages = append(stages, threshold)
		}
	}
	return stages
}

func unsentStages(stages, sent []int) []int {
	unsent := make([]int, 0)
	for _, stage := range stages {
		if !slices.Contains(sent, stage) {
			unsent = append(unsent, stage)
		}
	}
	return unsent
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
