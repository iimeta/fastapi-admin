package group

import (
	"context"
	"regexp"
	"time"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
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
)

type sGroup struct{}

func init() {
	service.RegisterGroup(New())
}

func New() service.IGroup {
	return &sGroup{}
}

// 新建分组
func (s *sGroup) Create(ctx context.Context, params model.GroupCreateReq) (err error) {

	var defaultGroup *entity.Group
	if params.IsDefault {
		defaultGroup, err = dao.Group.FindOne(ctx, bson.M{"is_default": true})
		params.Weight = 99999
	} else {
		if groups, err := dao.Group.Find(ctx, bson.M{"weight": params.Weight}, &dao.FindOptions{SortFields: []string{"-is_default"}}); err == nil && len(groups) > 0 {
			if len(groups) > 1 || !groups[0].IsDefault {
				return errors.Newf("排序: %d 已被占用, 请重新输入排序", params.Weight)
			}
		}
	}

	id, err := dao.Group.Insert(ctx, &do.Group{
		Name:               params.Name,
		Discount:           params.Discount,
		Models:             params.Models,
		IsEnableModelAgent: params.IsEnableModelAgent,
		LbStrategy:         params.LbStrategy,
		ModelAgents:        params.ModelAgents,
		IsDefault:          params.IsDefault,
		IsLimitQuota:       params.IsLimitQuota,
		Quota:              params.Quota,
		UsedQuota:          params.UsedQuota,
		IsEnableForward:    params.IsEnableForward,
		ForwardConfig:      params.ForwardConfig,
		IsPublic:           params.IsPublic,
		Weight:             params.Weight,
		ExpiresAt:          util.ConvTimestampMilli(params.ExpiresAt),
		Remark:             params.Remark,
		Status:             params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	group, err := dao.Group.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if group.IsLimitQuota {
		if _, err = redis.HSetInt(ctx, consts.API_GROUP_USAGE_KEY, group.Id, group.Quota); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
		Action:  consts.ACTION_CREATE,
		NewData: group,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if defaultGroup != nil {

		group, err := dao.Group.FindOneAndUpdateById(ctx, defaultGroup.Id, bson.M{
			"is_default": false,
		})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
			Action:  consts.ACTION_UPDATE,
			OldData: defaultGroup,
			NewData: group,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if params.IsPublic {

		resellerList, err := dao.Reseller.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.Reseller.UpdateMany(ctx, bson.M{}, bson.M{
			"$push": bson.M{
				"groups": id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, reseller := range resellerList {

			newResellerData := *reseller

			newResellerData.Groups = append(newResellerData.Groups, id)

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: reseller,
				NewData: newResellerData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		userList, err := dao.User.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.User.UpdateMany(ctx, bson.M{}, bson.M{
			"$push": bson.M{
				"groups": id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, user := range userList {

			newUserData := *user

			newUserData.Groups = append(newUserData.Groups, id)

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: user,
				NewData: newUserData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 更新分组
func (s *sGroup) Update(ctx context.Context, params model.GroupUpdateReq) error {

	oldData, err := dao.Group.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	var defaultGroup *entity.Group
	if params.IsDefault && !oldData.IsDefault {
		defaultGroup, err = dao.Group.FindOne(ctx, bson.M{"is_default": true})
		params.Weight = 99999
	} else if !params.IsDefault {
		if groups, err := dao.Group.Find(ctx, bson.M{"weight": params.Weight}, &dao.FindOptions{SortFields: []string{"-is_default"}}); err == nil && len(groups) > 0 {
			if len(groups) > 2 || (len(groups) > 1 && ((!groups[0].IsDefault && groups[0].Id != oldData.Id) || groups[1].Id != oldData.Id)) || (len(groups) == 1 && groups[0].Id != oldData.Id) {
				return errors.Newf("排序: %d 已被占用, 请重新输入排序", params.Weight)
			}
		}
	}

	group := &do.Group{
		Name:               params.Name,
		Discount:           params.Discount,
		Models:             params.Models,
		IsEnableModelAgent: params.IsEnableModelAgent,
		LbStrategy:         params.LbStrategy,
		ModelAgents:        params.ModelAgents,
		IsDefault:          params.IsDefault,
		IsLimitQuota:       params.IsLimitQuota,
		Quota:              params.Quota,
		UsedQuota:          params.UsedQuota,
		IsEnableForward:    params.IsEnableForward,
		ForwardConfig:      params.ForwardConfig,
		IsPublic:           params.IsPublic,
		Weight:             params.Weight,
		ExpiresAt:          util.ConvTimestampMilli(params.ExpiresAt),
		Remark:             params.Remark,
		Status:             params.Status,
	}

	newData, err := dao.Group.FindOneAndUpdateById(ctx, params.Id, group)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if newData.IsLimitQuota {
		if _, err = redis.HSetInt(ctx, consts.API_GROUP_USAGE_KEY, newData.Id, newData.Quota); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	// 旧数据是公开, 新数据改为了私有
	if oldData.IsPublic && !newData.IsPublic {

		resellerList, err := dao.Reseller.Find(ctx, bson.M{"groups": bson.M{"$in": []string{params.Id}}})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.Reseller.UpdateMany(ctx, bson.M{"groups": bson.M{"$in": []string{params.Id}}}, bson.M{
			"$pull": bson.M{
				"groups": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, reseller := range resellerList {

			newResellerData := *reseller

			for i, id := range newResellerData.Groups {
				if id == params.Id {
					newResellerData.Groups = append(newResellerData.Groups[:i], newResellerData.Groups[i+1:]...)
					break
				}
			}

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: reseller,
				NewData: newResellerData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		userList, err := dao.User.Find(ctx, bson.M{"groups": bson.M{"$in": []string{params.Id}}})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.User.UpdateMany(ctx, bson.M{"groups": bson.M{"$in": []string{params.Id}}}, bson.M{
			"$pull": bson.M{
				"groups": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, user := range userList {

			newUserData := *user

			for i, id := range newUserData.Groups {
				if id == params.Id {
					newUserData.Groups = append(newUserData.Groups[:i], newUserData.Groups[i+1:]...)
					break
				}
			}

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: user,
				NewData: newUserData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

	} else if !oldData.IsPublic && newData.IsPublic { // 旧数据是私有, 新数据改为了公开

		resellerList, err := dao.Reseller.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.Reseller.UpdateMany(ctx, bson.M{}, bson.M{
			"$addToSet": bson.M{
				"groups": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, reseller := range resellerList {

			newResellerData := *reseller

			newResellerData.Groups = gset.NewStrSetFrom(append(newResellerData.Groups, params.Id)).Slice()

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: reseller,
				NewData: newResellerData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		userList, err := dao.User.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.User.UpdateMany(ctx, bson.M{}, bson.M{
			"$addToSet": bson.M{
				"groups": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, user := range userList {

			newUserData := *user

			newUserData.Groups = gset.NewStrSetFrom(append(newUserData.Groups, params.Id)).Slice()

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: user,
				NewData: newUserData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if defaultGroup != nil {

		group, err := dao.Group.FindOneAndUpdateById(ctx, defaultGroup.Id, bson.M{
			"is_default": false,
		})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
			Action:  consts.ACTION_UPDATE,
			OldData: defaultGroup,
			NewData: group,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

// 更改过期时间
func (s *sGroup) ChangeExpire(ctx context.Context, params model.GroupChangeExpireReq) error {

	oldData, err := dao.Group.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	newData, err := dao.Group.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"expires_at": util.ConvTimestampMilli(params.ExpiresAt),
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改分组公开状态
func (s *sGroup) ChangePublic(ctx context.Context, params model.GroupChangePublicReq) error {

	oldData, err := dao.Group.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	newData, err := dao.Group.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"is_public": params.IsPublic,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 旧数据是公开, 新数据改为了私有
	if oldData.IsPublic && !newData.IsPublic {

		resellerList, err := dao.Reseller.Find(ctx, bson.M{"groups": bson.M{"$in": []string{params.Id}}})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.Reseller.UpdateMany(ctx, bson.M{"groups": bson.M{"$in": []string{params.Id}}}, bson.M{
			"$pull": bson.M{
				"groups": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, reseller := range resellerList {

			newResellerData := *reseller

			for i, id := range newResellerData.Groups {
				if id == params.Id {
					newResellerData.Groups = append(newResellerData.Groups[:i], newResellerData.Groups[i+1:]...)
					break
				}
			}

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: reseller,
				NewData: newResellerData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		userList, err := dao.User.Find(ctx, bson.M{"groups": bson.M{"$in": []string{params.Id}}})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.User.UpdateMany(ctx, bson.M{"groups": bson.M{"$in": []string{params.Id}}}, bson.M{
			"$pull": bson.M{
				"groups": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, user := range userList {

			newUserData := *user

			for i, id := range newUserData.Groups {
				if id == params.Id {
					newUserData.Groups = append(newUserData.Groups[:i], newUserData.Groups[i+1:]...)
					break
				}
			}

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: user,
				NewData: newUserData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

	} else if !oldData.IsPublic && newData.IsPublic { // 旧数据是私有, 新数据改为了公开

		resellerList, err := dao.Reseller.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.Reseller.UpdateMany(ctx, bson.M{}, bson.M{
			"$addToSet": bson.M{
				"groups": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, reseller := range resellerList {

			newResellerData := *reseller

			newResellerData.Groups = gset.NewStrSetFrom(append(newResellerData.Groups, params.Id)).Slice()

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_RESELLER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: reseller,
				NewData: newResellerData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		userList, err := dao.User.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if err = dao.User.UpdateMany(ctx, bson.M{}, bson.M{
			"$addToSet": bson.M{
				"groups": params.Id,
			},
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		for _, user := range userList {

			newUserData := *user

			newUserData.Groups = gset.NewStrSetFrom(append(newUserData.Groups, params.Id)).Slice()

			if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
				Action:  consts.ACTION_UPDATE,
				OldData: user,
				NewData: newUserData,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 更改分组状态
func (s *sGroup) ChangeStatus(ctx context.Context, params model.GroupChangeStatusReq) error {

	group, err := dao.Group.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: group,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除分组
func (s *sGroup) Delete(ctx context.Context, id string) error {

	group, err := dao.Group.FindOneAndDeleteById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	resellers, err := dao.Reseller.Find(ctx, bson.M{"groups": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, reseller := range resellers {

		resellerModelsReq := model.ResellerPermissionsReq{
			UserId: reseller.UserId,
			Models: reseller.Models,
			Groups: []string{},
		}

		for _, g := range reseller.Groups {
			if g != id {
				resellerModelsReq.Groups = append(resellerModelsReq.Groups, g)
			}
		}

		if err = service.AdminReseller().Permissions(ctx, resellerModelsReq); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	users, err := dao.User.Find(ctx, bson.M{"groups": bson.M{"$in": []string{id}}})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, user := range users {

		userPermissionsReq := model.UserPermissionsReq{
			UserId: user.UserId,
			Models: user.Models,
			Groups: []string{},
		}

		for _, g := range user.Groups {
			if g != id {
				userPermissionsReq.Groups = append(userPermissionsReq.Groups, g)
			}
		}

		if err = service.AdminUser().Permissions(ctx, userPermissionsReq); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	apps, err := dao.App.Find(ctx, bson.M{"group": id})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, app := range apps {

		appGroupReq := model.AppGroupReq{
			AppId:       app.AppId,
			IsBindGroup: app.IsBindGroup,
			Group:       "",
		}

		if err = service.App().Group(ctx, appGroupReq); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	appKeys, err := dao.AppKey.Find(ctx, bson.M{"group": id})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	for _, key := range appKeys {

		keyGroupReq := model.AppKeyGroupReq{
			Id:          key.Id,
			IsBindGroup: key.IsBindGroup,
			Group:       "",
		}

		if err = service.AppKey().Group(ctx, keyGroupReq); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_GROUP, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: group,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.HDel(ctx, consts.API_GROUP_USAGE_KEY, group.Id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 分组详情
func (s *sGroup) Detail(ctx context.Context, id string) (*model.Group, error) {

	group, err := dao.Group.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelNames, err := service.Model().ModelNames(ctx, group.Models)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelAgentNames := make([]string, 0)

	if len(group.ModelAgents) > 0 && service.Session().IsAdminRole(ctx) {

		modelAgentList, err := dao.ModelAgent.Find(ctx, bson.M{"_id": bson.M{"$in": group.ModelAgents}})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		for _, modelAgent := range modelAgentList {
			modelAgentNames = append(modelAgentNames, modelAgent.Name)
		}
	}

	detail := &model.Group{
		Id:                 group.Id,
		Name:               group.Name,
		Discount:           group.Discount,
		Models:             group.Models,
		ModelNames:         modelNames,
		IsEnableModelAgent: group.IsEnableModelAgent,
		LbStrategy:         group.LbStrategy,
		ModelAgentNames:    modelAgentNames,
		ModelAgents:        group.ModelAgents,
		IsDefault:          group.IsDefault,
		IsLimitQuota:       group.IsLimitQuota,
		Quota:              group.Quota,
		UsedQuota:          group.UsedQuota,
		IsEnableForward:    group.IsEnableForward,
		ForwardConfig:      group.ForwardConfig,
		IsPublic:           group.IsPublic,
		Weight:             group.Weight,
		ExpiresAt:          util.FormatDateTime(group.ExpiresAt),
		Remark:             group.Remark,
		Status:             group.Status,
		Creator:            group.Creator,
		Updater:            group.Updater,
		CreatedAt:          util.FormatDateTime(group.CreatedAt),
		UpdatedAt:          util.FormatDateTime(group.UpdatedAt),
	}

	if detail.ForwardConfig != nil {

		if detail.ForwardConfig.TargetModel != "" {
			modelNames, err := service.Model().ModelNames(ctx, []string{detail.ForwardConfig.TargetModel})
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			detail.ForwardConfig.TargetModelName = modelNames[0]
		}

		if detail.ForwardConfig.DecisionModel != "" {
			modelNames, err := service.Model().ModelNames(ctx, []string{detail.ForwardConfig.DecisionModel})
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			detail.ForwardConfig.DecisionModelName = modelNames[0]
		}

		if detail.ForwardConfig.TargetModels != nil && len(detail.ForwardConfig.TargetModels) > 0 {
			modelNames, err := service.Model().ModelNames(ctx, detail.ForwardConfig.TargetModels)
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			detail.ForwardConfig.TargetModelNames = modelNames
		}
	}

	return detail, nil
}

// 分组分页列表
func (s *sGroup) Page(ctx context.Context, params model.GroupPageReq) (*model.GroupPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if service.Session().IsResellerRole(ctx) {

		reseller, err := service.Reseller().GetResellerByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(reseller.Groups) == 0 {
			return nil, nil
		}

		filter["_id"] = bson.M{
			"$in": reseller.Groups,
		}
	}

	if service.Session().IsUserRole(ctx) {

		user, err := service.User().GetUserByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(user.Groups) == 0 {
			return nil, nil
		}

		filter["_id"] = bson.M{
			"$in": user.Groups,
		}
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

	if len(params.ModelAgents) > 0 {
		filter["model_agents"] = bson.M{
			"$in": params.ModelAgents,
		}
	}

	if params.Remark != "" {
		filter["remark"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Remark),
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if len(params.ExpiresAt) > 0 {
		gte := gtime.NewFromStrFormat(params.ExpiresAt[0], time.DateOnly).StartOfDay().TimestampMilli()
		lte := gtime.NewFromStrLayout(params.ExpiresAt[1], time.DateOnly).EndOfDay(true).TimestampMilli()
		filter["expires_at"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.Group.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-is_default", "-weight", "-updated_at"}})
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

	items := make([]*model.Group, 0)
	for _, result := range results {

		modelNames := make([]string, 0)
		for _, id := range result.Models {
			if modelMap[id] != nil {
				modelNames = append(modelNames, modelMap[id].Name)
			}
		}

		group := &model.Group{
			Id:         result.Id,
			Name:       result.Name,
			Discount:   result.Discount,
			Models:     result.Models,
			ModelNames: modelNames,
			IsDefault:  result.IsDefault,
			Weight:     result.Weight,
			ExpiresAt:  util.FormatDateTime(result.ExpiresAt),
			Remark:     result.Remark,
			Status:     result.Status,
			CreatedAt:  util.FormatDateTime(result.CreatedAt),
			UpdatedAt:  util.FormatDateTime(result.UpdatedAt),
		}

		if service.Session().IsAdminRole(ctx) {

			group.UsedQuota = result.UsedQuota
			group.IsPublic = result.IsPublic

			if result.IsEnableModelAgent {

				modelAgentNames := make([]string, 0)
				for _, id := range result.ModelAgents {
					if modelAgentMap[id] != nil {
						modelAgentNames = append(modelAgentNames, modelAgentMap[id].Name)
					}
				}

				group.IsEnableModelAgent = result.IsEnableModelAgent
				group.LbStrategy = result.LbStrategy
				group.ModelAgents = result.ModelAgents
				group.ModelAgentNames = modelAgentNames
			}
		}

		items = append(items, group)
	}

	return &model.GroupPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 分组列表
func (s *sGroup) List(ctx context.Context, params model.GroupListReq) ([]*model.Group, error) {

	filter := bson.M{}

	if service.Session().IsResellerRole(ctx) {

		reseller, err := service.Reseller().GetResellerByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(reseller.Groups) == 0 {
			return nil, nil
		}

		filter["_id"] = bson.M{
			"$in": reseller.Groups,
		}
	}

	if service.Session().IsUserRole(ctx) {

		user, err := service.User().GetUserByUserId(ctx, service.Session().GetUserId(ctx))
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(user.Groups) == 0 {
			return nil, nil
		}

		filter["_id"] = bson.M{
			"$in": user.Groups,
		}
	}

	results, err := dao.Group.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"status", "-is_default", "-weight", "-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Group, 0)
	for _, result := range results {
		items = append(items, &model.Group{
			Id:     result.Id,
			Name:   result.Name,
			Models: result.Models,
			Remark: result.Remark,
			Status: result.Status,
		})
	}

	return items, nil
}

// 分组批量操作
func (s *sGroup) BatchOperate(ctx context.Context, params model.GroupBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_STATUS:
		for _, id := range params.Ids {
			if err := s.ChangeStatus(ctx, model.GroupChangeStatusReq{
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

// 公开的分组Ids
func (s *sGroup) PublicGroups(ctx context.Context) ([]string, error) {

	results, err := dao.Group.Find(ctx, bson.M{"is_public": true}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	ids := make([]string, 0)
	for _, result := range results {
		ids = append(ids, result.Id)
	}

	return ids, nil
}

// 根据分组Ids查询分组名称
func (s *sGroup) GroupNames(ctx context.Context, groups []string) ([]string, error) {

	if groups == nil || len(groups) == 0 {
		return nil, nil
	}

	results, err := dao.Group.Find(ctx, bson.M{"_id": bson.M{"$in": groups}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	groupNames := make([]string, 0)

	groupMap := util.ToMap(results, func(t *entity.Group) string {
		return t.Id
	})

	for _, id := range groups {
		if groupMap[id] != nil {
			groupNames = append(groupNames, groupMap[id].Name)
		}
	}

	return groupNames, nil
}

// 根据分组Ids获取模型Ids
func (s *sGroup) GetModelsByGroups(ctx context.Context, groups ...string) ([]string, error) {

	if groups == nil || len(groups) == 0 {
		return nil, nil
	}

	filter := bson.M{
		"_id": bson.M{"$in": groups},
	}

	results, err := dao.Group.Find(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelSet := gset.NewStrSet()
	for _, result := range results {
		modelSet.Add(result.Models...)
	}

	return modelSet.Slice(), nil
}

// 根据模型Ids获取分组Ids
func (s *sGroup) GetGroupsByModels(ctx context.Context, models ...string) ([]string, error) {

	filter := bson.M{
		"models": bson.M{"$in": models},
	}

	results, err := dao.Group.Find(ctx, filter)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	groups := make([]string, 0)
	for _, result := range results {
		groups = append(groups, result.Id)
	}

	return groups, nil
}
