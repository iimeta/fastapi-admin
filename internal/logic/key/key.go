package key

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
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

	for _, key := range keys {
		if _, err := dao.Key.Insert(ctx, &do.Key{
			Corp:   params.Corp,
			Key:    key,
			Type:   2,
			Models: params.Models,
			Remark: params.Remark,
			Status: params.Status,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

// 更新密钥
func (s *sKey) Update(ctx context.Context, params model.KeyUpdateReq) error {

	if err := dao.Key.UpdateById(ctx, params.Id, &do.Key{
		Corp:   params.Corp,
		Key:    params.Key,
		Models: params.Models,
		Remark: params.Remark,
		Status: params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除密钥
func (s *sKey) Delete(ctx context.Context, id string) error {

	if err := dao.Key.DeleteById(ctx, id); err != nil {
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

	return &model.Key{
		Id:           key.Id,
		AppId:        key.AppId,
		Corp:         key.Corp,
		Key:          key.Key,
		Type:         key.Type,
		Models:       key.Models,
		IsLimitQuota: key.IsLimitQuota,
		Quota:        key.Quota,
		IpWhitelist:  key.IpWhitelist,
		IpBlacklist:  key.IpBlacklist,
		Remark:       key.Remark,
		Status:       key.Status,
		Creator:      key.Creator,
		Updater:      key.Updater,
		CreatedAt:    util.FormatDatetime(key.CreatedAt),
		UpdatedAt:    util.FormatDatetime(key.UpdatedAt),
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

	results, err := dao.Key.FindByPage(ctx, paging, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Key, 0)
	for _, result := range results {
		items = append(items, &model.Key{
			Id:           result.Id,
			AppId:        result.AppId,
			Corp:         result.Corp,
			Key:          result.Key,
			Type:         result.Type,
			Models:       result.Models,
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
			Id:        result.Id,
			Corp:      result.Corp,
			Key:       result.Key,
			Type:      result.Type,
			Models:    result.Models,
			Remark:    result.Remark,
			Status:    result.Status,
			Creator:   result.Creator,
			Updater:   result.Updater,
			CreatedAt: util.FormatDatetime(result.CreatedAt),
			UpdatedAt: util.FormatDatetime(result.UpdatedAt),
		})
	}

	return items, nil
}
