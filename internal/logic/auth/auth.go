package auth

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/core"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/crypto"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{}
}

// 注册接口
func (s *sAuth) Register(ctx context.Context, params model.RegisterReq) error {

	// 验证验证码是否正确
	if !service.Common().VerifyCode(ctx, consts.CHANNEL_REGISTER, params.Account, params.Code) {
		return errors.New("验证码填写错误")
	}

	if dao.User.IsAccountExist(ctx, params.Account) {
		return errors.New(params.Account + " 账号已存在")
	}

	salt := grand.Letters(8)

	user := &do.User{
		UserId:    core.IncrUserId(ctx),
		Email:     params.Account,
		Nickname:  params.Nickname,
		CreatedAt: gtime.Timestamp(),
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
	}); err != nil {
		return err
	}

	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	_ = service.Common().DelCode(ctx, consts.CHANNEL_REGISTER, params.Account)

	return nil
}

// 登录接口
func (s *sAuth) Login(ctx context.Context, params model.LoginReq) (res *model.LoginRes, err error) {

	defer func() {
		if err != nil {
			val, _ := redis.Incr(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account))
			if val == 1 {
				_, _ = redis.Expire(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account), 30*60) // 锁定30分钟
			}
		} else {
			_, _ = redis.Del(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account))
		}
	}()

	val, err := redis.GetInt(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account))
	if err == nil && val >= 5 {
		return nil, errors.New("登录失败次数过多, 请稍后再试")
	}

	accountInfo, err := dao.User.FindAccount(ctx, params.Account)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("账号或密码不正确")
		}
		logger.Error(ctx, err)
		return nil, err
	}

	if !crypto.VerifyPassword(accountInfo.Password, params.Password+accountInfo.Salt) {
		return nil, errors.New("账号或密码不正确")
	}

	ip := g.RequestFromCtx(ctx).GetClientIp()

	// 记录登录ip和时间
	if err = dao.Account.UpdateById(ctx, accountInfo.Id, bson.M{
		"last_login_ip":   ip,
		"last_login_time": gtime.Timestamp(),
	}); err != nil {
		logger.Error(ctx, err)
	}

	user, err := dao.User.FindById(ctx, accountInfo.Uid)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("用户不存在或已被禁用")
		}
		logger.Error(ctx, err)
		return nil, err
	}

	token := grand.Letters(32)

	_, _ = redis.Set(ctx, fmt.Sprintf(consts.USER_SESSION, token), gjson.MustEncodeString(user))

	return &model.LoginRes{
		Type:        "Bearer",
		AccessToken: token,
		ExpiresIn:   7200,
	}, nil
}

// 退出登录接口
func (s *sAuth) Logout(ctx context.Context) error {

	return nil
}

// 账号找回接口
func (s *sAuth) Forget(ctx context.Context, params model.ForgetReq) error {

	// 验证验证码是否正确
	if !service.Common().VerifyCode(ctx, consts.CHANNEL_FORGET_ACCOUNT, params.Account, params.Code) {
		return errors.New("验证码填写错误")
	}

	account, err := dao.User.FindAccount(ctx, params.Account)
	if err != nil || account.Id == "" {
		return errors.New(params.Account + " 账号不存在")
	}

	if err = dao.User.ChangePasswordByUserId(ctx, account.UserId, params.Password); err != nil {
		logger.Error(ctx, err)
		return errors.New("找回密码失败")
	}

	_ = service.Common().DelCode(ctx, consts.CHANNEL_FORGET_ACCOUNT, params.Account)

	return nil
}
