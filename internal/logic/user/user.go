package user

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

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// 用户更新信息
func (s *sUser) UpdateInfo(ctx context.Context, params model.UserUpdateInfoReq) error {

	if err := dao.User.UpdateById(ctx, service.Session().GetUid(ctx), bson.M{
		"name": params.Name,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	user := service.Session().GetUser(ctx)
	user.Name = params.Name

	if err := service.Session().UpdateUserSession(ctx, user); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 用户更改密码
func (s *sUser) ChangePassword(ctx context.Context, params model.UserChangePasswordReq) (err error) {

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

	user, err := dao.User.FindUserByUserId(ctx, uid)
	if err != nil || user.Id == "" {
		return errors.New("用户不存在")
	}

	account, err := dao.User.FindAccountByUserId(ctx, user.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return errors.New("账号信息有误")
	}

	if !crypto.VerifyPassword(account.Password, params.OldPassword+account.Salt) {
		return errors.New("登录密码有误, 请重新输入")
	}

	if err = dao.User.ChangePasswordByUserId(ctx, uid, params.NewPassword); err != nil {
		logger.Error(ctx, err)
		return errors.New("修改密码失败")
	}

	return nil
}

// 用户更改邮箱
func (s *sUser) ChangeEmail(ctx context.Context, params model.UserChangeEmailReq) error {

	if !service.Common().VerifyCode(ctx, consts.SCENE_CHANGE_EMAIL, params.Email, params.Code) {
		return errors.New("邮件验证码填写错误")
	}

	user, err := dao.User.FindUserByUserId(ctx, service.Session().GetUserId(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	account, err := dao.User.FindAccountByUserId(ctx, user.UserId)
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

	if user.Email == params.Email {
		return errors.New("邮箱与原邮箱一致无需修改")
	}

	if dao.User.IsAccountExist(ctx, params.Email) {
		return errors.New(params.Email + " 邮箱已被其它账号使用")
	}

	if err = dao.User.UpdateById(ctx, user.Id, bson.M{
		"email": params.Email,
	}); err != nil {
		logger.Error(ctx, err)
		return errors.New("邮箱修改失败")
	}

	if account.Account == user.Email {
		if err = dao.User.ChangeAccountById(ctx, account.Id, params.Email); err != nil {
			logger.Error(ctx, err)
			return err
		}
	} else {

		accountInfo, err := dao.User.FindAccount(ctx, user.Email)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
			return err
		}

		if accountInfo != nil {
			if err = dao.User.ChangeAccountById(ctx, accountInfo.Id, params.Email); err != nil {
				logger.Error(ctx, err)
				return err
			}
		} else {
			if _, err = dao.User.CreateAccount(ctx, &do.Account{
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

	userSession := service.Session().GetUser(ctx)
	userSession.Email = params.Email

	if err = service.Session().UpdateUserSession(ctx, userSession); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 用户更改头像
func (s *sUser) ChangeAvatar(ctx context.Context, file *ghttp.UploadFile) error {

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

	if err = dao.User.UpdateById(ctx, service.Session().GetUid(ctx), bson.M{
		"avatar": path + filename,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	user := service.Session().GetUser(ctx)
	user.Avatar = path + filename

	if err = service.Session().UpdateUserSession(ctx, user); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 根据userId获取用户信息
func (s *sUser) GetUserByUserId(ctx context.Context, userId int) (*model.User, error) {

	user, err := dao.User.FindUserByUserId(ctx, userId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.User{
		Id:                     user.Id,
		UserId:                 user.UserId,
		Name:                   user.Name,
		Avatar:                 user.Avatar,
		Email:                  user.Email,
		Phone:                  user.Phone,
		Quota:                  user.Quota,
		UsedQuota:              user.UsedQuota,
		QuotaExpiresAt:         util.FormatDateTime(user.QuotaExpiresAt),
		Models:                 user.Models,
		Groups:                 user.Groups,
		QuotaWarning:           user.QuotaWarning,
		WarningThreshold:       user.WarningThreshold,
		ExpireWarningThreshold: user.ExpireWarningThreshold,
		Remark:                 user.Remark,
		Status:                 user.Status,
		Rid:                    user.Rid,
		CreatedAt:              util.FormatDateTime(user.CreatedAt),
		UpdatedAt:              util.FormatDateTime(user.UpdatedAt),
	}, nil
}
