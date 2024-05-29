package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/core"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/crypto"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

type sAdminUser struct{}

func init() {
	service.RegisterAdminUser(New())
}

func New() service.IAdminUser {
	return &sAdminUser{}
}

// 新建用户
func (s *sAdminUser) Create(ctx context.Context, params model.UserCreateReq) error {

	if dao.User.IsAccountExist(ctx, params.Account) {
		return errors.New(params.Account + " 账号已存在")
	}

	models, err := service.Model().PublicModels(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	salt := grand.Letters(8)
	id := util.GenerateId()

	user := &do.User{
		Id:      id,
		UserId:  core.IncrUserId(ctx),
		Name:    params.Name,
		Email:   params.Account,
		Quota:   params.Quota,
		Models:  models,
		Remark:  params.Remark,
		Status:  1,
		Creator: id,
	}

	uid, err := dao.User.Insert(ctx, user)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.User.CreateAccount(ctx, &do.Account{
		Uid:      uid,
		UserId:   user.UserId,
		Account:  params.Account,
		Password: crypto.EncryptPassword(params.Password + salt),
		Salt:     salt,
		Status:   1,
		Creator:  uid,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_USAGE_KEY, user.UserId), consts.USER_QUOTA_FIELD, int64(params.Quota)); err != nil {
		logger.Error(ctx, err)
		return err
	}

	newData, err := dao.User.FindById(ctx, uid)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_CREATE,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新用户
func (s *sAdminUser) Update(ctx context.Context, params model.UserUpdateReq) error {

	oldData, err := dao.User.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	newData, err := dao.User.FindOneAndUpdateById(ctx, params.Id, &do.User{
		Name:   params.Name,
		Models: params.Models,
		Quota:  params.Quota,
		Remark: params.Remark,
		Status: params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改用户状态
func (s *sAdminUser) ChangeStatus(ctx context.Context, params model.UserChangeStatusReq) error {

	user, err := dao.User.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.Account.UpdateMany(ctx, bson.M{"user_id": user.UserId}, bson.M{
		"status": params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: user,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除用户
func (s *sAdminUser) Delete(ctx context.Context, id string) error {

	user, err := dao.User.FindOneAndDeleteById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = dao.Account.DeleteMany(ctx, bson.M{"user_id": user.UserId}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: user,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 用户详情
func (s *sAdminUser) Detail(ctx context.Context, id string) (*model.User, error) {

	user, err := dao.User.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	modelNames, err := service.Model().ModelNames(ctx, user.Models)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.User{
		Id:         user.Id,
		UserId:     user.UserId,
		Name:       user.Name,
		Phone:      user.Phone,
		Email:      user.Email,
		Quota:      user.Quota,
		UsedQuota:  user.UsedQuota,
		Models:     user.Models,
		ModelNames: modelNames,
		Remark:     user.Remark,
		Status:     user.Status,
		CreatedAt:  util.FormatDateTime(user.CreatedAt),
		UpdatedAt:  util.FormatDateTime(user.UpdatedAt),
	}, nil
}

// 用户分页列表
func (s *sAdminUser) Page(ctx context.Context, params model.UserPageReq) (*model.UserPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.Name != "" {
		filter["name"] = bson.M{
			"$regex": params.Name,
		}
	}

	if params.Email != "" {
		filter["email"] = bson.M{
			"$regex": params.Email,
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	results, err := dao.User.FindByPage(ctx, paging, filter, "status", "-user_id", "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.User, 0)
	for _, result := range results {

		items = append(items, &model.User{
			Id:        result.Id,
			UserId:    result.UserId,
			Name:      result.Name,
			Email:     result.Email,
			Phone:     result.Phone,
			Quota:     result.Quota,
			UsedQuota: result.UsedQuota,
			Models:    result.Models,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt: util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.UserPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 用户列表
func (s *sAdminUser) List(ctx context.Context, params model.UserListReq) ([]*model.User, error) {

	filter := bson.M{}

	results, err := dao.User.Find(ctx, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.User, 0)
	for _, result := range results {
		items = append(items, &model.User{
			Id:        result.Id,
			UserId:    result.UserId,
			Name:      result.Name,
			Email:     result.Email,
			Phone:     result.Phone,
			Quota:     result.Quota,
			UsedQuota: result.UsedQuota,
			Models:    result.Models,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt: util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return items, nil
}

// 授予用户额度
func (s *sAdminUser) GrantQuota(ctx context.Context, params model.UserGrantQuotaReq) error {

	oldData, err := dao.User.FindOne(ctx, bson.M{"user_id": params.UserId})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	newData, err := dao.User.FindOneAndUpdate(ctx, bson.M{"user_id": params.UserId}, bson.M{
		"$inc": bson.M{
			"quota": params.Quota,
		},
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_USAGE_KEY, params.UserId), consts.USER_QUOTA_FIELD, int64(params.Quota)); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 用户模型权限
func (s *sAdminUser) Models(ctx context.Context, params model.UserModelsReq) error {

	oldData, err := dao.User.FindOne(ctx, bson.M{"user_id": params.UserId})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	newData, err := dao.User.FindOneAndUpdate(ctx, bson.M{"user_id": params.UserId}, bson.M{
		"models": params.Models,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_USER, model.PubMessage{
		Action:  consts.ACTION_MODELS,
		OldData: oldData,
		NewData: newData,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
