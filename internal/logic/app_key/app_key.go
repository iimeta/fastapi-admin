package app_key

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/logic/common"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

type sAppKey struct {
	checkRedsync *redsync.Redsync
}

func init() {
	service.RegisterAppKey(New())
}

func New() service.IAppKey {
	return &sAppKey{
		checkRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}

// 新建应用密钥
func (s *sAppKey) Create(ctx context.Context, params model.AppKeyCreateReq) (string, error) {

	userId := service.Session().GetUserId(ctx)

	if service.Session().IsResellerRole(ctx) || service.Session().IsAdminRole(ctx) {

		if params.UserId == 0 {
			app, err := dao.App.FindByAppId(ctx, params.AppId)
			if err != nil {
				logger.Error(ctx, err)
				return "", err
			}
			params.UserId = app.UserId
		}

		userId = params.UserId
	}

	key := util.NewKey(config.Cfg.Core.SecretKeyPrefix, 51, gconv.String(userId), gconv.String(params.AppId))

	u, a, err := common.ParseSecretKey(ctx, key)
	if err != nil {
		logger.Error(ctx, err)
		return "", errors.New("创建密钥异常, 请重试...")
	}

	if u != userId || a != params.AppId {
		return "", errors.New("创建密钥异常, 请重试...")
	}

	return key, nil
}

// 应用密钥配置
func (s *sAppKey) Config(ctx context.Context, params model.AppKeyConfigReq) (k string, err error) {

	var (
		keyInfo *entity.AppKey
		oldData *entity.AppKey
		action  = consts.ACTION_CREATE
		key     = &do.AppKey{
			UserId:              service.Session().GetUserId(ctx),
			AppId:               params.AppId,
			Key:                 params.Key,
			BillingMethods:      params.BillingMethods,
			Models:              params.Models,
			IsLimitQuota:        params.IsLimitQuota,
			Quota:               common.ConvQuotaUnit(params.Quota),
			QuotaExpiresRule:    params.QuotaExpiresRule,
			QuotaExpiresAt:      util.ConvTimestampMilli(params.QuotaExpiresAt),
			QuotaExpiresMinutes: params.QuotaExpiresMinutes,
			IsBindGroup:         params.IsBindGroup,
			Group:               params.Group,
			IpWhitelist:         gstr.Split(gstr.Trim(params.IpWhitelist), "\n"),
			IpBlacklist:         gstr.Split(gstr.Trim(params.IpBlacklist), "\n"),
			Remark:              params.Remark,
			Status:              params.Status,
		}
	)

	if service.Session().IsResellerRole(ctx) || service.Session().IsAdminRole(ctx) {

		if params.UserId == 0 && params.Id == "" {
			app, err := dao.App.FindByAppId(ctx, params.AppId)
			if err != nil {
				logger.Error(ctx, err)
				return "", err
			}
			params.UserId = app.UserId
			key.Rid = app.Rid
		}

		key.UserId = params.UserId
	}

	if params.Id != "" {

		action = consts.ACTION_UPDATE
		if oldData, err = dao.AppKey.FindById(ctx, params.Id); err != nil {
			logger.Error(ctx, err)
			return "", err
		}

		if service.Session().IsResellerRole(ctx) && oldData.Rid != service.Session().GetRid(ctx) {
			return "", errors.New("Unauthorized")
		}

		if service.Session().IsUserRole(ctx) && oldData.UserId != service.Session().GetUserId(ctx) {
			return "", errors.New("Unauthorized")
		}

		key.AppId = 0
		key.Key = ""
		if keyInfo, err = dao.AppKey.FindOneAndUpdateById(ctx, params.Id, key); err != nil {
			logger.Error(ctx, err)
			return k, err
		}

	} else {

		userId, appId, err := common.ParseSecretKey(ctx, key.Key)
		if err != nil {
			logger.Error(ctx, err)
			return "", err
		}

		if userId != key.UserId || appId != key.AppId {
			return "", errors.New("Unauthorized")
		}

		id, err := dao.AppKey.Insert(ctx, key)
		if err != nil {
			logger.Error(ctx, err)
			return k, err
		}

		keyInfo = &entity.AppKey{
			Id:                  id,
			UserId:              key.UserId,
			AppId:               key.AppId,
			Key:                 key.Key,
			BillingMethods:      key.BillingMethods,
			Models:              key.Models,
			IsLimitQuota:        key.IsLimitQuota,
			Quota:               key.Quota,
			UsedQuota:           key.UsedQuota,
			QuotaExpiresRule:    key.QuotaExpiresRule,
			QuotaExpiresAt:      key.QuotaExpiresAt,
			QuotaExpiresMinutes: key.QuotaExpiresMinutes,
			IsBindGroup:         key.IsBindGroup,
			Group:               key.Group,
			IpWhitelist:         key.IpWhitelist,
			IpBlacklist:         key.IpBlacklist,
			Remark:              key.Remark,
			Status:              key.Status,
			Rid:                 key.Rid,
		}
	}

	app, err := dao.App.FindByAppId(ctx, keyInfo.AppId)
	if err != nil {
		logger.Error(ctx, err)
		return k, err
	}

	fields := g.Map{
		fmt.Sprintf(consts.KEY_QUOTA_FIELD, keyInfo.AppId, keyInfo.Key):          key.Quota,
		fmt.Sprintf(consts.KEY_IS_LIMIT_QUOTA_FIELD, keyInfo.AppId, keyInfo.Key): key.IsLimitQuota,
	}

	if _, err = redis.HSet(ctx, fmt.Sprintf(consts.API_USER_USAGE_KEY, app.UserId), fields); err != nil {
		logger.Error(ctx, err)
		return k, err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP_KEY, model.PubMessage{
		Action:  action,
		OldData: oldData,
		NewData: keyInfo,
	}); err != nil {
		logger.Error(ctx, err)
		return k, err
	}

	return keyInfo.Key, err
}

// 更改应用密钥状态
func (s *sAppKey) ChangeStatus(ctx context.Context, params model.AppKeyChangeStatusReq) error {

	if service.Session().IsResellerRole(ctx) {

		key, err := dao.AppKey.FindById(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if key.Rid != service.Session().GetRid(ctx) {
			return errors.New("Unauthorized")
		}
	}

	if service.Session().IsUserRole(ctx) {

		key, err := dao.AppKey.FindById(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if key.UserId != service.Session().GetUserId(ctx) {
			return errors.New("Unauthorized")
		}
	}

	key, err := dao.AppKey.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP_KEY, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: key,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除应用密钥
func (s *sAppKey) Delete(ctx context.Context, id string) error {

	if service.Session().IsResellerRole(ctx) {

		key, err := dao.AppKey.FindById(ctx, id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if key.Rid != service.Session().GetRid(ctx) {
			return errors.New("Unauthorized")
		}
	}

	if service.Session().IsUserRole(ctx) {

		key, err := dao.AppKey.FindById(ctx, id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if key.UserId != service.Session().GetUserId(ctx) {
			return errors.New("Unauthorized")
		}
	}

	key, err := dao.AppKey.FindOneAndDeleteById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP_KEY, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: key,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 应用密钥详情
func (s *sAppKey) Detail(ctx context.Context, id string) (*model.AppKey, error) {

	key, err := dao.AppKey.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if service.Session().IsResellerRole(ctx) && key.Rid != service.Session().GetRid(ctx) {
		return nil, errors.New("Unauthorized")
	}

	if service.Session().IsUserRole(ctx) && key.UserId != service.Session().GetUserId(ctx) {
		return nil, errors.New("Unauthorized")
	}

	modelNames, err := service.Model().ModelNames(ctx, key.Models)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	groupName := ""
	if key.IsBindGroup && key.Group != "" {
		group, err := dao.Group.FindById(ctx, key.Group)
		if err != nil {
			logger.Error(ctx, err)
		} else {
			groupName = group.Name
		}
	}

	return &model.AppKey{
		Id:                  key.Id,
		UserId:              key.UserId,
		AppId:               key.AppId,
		Key:                 key.Key,
		BillingMethods:      key.BillingMethods,
		Models:              key.Models,
		ModelNames:          modelNames,
		IsLimitQuota:        key.IsLimitQuota,
		Quota:               common.ConvQuotaUnitReverse(key.Quota),
		UsedQuota:           common.ConvQuotaUnitReverse(key.UsedQuota),
		QuotaExpiresRule:    key.QuotaExpiresRule,
		QuotaExpiresAt:      util.FormatDateTime(key.QuotaExpiresAt),
		QuotaExpiresMinutes: key.QuotaExpiresMinutes,
		IsBindGroup:         key.IsBindGroup,
		Group:               key.Group,
		GroupName:           groupName,
		IpWhitelist:         key.IpWhitelist,
		IpBlacklist:         key.IpBlacklist,
		Remark:              key.Remark,
		Status:              key.Status,
		Creator:             key.Creator,
		Updater:             key.Updater,
		CreatedAt:           util.FormatDateTime(key.CreatedAt),
		UpdatedAt:           util.FormatDateTime(key.UpdatedAt),
	}, nil
}

// 应用密钥分页列表
func (s *sAppKey) Page(ctx context.Context, params model.AppKeyPageReq) (*model.AppKeyPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
	} else if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	if params.Key != "" {
		filter["key"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Key),
		}
	}

	if len(params.Models) > 0 {
		filter["models"] = bson.M{
			"$in": params.Models,
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if params.Quota != 0 {
		filter["is_limit_quota"] = true
		filter["quota"] = bson.M{
			"$lte": params.Quota * consts.QUOTA_DEFAULT_UNIT,
		}
	}

	if len(params.QuotaExpiresAt) > 0 {
		gte := gtime.NewFromStrFormat(params.QuotaExpiresAt[0], time.DateOnly).StartOfDay().TimestampMilli()
		lte := gtime.NewFromStrLayout(params.QuotaExpiresAt[1], time.DateOnly).EndOfDay(true).TimestampMilli()
		filter["quota_expires_at"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	if params.Remark != "" {
		filter["remark"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Remark),
		}
	}

	results, err := dao.AppKey.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-updated_at", "key"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.AppKey, 0)
	for _, result := range results {
		items = append(items, &model.AppKey{
			Id:                  result.Id,
			UserId:              result.UserId,
			AppId:               result.AppId,
			Key:                 util.Desensitize(result.Key),
			BillingMethods:      result.BillingMethods,
			Models:              result.Models,
			IsLimitQuota:        result.IsLimitQuota,
			Quota:               common.ConvQuotaUnitReverse(result.Quota),
			UsedQuota:           common.ConvQuotaUnitReverse(result.UsedQuota),
			QuotaExpiresRule:    result.QuotaExpiresRule,
			QuotaExpiresAt:      util.FormatDateTime(result.QuotaExpiresAt),
			QuotaExpiresMinutes: result.QuotaExpiresMinutes,
			IsBindGroup:         result.IsBindGroup,
			Group:               result.Group,
			IpWhitelist:         result.IpWhitelist,
			IpBlacklist:         result.IpBlacklist,
			Remark:              result.Remark,
			Status:              result.Status,
			CreatedAt:           util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:           util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.AppKeyPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 应用密钥模型权限
func (s *sAppKey) Models(ctx context.Context, params model.AppKeyModelsReq) error {

	oldData, err := dao.AppKey.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if service.Session().IsResellerRole(ctx) && oldData.Rid != service.Session().GetRid(ctx) {
		return errors.New("Unauthorized")
	}

	if service.Session().IsUserRole(ctx) && oldData.UserId != service.Session().GetUserId(ctx) {
		return errors.New("Unauthorized")
	}

	newData, err := dao.AppKey.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"models": params.Models,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP_KEY, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 应用密钥绑定分组
func (s *sAppKey) Group(ctx context.Context, params model.AppKeyGroupReq) error {

	oldData, err := dao.AppKey.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if service.Session().IsResellerRole(ctx) && oldData.Rid != service.Session().GetRid(ctx) {
		return errors.New("Unauthorized")
	}

	if service.Session().IsUserRole(ctx) && oldData.UserId != service.Session().GetUserId(ctx) {
		return errors.New("Unauthorized")
	}

	newData, err := dao.AppKey.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"is_bind_group": params.IsBindGroup,
		"group":         params.Group,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP_KEY, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 应用密钥批量操作
func (s *sAppKey) BatchOperate(ctx context.Context, params model.AppKeyBatchOperateReq) (keys string, err error) {

	userId := service.Session().GetUserId(ctx)

	switch params.Action {
	case consts.ACTION_CREATE:

		app, err := dao.App.FindByAppId(ctx, params.AppId)
		if err != nil {
			logger.Error(ctx, err)
			return "", errors.New("应用ID不存在, 请重新输入")
		}

		if service.Session().IsResellerRole(ctx) || service.Session().IsAdminRole(ctx) {

			if params.UserId == 0 {
				params.UserId = app.UserId
			}

			userId = params.UserId

			user, err := service.User().GetUserByUserId(ctx, userId)
			if err != nil {
				logger.Error(ctx, err)
				return "", errors.New("用户ID不存在, 请重新输入")
			}

			if user.UserId != app.UserId {
				return "", errors.New("用户ID与应用ID不匹配, 请核对后重新输入")
			}
		}

		for i := 0; i < params.N; i++ {

			createKey, err := s.Create(ctx, model.AppKeyCreateReq{UserId: userId, AppId: params.AppId})
			if err != nil {
				logger.Error(ctx, err)
				return keys, err
			}

			key := &do.AppKey{
				UserId:              userId,
				AppId:               params.AppId,
				Key:                 createKey,
				BillingMethods:      params.BillingMethods,
				Models:              params.Models,
				IsLimitQuota:        params.IsLimitQuota,
				Quota:               common.ConvQuotaUnit(params.Quota),
				QuotaExpiresRule:    params.QuotaExpiresRule,
				QuotaExpiresAt:      util.ConvTimestampMilli(params.QuotaExpiresAt),
				QuotaExpiresMinutes: params.QuotaExpiresMinutes,
				IsBindGroup:         params.IsBindGroup,
				Group:               params.Group,
				IpWhitelist:         gstr.Split(gstr.Trim(params.IpWhitelist), "\n"),
				IpBlacklist:         gstr.Split(gstr.Trim(params.IpBlacklist), "\n"),
				Remark:              params.Remark,
				Status:              1,
			}

			keys += key.Key + "\n"

			if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

				id, err := dao.AppKey.Insert(ctx, key)
				if err != nil {
					logger.Error(ctx, err)
					return
				}

				keyInfo := &entity.AppKey{
					Id:                  id,
					UserId:              key.UserId,
					AppId:               key.AppId,
					Key:                 key.Key,
					BillingMethods:      key.BillingMethods,
					Models:              key.Models,
					IsLimitQuota:        key.IsLimitQuota,
					Quota:               key.Quota,
					UsedQuota:           key.UsedQuota,
					QuotaExpiresRule:    key.QuotaExpiresRule,
					QuotaExpiresAt:      key.QuotaExpiresAt,
					QuotaExpiresMinutes: key.QuotaExpiresMinutes,
					IsBindGroup:         key.IsBindGroup,
					Group:               key.Group,
					IpWhitelist:         key.IpWhitelist,
					IpBlacklist:         key.IpBlacklist,
					Remark:              key.Remark,
					Status:              key.Status,
					Rid:                 key.Rid,
				}

				fields := g.Map{
					fmt.Sprintf(consts.KEY_QUOTA_FIELD, keyInfo.AppId, keyInfo.Key):          key.Quota,
					fmt.Sprintf(consts.KEY_IS_LIMIT_QUOTA_FIELD, keyInfo.AppId, keyInfo.Key): key.IsLimitQuota,
				}

				if _, err = redis.HSet(ctx, fmt.Sprintf(consts.API_USER_USAGE_KEY, app.UserId), fields); err != nil {
					logger.Error(ctx, err)
				}

				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP_KEY, model.PubMessage{
					Action:  consts.ACTION_CREATE,
					NewData: keyInfo,
				}); err != nil {
					logger.Error(ctx, err)
				}

			}, nil); err != nil {
				logger.Error(ctx, err)
			}
		}

		if len(keys) > 0 {
			keys = keys[:len(keys)-1]
		}

	case consts.ACTION_UPDATE:

		for _, id := range params.Ids {
			if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

				if _, err := s.Config(ctx, model.AppKeyConfigReq{
					Id:                  id,
					BillingMethods:      params.BillingMethods,
					Models:              params.Models,
					IsLimitQuota:        params.IsLimitQuota,
					Quota:               params.Quota,
					QuotaExpiresRule:    params.QuotaExpiresRule,
					QuotaExpiresAt:      params.QuotaExpiresAt,
					QuotaExpiresMinutes: params.QuotaExpiresMinutes,
					IsBindGroup:         params.IsBindGroup,
					Group:               params.Group,
					IpWhitelist:         params.IpWhitelist,
					IpBlacklist:         params.IpBlacklist,
					Remark:              params.Remark,
					Status:              params.Status,
				}); err != nil {
					logger.Error(ctx, err)
				}

			}, nil); err != nil {
				logger.Error(ctx, err)
			}
		}

	case consts.ACTION_STATUS:
		for _, id := range params.Ids {
			if err := s.ChangeStatus(ctx, model.AppKeyChangeStatusReq{
				Id:     id,
				Status: gconv.Int(params.Value),
			}); err != nil {
				logger.Error(ctx, err)
			}
		}
	case consts.ACTION_DELETE:
		for _, id := range params.Ids {
			if err := s.Delete(ctx, id); err != nil {
				logger.Error(ctx, err)
			}
		}
	case consts.ACTION_ALL_UPDATE, consts.ACTION_ALL_STATUS, consts.ACTION_ALL_DELETE:

		filter := bson.M{}

		if service.Session().IsResellerRole(ctx) {
			filter["rid"] = service.Session().GetRid(ctx)
		}

		if service.Session().IsUserRole(ctx) {
			filter["user_id"] = service.Session().GetUserId(ctx)
		} else if params.QueryParams.UserId != 0 {
			filter["user_id"] = params.QueryParams.UserId
		}

		if params.QueryParams.AppId != 0 {
			filter["app_id"] = params.QueryParams.AppId
		}

		if params.QueryParams.Key != "" {
			filter["key"] = bson.M{
				"$regex": regexp.QuoteMeta(params.QueryParams.Key),
			}
		}

		if len(params.QueryParams.Models) > 0 {
			filter["models"] = bson.M{
				"$in": params.QueryParams.Models,
			}
		}

		if params.QueryParams.Status != 0 {
			filter["status"] = params.QueryParams.Status
		}

		if params.QueryParams.Quota != 0 {
			filter["is_limit_quota"] = true
			filter["quota"] = bson.M{
				"$lte": params.QueryParams.Quota * consts.QUOTA_DEFAULT_UNIT,
			}
		}

		if len(params.QueryParams.QuotaExpiresAt) > 0 {
			gte := gtime.NewFromStrFormat(params.QueryParams.QuotaExpiresAt[0], time.DateOnly).StartOfDay().TimestampMilli()
			lte := gtime.NewFromStrLayout(params.QueryParams.QuotaExpiresAt[1], time.DateOnly).EndOfDay(true).TimestampMilli()
			filter["quota_expires_at"] = bson.M{
				"$gte": gte,
				"$lte": lte,
			}
		}

		if params.QueryParams.Remark != "" {
			filter["remark"] = bson.M{
				"$regex": regexp.QuoteMeta(params.QueryParams.Remark),
			}
		}

		keys, err := dao.AppKey.Find(ctx, filter)
		if err != nil {
			logger.Error(ctx, err)
			return "", err
		}

		switch params.Action {
		case consts.ACTION_ALL_UPDATE:

			for _, key := range keys {
				if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

					if _, err := s.Config(ctx, model.AppKeyConfigReq{
						Id:                  key.Id,
						UserId:              key.UserId,
						AppId:               key.AppId,
						BillingMethods:      params.BillingMethods,
						Models:              params.Models,
						IsLimitQuota:        params.IsLimitQuota,
						Quota:               params.Quota,
						QuotaExpiresRule:    params.QuotaExpiresRule,
						QuotaExpiresAt:      params.QuotaExpiresAt,
						QuotaExpiresMinutes: params.QuotaExpiresMinutes,
						IsBindGroup:         params.IsBindGroup,
						Group:               params.Group,
						IpWhitelist:         params.IpWhitelist,
						IpBlacklist:         params.IpBlacklist,
						Remark:              params.Remark,
						Status:              params.Status,
					}); err != nil {
						logger.Error(ctx, err)
					}

				}, nil); err != nil {
					logger.Error(ctx, err)
				}
			}

		case consts.ACTION_ALL_STATUS:

			for _, key := range keys {
				if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

					if err = s.ChangeStatus(ctx, model.AppKeyChangeStatusReq{
						Id:     key.Id,
						Status: gconv.Int(params.Value),
					}); err != nil {
						logger.Error(ctx, err)
					}

				}, nil); err != nil {
					logger.Error(ctx, err)
				}
			}

		case consts.ACTION_ALL_DELETE:

			for _, key := range keys {
				if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

					if err = s.Delete(ctx, key.Id); err != nil {
						logger.Error(ctx, err)
					}

				}, nil); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}

	return keys, err
}

// 应用密钥导出
func (s *sAppKey) Export(ctx context.Context, params model.AppKeyExportReq) (string, error) {

	filter := bson.M{}

	if len(params.Ids) > 0 {
		filter = bson.M{"_id": bson.M{"$in": params.Ids}}
	} else {
		if params.AppId != 0 {
			filter["app_id"] = params.AppId
		}
	}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
	} else {
		if params.UserId != 0 {
			filter["user_id"] = params.UserId
		}
	}

	results, err := dao.AppKey.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	apps, err := service.App().List(ctx, model.AppListReq{UserId: params.UserId, AppId: params.AppId})
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	appMap := util.ToMap(apps, func(t *model.App) int {
		return t.AppId
	})

	colFieldMap := make(map[string]string)
	colFieldMap["应用ID"] = "AppId"
	colFieldMap["应用名称"] = "AppName"
	colFieldMap["应用密钥"] = "Key"
	colFieldMap["额度"] = "Quota"
	colFieldMap["额度过期时间"] = "QuotaExpiresAt"
	colFieldMap["备注"] = "Remark"

	var titleCols []string
	if service.Session().IsUserRole(ctx) {
		titleCols = append(titleCols, "应用ID", "应用名称", "应用密钥", "额度", "额度过期时间", "备注")
	} else {
		titleCols = append(titleCols, "用户ID", "应用ID", "应用名称", "应用密钥", "额度", "额度过期时间", "备注")
		colFieldMap["用户ID"] = "UserId"
	}

	filePath := fmt.Sprintf("./resource/export/app_key_%d.xlsx", gtime.TimestampMilli())

	values := make([]any, 0)
	for _, result := range results {

		appKeyExport := &model.AppKeyExport{
			UserId:         result.UserId,
			AppId:          result.AppId,
			Key:            result.Key,
			Quota:          "不限",
			QuotaExpiresAt: util.FormatDateTime(result.QuotaExpiresAt),
			Remark:         result.Remark,
		}

		if result.Quota > 0 {
			appKeyExport.Quota = gconv.String(common.ConvQuotaUnitReverse(result.Quota))
		}

		if app, ok := appMap[result.AppId]; ok {
			appKeyExport.AppName = app.Name
		}

		values = append(values, appKeyExport)
	}

	if err = util.ExcelExport("应用密钥", titleCols, colFieldMap, values, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}
