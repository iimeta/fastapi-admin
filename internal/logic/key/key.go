package key

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type sKey struct{}

func init() {
	service.RegisterKey(New())
}

func New() service.IKey {
	return &sKey{}
}

// 新建密钥
func (s *sKey) Create(ctx context.Context, params model.KeyCreateReq, isModelAgent bool) error {

	keys := gstr.Split(gstr.Trim(params.Key), "\n")

	keyList, err := dao.Key.Find(ctx, bson.M{"key": bson.M{"$in": keys}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 异步处理, 存在一定程度的延迟性
	if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

		keyMap := util.ToMap(keyList, func(t *entity.Key) string {
			return t.Key
		})

		for _, k := range keys {

			if len(gstr.Trim(k)) == 0 {
				continue
			}

			key := keyMap[k]

			if key == nil {

				id, err := dao.Key.Insert(ctx, &do.Key{
					Corp:         params.Corp,
					Key:          gstr.Trim(k),
					Type:         2,
					LbStrategy:   params.LbStrategy,
					Weight:       params.Weight,
					Models:       params.Models,
					ModelAgents:  params.ModelAgents,
					IsAgentsOnly: params.IsAgentsOnly,
					Remark:       params.Remark,
					Status:       params.Status,
				})

				if err != nil {
					logger.Error(ctx, err)
				}

				key, err := dao.Key.FindById(ctx, id)
				if err != nil {
					logger.Error(ctx, err)
				}

				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_KEY, model.PubMessage{
					Action:  consts.ACTION_CREATE,
					NewData: key,
				}); err != nil {
					logger.Error(ctx, err)
				}

			} else {

				modelSet := gset.NewStrSet()
				modelSet.Add(key.Models...)
				modelSet.Add(params.Models...)

				modelAgentSet := gset.NewStrSet()
				modelAgentSet.Add(key.ModelAgents...)
				modelAgentSet.Add(params.ModelAgents...)

				if err := s.Update(ctx, model.KeyUpdateReq{
					Id:           key.Id,
					Corp:         params.Corp,
					Key:          key.Key,
					LbStrategy:   key.LbStrategy,
					Weight:       key.Weight,
					Models:       modelSet.Slice(),
					ModelAgents:  modelAgentSet.Slice(),
					IsAgentsOnly: params.IsAgentsOnly,
				}, isModelAgent); err != nil {
					logger.Error(ctx, err)
				}
			}
		}
	}, nil); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新密钥
func (s *sKey) Update(ctx context.Context, params model.KeyUpdateReq, isModelAgent bool) error {

	oldData, err := dao.Key.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if service.Session().IsUserRole(ctx) && oldData.UserId != service.Session().GetUserId(ctx) {
		return errors.New("Unauthorized")
	}

	if isModelAgent {
		params.Remark = oldData.Remark
		params.Status = oldData.Status
	}

	key, err := dao.Key.FindOneAndUpdateById(ctx, params.Id, &do.Key{
		Corp:               params.Corp,
		Key:                gstr.Trim(params.Key),
		LbStrategy:         params.LbStrategy,
		Weight:             params.Weight,
		Models:             params.Models,
		ModelAgents:        params.ModelAgents,
		IsAgentsOnly:       params.IsAgentsOnly,
		Remark:             params.Remark,
		Status:             params.Status,
		IsAutoDisabled:     oldData.IsAutoDisabled,
		AutoDisabledReason: oldData.AutoDisabledReason,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_KEY, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: key,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改密钥状态
func (s *sKey) ChangeStatus(ctx context.Context, params model.KeyChangeStatusReq) error {

	if service.Session().IsUserRole(ctx) {

		key, err := dao.Key.FindById(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if key.UserId != service.Session().GetUserId(ctx) {
			return errors.New("Unauthorized")
		}
	}

	key, err := dao.Key.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"status":               params.Status,
		"is_auto_disabled":     false,
		"auto_disabled_reason": "",
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	channel := consts.CHANGE_CHANNEL_KEY

	if key.Type == 1 {
		channel = consts.CHANGE_CHANNEL_APP_KEY
	}

	if _, err = redis.Publish(ctx, channel, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: key,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除密钥
func (s *sKey) Delete(ctx context.Context, id string) error {

	if service.Session().IsUserRole(ctx) {

		key, err := dao.Key.FindById(ctx, id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if key.UserId != service.Session().GetUserId(ctx) {
			return errors.New("Unauthorized")
		}
	}

	key, err := dao.Key.FindOneAndDeleteById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	channel := consts.CHANGE_CHANNEL_KEY

	if key.Type == 1 {
		channel = consts.CHANGE_CHANNEL_APP_KEY
	}

	if _, err = redis.Publish(ctx, channel, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: key,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 密钥详情
func (s *sKey) Detail(ctx context.Context, id string) (*model.Key, error) {

	key, err := dao.Key.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if service.Session().IsUserRole(ctx) && key.UserId != service.Session().GetUserId(ctx) {
		return nil, errors.New("Unauthorized")
	}

	corpName := key.Corp
	if key.Corp != "" {
		if corp, err := dao.Corp.FindById(ctx, key.Corp); err == nil && corp != nil {
			corpName = corp.Name
		}
	}

	modelNames, err := service.Model().ModelNames(ctx, key.Models)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelAgentNames := make([]string, 0)

	if len(key.ModelAgents) > 0 {

		modelAgentList, err := dao.ModelAgent.Find(ctx, bson.M{"_id": bson.M{"$in": key.ModelAgents}})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		for _, modelAgent := range modelAgentList {
			modelAgentNames = append(modelAgentNames, modelAgent.Name)
		}
	}

	return &model.Key{
		Id:                 key.Id,
		UserId:             key.UserId,
		AppId:              key.AppId,
		Corp:               key.Corp,
		CorpName:           corpName,
		Key:                key.Key,
		Type:               key.Type,
		LbStrategy:         key.LbStrategy,
		Weight:             key.Weight,
		Models:             key.Models,
		ModelNames:         modelNames,
		ModelAgents:        key.ModelAgents,
		ModelAgentNames:    modelAgentNames,
		IsAgentsOnly:       key.IsAgentsOnly,
		IsLimitQuota:       key.IsLimitQuota,
		Quota:              key.Quota,
		UsedQuota:          key.UsedQuota,
		QuotaExpiresAt:     util.FormatDateTime(key.QuotaExpiresAt),
		IpWhitelist:        key.IpWhitelist,
		IpBlacklist:        key.IpBlacklist,
		Remark:             key.Remark,
		Status:             key.Status,
		IsAutoDisabled:     key.IsAutoDisabled,
		AutoDisabledReason: key.AutoDisabledReason,
		Creator:            key.Creator,
		Updater:            key.Updater,
		CreatedAt:          util.FormatDateTime(key.CreatedAt),
		UpdatedAt:          util.FormatDateTime(key.UpdatedAt),
	}, nil
}

// 密钥分页列表
func (s *sKey) Page(ctx context.Context, params model.KeyPageReq) (*model.KeyPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{
		"type": params.Type,
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
	} else if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.Corp != "" {
		filter["corp"] = params.Corp
	}

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	if params.Key != "" {
		filter["key"] = bson.M{
			"$regex": params.Key,
		}
	}

	if len(params.Models) > 0 {
		filter["models"] = bson.M{
			"$in": params.Models,
		}
	}

	if len(params.ModelAgents) > 0 {
		filter["model_agents"] = bson.M{
			"$in": params.ModelAgents,
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if params.Quota != 0 {
		filter["is_limit_quota"] = true
		filter["quota"] = bson.M{
			"$lte": params.Quota * consts.QUOTA_USD_UNIT,
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
			"$regex": params.Remark,
		}
	}

	results, err := dao.Key.FindByPage(ctx, paging, filter, "", "status", "-updated_at", "key")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	corps, err := dao.Corp.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	corpMap := util.ToMap(corps, func(t *entity.Corp) string {
		return t.Id
	})

	models, err := service.Model().List(ctx, model.ModelListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelMap := util.ToMap(models, func(t *model.Model) string {
		return t.Id
	})

	modelAgentMap := make(map[string]*entity.ModelAgent)
	if service.Session().IsAdminRole(ctx) {

		modelAgentResults, err := dao.ModelAgent.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		modelAgentMap = util.ToMap(modelAgentResults, func(t *entity.ModelAgent) string {
			return t.Id
		})
	}

	items := make([]*model.Key, 0)
	for _, result := range results {

		corpName := result.Corp
		if corpMap[result.Corp] != nil {
			corpName = corpMap[result.Corp].Name
		}

		modelNames := make([]string, 0)
		for _, id := range result.Models {
			if modelMap[id] != nil {
				modelNames = append(modelNames, modelMap[id].Name)
			}
		}

		modelAgentNames := make([]string, 0)
		for _, id := range result.ModelAgents {
			if modelAgentMap[id] != nil {
				modelAgentNames = append(modelAgentNames, modelAgentMap[id].Name)
			}
		}

		items = append(items, &model.Key{
			Id:                 result.Id,
			UserId:             result.UserId,
			AppId:              result.AppId,
			Corp:               result.Corp,
			CorpName:           corpName,
			Key:                util.Desensitize(result.Key),
			Type:               result.Type,
			LbStrategy:         result.LbStrategy,
			Weight:             result.Weight,
			Models:             result.Models,
			ModelNames:         modelNames,
			ModelAgents:        result.ModelAgents,
			ModelAgentNames:    modelAgentNames,
			IsAgentsOnly:       result.IsAgentsOnly,
			IsLimitQuota:       result.IsLimitQuota,
			Quota:              result.Quota,
			UsedQuota:          result.UsedQuota,
			QuotaExpiresAt:     util.FormatDateTime(result.QuotaExpiresAt),
			IpWhitelist:        result.IpWhitelist,
			IpBlacklist:        result.IpBlacklist,
			Remark:             result.Remark,
			Status:             result.Status,
			IsAutoDisabled:     result.IsAutoDisabled,
			AutoDisabledReason: result.AutoDisabledReason,
			CreatedAt:          util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:          util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.KeyPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 密钥列表
func (s *sKey) List(ctx context.Context, params model.KeyListReq) ([]*model.Key, error) {

	filter := bson.M{
		"type": 2,
	}

	if params.Corp != "" {
		filter["corp"] = params.Corp
	}

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	results, err := dao.Key.Find(ctx, filter, "status", "-updated_at", "key")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Key, 0)
	for _, result := range results {
		items = append(items, &model.Key{
			Id:     result.Id,
			Corp:   result.Corp,
			Key:    util.Desensitize(result.Key),
			Type:   result.Type,
			Remark: result.Remark,
			Status: result.Status,
		})
	}

	return items, nil
}

// 密钥批量操作
func (s *sKey) BatchOperate(ctx context.Context, params model.KeyBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_STATUS:
		for _, id := range params.Ids {
			if err := s.ChangeStatus(ctx, model.KeyChangeStatusReq{
				Id:     id,
				Status: gconv.Int(params.Value),
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	case consts.ACTION_DELETE:
		for _, id := range params.Ids {
			if err := s.Delete(ctx, id); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 根据Keys查询密钥详情列表
func (s *sKey) DetailListByKey(ctx context.Context, keys []string) ([]*entity.Key, error) {

	filter := bson.M{
		"type": 2,
		"key": bson.M{
			"$in": keys,
		},
		"status": 1,
	}

	results, err := dao.Key.Find(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return results, nil
}

// 密钥模型权限
func (s *sKey) Models(ctx context.Context, params model.KeyModelsReq) error {

	oldData, err := dao.Key.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if service.Session().IsUserRole(ctx) && oldData.UserId != service.Session().GetUserId(ctx) {
		return errors.New("Unauthorized")
	}

	newData, err := dao.Key.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"models": params.Models,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	channel := consts.CHANGE_CHANNEL_KEY

	if newData.Type == 1 {
		channel = consts.CHANGE_CHANNEL_APP_KEY
	}

	if _, err = redis.Publish(ctx, channel, model.PubMessage{
		Action:  consts.ACTION_MODELS,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
