package app

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/core"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type sApp struct{}

func init() {
	service.RegisterApp(New())
}

func New() service.IApp {
	return &sApp{}
}

// 新建应用
func (s *sApp) Create(ctx context.Context, params model.AppCreateReq) (string, error) {

	userId := service.Session().GetUserId(ctx)
	rid := service.Session().GetRid(ctx)

	if service.Session().IsResellerRole(ctx) || service.Session().IsAdminRole(ctx) {

		if params.UserId == 0 {
			return "", errors.New("请输入用户ID")
		}

		user, err := service.User().GetUserByUserId(ctx, params.UserId)
		if err != nil {
			logger.Error(ctx, err)
			return "", errors.New("用户ID不存在, 请重新输入")
		}

		userId = user.UserId
		rid = user.Rid
	}

	appId := core.IncrAppId(ctx)

	if _, err := dao.App.Insert(ctx, &do.App{
		AppId:          appId,
		Name:           params.Name,
		Models:         params.Models,
		IsLimitQuota:   params.IsLimitQuota,
		Quota:          common.ConvQuotaUnit(params.Quota),
		QuotaExpiresAt: util.ConvTimestampMilli(params.QuotaExpiresAt),
		IsBindGroup:    params.IsBindGroup,
		Group:          params.Group,
		IpWhitelist:    gstr.Split(gstr.Trim(params.IpWhitelist), "\n"),
		IpBlacklist:    gstr.Split(gstr.Trim(params.IpBlacklist), "\n"),
		Remark:         params.Remark,
		Status:         params.Status,
		UserId:         userId,
		Rid:            rid,
	}); err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if params.IsCreateKey {

		key, err := service.AppKey().Create(ctx, model.AppKeyCreateReq{UserId: userId, AppId: appId})
		if err != nil {
			logger.Error(ctx, err)
			return "", err
		}

		if _, err = service.AppKey().Config(ctx, model.AppKeyConfigReq{UserId: userId, AppId: appId, Key: key, BillingMethods: []int{1}, Status: 1}); err != nil {
			logger.Error(ctx, err)
			return "", err
		}

		return key, nil
	}

	return "", nil
}

// 更新应用
func (s *sApp) Update(ctx context.Context, params model.AppUpdateReq) error {

	oldData, err := dao.App.FindById(ctx, params.Id)
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

	app, err := dao.App.FindOneAndUpdateById(ctx, params.Id, &do.App{
		Name:           params.Name,
		Models:         params.Models,
		IsLimitQuota:   params.IsLimitQuota,
		Quota:          common.ConvQuotaUnit(params.Quota),
		QuotaExpiresAt: util.ConvTimestampMilli(params.QuotaExpiresAt),
		IsBindGroup:    params.IsBindGroup,
		Group:          params.Group,
		IpWhitelist:    gstr.Split(gstr.Trim(params.IpWhitelist), "\n"),
		IpBlacklist:    gstr.Split(gstr.Trim(params.IpBlacklist), "\n"),
		Remark:         params.Remark,
		Status:         params.Status,
	})

	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	fields := g.Map{
		fmt.Sprintf(consts.APP_QUOTA_FIELD, app.AppId):          app.Quota,
		fmt.Sprintf(consts.APP_IS_LIMIT_QUOTA_FIELD, app.AppId): app.IsLimitQuota,
	}

	if _, err = redis.HSet(ctx, fmt.Sprintf(consts.API_USER_USAGE_KEY, app.UserId), fields); err != nil {
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

	if service.Session().IsResellerRole(ctx) {

		app, err := dao.App.FindById(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if app.Rid != service.Session().GetRid(ctx) {
			return errors.New("Unauthorized")
		}
	}

	if service.Session().IsUserRole(ctx) {

		app, err := dao.App.FindById(ctx, params.Id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if app.UserId != service.Session().GetUserId(ctx) {
			return errors.New("Unauthorized")
		}
	}

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

	if service.Session().IsResellerRole(ctx) {

		app, err := dao.App.FindById(ctx, id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if app.Rid != service.Session().GetRid(ctx) {
			return errors.New("Unauthorized")
		}
	}

	if service.Session().IsUserRole(ctx) {

		app, err := dao.App.FindById(ctx, id)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if app.UserId != service.Session().GetUserId(ctx) {
			return errors.New("Unauthorized")
		}
	}

	app, err := dao.App.FindOneAndDeleteById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	keys, err := dao.Key.Find(ctx, bson.M{"app_id": app.AppId})
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

	for _, key := range keys {
		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP_KEY, model.PubMessage{
			Action:  consts.ACTION_DELETE,
			OldData: key,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
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

	if service.Session().IsResellerRole(ctx) && app.Rid != service.Session().GetRid(ctx) {
		return nil, errors.New("Unauthorized")
	}

	if service.Session().IsUserRole(ctx) && app.UserId != service.Session().GetUserId(ctx) {
		return nil, errors.New("Unauthorized")
	}

	modelNames, err := service.Model().ModelNames(ctx, app.Models)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	groupName := ""
	if app.IsBindGroup && app.Group != "" {
		group, err := dao.Group.FindById(ctx, app.Group)
		if err != nil {
			logger.Error(ctx, err)
		} else {
			groupName = group.Name
		}
	}

	return &model.App{
		Id:             app.Id,
		AppId:          app.AppId,
		Name:           app.Name,
		Models:         app.Models,
		ModelNames:     modelNames,
		IsLimitQuota:   app.IsLimitQuota,
		Quota:          common.ConvQuotaUnitReverse(app.Quota),
		UsedQuota:      common.ConvQuotaUnitReverse(app.UsedQuota),
		QuotaExpiresAt: util.FormatDateTime(app.QuotaExpiresAt),
		IsBindGroup:    app.IsBindGroup,
		Group:          app.Group,
		GroupName:      groupName,
		IpWhitelist:    app.IpWhitelist,
		IpBlacklist:    app.IpBlacklist,
		Remark:         app.Remark,
		Status:         app.Status,
		UserId:         app.UserId,
		CreatedAt:      util.FormatDateTime(app.CreatedAt),
		UpdatedAt:      util.FormatDateTime(app.UpdatedAt),
	}, nil
}

// 应用分页列表
func (s *sApp) Page(ctx context.Context, params model.AppPageReq) (*model.AppPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.AppKey != "" {

		userId, appId, err := common.ParseSecretKey(ctx, params.AppKey)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		params.UserId = userId
		params.AppId = appId
	}

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

	if params.Name != "" {
		filter["name"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Name),
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

	if len(params.QuotaExpiresAt) > 0 {
		gte := gtime.NewFromStrFormat(params.QuotaExpiresAt[0], time.DateOnly).StartOfDay().TimestampMilli()
		lte := gtime.NewFromStrLayout(params.QuotaExpiresAt[1], time.DateOnly).EndOfDay(true).TimestampMilli()
		filter["quota_expires_at"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.App.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-updated_at"}})
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
			Id:             result.Id,
			AppId:          result.AppId,
			Name:           result.Name,
			Models:         result.Models,
			ModelNames:     modelNames,
			IsLimitQuota:   result.IsLimitQuota,
			Quota:          common.ConvQuotaUnitReverse(result.Quota),
			UsedQuota:      common.ConvQuotaUnitReverse(result.UsedQuota),
			QuotaExpiresAt: util.FormatDateTime(result.QuotaExpiresAt),
			IsBindGroup:    result.IsBindGroup,
			Group:          result.Group,
			Status:         result.Status,
			UserId:         result.UserId,
			CreatedAt:      util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:      util.FormatDateTimeMonth(result.UpdatedAt),
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

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	results, err := dao.App.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"status", "-updated_at"}})
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
			IsLimitQuota: result.IsLimitQuota,
			Quota:        common.ConvQuotaUnitReverse(result.Quota),
			UsedQuota:    common.ConvQuotaUnitReverse(result.UsedQuota),
			Status:       result.Status,
		})
	}

	return items, nil
}

// 应用模型权限
func (s *sApp) Models(ctx context.Context, params model.AppModelsReq) error {

	oldData, err := dao.App.FindOne(ctx, bson.M{"app_id": params.AppId})
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

	newData, err := dao.App.FindOneAndUpdate(ctx, bson.M{"app_id": params.AppId}, bson.M{
		"models": params.Models,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 应用绑定分组
func (s *sApp) Group(ctx context.Context, params model.AppGroupReq) error {

	oldData, err := dao.App.FindOne(ctx, bson.M{"app_id": params.AppId})
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

	newData, err := dao.App.FindOneAndUpdate(ctx, bson.M{"app_id": params.AppId}, bson.M{
		"is_bind_group": params.IsBindGroup,
		"group":         params.Group,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_APP, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 应用批量操作
func (s *sApp) BatchOperate(ctx context.Context, params model.AppBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_STATUS:
		for _, id := range params.Ids {
			if err := s.ChangeStatus(ctx, model.AppChangeStatusReq{
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
