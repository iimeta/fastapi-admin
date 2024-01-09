package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/crypto"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// 用户信息
func (s *sUser) Info(ctx context.Context) (*model.UserInfoRes, error) {

	user, err := dao.User.FindUserByUserId(ctx, service.Session().GetUserId(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.UserInfoRes{
		Id:     gconv.String(user.UserId),
		Phone:  user.Phone,
		Name:   user.Name,
		Avatar: user.Avatar,
		Gender: user.Gender,
		Email:  user.Email,
	}, nil
}

// 修改用户信息
func (s *sUser) ChangeDetail(ctx context.Context, params model.UserDetailUpdateReq) error {

	if params.Birthday != "" {
		if !util.IsDateFormat(params.Birthday) {
			return errors.New("birthday 格式错误")
		}
	}

	if err := dao.User.UpdateOne(ctx, bson.M{"user_id": service.Session().GetUserId(ctx)}, &do.User{
		Name:   strings.TrimSpace(strings.Replace(params.Name, " ", "", -1)),
		Avatar: params.Avatar,
		Gender: params.Gender,
	}); err != nil {
		logger.Error(ctx, err)
		return errors.New("个人信息修改失败")
	}

	return nil
}

// 修改密码接口
func (s *sUser) ChangePassword(ctx context.Context, params model.UserPasswordUpdateReq) (err error) {

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

	val, err := redis.GetInt(ctx, fmt.Sprintf(consts.LOCK_CHANGE_PASSWORD, uid))
	if err == nil && val >= 5 {
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

// 用户设置
func (s *sUser) Setting(ctx context.Context) (*model.UserSettingRes, error) {

	user, err := dao.User.FindUserByUserId(ctx, service.Session().GetUserId(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.UserSettingRes{
		User: &model.User{
			UserId: user.UserId,
			Name:   user.Name,
			Avatar: user.Avatar,
			Gender: user.Gender,
			Phone:  user.Phone,
			Email:  user.Email,
		},
		Setting: &model.SettingInfo{},
	}, nil
}

// 换绑手机号
func (s *sUser) ChangePhone(ctx context.Context, params model.UserPhoneUpdateReq) error {

	if !service.Common().VerifyCode(ctx, consts.CHANNEL_CHANGE_MOBILE, params.Phone, params.Code) {
		return errors.New("短信验证码填写错误")
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

	if user.Phone == params.Phone {
		return errors.New("手机号与原手机号一致无需修改")
	}

	if dao.User.IsAccountExist(ctx, params.Phone) {
		return errors.New(params.Phone + " 手机号已被其它账号使用")
	}

	if err = dao.User.UpdateById(ctx, user.Id, bson.M{
		"phone": params.Phone,
	}); err != nil {
		logger.Error(ctx, err)
		return errors.New("手机号修改失败")
	}

	if account.Account == user.Phone {
		if err = dao.User.ChangeAccountById(ctx, account.Id, params.Phone); err != nil {
			logger.Error(ctx, err)
			return err
		}
	} else {

		accountInfo, err := dao.User.FindAccount(ctx, user.Phone)
		if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
			return err
		}

		if accountInfo != nil {
			if err = dao.User.ChangeAccountById(ctx, accountInfo.Id, params.Phone); err != nil {
				logger.Error(ctx, err)
				return err
			}
		} else {
			if _, err := dao.User.CreateAccount(ctx, &do.Account{
				Uid:      account.Uid,
				UserId:   account.UserId,
				Account:  params.Phone,
				Password: account.Password,
				Salt:     account.Salt,
				Status:   1,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 换绑邮箱
func (s *sUser) ChangeEmail(ctx context.Context, params model.UserEmailUpdateReq) error {

	if !service.Common().VerifyCode(ctx, consts.CHANNEL_CHANGE_EMAIL, params.Email, params.Code) {
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
			if _, err := dao.User.CreateAccount(ctx, &do.Account{
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

	return nil
}

// 根据userId获取用户信息
func (s *sUser) GetUserById(ctx context.Context, userId int) (*model.User, error) {

	user, err := dao.User.FindUserByUserId(ctx, userId)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.User{
		Id:        user.Id,
		UserId:    user.UserId,
		Phone:     user.Phone,
		Name:      user.Name,
		Avatar:    user.Avatar,
		Gender:    user.Gender,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
