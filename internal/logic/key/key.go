package key

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	mcommon "github.com/iimeta/fastapi-admin/v2/internal/model/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/crypto"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

type sKey struct {
	checkRedsync *redsync.Redsync
}

func init() {
	service.RegisterKey(New())
}

func New() service.IKey {
	return &sKey{
		checkRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
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
					ProviderId:     params.ProviderId,
					Key:            gstr.Trim(k),
					Weight:         params.Weight,
					Models:         params.Models,
					ModelAgents:    params.ModelAgents,
					IsAgentsOnly:   params.IsAgentsOnly,
					IsNeverDisable: params.IsNeverDisable,
					Remark:         params.Remark,
					Status:         params.Status,
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
					Id:             key.Id,
					ProviderId:     params.ProviderId,
					Key:            key.Key,
					Weight:         key.Weight,
					Models:         modelSet.Slice(),
					ModelAgents:    modelAgentSet.Slice(),
					IsAgentsOnly:   params.IsAgentsOnly,
					IsNeverDisable: params.IsNeverDisable,
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

	if isModelAgent {
		params.Remark = oldData.Remark
		params.Status = oldData.Status
	}

	key, err := dao.Key.FindOneAndUpdateById(ctx, params.Id, &do.Key{
		ProviderId:         params.ProviderId,
		Key:                gstr.Trim(params.Key),
		Weight:             params.Weight,
		Models:             params.Models,
		ModelAgents:        params.ModelAgents,
		IsAgentsOnly:       params.IsAgentsOnly,
		IsNeverDisable:     params.IsNeverDisable,
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

	key, err := dao.Key.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"status":               params.Status,
		"is_auto_disabled":     false,
		"auto_disabled_reason": "",
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_KEY, model.PubMessage{
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

	key, err := dao.Key.FindOneAndDeleteById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_KEY, model.PubMessage{
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

	providerName := key.ProviderId
	if key.ProviderId != "" {
		if provider, err := dao.Provider.FindById(ctx, key.ProviderId); err == nil && provider != nil {
			providerName = provider.Name
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
		ProviderId:         key.ProviderId,
		ProviderName:       providerName,
		Key:                key.Key,
		Weight:             key.Weight,
		Models:             key.Models,
		ModelNames:         modelNames,
		ModelAgents:        key.ModelAgents,
		ModelAgentNames:    modelAgentNames,
		IsAgentsOnly:       key.IsAgentsOnly,
		IsNeverDisable:     key.IsNeverDisable,
		UsedQuota:          common.ConvQuotaUnitReverse(key.UsedQuota),
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

	filter := bson.M{}

	if params.ProviderId != "" {
		filter["provider_id"] = params.ProviderId
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

	if len(params.ModelAgents) > 0 {
		filter["model_agents"] = bson.M{
			"$in": params.ModelAgents,
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if params.Remark != "" {
		filter["$or"] = bson.A{
			bson.M{"remark": bson.M{
				"$regex": regexp.QuoteMeta(params.Remark),
			}},
			bson.M{"auto_disabled_reason": bson.M{
				"$regex": regexp.QuoteMeta(params.Remark),
			}},
		}
	}

	results, err := dao.Key.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-updated_at", "key"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	providerMap := make(map[string]*entity.Provider)
	modelAgentMap := make(map[string]*entity.ModelAgent)

	providers, err := dao.Provider.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	providerMap = util.ToMap(providers, func(t *entity.Provider) string {
		return t.Id
	})

	modelAgentResults, err := dao.ModelAgent.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelAgentMap = util.ToMap(modelAgentResults, func(t *entity.ModelAgent) string {
		return t.Id
	})

	items := make([]*model.Key, 0)
	for _, result := range results {

		key := &model.Key{
			Id:           result.Id,
			ProviderId:   result.ProviderId,
			Key:          util.Desensitize(result.Key),
			Weight:       result.Weight,
			Models:       result.Models,
			ModelAgents:  result.ModelAgents,
			IsAgentsOnly: result.IsAgentsOnly,
			UsedQuota:    common.ConvQuotaUnitReverse(result.UsedQuota),
			Remark:       result.Remark,
			Status:       result.Status,
			CreatedAt:    util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:    util.FormatDateTimeMonth(result.UpdatedAt),
		}

		providerName := result.ProviderId
		if providerMap[result.ProviderId] != nil {
			providerName = providerMap[result.ProviderId].Name
		}
		key.ProviderName = providerName

		modelAgentNames := make([]string, 0)
		for _, id := range result.ModelAgents {
			if modelAgentMap[id] != nil {
				modelAgentNames = append(modelAgentNames, modelAgentMap[id].Name)
			}
		}
		key.ModelAgentNames = modelAgentNames

		items = append(items, key)
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

	filter := bson.M{}

	if params.ProviderId != "" {
		filter["provider_id"] = params.ProviderId
	}

	results, err := dao.Key.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"status", "-updated_at", "key"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Key, 0)
	for _, result := range results {
		items = append(items, &model.Key{
			Id:         result.Id,
			ProviderId: result.ProviderId,
			Key:        util.Desensitize(result.Key),
			Remark:     result.Remark,
			Status:     result.Status,
		})
	}

	return items, nil
}

// 密钥批量操作
func (s *sKey) BatchOperate(ctx context.Context, params model.KeyBatchOperateReq) error {

	// 异步处理, 存在一定程度的延迟性
	if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

		switch params.Action {
		case consts.ACTION_STATUS:
			for _, id := range params.Ids {
				if err := s.ChangeStatus(ctx, model.KeyChangeStatusReq{
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
		case consts.ACTION_ALL_STATUS:

			filter := bson.M{}

			if params.ProviderId != "" {
				filter["provider_id"] = params.ProviderId
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

			if len(params.ModelAgents) > 0 {
				filter["model_agents"] = bson.M{
					"$in": params.ModelAgents,
				}
			}

			if params.Status != 0 {
				filter["status"] = params.Status
			}

			if params.Remark != "" {
				filter["remark"] = bson.M{
					"$regex": regexp.QuoteMeta(params.Remark),
				}
			}

			keys, err := dao.Key.Find(ctx, filter)
			if err != nil {
				logger.Error(ctx, err)
				return
			}

			for _, key := range keys {
				if err = s.ChangeStatus(ctx, model.KeyChangeStatusReq{
					Id:     key.Id,
					Status: gconv.Int(params.Value),
				}); err != nil {
					logger.Error(ctx, err)
				}
			}

		case consts.ACTION_ALL_DELETE:

			filter := bson.M{}

			if params.ProviderId != "" {
				filter["provider_id"] = params.ProviderId
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

			if len(params.ModelAgents) > 0 {
				filter["model_agents"] = bson.M{
					"$in": params.ModelAgents,
				}
			}

			if params.Status != 0 {
				filter["status"] = params.Status
			}

			if params.Remark != "" {
				filter["remark"] = bson.M{
					"$regex": regexp.QuoteMeta(params.Remark),
				}
			}

			keys, err := dao.Key.Find(ctx, filter)
			if err != nil {
				logger.Error(ctx, err)
				return
			}

			for _, key := range keys {
				if err = s.Delete(ctx, key.Id); err != nil {
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

// 根据Keys查询密钥详情列表
func (s *sKey) DetailListByKey(ctx context.Context, keys []string) ([]*entity.Key, error) {

	filter := bson.M{
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

	newData, err := dao.Key.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"models": params.Models,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_KEY, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 检查任务
func (s *sKey) CheckTask(ctx context.Context, enableError mcommon.EnableError) {

	logger.Infof(ctx, "sKey CheckTask enableError: %s start", enableError.Error)

	now := gtime.TimestampMilli()

	mutex := s.checkRedsync.NewMutex(fmt.Sprintf(consts.TASK_CHECK_LOCK_KEY, crypto.SM3(enableError.Error)), redsync.WithExpiry(enableError.EnableTime*time.Second))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sKey CheckTask enableError: "+enableError.Error, err)
		logger.Debugf(ctx, "sKey CheckTask enableError: %s end time: %d", enableError.Error, gtime.TimestampMilli()-now)
		return
	}
	logger.Debugf(ctx, "sKey CheckTask enableError: %s lock", enableError.Error)

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debugf(ctx, "sKey CheckTask enableError: %s unlock", enableError.Error)
		}
		logger.Debugf(ctx, "sKey CheckTask enableError: %s end time: %d", enableError.Error, gtime.TimestampMilli()-now)
	}()

	keys, err := dao.Key.Find(ctx, bson.M{
		"status":           2,
		"is_auto_disabled": true,
		"auto_disabled_reason": bson.M{
			"$regex": regexp.QuoteMeta(enableError.Error),
		},
		"updated_at": bson.M{
			"$lte": gtime.TimestampMilli() - (enableError.EnableTime * time.Second).Milliseconds(),
		},
	})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	modelAgentSet := gset.NewStrSet()
	for _, key := range keys {
		if err = s.ChangeStatus(ctx, model.KeyChangeStatusReq{
			Id:     key.Id,
			Status: 1,
		}); err != nil {
			logger.Error(ctx, err)
		}
		modelAgentSet.Add(key.ModelAgents...)
	}

	modelAgents, err := dao.ModelAgent.Find(ctx, bson.M{
		"status":           2,
		"is_auto_disabled": true,
		"_id": bson.M{
			"$in": modelAgentSet.Slice(),
		},
	})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	for _, modelAgent := range modelAgents {
		if err = service.ModelAgent().ChangeStatus(ctx, model.ModelAgentChangeStatusReq{
			Id:     modelAgent.Id,
			Status: 1,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}

	if _, err = redis.Set(ctx, fmt.Sprintf(consts.TASK_CHECK_END_TIME_KEY, crypto.SM3(enableError.Error)), gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
