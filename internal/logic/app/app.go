package app

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/core"
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
		AppId:        core.IncrAppId(ctx),
		Name:         params.Name,
		Type:         params.Type,
		Models:       params.Models,
		IsLimitQuota: params.IsLimitQuota,
		Quota:        params.Quota,
		IpWhitelist:  gstr.Split(gstr.Trim(params.IpWhitelist), "\n"),
		IpBlacklist:  gstr.Split(gstr.Trim(params.IpBlacklist), "\n"),
		Remark:       params.Remark,
		Status:       params.Status,
		UserId:       service.Session().GetUserId(ctx),
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新应用
func (s *sApp) Update(ctx context.Context, params model.AppUpdateReq) error {

	oldData, err := dao.App.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	app, err := dao.App.FindOneAndUpdateById(ctx, params.Id, &do.App{
		Name:         params.Name,
		Type:         params.Type,
		Models:       params.Models,
		IsLimitQuota: params.IsLimitQuota,
		Quota:        params.Quota,
		IpWhitelist:  gstr.Split(gstr.Trim(params.IpWhitelist), "\n"),
		IpBlacklist:  gstr.Split(gstr.Trim(params.IpBlacklist), "\n"),
		Remark:       params.Remark,
		Status:       params.Status,
	})

	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	fields := g.Map{
		fmt.Sprintf(consts.APP_TOTAL_TOKENS_FIELD, app.AppId):   app.Quota,
		fmt.Sprintf(consts.APP_IS_LIMIT_QUOTA_FIELD, app.AppId): app.IsLimitQuota,
	}

	if _, err = redis.HSet(ctx, fmt.Sprintf(consts.API_USAGE_KEY, app.UserId), fields); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: app,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改应用状态
func (s *sApp) ChangeStatus(ctx context.Context, params model.AppChangeStatusReq) error {

	app, err := dao.App.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: app,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除应用
func (s *sApp) Delete(ctx context.Context, id string) error {

	app, err := dao.App.FindOneAndDeleteById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.Key.DeleteMany(ctx, bson.M{"app_id": app.AppId}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: app,
	}); err != nil {
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

	modelNames, err := service.Model().ModelNames(ctx, app.Models)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.App{
		Id:           app.Id,
		AppId:        app.AppId,
		Name:         app.Name,
		Type:         app.Type,
		Models:       app.Models,
		ModelNames:   modelNames,
		IsLimitQuota: app.IsLimitQuota,
		Quota:        app.Quota,
		IpWhitelist:  app.IpWhitelist,
		IpBlacklist:  app.IpBlacklist,
		Remark:       app.Remark,
		Status:       app.Status,
		Creator:      app.Creator,
		Updater:      app.Updater,
		CreatedAt:    util.FormatDatetime(app.CreatedAt),
		UpdatedAt:    util.FormatDatetime(app.UpdatedAt),
	}, nil
}

// 应用分页列表
func (s *sApp) Page(ctx context.Context, params model.AppPageReq) (*model.AppPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	if params.Name != "" {
		filter["name"] = params.Name
	}

	if len(params.Models) > 0 {
		filter["models"] = bson.M{
			"$in": params.Models,
		}
	}

	results, err := dao.App.FindByPage(ctx, paging, filter, "-updated_at")
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

	items := make([]*model.App, 0)
	for _, result := range results {

		modelNames := make([]string, 0)
		for _, id := range result.Models {
			if modelMap[id] != nil {
				modelNames = append(modelNames, modelMap[id].Name)
			}
		}

		items = append(items, &model.App{
			Id:           result.Id,
			AppId:        result.AppId,
			Name:         result.Name,
			Type:         result.Type,
			Models:       result.Models,
			ModelNames:   modelNames,
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

	return &model.AppPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 应用列表
func (s *sApp) List(ctx context.Context, params model.AppListReq) ([]*model.App, error) {

	filter := bson.M{}

	results, err := dao.App.Find(ctx, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.App, 0)
	for _, result := range results {
		items = append(items, &model.App{
			Id:           result.Id,
			AppId:        result.AppId,
			Name:         result.Name,
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

	return items, nil
}

// 新建应用密钥
func (s *sApp) CreateKey(ctx context.Context, params model.AppCreateKeyReq) (string, error) {
	// 警告: 固定前缀, 修改请慎重, 可能会引发不可预知问题!!!
	return util.NewKey("sk-FastAPI", 51, gconv.String(service.Session().GetUserId(ctx)), gconv.String(params.AppId)), nil
}

// 应用密钥配置
func (s *sApp) KeyConfig(ctx context.Context, params model.AppKeyConfigReq) (err error) {

	key := &do.Key{
		UserId:       service.Session().GetUserId(ctx), // todo
		AppId:        params.AppId,
		Key:          params.Key,
		IsLimitQuota: params.IsLimitQuota,
		Quota:        params.Quota,
		Type:         1,
		Models:       params.Models,
		IpWhitelist:  gstr.Split(gstr.Trim(params.IpWhitelist), "\n"),
		IpBlacklist:  gstr.Split(gstr.Trim(params.IpBlacklist), "\n"),
		Remark:       params.Remark,
		Status:       params.Status,
	}

	var keyInfo *entity.Key
	var oldData *entity.Key
	action := consts.ACTION_CREATE

	if params.Id != "" {

		action = consts.ACTION_UPDATE
		if oldData, err = dao.Key.FindById(ctx, params.Id); err != nil {
			logger.Error(ctx, err)
			return err
		}

		key.AppId = 0
		key.Key = ""
		if keyInfo, err = dao.Key.FindOneAndUpdateById(ctx, params.Id, key); err != nil {
			logger.Error(ctx, err)
			return err
		}

	} else {

		id, err := dao.Key.Insert(ctx, key)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		keyInfo = &entity.Key{
			Id:           id,
			UserId:       key.UserId,
			AppId:        key.AppId,
			Key:          key.Key,
			IsLimitQuota: key.IsLimitQuota,
			Quota:        key.Quota,
			Type:         key.Type,
			Models:       key.Models,
			IpWhitelist:  key.IpWhitelist,
			IpBlacklist:  key.IpBlacklist,
			Remark:       key.Remark,
			Status:       key.Status,
		}
	}

	app, err := dao.App.FindByAppId(ctx, keyInfo.AppId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	fields := g.Map{
		fmt.Sprintf(consts.KEY_TOTAL_TOKENS_FIELD, keyInfo.AppId, keyInfo.Key):   key.Quota,
		fmt.Sprintf(consts.KEY_IS_LIMIT_QUOTA_FIELD, keyInfo.AppId, keyInfo.Key): key.IsLimitQuota,
	}

	if _, err = redis.HSet(ctx, fmt.Sprintf(consts.API_USAGE_KEY, app.UserId), fields); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP_KEY, model.PubMessage{
		Action:  action,
		OldData: oldData,
		NewData: keyInfo,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
