package auth

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{}
}

// 注册接口
func (s *sAuth) Register(ctx context.Context, params model.RegisterReq, channel ...string) error {

	if len(channel) == 0 {
		channel = []string{consts.CHANNEL_REGISTER}
	}

	// 验证验证码是否正确
	if !service.Common().VerifyCode(ctx, channel[0], params.Account, params.Code) {
		return errors.New("验证码填写错误")
	}

	if dao.User.IsAccountExist(ctx, params.Account) {
		return errors.New(params.Account + " 账号已存在")
	}

	salt := grand.Letters(8)

	user := &do.User{
		UserId:    core.IncrUserId(ctx),
		Email:     params.Account,
		Name:      params.Account,
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

	r := g.RequestFromCtx(ctx)
	ip := r.GetClientIp()
	token := ""

	if params.Channel == consts.USER_CHANNEL {

		if params.Method == consts.METHOD_CODE {
			// 验证验证码是否正确
			if !service.Common().VerifyCode(ctx, consts.CHANNEL_LOGIN, params.Account, params.Code) {
				return nil, errors.New("验证码填写错误")
			}
		}

		accountInfo, err := dao.User.FindAccount(ctx, params.Account)

		if params.Method == consts.METHOD_ACCOUNT {

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

		} else if params.Method == consts.METHOD_CODE {

			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {

					if err = s.Register(ctx, model.RegisterReq{
						Account:  params.Account,
						Password: grand.Letters(8),
						Terminal: params.Terminal,
						Code:     params.Code,
					}, consts.CHANNEL_LOGIN); err != nil {
						logger.Error(ctx, err)
						return nil, err
					}

					accountInfo, err = dao.User.FindAccount(ctx, params.Account)
					if err != nil {
						logger.Error(ctx, err)
						return nil, err
					}

				} else {
					logger.Error(ctx, err)
					return nil, err
				}
			}

		} else {
			return nil, errors.New("账号或密码不正确")
		}

		user, err := dao.User.FindById(ctx, accountInfo.Uid)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errors.New("用户不存在或已被禁用")
			}
			logger.Error(ctx, err)
			return nil, err
		}

		r.SetCtxVar("uid", user.Id)

		// 记录登录IP和时间
		if err = dao.Account.UpdateById(gctx.WithCtx(r.GetCtx()), accountInfo.Id, bson.M{
			"login_ip":   ip,
			"login_time": gtime.Timestamp(),
		}); err != nil {
			logger.Error(ctx, err)
		}

		token, err = s.GenUserToken(ctx, &model.User{
			Id:        user.Id,
			UserId:    user.UserId,
			Name:      user.Name,
			Avatar:    user.Avatar,
			Gender:    user.Gender,
			Email:     user.Email,
			Phone:     user.Phone,
			CreatedAt: util.FormatDatetime(user.CreatedAt),
			UpdatedAt: util.FormatDatetime(user.UpdatedAt),
		}, true)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

	} else if params.Channel == consts.ADMIN_CHANNEL {

		admin, err := dao.SysAdmin.FindOne(ctx, bson.M{"account": params.Account})
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {

				count, err := dao.SysAdmin.EstimatedDocumentCount(ctx)
				if err != nil {
					logger.Error(ctx, err)
					return nil, err
				}

				// 初次登录自动创建账号
				if count == 0 {

					if err = service.SysAdmin().Create(ctx, model.SysAdminCreateReq{
						Name:     params.Account,
						Account:  params.Account,
						Password: params.Password,
					}); err != nil {
						logger.Error(ctx, err)
						return nil, err
					}

					admin, err = dao.SysAdmin.FindOne(ctx, bson.M{"account": params.Account})
					if err != nil {
						logger.Error(ctx, err)
						return nil, err
					}

				} else {
					return nil, errors.New("账号或密码不正确")
				}
			} else {
				logger.Error(ctx, err)
				return nil, err
			}
		}

		if !crypto.VerifyPassword(admin.Password, params.Password+admin.Salt) {
			return nil, errors.New("账号或密码不正确")
		}

		r.SetCtxVar("uid", admin.Id)

		// 记录登录ip和时间
		if err = dao.SysAdmin.UpdateById(gctx.WithCtx(r.GetCtx()), admin.Id, bson.M{
			"login_ip":   ip,
			"login_time": gtime.Timestamp(),
		}); err != nil {
			logger.Error(ctx, err)
		}

		token, err = s.GenAdminToken(ctx, &model.SysAdmin{
			Id:        admin.Id,
			Name:      admin.Name,
			Avatar:    admin.Avatar,
			Gender:    admin.Gender,
			Phone:     admin.Phone,
			Email:     admin.Email,
			Account:   admin.Account,
			Remark:    admin.Remark,
			Status:    admin.Status,
			CreatedAt: admin.CreatedAt,
			UpdatedAt: admin.UpdatedAt,
		}, true)
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	return &model.LoginRes{
		Type:      "Bearer",
		Token:     token,
		ExpiresIn: 7200,
	}, nil
}

// 退出登录接口
func (s *sAuth) Logout(ctx context.Context) error {

	token := g.RequestFromCtx(ctx).GetHeader("Authorization")
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

	key := fmt.Sprintf(consts.USER_SESSION, token)

	if gstr.HasPrefix(token, consts.ADMIN_TOKEN_PREFIX) {
		key = fmt.Sprintf(consts.ADMIN_SESSION, token)
	}

	if _, err := redis.Del(ctx, key); err != nil {
		logger.Error(ctx, err)
		return err
	}

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

// 生成用户Token
func (s *sAuth) GenUserToken(ctx context.Context, user *model.User, isSaveSession bool) (token string, err error) {

	token = util.NewKey(consts.USER_TOKEN_PREFIX, 32, gconv.String(user.UserId))

	if isSaveSession {
		err = redis.SetEX(ctx, fmt.Sprintf(consts.USER_SESSION, token), gjson.MustEncodeString(user), 7200)
	}

	return token, err
}

// 根据Token获取用户信息
func (s *sAuth) GetUserByToken(ctx context.Context, token string) (*model.User, error) {

	reply, err := redis.Get(ctx, fmt.Sprintf(consts.USER_SESSION, token))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if reply == nil || reply.IsNil() {
		return nil, errors.New("session is nil")
	}

	user := new(model.User)

	err = reply.Struct(&user)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return user, nil
}

// 生成管理员Token
func (s *sAuth) GenAdminToken(ctx context.Context, admin *model.SysAdmin, isSaveSession bool) (token string, err error) {

	token = util.NewKey(consts.ADMIN_TOKEN_PREFIX, 32, admin.Id)

	if isSaveSession {
		err = redis.SetEX(ctx, fmt.Sprintf(consts.ADMIN_SESSION, token), gjson.MustEncodeString(admin), 7200)
	}

	return token, err
}

// 根据Token获取管理员信息
func (s *sAuth) GetAdminByToken(ctx context.Context, token string) (*model.SysAdmin, error) {

	reply, err := redis.Get(ctx, fmt.Sprintf(consts.ADMIN_SESSION, token))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if reply == nil || reply.IsNil() {
		return nil, errors.New("session is nil")
	}

	admin := new(model.SysAdmin)

	err = reply.Struct(&admin)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return admin, nil
}
