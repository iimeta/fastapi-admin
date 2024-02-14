package key

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

type sKey struct{}

func init() {
	service.RegisterKey(New())
}

func New() service.IKey {
	return &sKey{}
}

// 新建密钥
func (s *sKey) Create(ctx context.Context, params model.KeyCreateReq) error {

	keys := gstr.Split(gstr.Trim(params.Key), "\n")

	keyList, err := dao.Key.Find(ctx, bson.M{"key": bson.M{"$in": keys}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	keyMap := util.ToMap(keyList, func(t *entity.Key) string {
		return t.Key
	})

	for _, k := range keys {

		key := keyMap[k]

		if key == nil {
			if _, err := dao.Key.Insert(ctx, &do.Key{
				Corp:        params.Corp,
				Key:         k,
				Type:        2,
				Models:      params.Models,
				ModelAgents: params.ModelAgents,
				Remark:      params.Remark,
				Status:      params.Status,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		} else {

			modelSet := gset.NewStrSet()
			modelSet.Add(key.Models...)
			modelSet.Add(params.Models...)

			modelAgentSet := gset.NewStrSet()
			modelAgentSet.Add(key.ModelAgents...)
			modelAgentSet.Add(params.ModelAgents...)

			if err := s.Update(ctx, model.KeyUpdateReq{
				Id:          key.Id,
				Corp:        params.Corp,
				Key:         params.Key,
				Models:      modelSet.Slice(),
				ModelAgents: modelAgentSet.Slice(),
				Remark:      params.Remark,
				Status:      params.Status,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 更新密钥
func (s *sKey) Update(ctx context.Context, params model.KeyUpdateReq) error {

	if err := dao.Key.UpdateById(ctx, params.Id, &do.Key{
		Corp:        params.Corp,
		Key:         params.Key,
		Models:      params.Models,
		ModelAgents: params.ModelAgents,
		Remark:      params.Remark,
		Status:      params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改密钥状态
func (s *sKey) ChangeStatus(ctx context.Context, params model.KeyChangeStatusReq) error {

	if err := dao.Key.UpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除密钥
func (s *sKey) Delete(ctx context.Context, id string) error {

	if _, err := dao.Key.DeleteById(ctx, id); err != nil {
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

	modelNames := make([]string, 0)
	if len(key.Models) > 0 {

		models, err := service.Model().List(ctx, model.ModelListReq{})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		modelMap := util.ToMap(models, func(t *model.Model) string {
			return t.Id
		})

		for _, id := range key.Models {
			if modelMap[id] != nil {
				modelNames = append(modelNames, modelMap[id].Name)
			}
		}
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
		Id:              key.Id,
		AppId:           key.AppId,
		Corp:            key.Corp,
		Key:             key.Key,
		Type:            key.Type,
		Models:          key.Models,
		ModelNames:      modelNames,
		ModelAgents:     key.ModelAgents,
		ModelAgentNames: modelAgentNames,
		IsLimitQuota:    key.IsLimitQuota,
		Quota:           key.Quota,
		IpWhitelist:     key.IpWhitelist,
		IpBlacklist:     key.IpBlacklist,
		Remark:          key.Remark,
		Status:          key.Status,
		Creator:         key.Creator,
		Updater:         key.Updater,
		CreatedAt:       util.FormatDatetime(key.CreatedAt),
		UpdatedAt:       util.FormatDatetime(key.UpdatedAt),
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

	if params.Corp != "" {
		filter["corp"] = params.Corp
	}

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	if params.Key != "" {
		filter["key"] = params.Key
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

	results, err := dao.Key.FindByPage(ctx, paging, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	models, err := service.Model().List(ctx, model.ModelListReq{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelMap := util.ToMap(models, func(t *model.Model) string {
		return t.Id
	})

	items := make([]*model.Key, 0)
	for _, result := range results {

		modelNames := make([]string, 0)
		for _, id := range result.Models {
			if modelMap[id] != nil {
				modelNames = append(modelNames, modelMap[id].Name)
			}
		}

		items = append(items, &model.Key{
			Id:           result.Id,
			AppId:        result.AppId,
			Corp:         result.Corp,
			Key:          result.Key,
			Type:         result.Type,
			Models:       result.Models,
			ModelNames:   modelNames,
			ModelAgents:  result.ModelAgents,
			IsLimitQuota: result.IsLimitQuota,
			Quota:        result.Quota,
			IpWhitelist:  result.IpWhitelist,
			IpBlacklist:  result.IpBlacklist,
			Remark:       result.Remark,
			Status:       result.Status,
			Creator:      result.Creator,
			Updater:      result.Updater,
			CreatedAt:    util.FormatDatetime(result.CreatedAt),
			UpdatedAt:    util.FormatDatetime(result.UpdatedAt),
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

	results, err := dao.Key.Find(ctx, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Key, 0)
	for _, result := range results {
		items = append(items, &model.Key{
			Id:          result.Id,
			Corp:        result.Corp,
			Key:         result.Key,
			Type:        result.Type,
			Models:      result.Models,
			ModelAgents: result.ModelAgents,
			Remark:      result.Remark,
			Status:      result.Status,
			Creator:     result.Creator,
			Updater:     result.Updater,
			CreatedAt:   util.FormatDatetime(result.CreatedAt),
			UpdatedAt:   util.FormatDatetime(result.UpdatedAt),
		})
	}

	return items, nil
}
