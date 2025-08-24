package reseller

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/crypto"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sReseller struct{}

func init() {
	service.RegisterReseller(New())
}

func New() service.IReseller {
	return &sReseller{}
}

// 代理商更新信息
func (s *sReseller) UpdateInfo(ctx context.Context, params model.UserUpdateInfoReq) error {

	if err := dao.Reseller.UpdateById(ctx, service.Session().GetUid(ctx), bson.M{
		"name": params.Name,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	reseller := service.Session().GetReseller(ctx)
	reseller.Name = params.Name

	if err := service.Session().UpdateResellerSession(ctx, reseller); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 代理商更改密码
func (s *sReseller) ChangePassword(ctx context.Context, params model.UserChangePasswordReq) (err error) {

	uid := service.Session().GetUserId(ctx)

	defer func() {
		if err != nil {
			val, _ := redis.Incr(ctx, fmt.Sprintf(consts.LOCK_CHANGE_PASSWORD, uid))
			if val == 1 {
				_, _ = redis.Expire(ctx, fmt.Sprintf(consts.LOCK_CHANGE_PASSWORD, uid), 30*60) // 锁定30分钟
			}
		} else {
			_, _ = redis.Del(ctx, fmt.Sprintf(consts.LOCK_CHANGE_PASSWORD, uid))
		}
	}()

	if val, err := redis.GetInt(ctx, fmt.Sprintf(consts.LOCK_CHANGE_PASSWORD, uid)); err == nil && val >= 5 {
		return errors.New("失败次数过多, 请稍后再试")
	}

	reseller, err := dao.Reseller.FindResellerByUserId(ctx, uid)
	if err != nil || reseller.Id == "" {
		return errors.New("代理商不存在")
	}

	account, err := dao.Reseller.FindAccountByUserId(ctx, reseller.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("账号信息有误")
	}

	if !crypto.VerifyPassword(account.Password, params.OldPassword+account.Salt) {
		return errors.New("登录密码有误, 请重新输入")
	}

	if err = dao.Reseller.ChangePasswordByUserId(ctx, uid, params.NewPassword); err != nil {
		logger.Error(ctx, err)
		return errors.New("修改密码失败")
	}

	return nil
}

// 代理商更改邮箱
func (s *sReseller) ChangeEmail(ctx context.Context, params model.UserChangeEmailReq) error {

	if !service.Common().VerifyCode(ctx, consts.SCENE_CHANGE_EMAIL, params.Email, params.Code) {
		return errors.New("邮件验证码填写错误")
	}

	reseller, err := dao.Reseller.FindResellerByUserId(ctx, service.Session().GetUserId(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	account, err := dao.Reseller.FindAccountByUserId(ctx, reseller.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("账号信息有误")
	}

	if !crypto.VerifyPassword(account.Password, params.Password+account.Salt) {
		return errors.New("登录密码有误, 请重新输入")
	}

	defer func() {
		_ = service.Common().DelCode(ctx, consts.SCENE_CHANGE_EMAIL, params.Email)
	}()

	if reseller.Email == params.Email {
		return errors.New("邮箱与原邮箱一致无需修改")
	}

	if dao.Reseller.IsAccountExist(ctx, params.Email) {
		return errors.New(params.Email + " 邮箱已被其它账号使用")
	}

	if err = dao.Reseller.UpdateById(ctx, reseller.Id, bson.M{
		"email": params.Email,
	}); err != nil {
		logger.Error(ctx, err)
		return errors.New("邮箱修改失败")
	}

	if account.Account == reseller.Email {
		if err = dao.Reseller.ChangeAccountById(ctx, account.Id, params.Email); err != nil {
			logger.Error(ctx, err)
			return err
		}
	} else {

		accountInfo, err := dao.Reseller.FindAccount(ctx, reseller.Email)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
			return err
		}

		if accountInfo != nil {
			if err = dao.Reseller.ChangeAccountById(ctx, accountInfo.Id, params.Email); err != nil {
				logger.Error(ctx, err)
				return err
			}
		} else {
			if _, err := dao.Reseller.CreateAccount(ctx, &do.ResellerAccount{
				Uid:      account.Uid,
				UserId:   account.UserId,
				Account:  params.Email,
				Password: account.Password,
				Salt:     account.Salt,
				Status:   1,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	resellerSession := service.Session().GetReseller(ctx)
	resellerSession.Email = params.Email

	if err := service.Session().UpdateResellerSession(ctx, resellerSession); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 代理商更改头像
func (s *sReseller) ChangeAvatar(ctx context.Context, file *ghttp.UploadFile) error {

	if file.Size > 1024*1024*8 {
		return errors.New("头像文件过大, 请更换或压缩后再上传")
	}

	root := "./resource"
	path := "/public/avatar/"

	filename, err := file.Save(root+path, true)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.Reseller.UpdateById(ctx, service.Session().GetUid(ctx), bson.M{
		"avatar": path + filename,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	reseller := service.Session().GetReseller(ctx)
	reseller.Avatar = path + filename

	if err = service.Session().UpdateResellerSession(ctx, reseller); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 根据userId获取代理商信息
func (s *sReseller) GetResellerByUserId(ctx context.Context, userId int) (*model.Reseller, error) {

	reseller, err := dao.Reseller.FindResellerByUserId(ctx, userId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.Reseller{
		Id:                     reseller.Id,
		UserId:                 reseller.UserId,
		Name:                   reseller.Name,
		Avatar:                 reseller.Avatar,
		Email:                  reseller.Email,
		Phone:                  reseller.Phone,
		Quota:                  reseller.Quota,
		UsedQuota:              reseller.UsedQuota,
		QuotaExpiresAt:         util.FormatDateTime(reseller.QuotaExpiresAt),
		Models:                 reseller.Models,
		Groups:                 reseller.Groups,
		QuotaWarning:           reseller.QuotaWarning,
		WarningThreshold:       reseller.WarningThreshold,
		ExpireWarningThreshold: reseller.ExpireWarningThreshold,
		Remark:                 reseller.Remark,
		Status:                 reseller.Status,
		CreatedAt:              util.FormatDateTime(reseller.CreatedAt),
		UpdatedAt:              util.FormatDateTime(reseller.UpdatedAt),
	}, nil
}
