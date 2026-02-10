package model_agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"text/template"
	"time"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	sdk "github.com/iimeta/fastapi-sdk/v2"
	sconsts "github.com/iimeta/fastapi-sdk/v2/consts"
	"github.com/iimeta/fastapi-sdk/v2/general"
	smodel "github.com/iimeta/fastapi-sdk/v2/model"
	"github.com/iimeta/fastapi-sdk/v2/options"
	sutil "github.com/iimeta/fastapi-sdk/v2/util"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type sModelAgent struct{}

func init() {
	service.RegisterModelAgent(New())
}

func New() service.IModelAgent {
	return &sModelAgent{}
}

// 新建模型代理
func (s *sModelAgent) Create(ctx context.Context, params model.ModelAgentCreateReq) (string, error) {

	if s.IsNameExist(ctx, params.Name) {
		return "", errors.Newf("模型代理名称 \"%s\" 已存在", params.Name)
	}

	id, err := dao.ModelAgent.Insert(ctx, &do.ModelAgent{
		ProviderId:           params.ProviderId,
		Name:                 gstr.Trim(params.Name),
		BaseUrl:              gstr.Trim(params.BaseUrl),
		Path:                 gstr.Trim(params.Path),
		Weight:               params.Weight,
		BillingMethods:       params.BillingMethods,
		Models:               params.Models,
		IsEnableModelReplace: params.IsEnableModelReplace,
		ReplaceModels:        params.ReplaceModels,
		TargetModels:         params.TargetModels,
		IsNeverDisable:       params.IsNeverDisable,
		LbStrategy:           params.LbStrategy,
		Remark:               params.Remark,
		Status:               params.Status,
	})

	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if params.Key != "" {
		if err = service.Key().Create(ctx, model.KeyCreateReq{
			ProviderId:     params.ProviderId,
			Key:            params.Key,
			Weight:         params.Weight,
			Models:         params.Models,
			ModelAgents:    []string{id},
			IsAgentsOnly:   params.IsAgentsOnly,
			IsNeverDisable: params.IsNeverDisableKey,
			Remark:         params.Remark,
			Status:         params.Status,
		}, true); err != nil {
			logger.Error(ctx, err)
			return "", err
		}
	}

	modelAgent, err := s.Detail(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
		Action:  consts.ACTION_CREATE,
		NewData: modelAgent,
	}); err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if len(params.Groups) > 0 {
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

			for _, groupId := range params.Groups {

				oldData, err := dao.Group.FindById(ctx, groupId)
				if err != nil {
					logger.Error(ctx, err)
					continue
				}

				newData, err := dao.Group.FindOneAndUpdateById(ctx, groupId, bson.M{
					"$push": bson.M{
						"model_agents": id,
					},
				})
				if err != nil {
					logger.Error(ctx, err)
					continue
				}

				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
					Action:  consts.ACTION_UPDATE,
					OldData: oldData,
					NewData: newData,
				}); err != nil {
					logger.Error(ctx, err)
				}
			}

		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	return id, nil
}

// 更新模型代理
func (s *sModelAgent) Update(ctx context.Context, params model.ModelAgentUpdateReq) error {

	if s.IsNameExist(ctx, params.Name, params.Id) {
		return errors.Newf("模型代理名称 \"%s\" 已存在", params.Name)
	}

	oldData, err := s.Detail(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.ModelAgent.UpdateById(ctx, params.Id, &do.ModelAgent{
		ProviderId:           params.ProviderId,
		Name:                 gstr.Trim(params.Name),
		BaseUrl:              gstr.Trim(params.BaseUrl),
		Path:                 gstr.Trim(params.Path),
		Weight:               params.Weight,
		BillingMethods:       params.BillingMethods,
		Models:               params.Models,
		IsEnableModelReplace: params.IsEnableModelReplace,
		ReplaceModels:        params.ReplaceModels,
		TargetModels:         params.TargetModels,
		IsNeverDisable:       params.IsNeverDisable,
		LbStrategy:           params.LbStrategy,
		Remark:               params.Remark,
		Status:               params.Status,
		IsAutoDisabled:       oldData.IsAutoDisabled,
		AutoDisabledReason:   oldData.AutoDisabledReason,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	var oldKeyList []*entity.Key

	if oldData.Key != "" {

		oldKeys := gstr.Split(gstr.Trim(oldData.Key), "\n")

		oldKeyList, err = service.Key().DetailListByKey(ctx, oldKeys)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if err = dao.Key.UpdateMany(ctx, bson.M{"model_agents": bson.M{"$in": []string{params.Id}}}, bson.M{
		"$pull": bson.M{
			"model_agents": params.Id,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if params.Key != "" {
		if err = service.Key().Create(ctx, model.KeyCreateReq{
			ProviderId:     params.ProviderId,
			Key:            params.Key,
			Weight:         params.Weight,
			Models:         params.Models,
			ModelAgents:    []string{params.Id},
			IsAgentsOnly:   params.IsAgentsOnly,
			IsNeverDisable: params.IsNeverDisableKey,
			Remark:         params.Remark,
			Status:         params.Status,
		}, true); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if oldKeyList != nil && len(oldKeyList) > 0 {

		oldKeyMap := util.ToMap(oldKeyList, func(t *entity.Key) string {
			return t.Key
		})

		newKeys := gstr.Split(gstr.Trim(params.Key), "\n")
		newKeyMap := util.ToMap(newKeys, func(t string) string {
			return t
		})

		updateKeys := make([]string, 0)
		for _, key := range oldKeyList {
			if newKeyMap[key.Key] == "" {
				updateKeys = append(updateKeys, key.Key)
			}
		}

		if len(updateKeys) > 0 {

			keys, err := service.Key().DetailListByKey(ctx, updateKeys)
			if err != nil {
				logger.Error(ctx, err)
				return err
			}

			for _, key := range keys {
				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_KEY, model.PubMessage{
					Action:  consts.ACTION_UPDATE,
					OldData: oldKeyMap[key.Key],
					NewData: key,
				}); err != nil {
					logger.Error(ctx, err)
					return err
				}
			}
		}
	}

	modelAgent, err := s.Detail(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: modelAgent,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	oldGroups := gset.NewStrSetFrom(oldData.Groups)
	newGroups := gset.NewStrSetFrom(params.Groups)

	if !oldGroups.Equal(newGroups) {
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

			allGroups := gset.NewStrSet()
			allGroups.Add(oldGroups.Slice()...)
			allGroups.Add(newGroups.Slice()...)

			for _, group := range allGroups.Slice() {

				// 新的有, 旧的没有, 说明新增了
				if newGroups.Contains(group) && !oldGroups.Contains(group) {

					oldData, err := dao.Group.FindById(ctx, group)
					if err != nil {
						logger.Error(ctx, err)
						continue
					}

					newData, err := dao.Group.FindOneAndUpdateById(ctx, group, bson.M{
						"$push": bson.M{
							"model_agents": params.Id,
						},
					})
					if err != nil {
						logger.Error(ctx, err)
						continue
					}

					if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
						Action:  consts.ACTION_UPDATE,
						OldData: oldData,
						NewData: newData,
					}); err != nil {
						logger.Error(ctx, err)
					}

				} else if oldGroups.Contains(group) && !newGroups.Contains(group) { // 旧的有, 新的没有, 说明移除了

					oldData, err := dao.Group.FindById(ctx, group)
					if err != nil {
						logger.Error(ctx, err)
						continue
					}

					newData, err := dao.Group.FindOneAndUpdateById(ctx, group, bson.M{
						"$pull": bson.M{
							"model_agents": params.Id,
						},
					})
					if err != nil {
						logger.Error(ctx, err)
						continue
					}

					if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
						Action:  consts.ACTION_UPDATE,
						OldData: oldData,
						NewData: newData,
					}); err != nil {
						logger.Error(ctx, err)
					}
				}
			}

		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	return nil
}

// 更改模型代理状态
func (s *sModelAgent) ChangeStatus(ctx context.Context, params model.ModelAgentChangeStatusReq) error {

	if err := dao.ModelAgent.UpdateById(ctx, params.Id, bson.M{
		"status":               params.Status,
		"is_auto_disabled":     false,
		"auto_disabled_reason": "",
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	modelAgent, err := s.Detail(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err := redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: modelAgent,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除模型代理
func (s *sModelAgent) Delete(ctx context.Context, id string) error {

	modelAgent, err := s.Detail(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.ModelAgent.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.Key.UpdateMany(ctx, bson.M{"model_agents": bson.M{"$in": []string{id}}}, bson.M{
		"$pull": bson.M{
			"model_agents": id,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: modelAgent,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	groups, err := dao.Group.Find(ctx, bson.M{"model_agents": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if len(groups) > 0 {
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

			for _, group := range groups {

				oldData, err := dao.Group.FindById(ctx, group)
				if err != nil {
					logger.Error(ctx, err)
					continue
				}

				newData, err := dao.Group.FindOneAndUpdateById(ctx, group, bson.M{
					"$pull": bson.M{
						"model_agents": id,
					},
				})
				if err != nil {
					logger.Error(ctx, err)
					continue
				}

				if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
					Action:  consts.ACTION_UPDATE,
					OldData: oldData,
					NewData: newData,
				}); err != nil {
					logger.Error(ctx, err)
				}
			}

		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	return nil
}

// 模型代理详情
func (s *sModelAgent) Detail(ctx context.Context, id string) (*model.ModelAgent, error) {

	modelAgent, err := dao.ModelAgent.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	providerName := modelAgent.ProviderId
	if provider, err := dao.Provider.FindById(ctx, modelAgent.ProviderId); err == nil && provider != nil {
		providerName = provider.Name
	}

	modelList, err := dao.Model.Find(ctx, bson.M{"_id": bson.M{"$in": modelAgent.Models}}, &dao.FindOptions{SortFields: []string{"-updated_at", "name"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelNames := make([]string, 0)
	for _, model := range modelList {
		modelNames = append(modelNames, model.Name)
	}

	fallbackModelList, err := dao.Model.Find(ctx, bson.M{"fallback_config.model_agent": id}, &dao.FindOptions{SortFields: []string{"-updated_at", "name"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	fallbackModels := make([]string, 0)
	fallbackModelNames := make([]string, 0)

	for _, model := range fallbackModelList {
		fallbackModels = append(fallbackModels, model.Id)
		fallbackModelNames = append(fallbackModelNames, model.Name)
	}

	keyList, err := dao.Key.Find(ctx, bson.M{"model_agents": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	keys := make([]string, 0)
	for _, key := range keyList {
		keys = append(keys, key.Key)
	}

	groups, err := dao.Group.Find(ctx, bson.M{"model_agents": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	groupIds := make([]string, 0)
	groupNames := make([]string, 0)
	for _, group := range groups {
		groupIds = append(groupIds, group.Id)
		groupNames = append(groupNames, group.Name)
	}

	return &model.ModelAgent{
		Id:                   modelAgent.Id,
		ProviderId:           modelAgent.ProviderId,
		ProviderName:         providerName,
		Name:                 modelAgent.Name,
		BaseUrl:              modelAgent.BaseUrl,
		Path:                 modelAgent.Path,
		Weight:               modelAgent.Weight,
		BillingMethods:       modelAgent.BillingMethods,
		Groups:               groupIds,
		GroupNames:           groupNames,
		Models:               modelAgent.Models,
		ModelNames:           modelNames,
		FallbackModels:       fallbackModels,
		FallbackModelNames:   fallbackModelNames,
		IsEnableModelReplace: modelAgent.IsEnableModelReplace,
		ReplaceModels:        modelAgent.ReplaceModels,
		TargetModels:         modelAgent.TargetModels,
		IsNeverDisable:       modelAgent.IsNeverDisable,
		LbStrategy:           modelAgent.LbStrategy,
		Key:                  gstr.Join(keys, "\n"),
		Remark:               modelAgent.Remark,
		Status:               modelAgent.Status,
		IsAutoDisabled:       modelAgent.IsAutoDisabled,
		AutoDisabledReason:   modelAgent.AutoDisabledReason,
		Creator:              modelAgent.Creator,
		Updater:              modelAgent.Updater,
		CreatedAt:            util.FormatDateTime(modelAgent.CreatedAt),
		UpdatedAt:            util.FormatDateTime(modelAgent.UpdatedAt),
	}, nil
}

// 模型代理分页列表
func (s *sModelAgent) Page(ctx context.Context, params model.ModelAgentPageReq) (*model.ModelAgentPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.ProviderId != "" {
		filter["provider_id"] = params.ProviderId
	}

	if params.Name != "" {
		filter["name"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Name),
		}
	}

	if params.BaseUrl != "" {
		filter["base_url"] = bson.M{
			"$regex": regexp.QuoteMeta(params.BaseUrl),
		}
	}

	if len(params.Models) > 0 {

		modelAgents, err := dao.ModelAgent.Find(ctx, bson.M{"models": bson.M{"$in": params.Models}})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(modelAgents) == 0 {
			return nil, nil
		}

		modelAgentIds := make([]string, 0)
		for _, modelAgent := range modelAgents {
			modelAgentIds = append(modelAgentIds, modelAgent.Id)
		}

		filter["_id"] = bson.M{
			"$in": modelAgentIds,
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

	results, err := dao.ModelAgent.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	providers, err := dao.Provider.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	providerMap := util.ToMap(providers, func(t *entity.Provider) string {
		return t.Id
	})

	groups, err := service.Group().List(ctx, model.GroupListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelList, err := service.Model().List(ctx, model.ModelListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelNameMap := make(map[string]string)

	fallbackModelMap := make(map[string][]string)
	fallbackModelNameMap := make(map[string][]string)

	for _, model := range modelList {

		modelNameMap[model.Id] = model.Name

		if model.IsEnableFallback && model.FallbackConfig.ModelAgent != "" {
			fallbackModelMap[model.FallbackConfig.ModelAgent] = append(fallbackModelMap[model.FallbackConfig.ModelAgent], model.Id)
			fallbackModelNameMap[model.FallbackConfig.ModelAgent] = append(fallbackModelNameMap[model.FallbackConfig.ModelAgent], model.Name)
		}
	}

	items := make([]*model.ModelAgent, 0)
	for _, result := range results {

		providerName := result.ProviderId
		if providerMap[result.ProviderId] != nil {
			providerName = providerMap[result.ProviderId].Name
		}

		groupIds := make([]string, 0)
		groupNames := make([]string, 0)
		for _, group := range groups {
			if slices.Contains(group.ModelAgents, result.Id) {
				groupIds = append(groupIds, group.Id)
				groupNames = append(groupNames, group.Name)
			}
		}

		modelNames := make([]string, 0)
		for _, model := range result.Models {
			modelNames = append(modelNames, modelNameMap[model])
		}

		items = append(items, &model.ModelAgent{
			Id:                 result.Id,
			ProviderId:         result.ProviderId,
			ProviderName:       providerName,
			Name:               result.Name,
			BaseUrl:            result.BaseUrl,
			Path:               result.Path,
			Weight:             result.Weight,
			BillingMethods:     result.BillingMethods,
			LbStrategy:         result.LbStrategy,
			Groups:             groupIds,
			GroupNames:         groupNames,
			Models:             result.Models,
			ModelNames:         modelNames,
			FallbackModels:     fallbackModelMap[result.Id],
			FallbackModelNames: fallbackModelNameMap[result.Id],
			Remark:             result.Remark,
			Status:             result.Status,
			CreatedAt:          util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:          util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.ModelAgentPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 模型代理列表
func (s *sModelAgent) List(ctx context.Context, params model.ModelAgentListReq) ([]*model.ModelAgent, error) {

	filter := bson.M{}

	results, err := dao.ModelAgent.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelList, err := service.Model().List(ctx, model.ModelListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelMap := make(map[string][]string)
	modelNameMap := make(map[string][]string)

	for _, model := range modelList {
		for _, id := range model.ModelAgents {
			modelMap[id] = append(modelMap[id], model.Id)
			modelNameMap[id] = append(modelNameMap[id], model.Name)
		}
	}

	items := make([]*model.ModelAgent, 0)
	for _, result := range results {
		items = append(items, &model.ModelAgent{
			Id:         result.Id,
			ProviderId: result.ProviderId,
			Name:       result.Name,
			BaseUrl:    result.BaseUrl,
			Path:       result.Path,
			Models:     modelMap[result.Id],
			ModelNames: modelNameMap[result.Id],
			Status:     result.Status,
		})
	}

	return items, nil
}

// 模型代理批量操作
func (s *sModelAgent) BatchOperate(ctx context.Context, params model.ModelAgentBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_STATUS:
		for _, id := range params.Ids {
			if err := s.ChangeStatus(ctx, model.ModelAgentChangeStatusReq{
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

// 模型代理名称是否存在
func (s *sModelAgent) IsNameExist(ctx context.Context, name string, id ...string) bool {

	modelAgent, err := dao.ModelAgent.FindOne(ctx, bson.M{"name": gstr.Trim(name)})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false
		}
		logger.Error(ctx, err)
		return true
	}

	if modelAgent != nil {
		if len(id) > 0 && modelAgent.Id == id[0] {
			return false
		}
		return true
	}

	return false
}

// 快速填入模型
func (s *sModelAgent) QuickFillModel(ctx context.Context, params model.ModelAgentQuickFillModelReq) ([]string, error) {

	if params.BaseUrl == "" {
		return nil, errors.New("请输入代理地址后重试")
	}

	keys := gstr.Split(gstr.Trim(params.Key), "\n")
	if len(keys) == 0 || keys[0] == "" {
		return nil, errors.New("请输入密钥后重试")
	}

	result := &model.ModelsRes{}
	if _, err := sutil.HttpGet(ctx, params.BaseUrl+"/models", g.MapStrStr{"Authorization": "Bearer " + keys[0]}, nil, &result, config.Cfg.Http.Timeout*time.Second, config.Cfg.Http.ProxyUrl, nil); err != nil {
		logger.Error(ctx, err)
		return nil, errors.New("获取数据异常, 请手动选择模型")
	}

	if len(result.Data) == 0 {
		return nil, errors.New("获取数据为空, 请手动选择模型")
	}

	list, err := service.Model().List(ctx, model.ModelListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	models := make([]string, 0)

	for _, data := range result.Data {
		for _, model := range list {
			if data.Id == model.Model {
				models = append(models, model.Id)
			}
		}
	}

	return models, nil
}

// 测试模型
func (s *sModelAgent) TestModel(ctx context.Context, params model.ModelAgentTestModelReq) (*model.ModelAgentTestModelRes, error) {

	modelAgent, err := dao.ModelAgent.FindById(ctx, params.ModelAgentId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	m, err := dao.Model.FindById(ctx, params.ModelId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	provider, err := dao.Provider.FindById(ctx, modelAgent.ProviderId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	options := &options.AdapterOptions{
		Provider:                provider.Code,
		Model:                   m.Model,
		Key:                     params.Key,
		BaseUrl:                 params.BaseUrl,
		Path:                    modelAgent.Path,
		Header:                  g.MapStrStr{consts.MODEL_AGENT_HEADER: params.ModelAgentId},
		Timeout:                 config.Cfg.Base.ShortTimeout * time.Second,
		ProxyUrl:                config.Cfg.Http.ProxyUrl,
		IsOfficialFormatRequest: provider.Code != sconsts.PROVIDER_OPENAI,
	}

	if params.TestMethod == 2 {

		keys, err := dao.Key.Find(ctx, bson.M{"model_agents": bson.M{"$in": []string{params.ModelAgentId}}})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		isAvailable := false
		for _, key := range keys {
			if key.Status == 1 {
				options.Key = key.Key
				isAvailable = true
			}
		}

		if !isAvailable {
			return nil, errors.New("此模型代理无可用密钥")
		}

		options.BaseUrl = modelAgent.BaseUrl

	} else if options.BaseUrl == "" {

		address, err := redis.HGetStr(ctx, consts.SERVERS_KEY, fmt.Sprintf("api:%s", util.GetLocalIp()))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if provider.Code == sconsts.PROVIDER_GOOGLE {
			options.BaseUrl = fmt.Sprintf("http://%s%s/v1beta", util.GetLocalIp(), address)
		} else {
			options.BaseUrl = fmt.Sprintf("http://%s%s/v1", util.GetLocalIp(), address)
		}
	}

	adapter := sdk.NewAdapter(ctx, options)

	if general, isGeneral := adapter.(*general.General); isGeneral {
		if general.Path == "" {
			switch m.Type {
			case 1, 100:
				general.Path = "/chat/completions"
			case 2:
				general.Path = "/images/generations"
			case 5:
				general.Path = "/audio/speech"
			case 7:
				general.Path = "/embeddings"
			case 8:
				general.Path = "/videos"
			default:
				general.Path = "/chat/completions"
			}
		}
	}

	var requestData string

	for _, test := range config.Cfg.Test.Tests {

		if test.Provider == provider.Code && test.ModelType == m.Type && test.Model == m.Model {
			requestData = test.RequestData
			break
		}

		if test.Provider == provider.Code && test.ModelType == m.Type {
			requestData = test.RequestData
		}

		if requestData == "" && test.Provider == sconsts.PROVIDER_OPENAI && test.ModelType == m.Type {
			requestData = test.RequestData
		}
	}

	if requestData == "" {
		return nil, errors.New("未找到此模型类型对应的请求数据, 请先配置相应的模型类型请求数据")
	}

	requestDataTmpl, err := template.New("data").Parse(requestData)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	var data bytes.Buffer
	if err = requestDataTmpl.Execute(&data, g.MapStrStr{"model": m.Model}); err != nil {
		return nil, err
	}

	modelAgentTestModelRes := &model.ModelAgentTestModelRes{
		TraceId: gtrace.GetTraceID(ctx),
	}

	switch m.Type {
	case 1, 100:

		response, err := adapter.ChatCompletions(ctx, data.Bytes())
		if err != nil {
			logger.Error(ctx, err)
			modelAgentTestModelRes.Error = err.Error()
		}

		modelAgentTestModelRes.Result = err == nil
		modelAgentTestModelRes.TotalTime = response.TotalTime

		if response.Error != nil && modelAgentTestModelRes.Error == "" {
			modelAgentTestModelRes.Result = false
			modelAgentTestModelRes.Error = response.Error.Error()
		}

	case 2:

		response, err := adapter.ImageGenerations(ctx, data.Bytes())
		if err != nil {
			logger.Error(ctx, err)
			modelAgentTestModelRes.Error = err.Error()
		}

		modelAgentTestModelRes.Result = err == nil
		modelAgentTestModelRes.TotalTime = response.TotalTime

	case 5:

		response, err := adapter.AudioSpeech(ctx, data.Bytes())
		if err != nil {
			logger.Error(ctx, err)
			modelAgentTestModelRes.Error = err.Error()
		}

		modelAgentTestModelRes.Result = err == nil
		modelAgentTestModelRes.TotalTime = response.TotalTime

	case 7:

		response, err := adapter.TextEmbeddings(ctx, data.Bytes())
		if err != nil {
			logger.Error(ctx, err)
			modelAgentTestModelRes.Error = err.Error()
		}

		modelAgentTestModelRes.Result = err == nil
		modelAgentTestModelRes.TotalTime = response.TotalTime

	case 8:

		request := smodel.VideoCreateRequest{}

		if err := json.Unmarshal(data.Bytes(), &request); err == nil {
			logger.Error(ctx, err)
			return nil, err
		}

		response, err := adapter.VideoCreate(ctx, request)
		if err != nil {
			logger.Error(ctx, err)
			modelAgentTestModelRes.Error = err.Error()
		}

		modelAgentTestModelRes.Result = err == nil
		modelAgentTestModelRes.TotalTime = response.TotalTime

	default:

		response, err := adapter.ChatCompletions(ctx, data.Bytes())
		if err != nil {
			logger.Error(ctx, err)
			modelAgentTestModelRes.Error = err.Error()
		}

		modelAgentTestModelRes.Result = err == nil
		modelAgentTestModelRes.TotalTime = response.TotalTime

		if response.Error != nil && modelAgentTestModelRes.Error == "" {
			modelAgentTestModelRes.Result = false
			modelAgentTestModelRes.Error = response.Error.Error()
		}
	}

	return modelAgentTestModelRes, nil
}
