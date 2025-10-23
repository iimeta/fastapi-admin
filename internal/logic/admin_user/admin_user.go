package admin_user

import (
	"context"
	"fmt"
	"math"
	"regexp"
	"slices"
	"time"

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

type sAdminUser struct{}

func init() {
	service.RegisterAdminUser(New())
}

func New() service.IAdminUser {
	return &sAdminUser{}
}

// 新建用户
func (s *sAdminUser) Create(ctx context.Context, params model.UserCreateReq) (err error) {

	if dao.User.IsAccountExist(ctx, params.Account) {
		return errors.New(params.Account + " 账号已存在")
	}

	if dao.User.IsAccountExist(ctx, params.Email) {
		return errors.New(params.Email + " 邮箱已被其它账号使用")
	}

	if len(params.Groups) == 0 {
		if service.Session().IsResellerRole(ctx) {
			params.Groups = service.Session().GetReseller(ctx).Groups
		} else {
			if params.Groups, err = service.Group().PublicGroups(ctx); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	var (
		salt = grand.Letters(8)
		id   = util.GenerateId()
		user = &do.User{
			Id:             id,
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

	if user.Quota > 0 && service.Session().IsResellerRole(ctx) {

		users, err := dao.User.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		totalQuota := user.Quota
		for _, user := range users {
			totalQuota += user.Quota
		}

		reseller, err := dao.Reseller.FindResellerByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if totalQuota > reseller.Quota {
			return errors.New("所有用户累计额度已超过账户额度")
		}
	}

	user.UserId = core.IncrUserId(ctx)

	uid, err := dao.User.Insert(ctx, user)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.User.CreateAccount(ctx, &do.Account{
		Uid:      uid,
		UserId:   user.UserId,
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
			UserId: user.UserId,
			Quota:  params.Quota,
			Type:   params.QuotaType,
			Status: 1,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_USER_USAGE_KEY, user.UserId), consts.USER_QUOTA_FIELD, int64(params.Quota)); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	newData, err := dao.User.FindById(ctx, uid)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_CREATE,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 创建默认应用
	if _, err = service.App().Create(ctx, model.AppCreateReq{
		UserId:      user.UserId,
		Name:        "默认应用",
		IsCreateKey: true,
		Status:      1,
	}); err != nil {
		logger.Error(ctx, err)
	}

	// 发送欢迎邮件
	if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

		if noticeTemplate, err := service.NoticeTemplate().GetNoticeTemplateByScene(ctx, consts.SCENE_NOTICE_REGISTER, []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL}); err != nil {
			logger.Error(ctx, err)
		} else {

			var (
				dialer     = email.NewDefaultDialer()
				siteConfig *entity.SiteConfig
			)

			if newData.Rid > 0 {

				isConfigEmail := false

				if siteConfig = service.SiteConfig().GetSiteConfigByDomain(ctx, g.RequestFromCtx(ctx).GetHost()); siteConfig != nil && siteConfig.Rid == newData.Rid && siteConfig.Host != "" {
					dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
					isConfigEmail = true
				}

				if !isConfigEmail {
					siteConfigs := service.SiteConfig().GetSiteConfigsByRid(ctx, newData.Rid)
					for _, siteConfig = range siteConfigs {
						if siteConfig != nil && siteConfig.Host != "" {
							dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
							isConfigEmail = true
							break
						}
					}
				}

				if !isConfigEmail {
					logger.Infof(ctx, "sAdminUser Create 因代理商: %d, 所有站点未配置邮箱, 不发送欢迎邮件", newData.Rid)
					return
				}

			} else {

				siteConfig = service.SiteConfig().GetSiteConfigByDomain(ctx, g.RequestFromCtx(ctx).GetHost())
				if siteConfig == nil {
					if siteConfig, err = dao.SiteConfig.FindOne(ctx, bson.M{"user_id": 1, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}}); err != nil {
						logger.Error(ctx, err)
					}
				}

				if siteConfig != nil && siteConfig.Host != "" {
					dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
				} else {
					logger.Infof(ctx, "sAdminUser Create 因站点 %s 未配置邮箱, 默认使用系统配置邮箱", g.RequestFromCtx(ctx).GetHost())
				}
			}

			data := common.GetVariableData(ctx, newData, nil, siteConfig, noticeTemplate.Variables)

			data["name"] = newData.Name
			data["account"] = params.Account
			data["quota"] = fmt.Sprintf("$%f", common.ConvQuota(newData.Quota))
			data["quota_expires_at"] = "无期限"
			if newData.QuotaExpiresAt > 0 {
				data["quota_expires_at"] = util.FormatDateTime(newData.QuotaExpiresAt)
			}

			if title, content, err := util.RenderTemplate(noticeTemplate.Title, noticeTemplate.Content, data); err != nil {
				logger.Error(ctx, err)
			} else {
				if err = email.SendMail(email.NewMessage([]string{newData.Email}, title, content), dialer); err != nil {
					logger.Errorf(ctx, "sAdminUser Create user: %d, email: %s, SendMail %s error: %v", newData.UserId, newData.Email, title, err)
				}
			}
		}
	}, nil); err != nil {
		logger.Error(ctx, err)
	}

	return nil
}

// 更新用户
func (s *sAdminUser) Update(ctx context.Context, params model.UserUpdateReq) error {

	oldData, err := dao.User.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if service.Session().IsResellerRole(ctx) && oldData.Rid != service.Session().GetRid(ctx) {
		return errors.New("Unauthorized")
	}

	newData, err := dao.User.FindOneAndUpdateById(ctx, params.Id, bson.M{
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

	account, err := dao.User.FindAccountByUserId(ctx, newData.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if account.Account != params.Account {
		if err = dao.Account.UpdateById(ctx, account.Id, bson.M{
			"account": params.Account,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if params.Password != "" {
		if err = dao.User.ChangePasswordByUserId(ctx, account.UserId, params.Password); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改用户额度过期时间
func (s *sAdminUser) ChangeQuotaExpire(ctx context.Context, params model.UserChangeQuotaExpireReq) error {

	oldData, err := dao.User.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if service.Session().IsResellerRole(ctx) && oldData.Rid != service.Session().GetRid(ctx) {
		return errors.New("Unauthorized")
	}

	newData, err := dao.User.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"quota_expires_at":      util.ConvTimestampMilli(params.QuotaExpiresAt),
		"expire_warning_notice": false,
		"expire_notice":         false,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改用户状态
func (s *sAdminUser) ChangeStatus(ctx context.Context, params model.UserChangeStatusReq) error {

	user, err := dao.User.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.Account.UpdateMany(ctx, bson.M{"user_id": user.UserId}, bson.M{
		"status": params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: user,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除用户
func (s *sAdminUser) Delete(ctx context.Context, params model.UserDeleteReq) error {

	if service.Session().IsResellerRole(ctx) {

		oldData, err := dao.User.FindById(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if oldData.Rid != service.Session().GetRid(ctx) {
			return errors.New("Unauthorized")
		}
	}

	user, err := dao.User.FindOneAndDeleteById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.Account.DeleteMany(ctx, bson.M{"user_id": user.UserId}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: user,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 删除应用数据
	if slices.Contains(params.Data, 2) {
		if apps, err := service.App().List(ctx, model.AppListReq{UserId: user.UserId}); err != nil {
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

			if _, err := dao.DealRecord.DeleteMany(ctx, bson.M{"user_id": user.UserId}); err != nil {
				logger.Error(ctx, err)
			}

		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	// 删除账单明细
	if slices.Contains(params.Data, 4) {
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

			if _, err := dao.StatisticsUser.DeleteMany(ctx, bson.M{"user_id": user.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.StatisticsApp.DeleteMany(ctx, bson.M{"user_id": user.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.StatisticsAppKey.DeleteMany(ctx, bson.M{"user_id": user.UserId}); err != nil {
				logger.Error(ctx, err)
			}

		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	// 删除日志数据
	if slices.Contains(params.Data, 5) {
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

			if _, err := dao.Chat.DeleteMany(ctx, bson.M{"user_id": user.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.Image.DeleteMany(ctx, bson.M{"user_id": user.UserId}); err != nil {
				logger.Error(ctx, err)
			}

			if _, err := dao.Audio.DeleteMany(ctx, bson.M{"user_id": user.UserId}); err != nil {
				logger.Error(ctx, err)
			}

		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	return nil
}

// 用户详情
func (s *sAdminUser) Detail(ctx context.Context, id string) (*model.User, error) {

	user, err := dao.User.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if service.Session().IsResellerRole(ctx) && user.Rid != service.Session().GetRid(ctx) {
		return nil, errors.New("Unauthorized")
	}

	account, err := dao.User.FindAccountByUserId(ctx, user.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	groupNames, err := service.Group().GroupNames(ctx, user.Groups)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.User{
		Id:                     user.Id,
		UserId:                 user.UserId,
		Account:                account.Account,
		Name:                   user.Name,
		Phone:                  user.Phone,
		Email:                  user.Email,
		Quota:                  user.Quota,
		UsedQuota:              user.UsedQuota,
		QuotaExpiresAt:         util.FormatDateTime(user.QuotaExpiresAt),
		Groups:                 user.Groups,
		GroupNames:             groupNames,
		QuotaWarning:           user.QuotaWarning,
		WarningThreshold:       user.WarningThreshold / consts.QUOTA_DEFAULT_UNIT,
		ExpireWarningThreshold: user.ExpireWarningThreshold,
		WarningNotice:          user.WarningNotice,
		ExhaustionNotice:       user.ExhaustionNotice,
		ExpireWarningNotice:    user.ExpireWarningNotice,
		ExpireNotice:           user.ExpireNotice,
		Remark:                 user.Remark,
		Status:                 user.Status,
		Rid:                    user.Rid,
		LoginIP:                account.LoginIP,
		LoginTime:              util.FormatDateTime(account.LoginTime),
		LoginDomain:            account.LoginDomain,
		CreatedAt:              util.FormatDateTime(user.CreatedAt),
		UpdatedAt:              util.FormatDateTime(user.UpdatedAt),
	}, nil
}

// 用户分页列表
func (s *sAdminUser) Page(ctx context.Context, params model.UserPageReq) (*model.UserPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

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
		account, err := dao.Account.FindOne(ctx, bson.M{"account": params.Account})
		if err != nil {
			return nil, nil
		}
		filter["user_id"] = account.UserId
	}

	if params.Quota != 0 {
		filter["quota"] = bson.M{
			"$lt": params.Quota * consts.QUOTA_DEFAULT_UNIT,
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

	results, err := dao.User.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-created_at", "-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	accountMap := make(map[int]*entity.Account)
	if len(results) > 0 {

		accounts, err := dao.Account.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		accountMap = util.ToMap(accounts, func(t *entity.Account) int {
			return t.UserId
		})
	}

	items := make([]*model.User, 0)
	for _, result := range results {

		items = append(items, &model.User{
			Id:             result.Id,
			UserId:         result.UserId,
			Name:           result.Name,
			Email:          result.Email,
			Phone:          result.Phone,
			Quota:          result.Quota,
			UsedQuota:      result.UsedQuota,
			QuotaExpiresAt: util.FormatDateTime(result.QuotaExpiresAt),
			Groups:         result.Groups,
			Account:        accountMap[result.UserId].Account,
			Remark:         result.Remark,
			Status:         result.Status,
			Rid:            result.Rid,
			CreatedAt:      util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:      util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.UserPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 用户列表
func (s *sAdminUser) List(ctx context.Context, params model.UserListReq) ([]*model.User, error) {

	filter := bson.M{}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	results, err := dao.User.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"status", "-created_at", "-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.User, 0)
	for _, result := range results {
		items = append(items, &model.User{
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

// 用户充值
func (s *sAdminUser) Recharge(ctx context.Context, params model.UserRechargeReq) error {

	oldData, err := dao.User.FindOne(ctx, bson.M{"user_id": params.UserId})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if service.Session().IsResellerRole(ctx) && oldData.Rid != service.Session().GetRid(ctx) {
		return errors.New("Unauthorized")
	}

	if params.QuotaType == 2 {
		params.Quota = -params.Quota
	}

	if params.Quota > 0 && oldData.Rid != 0 {

		users, err := dao.User.Find(ctx, bson.M{"rid": oldData.Rid})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		totalQuota := params.Quota
		for _, user := range users {
			if user.UserId == oldData.UserId {
				totalQuota += user.Quota
			} else if user.Quota > 0 {
				totalQuota += user.Quota
			}
		}

		reseller, err := dao.Reseller.FindResellerByUserId(ctx, oldData.Rid)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if oldData.Quota+params.Quota > 0 && totalQuota > reseller.Quota {
			return errors.New("所有用户累计额度已超过账户额度")
		}
	}

	newData, err := dao.User.FindOneAndUpdate(ctx, bson.M{"user_id": params.UserId}, bson.M{
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

	if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_USER_USAGE_KEY, params.UserId), consts.USER_QUOTA_FIELD, int64(params.Quota)); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 交易记录
	if _, err = dao.DealRecord.Insert(ctx, &do.DealRecord{
		UserId: params.UserId,
		Quota:  params.Quota,
		Type:   params.QuotaType,
		Status: 1,
		Rid:    newData.Rid,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if params.IsSendNotice {

		if err = email.Verify(newData.Email); err != nil {
			logger.Infof(ctx, "sAdminUser Recharge user: %d, error: %v", newData.UserId, err)
			return nil
		} else {
			if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

				var (
					dialer     = email.NewDefaultDialer()
					siteConfig *entity.SiteConfig
				)

				account, err := dao.Account.FindOne(ctx, bson.M{"user_id": newData.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
				if err != nil {
					logger.Error(ctx, err)
					return
				}

				if account == nil {
					logger.Infof(ctx, "sAdminUser Recharge user: %d, 因无可用账号, 不发送充值通知邮件", newData.UserId)
					return
				}

				if newData.Rid > 0 {

					isConfigEmail := false

					if siteConfig = service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain); siteConfig != nil && siteConfig.Rid == newData.Rid && siteConfig.Host != "" {
						dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
						isConfigEmail = true
					}

					if !isConfigEmail {
						siteConfigs := service.SiteConfig().GetSiteConfigsByRid(ctx, newData.Rid)
						for _, siteConfig = range siteConfigs {
							if siteConfig != nil && siteConfig.Host != "" {
								dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
								isConfigEmail = true
								break
							}
						}
					}

					if !isConfigEmail {
						logger.Infof(ctx, "sAdminUser Recharge 因代理商: %d, 所有站点未配置邮箱, 不发送充值通知邮件", newData.Rid)
						return
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
						logger.Infof(ctx, "sAdminUser Recharge 因站点 %s 未配置邮箱, 默认使用系统配置邮箱", account.LoginDomain)
					}
				}

				noticeTemplate, err := service.NoticeTemplate().GetNoticeTemplateByScene(ctx, consts.SCENE_QUOTA_RECHARGE, []string{consts.NOTICE_CHANNEL_WEB, consts.NOTICE_CHANNEL_EMAIL})
				if err != nil {
					logger.Error(ctx, err)
					return
				}

				data := common.GetVariableData(ctx, newData, nil, siteConfig, noticeTemplate.Variables)

				data["quota_type"] = consts.QUOTA_TYPE[params.QuotaType]
				data["name"] = newData.Name

				if params.Quota < 0 {
					data["recharge_quota"] = fmt.Sprintf("-$%f", common.ConvQuota(int(math.Abs(float64(params.Quota)))))
				} else {
					data["recharge_quota"] = fmt.Sprintf("$%f", common.ConvQuota(params.Quota))
				}

				if newData.Quota < 0 {
					data["quota"] = fmt.Sprintf("-$%f", common.ConvQuota(int(math.Abs(float64(newData.Quota)))))
				} else {
					data["quota"] = fmt.Sprintf("$%f", common.ConvQuota(newData.Quota))
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
					logger.Errorf(ctx, "sAdminUser Recharge user: %d, email: %s, SendMailTask %s error: %v", newData.UserId, newData.Email, title, err)
					return
				}

				logger.Infof(ctx, "sAdminUser Recharge user: %d, email: %s, SendMailTask %s success", newData.UserId, newData.Email, title)

			}, nil); err != nil {
				logger.Error(ctx, err)
			}
		}
	}

	return nil
}

// 用户权限
func (s *sAdminUser) Permissions(ctx context.Context, userId int, groups []string) error {

	oldData, err := dao.User.FindOne(ctx, bson.M{"user_id": userId})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if service.Session().IsResellerRole(ctx) && oldData.Rid != service.Session().GetRid(ctx) {
		return errors.New("Unauthorized")
	}

	newData, err := dao.User.FindOneAndUpdate(ctx, bson.M{"user_id": userId}, bson.M{
		"groups": groups,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 用户批量操作
func (s *sAdminUser) BatchOperate(ctx context.Context, params model.UserBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_RECHARGE:

		users, err := dao.User.FindByIds(ctx, params.Ids)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, user := range users {

			quotaExpiresAt := params.QuotaExpiresAt
			if quotaExpiresAt == "" {
				quotaExpiresAt = util.FormatDateTime(user.QuotaExpiresAt)
			}

			if err := s.Recharge(ctx, model.UserRechargeReq{
				UserId:         user.UserId,
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
			if err := s.ChangeStatus(ctx, model.UserChangeStatusReq{
				Id:     id,
				Status: gconv.Int(params.Value),
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	case consts.ACTION_DELETE:
		for _, id := range params.Ids {
			if err := s.Delete(ctx, model.UserDeleteReq{Id: id, Data: params.Data}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}
