package admin_reseller

import (
	"context"
	"fmt"
	"math"
	"regexp"
	"slices"
	"time"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/core"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/logic/common"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/crypto"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/email"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

type sAdminReseller struct{}

func init() {
	service.RegisterAdminReseller(New())
}

func New() service.IAdminReseller {
	return &sAdminReseller{}
}

// 新建代理商
func (s *sAdminReseller) Create(ctx context.Context, params model.ResellerCreateReq) (err error) {

	if dao.Reseller.IsAccountExist(ctx, params.Account) {
		return errors.New(params.Account + " 账号已存在")
	}

	if dao.Reseller.IsAccountExist(ctx, params.Email) {
		return errors.New(params.Email + " 邮箱已被其它账号使用")
	}

	if len(params.Groups) == 0 {
		if params.Groups, err = service.Group().PublicGroups(ctx); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	var (
		salt     = grand.Letters(8)
		id       = util.GenerateId()
		reseller = &do.Reseller{
			Id:             id,
			UserId:         core.IncrResellerId(ctx),
			Name:           params.Name,
			Email:          params.Email,
			Quota:          params.Quota,
			QuotaExpiresAt: util.ConvTimestampMilli(params.QuotaExpiresAt),
			Groups:         params.Groups,
			Remark:         params.Remark,
			Status:         1,
			Creator:        id,
		}
	)

	uid, err := dao.Reseller.Insert(ctx, reseller)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.Reseller.CreateAccount(ctx, &do.ResellerAccount{
		Uid:      uid,
		UserId:   reseller.UserId,
		Account:  params.Account,
		Password: crypto.EncryptPassword(params.Password + salt),
		Salt:     salt,
		Status:   1,
		Creator:  uid,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if params.Quota != 0 {

		// 交易记录
		if _, err = dao.DealRecord.Insert(ctx, &do.DealRecord{
			UserId: reseller.UserId,
			Quota:  params.Quota,
			Type:   params.QuotaType,
			Status: 1,
			Rid:    reseller.UserId,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_RESELLER_USAGE_KEY, reseller.UserId), consts.RESELLER_QUOTA_FIELD, int64(params.Quota)); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	newData, err := dao.Reseller.FindById(ctx, uid)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
		Action:  consts.ACTION_CREATE,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 发送欢迎邮件
	if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

		if noticeTemplate, err := service.NoticeTemplate().GetNoticeTemplateByScene(ctx, consts.SCENE_NOTICE_REGISTER, []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL}); err != nil {
			logger.Error(ctx, err)
		} else {

			dialer := email.NewDefaultDialer()

			siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, g.RequestFromCtx(ctx).GetHost())
			if siteConfig == nil {
				if siteConfig, err = dao.SiteConfig.FindOne(ctx, bson.M{"user_id": 1, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}}); err != nil {
					logger.Error(ctx, err)
				}
			}

			if siteConfig != nil && siteConfig.Host != "" {
				dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
			} else {
				logger.Infof(ctx, "sAdminReseller Create 因站点 %s 未配置邮箱, 默认使用系统配置邮箱", g.RequestFromCtx(ctx).GetHost())
			}

			data := common.GetVariableData(ctx, nil, newData, siteConfig, noticeTemplate.Variables)

			data["name"] = newData.Name
			data["account"] = params.Account
			data["quota"] = fmt.Sprintf("$%f", util.Round(float64(newData.Quota)/consts.QUOTA_USD_UNIT, 6))
			data["quota_expires_at"] = "无期限"
			if newData.QuotaExpiresAt > 0 {
				data["quota_expires_at"] = util.FormatDateTime(newData.QuotaExpiresAt)
			}

			if title, content, err := util.RenderTemplate(noticeTemplate.Title, noticeTemplate.Content, data); err != nil {
				logger.Error(ctx, err)
			} else {
				if err = email.SendMail(email.NewMessage([]string{newData.Email}, title, content), dialer); err != nil {
					logger.Errorf(ctx, "sAdminReseller Create reseller: %d, email: %s, SendMail %s error: %v", newData.UserId, newData.Email, title, err)
				}
			}
		}
	}, nil); err != nil {
		logger.Error(ctx, err)
	}

	return nil
}

// 更新代理商
func (s *sAdminReseller) Update(ctx context.Context, params model.ResellerUpdateReq) error {

	oldData, err := dao.Reseller.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	newData, err := dao.Reseller.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"name":                  params.Name,
		"email":                 params.Email,
		"quota_expires_at":      util.ConvTimestampMilli(params.QuotaExpiresAt),
		"groups":                params.Groups,
		"remark":                params.Remark,
		"status":                params.Status,
		"expire_warning_notice": false,
		"expire_notice":         false,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	account, err := dao.Reseller.FindAccountByUserId(ctx, newData.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if account.Account != params.Account {
		if err = dao.ResellerAccount.UpdateById(ctx, account.Id, bson.M{
			"account": params.Account,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if params.Password != "" {
		if err = dao.Reseller.ChangePasswordByUserId(ctx, account.UserId, params.Password); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if !slices.Equal(oldData.Groups, newData.Groups) {
		if err = s.Permissions(ctx, newData.UserId, oldData.Groups, newData.Groups); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

// 更改代理商额度过期时间
func (s *sAdminReseller) ChangeQuotaExpire(ctx context.Context, params model.ResellerChangeQuotaExpireReq) error {

	oldData, err := dao.Reseller.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	newData, err := dao.Reseller.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"quota_expires_at":      util.ConvTimestampMilli(params.QuotaExpiresAt),
		"expire_warning_notice": false,
		"expire_notice":         false,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改代理商状态
func (s *sAdminReseller) ChangeStatus(ctx context.Context, params model.ResellerChangeStatusReq) error {

	reseller, err := dao.Reseller.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.ResellerAccount.UpdateMany(ctx, bson.M{"user_id": reseller.UserId}, bson.M{
		"status": params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: reseller,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除代理商
func (s *sAdminReseller) Delete(ctx context.Context, params model.ResellerDeleteReq) error {

	reseller, err := dao.Reseller.FindOneAndDeleteById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.ResellerAccount.DeleteMany(ctx, bson.M{"user_id": reseller.UserId}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: reseller,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 删除用户数据
	if slices.Contains(params.Data, 1) {

		if users, err := dao.User.Find(ctx, bson.M{"rid": reseller.UserId}); err != nil {
			logger.Error(ctx, err)
		} else {

			if _, err = dao.User.DeleteMany(ctx, bson.M{"rid": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err = dao.Account.DeleteMany(ctx, bson.M{"rid": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			for _, user := range users {
				if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

					if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
						Action:  consts.ACTION_DELETE,
						OldData: user,
					}); err != nil {
						logger.Error(ctx, err)
					}

				}, nil); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	// 删除应用数据
	if slices.Contains(params.Data, 2) {
		if apps, err := dao.App.Find(ctx, bson.M{"rid": reseller.UserId}); err != nil {
			logger.Error(ctx, err)
		} else {
			for _, app := range apps {
				if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

					if err = service.App().Delete(ctx, app.Id); err != nil {
						logger.Error(ctx, err)
					}

				}, nil); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	// 删除交易记录
	if slices.Contains(params.Data, 3) {
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

			if _, err := dao.DealRecord.DeleteMany(ctx, bson.M{"user_id": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.DealRecord.DeleteMany(ctx, bson.M{"rid": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	// 删除账单明细
	if slices.Contains(params.Data, 4) {
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

			if _, err := dao.StatisticsUser.DeleteMany(ctx, bson.M{"user_id": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.StatisticsApp.DeleteMany(ctx, bson.M{"user_id": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.StatisticsAppKey.DeleteMany(ctx, bson.M{"user_id": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.StatisticsUser.DeleteMany(ctx, bson.M{"rid": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.StatisticsApp.DeleteMany(ctx, bson.M{"rid": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.StatisticsAppKey.DeleteMany(ctx, bson.M{"rid": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	// 删除日志数据
	if slices.Contains(params.Data, 5) {
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

			if _, err := dao.Chat.DeleteMany(ctx, bson.M{"rid": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.Image.DeleteMany(ctx, bson.M{"rid": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.Audio.DeleteMany(ctx, bson.M{"rid": reseller.UserId}); err != nil {
				logger.Error(ctx, err)
			}

		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	return nil
}

// 代理商详情
func (s *sAdminReseller) Detail(ctx context.Context, id string) (*model.Reseller, error) {

	reseller, err := dao.Reseller.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	account, err := dao.Reseller.FindAccountByUserId(ctx, reseller.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	groupNames, err := service.Group().GroupNames(ctx, reseller.Groups)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	users, err := dao.User.Find(ctx, bson.M{"rid": reseller.UserId})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	allocatedQuota := 0

	for _, user := range users {

		if user.Quota > 0 {
			allocatedQuota += user.Quota
		}

		allocatedQuota += user.UsedQuota
	}

	toBeAllocated := reseller.Quota + reseller.UsedQuota - allocatedQuota

	return &model.Reseller{
		Id:                     reseller.Id,
		UserId:                 reseller.UserId,
		Account:                account.Account,
		Name:                   reseller.Name,
		Phone:                  reseller.Phone,
		Email:                  reseller.Email,
		Quota:                  reseller.Quota,
		UsedQuota:              reseller.UsedQuota,
		AllocatedQuota:         allocatedQuota,
		ToBeAllocated:          toBeAllocated,
		QuotaExpiresAt:         util.FormatDateTime(reseller.QuotaExpiresAt),
		Groups:                 reseller.Groups,
		GroupNames:             groupNames,
		QuotaWarning:           reseller.QuotaWarning,
		WarningThreshold:       reseller.WarningThreshold / consts.QUOTA_USD_UNIT,
		ExpireWarningThreshold: reseller.ExpireWarningThreshold,
		WarningNotice:          reseller.WarningNotice,
		ExhaustionNotice:       reseller.ExhaustionNotice,
		ExpireWarningNotice:    reseller.ExpireWarningNotice,
		ExpireNotice:           reseller.ExpireNotice,
		Remark:                 reseller.Remark,
		Status:                 reseller.Status,
		LoginIP:                account.LoginIP,
		LoginTime:              util.FormatDateTime(account.LoginTime),
		LoginDomain:            account.LoginDomain,
		CreatedAt:              util.FormatDateTime(reseller.CreatedAt),
		UpdatedAt:              util.FormatDateTime(reseller.UpdatedAt),
	}, nil
}

// 代理商分页列表
func (s *sAdminReseller) Page(ctx context.Context, params model.ResellerPageReq) (*model.ResellerPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.Name != "" {
		filter["$or"] = bson.A{
			bson.M{"name": bson.M{
				"$regex": regexp.QuoteMeta(params.Name),
			}},
			bson.M{"email": bson.M{
				"$regex": regexp.QuoteMeta(params.Name),
			}},
		}
	}

	if params.Account != "" && params.UserId == 0 {
		account, err := dao.ResellerAccount.FindOne(ctx, bson.M{"account": params.Account})
		if err != nil {
			return nil, nil
		}
		filter["user_id"] = account.UserId
	}

	if params.Quota != 0 {
		filter["quota"] = bson.M{
			"$lt": params.Quota * consts.QUOTA_USD_UNIT,
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if len(params.QuotaExpiresAt) > 0 {
		gte := gtime.NewFromStrFormat(params.QuotaExpiresAt[0], time.DateOnly).StartOfDay().TimestampMilli()
		lte := gtime.NewFromStrLayout(params.QuotaExpiresAt[1], time.DateOnly).EndOfDay(true).TimestampMilli()
		filter["quota_expires_at"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.Reseller.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-user_id", "-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	accountMap := make(map[int]*entity.ResellerAccount)
	if len(results) > 0 {

		accounts, err := dao.ResellerAccount.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		accountMap = util.ToMap(accounts, func(t *entity.ResellerAccount) int {
			return t.UserId
		})
	}

	items := make([]*model.Reseller, 0)
	for _, result := range results {

		users, err := dao.User.Find(ctx, bson.M{"rid": result.UserId})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		allocatedQuota := 0

		for _, user := range users {

			if user.Quota > 0 {
				allocatedQuota += user.Quota
			}

			allocatedQuota += user.UsedQuota
		}

		toBeAllocated := result.Quota + result.UsedQuota - allocatedQuota

		items = append(items, &model.Reseller{
			Id:             result.Id,
			UserId:         result.UserId,
			Name:           result.Name,
			Email:          result.Email,
			Phone:          result.Phone,
			Quota:          result.Quota,
			UsedQuota:      result.UsedQuota,
			AllocatedQuota: allocatedQuota,
			ToBeAllocated:  toBeAllocated,
			QuotaExpiresAt: util.FormatDateTime(result.QuotaExpiresAt),
			Groups:         result.Groups,
			Account:        accountMap[result.UserId].Account,
			Remark:         result.Remark,
			Status:         result.Status,
			CreatedAt:      util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:      util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.ResellerPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 代理商列表
func (s *sAdminReseller) List(ctx context.Context, params model.ResellerListReq) ([]*model.Reseller, error) {

	filter := bson.M{}

	results, err := dao.Reseller.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Reseller, 0)
	for _, result := range results {
		items = append(items, &model.Reseller{
			Id:        result.Id,
			UserId:    result.UserId,
			Name:      result.Name,
			Email:     result.Email,
			Phone:     result.Phone,
			Quota:     result.Quota,
			UsedQuota: result.UsedQuota,
			Groups:    result.Groups,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt: util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return items, nil
}

// 代理商充值
func (s *sAdminReseller) Recharge(ctx context.Context, params model.ResellerRechargeReq) error {

	oldData, err := dao.Reseller.FindOne(ctx, bson.M{"user_id": params.UserId})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if params.QuotaType == 2 {
		params.Quota = -params.Quota
	}

	newData, err := dao.Reseller.FindOneAndUpdate(ctx, bson.M{"user_id": params.UserId}, bson.M{
		"$inc": bson.M{
			"quota": params.Quota,
		},
		"quota_expires_at":      util.ConvTimestampMilli(params.QuotaExpiresAt),
		"warning_notice":        false,
		"exhaustion_notice":     false,
		"expire_warning_notice": false,
		"expire_notice":         false,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_RESELLER_USAGE_KEY, params.UserId), consts.RESELLER_QUOTA_FIELD, int64(params.Quota)); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 交易记录
	if _, err = dao.DealRecord.Insert(ctx, &do.DealRecord{
		UserId: params.UserId,
		Quota:  params.Quota,
		Type:   params.QuotaType,
		Status: 1,
		Rid:    params.UserId,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if params.IsSendNotice {

		if err = email.Verify(newData.Email); err != nil {
			logger.Infof(ctx, "sAdminReseller Recharge reseller: %d, error: %v", newData.UserId, err)
			return nil
		} else {
			if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

				var (
					dialer     = email.NewDefaultDialer()
					siteConfig *entity.SiteConfig
				)

				account, err := dao.ResellerAccount.FindOne(ctx, bson.M{"user_id": newData.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
				if err != nil {
					logger.Error(ctx, err)
					return
				}

				if account == nil {
					logger.Infof(ctx, "sAdminReseller Recharge reseller: %d, 因无可用账号, 不发送充值通知邮件", newData.UserId)
					return
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
					logger.Infof(ctx, "sAdminReseller Recharge 因站点 %s 未配置邮箱, 默认使用系统配置邮箱", account.LoginDomain)
				}

				noticeTemplate, err := service.NoticeTemplate().GetNoticeTemplateByScene(ctx, consts.SCENE_QUOTA_RECHARGE, []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL})
				if err != nil {
					logger.Error(ctx, err)
					return
				}

				data := common.GetVariableData(ctx, nil, newData, siteConfig, noticeTemplate.Variables)

				data["quota_type"] = consts.QUOTA_TYPE[params.QuotaType]
				data["name"] = newData.Name

				if params.Quota < 0 {
					data["recharge_quota"] = fmt.Sprintf("-$%f", util.Round(math.Abs(float64(params.Quota))/consts.QUOTA_USD_UNIT, 6))
				} else {
					data["recharge_quota"] = fmt.Sprintf("$%f", util.Round(float64(params.Quota)/consts.QUOTA_USD_UNIT, 6))
				}

				if newData.Quota < 0 {
					data["quota"] = fmt.Sprintf("-$%f", util.Round(math.Abs(float64(newData.Quota))/consts.QUOTA_USD_UNIT, 6))
				} else {
					data["quota"] = fmt.Sprintf("$%f", util.Round(float64(newData.Quota)/consts.QUOTA_USD_UNIT, 6))
				}

				data["quota_expires_at"] = "无期限"
				if newData.QuotaExpiresAt > 0 {
					data["quota_expires_at"] = util.FormatDateTime(newData.QuotaExpiresAt)
				}

				title, content, err := util.RenderTemplate(noticeTemplate.Title, noticeTemplate.Content, data)
				if err != nil {
					logger.Error(ctx, err)
					return
				}

				if err = email.SendMailTask(ctx, email.NewMessage([]string{newData.Email}, title, content), dialer); err != nil {
					logger.Errorf(ctx, "sAdminReseller Recharge reseller: %d, email: %s, SendMailTask %s error: %v", newData.UserId, newData.Email, title, err)
					return
				}

				logger.Infof(ctx, "sAdminReseller Recharge reseller: %d, email: %s, SendMailTask %s success", newData.UserId, newData.Email, title)

			}, nil); err != nil {
				logger.Error(ctx, err)
			}
		}
	}

	return nil
}

// 代理商权限
func (s *sAdminReseller) Permissions(ctx context.Context, userId int, oldGroups, newGroups []string) error {

	addGroups := make([]string, 0)
	delGroups := make([]string, 0)

	for _, g := range newGroups {
		if !slices.Contains(oldGroups, g) {
			addGroups = append(addGroups, g)
		}
	}

	for _, g := range oldGroups {
		if !slices.Contains(newGroups, g) {
			delGroups = append(delGroups, g)
		}
	}

	users, err := dao.User.Find(ctx, bson.M{"rid": userId})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, user := range users {

		if len(addGroups) > 0 {
			groupSet := gset.NewStrSetFrom(user.Groups)
			groupSet.Add(addGroups...)
			user.Groups = groupSet.Slice()
		}

		groups := make([]string, 0)

		if len(delGroups) > 0 {
			for _, g := range user.Groups {
				if !slices.Contains(delGroups, g) {
					groups = append(groups, g)
				}
			}
		} else {
			groups = user.Groups
		}

		if err = service.AdminUser().Permissions(ctx, user.UserId, groups); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

// 代理商批量操作
func (s *sAdminReseller) BatchOperate(ctx context.Context, params model.ResellerBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_RECHARGE:

		resellers, err := dao.Reseller.FindByIds(ctx, params.Ids)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, reseller := range resellers {

			quotaExpiresAt := params.QuotaExpiresAt
			if quotaExpiresAt == "" {
				quotaExpiresAt = util.FormatDateTime(reseller.QuotaExpiresAt)
			}

			if err := s.Recharge(ctx, model.ResellerRechargeReq{
				UserId:         reseller.UserId,
				Quota:          gconv.Int(params.Value),
				QuotaType:      params.QuotaType,
				QuotaExpiresAt: quotaExpiresAt,
				IsSendNotice:   params.IsSendNotice,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	case consts.ACTION_STATUS:
		for _, id := range params.Ids {
			if err := s.ChangeStatus(ctx, model.ResellerChangeStatusReq{
				Id:     id,
				Status: gconv.Int(params.Value),
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	case consts.ACTION_DELETE:
		for _, id := range params.Ids {
			if err := s.Delete(ctx, model.ResellerDeleteReq{Id: id, Data: params.Data}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}
