package provider

import (
	"context"
	"regexp"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type sProvider struct{}

func init() {
	service.RegisterProvider(New())
}

func New() service.IProvider {
	return &sProvider{}
}

// 新建提供商
func (s *sProvider) Create(ctx context.Context, params model.ProviderCreateReq) (string, error) {

	if params.Sort == 0 {
		if provider, err := dao.Provider.FindOne(ctx, bson.M{}, &dao.FindOptions{SortFields: []string{"-sort"}}); err == nil && provider != nil {
			params.Sort = provider.Sort + 1
		} else {
			params.Sort = 1
		}
	}

	id, err := dao.Provider.Insert(ctx, &do.Provider{
		Name:     gstr.Trim(params.Name),
		Code:     gstr.Trim(params.Code),
		Sort:     params.Sort,
		IsPublic: params.IsPublic,
		Remark:   params.Remark,
		Status:   params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	provider, err := dao.Provider.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_PROVIDER, model.PubMessage{
		Action:  consts.ACTION_CREATE,
		NewData: provider,
	}); err != nil {
		logger.Error(ctx, err)
	}

	return id, nil
}

// 更新提供商
func (s *sProvider) Update(ctx context.Context, params model.ProviderUpdateReq) error {

	if params.Sort == 0 {
		if provider, err := dao.Provider.FindOne(ctx, bson.M{}, &dao.FindOptions{SortFields: []string{"-sort"}}); err == nil && provider != nil {
			params.Sort = provider.Sort + 1
		} else {
			params.Sort = 1
		}
	}

	oldData, err := dao.Provider.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	provider, err := dao.Provider.FindOneAndUpdateById(ctx, params.Id, &do.Provider{
		Name:     gstr.Trim(params.Name),
		Code:     gstr.Trim(params.Code),
		Sort:     params.Sort,
		IsPublic: params.IsPublic,
		Remark:   params.Remark,
		Status:   params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_PROVIDER, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: provider,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改提供商公开状态
func (s *sProvider) ChangePublic(ctx context.Context, params model.ProviderChangePublicReq) error {

	if err := dao.Provider.UpdateById(ctx, params.Id, bson.M{
		"is_public": params.IsPublic,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改提供商状态
func (s *sProvider) ChangeStatus(ctx context.Context, params model.ProviderChangeStatusReq) error {

	provider, err := dao.Provider.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_PROVIDER, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: provider,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除提供商
func (s *sProvider) Delete(ctx context.Context, id string) error {

	provider, err := dao.Provider.FindOneAndDeleteById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_PROVIDER, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: provider,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 提供商详情
func (s *sProvider) Detail(ctx context.Context, id string) (*model.Provider, error) {

	provider, err := dao.Provider.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.Provider{
		Id:        provider.Id,
		Name:      provider.Name,
		Code:      provider.Code,
		Sort:      provider.Sort,
		IsPublic:  provider.IsPublic,
		Remark:    provider.Remark,
		Status:    provider.Status,
		Creator:   provider.Creator,
		Updater:   provider.Updater,
		CreatedAt: util.FormatDateTime(provider.CreatedAt),
		UpdatedAt: util.FormatDateTime(provider.UpdatedAt),
	}, nil
}

// 提供商分页列表
func (s *sProvider) Page(ctx context.Context, params model.ProviderPageReq) (*model.ProviderPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.Name != "" {
		filter["name"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Name),
		}
	}

	if params.Code != "" {
		filter["code"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Code),
		}
	}

	if params.IsPublic != "" {
		filter["is_public"] = gconv.Bool(params.IsPublic)
	}

	if params.Remark != "" {
		filter["remark"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Remark),
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if len(params.UpdatedAt) > 0 {
		gte := gtime.NewFromStrFormat(params.UpdatedAt[0], time.DateOnly).StartOfDay().TimestampMilli()
		lte := gtime.NewFromStrLayout(params.UpdatedAt[1], time.DateOnly).EndOfDay(true).TimestampMilli()
		filter["updated_at"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.Provider.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Provider, 0)
	for _, result := range results {
		items = append(items, &model.Provider{
			Id:        result.Id,
			Name:      result.Name,
			Code:      result.Code,
			Sort:      result.Sort,
			IsPublic:  result.IsPublic,
			Remark:    result.Remark,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt: util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.ProviderPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 提供商列表
func (s *sProvider) List(ctx context.Context, params model.ProviderListReq) ([]*model.Provider, error) {

	filter := bson.M{
		"status": 1,
	}

	if service.Session().IsUserRole(ctx) {
		filter["is_public"] = true
	}

	results, err := dao.Provider.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"sort", "status", "-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Provider, 0)
	for _, result := range results {
		items = append(items, &model.Provider{
			Id:   result.Id,
			Name: result.Name,
			Code: result.Code,
		})
	}

	return items, nil
}

// 提供商批量操作
func (s *sProvider) BatchOperate(ctx context.Context, params model.ProviderBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_STATUS:
		for _, id := range params.Ids {
			if err := s.ChangeStatus(ctx, model.ProviderChangeStatusReq{
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
