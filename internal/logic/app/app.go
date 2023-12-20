package model

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/internal/core"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
)

type sApp struct{}

func init() {
	service.RegisterApp(New())
}

func New() service.IApp {
	return &sApp{}
}

// 新建应用
func (s *sApp) Create(ctx context.Context, params model.AppCreateReq) error {

	if _, err := dao.App.Insert(ctx, &do.App{
		AppId:       core.IncrAppId(ctx),
		Name:        params.Name,
		Type:        params.Type,
		Models:      params.Models,
		IpWhitelist: gstr.Split(params.IpWhitelist, "\n"),
		IpBlacklist: gstr.Split(params.IpBlacklist, "\n"),
		Remark:      params.Remark,
		Status:      params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新应用
func (s *sApp) Update(ctx context.Context, params model.AppUpdateReq) error {

	if err := dao.App.UpdateById(ctx, params.Id, &do.App{
		Name:        params.Name,
		Type:        params.Type,
		Models:      params.Models,
		IpWhitelist: params.IpWhitelist,
		IpBlacklist: params.IpBlacklist,
		Remark:      params.Remark,
		Status:      params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除应用
func (s *sApp) Delete(ctx context.Context, id string) error {

	if err := dao.App.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 应用详情
func (s *sApp) Detail(ctx context.Context, id string) (*model.App, error) {

	app, err := dao.App.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.App{
		Id:          app.Id,
		AppId:       app.AppId,
		Name:        app.Name,
		Type:        app.Type,
		Models:      app.Models,
		IpWhitelist: app.IpWhitelist,
		IpBlacklist: app.IpBlacklist,
		Remark:      app.Remark,
		Status:      app.Status,
		Creator:     app.Creator,
		Updater:     app.Updater,
		CreatedAt:   app.CreatedAt,
		UpdatedAt:   app.UpdatedAt,
	}, nil
}

// 应用分页列表
func (s *sApp) Page(ctx context.Context, params model.AppPageReq) (*model.AppPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	results, err := dao.App.FindByPage(ctx, paging, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.App, 0)
	for _, result := range results {
		items = append(items, &model.App{
			Id:          result.Id,
			AppId:       result.AppId,
			Name:        result.Name,
			Type:        result.Type,
			Models:      result.Models,
			IpWhitelist: result.IpWhitelist,
			IpBlacklist: result.IpBlacklist,
			Remark:      result.Remark,
			Status:      result.Status,
			Creator:     result.Creator,
			Updater:     result.Updater,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		})
	}

	return &model.AppPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 新建应用密钥
func (s *sApp) CreateKey(ctx context.Context, params model.AppCreateKeyReq) error {

	if _, err := dao.Key.Insert(ctx, &do.Key{
		AppId:  params.AppId,
		Key:    params.Key,
		Type:   1,
		Models: params.Models,
		Remark: params.Remark,
		Status: params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
