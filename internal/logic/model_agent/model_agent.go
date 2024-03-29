package model

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sModelAgent struct{}

func init() {
	service.RegisterModelAgent(New())
}

func New() service.IModelAgent {
	return &sModelAgent{}
}

// 新建模型代理
func (s *sModelAgent) Create(ctx context.Context, params model.ModelAgentCreateReq) error {

	if s.IsNameExist(ctx, params.Name) {
		return errors.Newf("模型代理名称 \"%s\" 已存在", params.Name)
	}

	id, err := dao.ModelAgent.Insert(ctx, &do.ModelAgent{
		Name:    gstr.Trim(params.Name),
		BaseUrl: params.BaseUrl,
		Path:    params.Path,
		Weight:  params.Weight,
		Remark:  params.Remark,
		Status:  params.Status,
	})

	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if len(params.Models) > 0 {
		if err = dao.Model.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": params.Models}}, bson.M{
			"$push": bson.M{
				"model_agents": id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if params.Key != "" {
		if err = service.Key().Create(ctx, model.KeyCreateReq{
			Corp:        "OpenAI",
			Key:         params.Key,
			Models:      params.Models,
			ModelAgents: []string{id},
			Remark:      params.Remark,
			Status:      params.Status,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	modelAgent, err := s.Detail(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
		Action:  consts.ACTION_CREATE,
		NewData: modelAgent,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, id := range modelAgent.Models {

		newData, err := dao.Model.FindById(ctx, id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_MODEL, model.PubMessage{
			Action:  consts.ACTION_UPDATE,
			NewData: newData,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
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
		Name:    gstr.Trim(params.Name),
		BaseUrl: params.BaseUrl,
		Path:    params.Path,
		Weight:  params.Weight,
		Remark:  params.Remark,
		Status:  params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := dao.Model.UpdateMany(ctx, bson.M{"model_agents": bson.M{"$in": []string{params.Id}}}, bson.M{
		"$pull": bson.M{
			"model_agents": params.Id,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if len(params.Models) > 0 {
		if err := dao.Model.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": params.Models}}, bson.M{
			"$addToSet": bson.M{
				"model_agents": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
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

	if err := dao.Key.UpdateMany(ctx, bson.M{"model_agents": bson.M{"$in": []string{params.Id}}}, bson.M{
		"$pull": bson.M{
			"model_agents": params.Id,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if params.Key != "" {
		if err := service.Key().Create(ctx, model.KeyCreateReq{
			Corp:        "OpenAI",
			Key:         params.Key,
			Models:      params.Models,
			ModelAgents: []string{params.Id},
			Remark:      params.Remark,
			Status:      params.Status,
		}); err != nil {
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

	modelSet := gset.NewStrSet()
	modelSet.Add(oldData.Models...)
	modelSet.Add(modelAgent.Models...)

	for _, id := range modelSet.Slice() {

		newData, err := dao.Model.FindById(ctx, id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_MODEL, model.PubMessage{
			Action:  consts.ACTION_UPDATE,
			NewData: newData,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

// 更改模型代理状态
func (s *sModelAgent) ChangeStatus(ctx context.Context, params model.ModelAgentChangeStatusReq) error {

	if err := dao.ModelAgent.UpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
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

	if err := dao.Model.UpdateMany(ctx, bson.M{"model_agents": bson.M{"$in": []string{id}}}, bson.M{
		"$pull": bson.M{
			"model_agents": id,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := dao.Key.UpdateMany(ctx, bson.M{"model_agents": bson.M{"$in": []string{id}}}, bson.M{
		"$pull": bson.M{
			"model_agents": id,
		},
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, id := range modelAgent.Models {

		newData, err := dao.Model.FindById(ctx, id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_MODEL, model.PubMessage{
			Action:  consts.ACTION_UPDATE,
			NewData: newData,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_AGENT, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: modelAgent,
	}); err != nil {
		logger.Error(ctx, err)
		return err
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

	modelList, err := dao.Model.Find(ctx, bson.M{"model_agents": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	models := make([]string, 0)
	modelNames := make([]string, 0)

	for _, model := range modelList {
		models = append(models, model.Id)
		modelNames = append(modelNames, model.Name)
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

	return &model.ModelAgent{
		Id:         modelAgent.Id,
		Name:       modelAgent.Name,
		BaseUrl:    modelAgent.BaseUrl,
		Path:       modelAgent.Path,
		Weight:     modelAgent.Weight,
		Models:     models,
		ModelNames: modelNames,
		Key:        gstr.Join(keys, "\n"),
		Remark:     modelAgent.Remark,
		Status:     modelAgent.Status,
		Creator:    modelAgent.Creator,
		Updater:    modelAgent.Updater,
		CreatedAt:  util.FormatDatetime(modelAgent.CreatedAt),
		UpdatedAt:  util.FormatDatetime(modelAgent.UpdatedAt),
	}, nil
}

// 模型代理分页列表
func (s *sModelAgent) Page(ctx context.Context, params model.ModelAgentPageReq) (*model.ModelAgentPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.Name != "" {
		filter["name"] = params.Name
	}

	if params.BaseUrl != "" {
		filter["base_url"] = params.BaseUrl
	}

	results, err := dao.ModelAgent.FindByPage(ctx, paging, filter, "-updated_at")
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
			Name:       result.Name,
			BaseUrl:    result.BaseUrl,
			Path:       result.Path,
			Weight:     result.Weight,
			Models:     modelMap[result.Id],
			ModelNames: modelNameMap[result.Id],
			Remark:     result.Remark,
			Status:     result.Status,
			Creator:    result.Creator,
			Updater:    result.Updater,
			CreatedAt:  util.FormatDatetime(result.CreatedAt),
			UpdatedAt:  util.FormatDatetime(result.UpdatedAt),
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

	results, err := dao.ModelAgent.Find(ctx, filter, "-updated_at")
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
			Name:       result.Name,
			BaseUrl:    result.BaseUrl,
			Path:       result.Path,
			Weight:     result.Weight,
			Models:     modelMap[result.Id],
			ModelNames: modelNameMap[result.Id],
			Remark:     result.Remark,
			Status:     result.Status,
			Creator:    result.Creator,
			Updater:    result.Updater,
			CreatedAt:  util.FormatDatetime(result.CreatedAt),
			UpdatedAt:  util.FormatDatetime(result.UpdatedAt),
		})
	}

	return items, nil
}

// 模型代理名称是否存在
func (s *sModelAgent) IsNameExist(ctx context.Context, name string, id ...string) bool {

	model, err := dao.ModelAgent.FindOne(ctx, bson.M{"name": gstr.Trim(name)})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false
		}
		logger.Error(ctx, err)
		return true
	}

	if model != nil {
		if len(id) > 0 && model.Id == id[0] {
			return false
		}
		return true
	}

	return false
}
