package admin_user

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/core"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/logic/common"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/crypto"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"regexp"
	"time"
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

	var (
		salt = grand.Letters(8)
		id   = util.GenerateId()
		user = &do.User{
			Id:             id,
			UserId:         core.IncrUserId(ctx),
			Name:           params.Name,
			Email:          params.Account,
			Quota:          params.Quota,
			QuotaExpiresAt: common.ConvQuotaExpiresAt(params.QuotaExpiresAt),
			Models:         models,
			Remark:         params.Remark,
			Status:         1,
			Creator:        id,
		}
	)

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

	if params.Quota != 0 {

		// 交易记录
		if _, err = dao.DealRecord.Insert(ctx, &do.DealRecord{
			UserId: user.UserId,
			Quota:  user.Quota,
			Status: 1,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_USAGE_KEY, user.UserId), consts.USER_QUOTA_FIELD, int64(params.Quota)); err != nil {
			logger.Error(ctx, err)
			return err
		}
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

	newData, err := dao.User.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"name":             params.Name,
		"quota_expires_at": common.ConvQuotaExpiresAt(params.QuotaExpiresAt),
		"remark":           params.Remark,
		"status":           params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	account, err := dao.User.FindAccountByUserId(ctx, newData.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if account.Account != params.Account {
		if err = dao.Account.UpdateById(ctx, account.Id, bson.M{
			"account": params.Account,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if params.Password != "" {
		if err = dao.User.ChangePasswordByUserId(ctx, account.UserId, params.Password); err != nil {
			logger.Error(ctx, err)
			return err
		}
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

// 更改用户额度过期时间
func (s *sAdminUser) ChangeQuotaExpire(ctx context.Context, params model.UserChangeQuotaExpireReq) error {

	oldData, err := dao.User.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	newData, err := dao.User.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"quota_expires_at": common.ConvQuotaExpiresAt(params.QuotaExpiresAt),
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

	account, err := dao.User.FindAccountByUserId(ctx, user.UserId)
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
		Id:             user.Id,
		UserId:         user.UserId,
		Account:        account.Account,
		Name:           user.Name,
		Phone:          user.Phone,
		Email:          user.Email,
		Quota:          user.Quota,
		UsedQuota:      user.UsedQuota,
		QuotaExpiresAt: util.FormatDateTime(user.QuotaExpiresAt),
		Models:         user.Models,
		ModelNames:     modelNames,
		Remark:         user.Remark,
		Status:         user.Status,
		LoginIP:        account.LoginIP,
		LoginTime:      util.FormatDateTime(account.LoginTime),
		CreatedAt:      util.FormatDateTime(user.CreatedAt),
		UpdatedAt:      util.FormatDateTime(user.UpdatedAt),
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
		filter["$or"] = bson.A{
			bson.M{"name": bson.M{
				"$regex": regexp.QuoteMeta(params.Name),
			}},
			bson.M{"email": bson.M{
				"$regex": regexp.QuoteMeta(params.Name),
			}},
		}
	}

	if params.Account != "" && params.UserId == 0 {
		account, err := dao.Account.FindOne(ctx, bson.M{"account": params.Account})
		if err != nil {
			return nil, nil
		}
		filter["user_id"] = account.UserId
	}

	if params.Quota != 0 {
		filter["quota"] = bson.M{
			"$lte": params.Quota * consts.QUOTA_USD_UNIT,
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

	results, err := dao.User.FindByPage(ctx, paging, filter, "", "status", "-user_id", "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	accountMap := make(map[int]*entity.Account)
	if len(results) > 0 {

		accounts, err := dao.Account.Find(ctx, bson.M{})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		accountMap = util.ToMap(accounts, func(t *entity.Account) int {
			return t.UserId
		})
	}

	items := make([]*model.User, 0)
	for _, result := range results {

		items = append(items, &model.User{
			Id:             result.Id,
			UserId:         result.UserId,
			Name:           result.Name,
			Email:          result.Email,
			Phone:          result.Phone,
			Quota:          result.Quota,
			UsedQuota:      result.UsedQuota,
			QuotaExpiresAt: util.FormatDateTime(result.QuotaExpiresAt),
			Models:         result.Models,
			Account:        accountMap[result.UserId].Account,
			Remark:         result.Remark,
			Status:         result.Status,
			CreatedAt:      util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:      util.FormatDateTimeMonth(result.UpdatedAt),
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
		"quota_expires_at":  common.ConvQuotaExpiresAt(params.QuotaExpiresAt),
		"warning_notice":    false,
		"exhaustion_notice": false,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_USAGE_KEY, params.UserId), consts.USER_QUOTA_FIELD, int64(params.Quota)); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 交易记录
	if _, err = dao.DealRecord.Insert(ctx, &do.DealRecord{
		UserId: params.UserId,
		Quota:  params.Quota,
		Status: 1,
	}); err != nil {
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
