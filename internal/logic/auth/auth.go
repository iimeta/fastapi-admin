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
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/core"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/cache"
	"github.com/iimeta/fastapi-admin/utility/crypto"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

type sAuth struct {
	tokenCache *cache.Cache // [token]User
}

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{
		tokenCache: cache.New(),
	}
}

// 注册
func (s *sAuth) Register(ctx context.Context, params model.RegisterReq, channel ...string) error {

	if len(channel) == 0 {
		channel = []string{consts.ACTION_REGISTER}
	}

	// 验证验证码是否正确
	if !service.Common().VerifyCode(ctx, channel[0], params.Account, params.Code) {
		return errors.New("验证码填写错误")
	}

	if params.Channel == consts.USER_CHANNEL && !config.Cfg.UserLoginRegister.EmailRegister {
		return errors.New("未开启用户注册, 请联系管理员")
	}

	if params.Channel == consts.RESELLER_CHANNEL && !config.Cfg.ResellerLoginRegister.EmailRegister {
		return errors.New("未开启代理商注册, 请联系管理员")
	}

	siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, params.Domain)
	if siteConfig != nil {

		if siteConfig.RegisterTips != "" {
			return errors.New(siteConfig.RegisterTips)
		}

		if len(siteConfig.SupportEmailSuffix) > 0 {

			isSupport := false
			for _, emailSuffix := range siteConfig.SupportEmailSuffix {
				if isSupport = gstr.HasSuffix(params.Account, emailSuffix); isSupport {
					break
				}
			}

			if !isSupport {
				return errors.Newf("邮箱仅支持 %s 后缀", siteConfig.SupportEmailSuffix)
			}
		}
	}

	if params.Channel == consts.USER_CHANNEL && dao.User.IsAccountExist(ctx, params.Account) {
		return errors.New(params.Account + " 账号已存在")
	}

	if params.Channel == consts.RESELLER_CHANNEL && dao.Reseller.IsAccountExist(ctx, params.Account) {
		return errors.New(params.Account + " 账号已存在")
	}

	models, err := service.Model().PublicModels(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	groups, err := service.Group().PublicGroups(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if params.Channel == consts.USER_CHANNEL {

		salt := grand.Letters(8)
		id := util.GenerateId()

		user := &do.User{
			Id:      id,
			UserId:  core.IncrUserId(ctx),
			Name:    params.Account,
			Email:   params.Account,
			Models:  models,
			Groups:  groups,
			Status:  1,
			Creator: id,
		}

		if siteConfig != nil && siteConfig.GrantQuota > 0 {

			user.Rid = siteConfig.Rid

			if user.Rid != 0 {

				tmpCtx := gctx.New()
				tmpCtx = context.WithValue(tmpCtx, consts.SESSION_RID, user.Rid)
				tmpCtx = context.WithValue(tmpCtx, consts.SESSION_USER_ID, user.Rid)
				tmpCtx = context.WithValue(tmpCtx, consts.SESSION_ROLE, consts.SESSION_RESELLER)
				tmpCtx = context.WithValue(tmpCtx, consts.SESSION_CREATOR, siteConfig.Creator)

				expense, err := service.Dashboard().Expense(tmpCtx)
				if err != nil {
					logger.Error(ctx, err)
				}

				if expense != nil && expense.ToBeAllocated >= siteConfig.GrantQuota {
					user.Quota = siteConfig.GrantQuota
					if siteConfig.QuotaExpiresAt > 0 {
						user.QuotaExpiresAt = gtime.Now().Add(time.Duration(siteConfig.QuotaExpiresAt) * time.Minute).TimestampMilli()
					}
				}

			} else {
				user.Quota = siteConfig.GrantQuota
				if siteConfig.QuotaExpiresAt > 0 {
					user.QuotaExpiresAt = gtime.Now().Add(time.Duration(siteConfig.QuotaExpiresAt) * time.Minute).TimestampMilli()
				}
			}
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
			Rid:      user.Rid,
			Creator:  uid,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		_ = service.Common().DelCode(ctx, consts.ACTION_REGISTER, params.Account)

		if user.Quota != 0 {

			// 交易记录
			if _, err = dao.DealRecord.Insert(ctx, &do.DealRecord{
				UserId: user.UserId,
				Quota:  user.Quota,
				Type:   3,
				Status: 1,
				Rid:    user.Rid,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}

			if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_USER_USAGE_KEY, user.UserId), consts.USER_QUOTA_FIELD, int64(user.Quota)); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}

		r := g.RequestFromCtx(ctx)
		r.SetCtxVar(consts.SESSION_USER_ID, user.UserId)
		r.SetCtxVar(consts.SESSION_RID, user.Rid)

		// 创建默认应用
		if _, err = service.App().Create(r.GetCtx(), model.AppCreateReq{
			Name:        "默认应用",
			IsCreateKey: true,
			Status:      1,
		}); err != nil {
			logger.Error(ctx, err)
		}

	} else if params.Channel == consts.RESELLER_CHANNEL {

		salt := grand.Letters(8)
		id := util.GenerateId()

		reseller := &do.Reseller{
			Id:      id,
			UserId:  core.IncrResellerId(ctx),
			Name:    params.Account,
			Email:   params.Account,
			Models:  models,
			Groups:  groups,
			Status:  1,
			Creator: id,
		}

		if siteConfig != nil && siteConfig.GrantQuota > 0 && siteConfig.Rid == 0 {
			reseller.Quota = siteConfig.GrantQuota
			if siteConfig.QuotaExpiresAt > 0 {
				reseller.QuotaExpiresAt = gtime.Now().Add(time.Duration(siteConfig.QuotaExpiresAt) * time.Minute).TimestampMilli()
			}
		}

		uid, err := dao.Reseller.Insert(ctx, reseller)
		if err != nil {
			logger.Error(ctx, err)
			return err
		}

		if _, err = dao.Reseller.CreateAccount(ctx, &do.ResellerAccount{
			Uid:      uid,
			UserId:   reseller.UserId,
			Account:  params.Account,
			Password: crypto.EncryptPassword(params.Password + salt),
			Salt:     salt,
			Status:   1,
			Creator:  uid,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		_ = service.Common().DelCode(ctx, consts.ACTION_REGISTER, params.Account)

		if reseller.Quota != 0 {

			// 交易记录
			if _, err = dao.DealRecord.Insert(ctx, &do.DealRecord{
				UserId: reseller.UserId,
				Quota:  reseller.Quota,
				Type:   3,
				Status: 1,
				Rid:    reseller.UserId,
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}

			if _, err = redis.HIncrBy(ctx, fmt.Sprintf(consts.API_RESELLER_USAGE_KEY, reseller.UserId), consts.RESELLER_QUOTA_FIELD, int64(reseller.Quota)); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 登录
func (s *sAuth) Login(ctx context.Context, params model.LoginReq) (res *model.LoginRes, err error) {

	defer func() {
		if err != nil {
			if val, _ := redis.Incr(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account)); val == 1 {
				_, _ = redis.Expire(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account), 30*60) // 锁定30分钟
			}
		} else {
			_, _ = redis.Del(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account))
		}
	}()

	if val, err := redis.GetInt(ctx, fmt.Sprintf(consts.LOCK_LOGIN, params.Account)); err == nil && val >= 5 {
		return nil, errors.New("登录失败次数过多, 请稍后再试")
	}

	if params.Method == consts.METHOD_CODE {

		// 验证验证码是否正确
		if !service.Common().VerifyCode(ctx, consts.ACTION_LOGIN, params.Account, params.Code) {
			return nil, errors.New("验证码填写错误")
		}

		if params.Channel == consts.USER_CHANNEL && !config.Cfg.UserLoginRegister.EmailLogin {
			return nil, errors.New("未开启邮箱登录, 请联系管理员")
		}

		if params.Channel == consts.RESELLER_CHANNEL && !config.Cfg.ResellerLoginRegister.EmailLogin {
			return nil, errors.New("未开启邮箱登录, 请联系管理员")
		}

		if params.Channel == consts.ADMIN_CHANNEL && !config.Cfg.AdminLogin.EmailLogin {
			return nil, errors.New("未开启邮箱登录, 请联系作者")
		}

	} else {
		if params.Channel == consts.USER_CHANNEL && !config.Cfg.UserLoginRegister.AccountLogin {
			return nil, errors.New("未开启账密登录, 请联系管理员")
		}

		if params.Channel == consts.RESELLER_CHANNEL && !config.Cfg.ResellerLoginRegister.AccountLogin {
			return nil, errors.New("未开启账密登录, 请联系管理员")
		}

		if params.Channel == consts.ADMIN_CHANNEL && !config.Cfg.AdminLogin.AccountLogin {
			return nil, errors.New("未开启账密登录, 请联系作者")
		}
	}

	r := g.RequestFromCtx(ctx)
	ip := r.GetClientIp()
	token := ""

	if params.Channel == consts.USER_CHANNEL {

		account, err := dao.User.FindAccount(ctx, params.Account)

		if params.Method == consts.METHOD_ACCOUNT {

			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					return nil, errors.New("账号或密码不正确")
				}
				logger.Error(ctx, err)
				return nil, err
			}

			if !crypto.VerifyPassword(account.Password, params.Password+account.Salt) {
				return nil, errors.New("账号或密码不正确")
			}

			if account.Status == 2 {
				return nil, errors.New("账号已被禁用")
			}

		} else if params.Method == consts.METHOD_CODE {

			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {

					if err = s.Register(ctx, model.RegisterReq{
						Account:  params.Account,
						Password: grand.Letters(8),
						Terminal: params.Terminal,
						Channel:  params.Channel,
						Code:     params.Code,
						Domain:   params.Domain,
					}, consts.ACTION_LOGIN); err != nil {
						logger.Error(ctx, err)
						return nil, err
					}

					if account, err = dao.User.FindAccount(ctx, params.Account); err != nil {
						logger.Error(ctx, err)
						return nil, err
					}

				} else {
					logger.Error(ctx, err)
					return nil, err
				}

			} else if account.Status == 2 {
				return nil, errors.New("账号已被禁用")
			}

		} else {
			return nil, errors.New("账号或密码不正确")
		}

		user, err := dao.User.FindOne(ctx, bson.M{"_id": account.Uid, "status": 1})
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errors.New("用户不存在或已被禁用")
			}
			logger.Error(ctx, err)
			return nil, err
		}

		r.SetCtxVar("uid", user.Id)

		// 记录登录IP和登录时间
		if err = dao.Account.UpdateById(r.GetCtx(), account.Id, bson.M{
			"login_ip":     ip,
			"login_time":   gtime.TimestampMilli(),
			"login_domain": params.Domain,
		}); err != nil {
			logger.Error(ctx, err)
		}

		if token, err = s.GenUserToken(ctx, &model.User{
			Id:        user.Id,
			UserId:    user.UserId,
			Name:      user.Name,
			Avatar:    user.Avatar,
			Email:     user.Email,
			Phone:     user.Phone,
			Quota:     user.Quota,
			UsedQuota: user.UsedQuota,
			Models:    user.Models,
			Groups:    user.Groups,
			Rid:       user.Rid,
			Account:   account.Account,
			CreatedAt: util.FormatDateTime(user.CreatedAt),
			UpdatedAt: util.FormatDateTime(user.UpdatedAt),
		}, true); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

	} else if params.Channel == consts.RESELLER_CHANNEL {

		accountInfo, err := dao.Reseller.FindAccount(ctx, params.Account)

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

			if accountInfo.Status == 2 {
				return nil, errors.New("账号已被禁用")
			}

		} else if params.Method == consts.METHOD_CODE {

			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {

					if err = s.Register(ctx, model.RegisterReq{
						Account:  params.Account,
						Password: grand.Letters(8),
						Terminal: params.Terminal,
						Channel:  params.Channel,
						Code:     params.Code,
						Domain:   params.Domain,
					}, consts.ACTION_LOGIN); err != nil {
						logger.Error(ctx, err)
						return nil, err
					}

					if accountInfo, err = dao.Reseller.FindAccount(ctx, params.Account); err != nil {
						logger.Error(ctx, err)
						return nil, err
					}

				} else {
					logger.Error(ctx, err)
					return nil, err
				}

			} else if accountInfo.Status == 2 {
				return nil, errors.New("账号已被禁用")
			}

		} else {
			return nil, errors.New("账号或密码不正确")
		}

		reseller, err := dao.Reseller.FindOne(ctx, bson.M{"_id": accountInfo.Uid, "status": 1})
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, errors.New("用户不存在或已被禁用")
			}
			logger.Error(ctx, err)
			return nil, err
		}

		r.SetCtxVar("uid", reseller.Id)

		// 记录登录IP和时间
		if err = dao.ResellerAccount.UpdateById(r.GetCtx(), accountInfo.Id, bson.M{
			"login_ip":     ip,
			"login_time":   gtime.TimestampMilli(),
			"login_domain": params.Domain,
		}); err != nil {
			logger.Error(ctx, err)
		}

		if token, err = s.GenResellerToken(ctx, &model.Reseller{
			Id:        reseller.Id,
			UserId:    reseller.UserId,
			Name:      reseller.Name,
			Avatar:    reseller.Avatar,
			Email:     reseller.Email,
			Phone:     reseller.Phone,
			Quota:     reseller.Quota,
			UsedQuota: reseller.UsedQuota,
			Models:    reseller.Models,
			Groups:    reseller.Groups,
			Account:   accountInfo.Account,
			CreatedAt: util.FormatDateTime(reseller.CreatedAt),
			UpdatedAt: util.FormatDateTime(reseller.UpdatedAt),
		}, true); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

	} else if params.Channel == consts.ADMIN_CHANNEL {

		admin, err := dao.SysAdmin.FindOne(ctx, bson.M{
			"$or": bson.A{
				bson.M{"account": params.Account},
				bson.M{"email": params.Account},
			},
		})
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {

				count, err := dao.SysAdmin.EstimatedDocumentCount(ctx)
				if err != nil {
					logger.Error(ctx, err)
					return nil, err
				}

				// 首次登录自动创建账号
				if count == 0 {

					if err = service.SysAdmin().Create(ctx, model.SysAdminCreateReq{
						Name:     params.Account,
						Account:  params.Account,
						Password: params.Password,
					}); err != nil {
						logger.Error(ctx, err)
						return nil, err
					}

					if admin, err = dao.SysAdmin.FindOne(ctx, bson.M{"account": params.Account}); err != nil {
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

		if params.Method == consts.METHOD_ACCOUNT && !crypto.VerifyPassword(admin.Password, params.Password+admin.Salt) {
			return nil, errors.New("账号或密码不正确")
		}

		if admin.Status == 2 {
			return nil, errors.New("账号已被禁用")
		}

		r.SetCtxVar("uid", admin.Id)

		// 记录登录IP和登录时间
		if err = dao.SysAdmin.UpdateById(r.GetCtx(), admin.Id, bson.M{
			"login_ip":     ip,
			"login_time":   gtime.TimestampMilli(),
			"login_domain": params.Domain,
		}); err != nil {
			logger.Error(ctx, err)
		}

		if token, err = s.GenAdminToken(ctx, &model.SysAdmin{
			Id:        admin.Id,
			UserId:    admin.UserId,
			Name:      admin.Name,
			Avatar:    admin.Avatar,
			Email:     admin.Email,
			Phone:     admin.Phone,
			Account:   admin.Account,
			Remark:    admin.Remark,
			Status:    admin.Status,
			CreatedAt: admin.CreatedAt,
			UpdatedAt: admin.UpdatedAt,
		}, true); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	time.Sleep(time.Duration(grand.N(150, 220)) * time.Millisecond)

	loginRes := &model.LoginRes{
		Type:      "Bearer",
		Token:     token,
		ExpiresIn: config.Cfg.UserLoginRegister.SessionExpire,
	}

	if params.Channel == consts.ADMIN_CHANNEL {
		loginRes.ExpiresIn = config.Cfg.AdminLogin.SessionExpire
	} else if params.Channel == consts.RESELLER_CHANNEL {
		loginRes.ExpiresIn = config.Cfg.ResellerLoginRegister.SessionExpire
	}

	return loginRes, nil
}

// 退出登录
func (s *sAuth) Logout(ctx context.Context) error {

	token := g.RequestFromCtx(ctx).GetHeader("Authorization")
	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer"))

	key := fmt.Sprintf(consts.USER_SESSION, token)

	if gstr.HasPrefix(token, consts.RESELLER_TOKEN_PREFIX) {
		key = fmt.Sprintf(consts.RESELLER_SESSION, token)
	} else if gstr.HasPrefix(token, consts.ADMIN_TOKEN_PREFIX) {
		key = fmt.Sprintf(consts.ADMIN_SESSION, token)
	}

	if _, err := redis.Del(ctx, key); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err := s.tokenCache.Remove(ctx, key); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 找回密码
func (s *sAuth) Forget(ctx context.Context, params model.ForgetReq) error {

	// 验证验证码是否正确
	if !service.Common().VerifyCode(ctx, consts.ACTION_FORGET_ACCOUNT, params.Account, params.Code) {
		return errors.New("验证码填写错误")
	}

	if params.Channel == consts.USER_CHANNEL {

		if !config.Cfg.UserLoginRegister.EmailRetrieve {
			return errors.New("未开启找回密码, 请联系管理员")
		}

		account, err := dao.User.FindAccount(ctx, params.Account)
		if err != nil || account.Id == "" {
			return errors.New(params.Account + " 账号不存在")
		}

		if err = dao.User.ChangePasswordByUserId(ctx, account.UserId, params.Password); err != nil {
			logger.Error(ctx, err)
			return errors.New("找回密码失败")
		}
	}

	if params.Channel == consts.RESELLER_CHANNEL {

		if !config.Cfg.ResellerLoginRegister.EmailRetrieve {
			return errors.New("未开启找回密码, 请联系管理员")
		}

		account, err := dao.Reseller.FindAccount(ctx, params.Account)
		if err != nil || account.Id == "" {
			return errors.New(params.Account + " 账号不存在")
		}

		if err = dao.Reseller.ChangePasswordByUserId(ctx, account.UserId, params.Password); err != nil {
			logger.Error(ctx, err)
			return errors.New("找回密码失败")
		}
	}

	if params.Channel == consts.ADMIN_CHANNEL {

		if !config.Cfg.AdminLogin.EmailRetrieve {
			return errors.New("未开启找回密码, 请联系作者")
		}

		admin, err := dao.SysAdmin.FindOne(ctx, bson.M{"email": params.Account})
		if err != nil {
			return errors.New(params.Account + " 账号不存在")
		}

		if err = dao.SysAdmin.ChangePassword(ctx, admin.Id, params.Password); err != nil {
			logger.Error(ctx, err)
			return errors.New("找回密码失败")
		}
	}

	_ = service.Common().DelCode(ctx, consts.ACTION_FORGET_ACCOUNT, params.Account)

	return nil
}

// 生成用户Token
func (s *sAuth) GenUserToken(ctx context.Context, user *model.User, isSaveSession bool) (token string, err error) {

	token = util.NewKey(consts.USER_TOKEN_PREFIX, 32, gconv.String(user.UserId))

	if isSaveSession {

		if err = redis.SetEX(ctx, fmt.Sprintf(consts.USER_SESSION, token), gjson.MustEncodeString(user), int64(config.Cfg.UserLoginRegister.SessionExpire)); err != nil {
			logger.Errorf(ctx, "GenUserToken key: %s, error: %v", fmt.Sprintf(consts.USER_SESSION, token), err)
			return
		}

		if err = s.tokenCache.Set(ctx, fmt.Sprintf(consts.USER_SESSION, token), user, time.Duration(config.Cfg.UserLoginRegister.SessionExpire)*time.Second); err != nil {
			logger.Errorf(ctx, "GenUserToken key: %s, error: %v", fmt.Sprintf(consts.USER_SESSION, token), err)
			return
		}
	}

	return token, err
}

// 根据Token获取用户信息
func (s *sAuth) GetUserByToken(ctx context.Context, token string) (*model.User, error) {

	if tokenCache := s.tokenCache.GetVal(ctx, fmt.Sprintf(consts.USER_SESSION, token)); tokenCache != nil {
		return tokenCache.(*model.User), nil
	}

	reply, err := redis.Get(ctx, fmt.Sprintf(consts.USER_SESSION, token))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if reply == nil || reply.IsNil() {
		return nil, errors.New("session is nil")
	}

	user := new(model.User)
	if err = reply.Struct(&user); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if ttl, err := redis.TTL(ctx, fmt.Sprintf(consts.USER_SESSION, token)); err != nil {
		logger.Error(ctx, err)
	} else {
		if err = s.tokenCache.Set(ctx, fmt.Sprintf(consts.USER_SESSION, token), user, time.Duration(ttl)*time.Second); err != nil {
			logger.Errorf(ctx, "GetUserByToken key: %s, error: %v", fmt.Sprintf(consts.USER_SESSION, token), err)
		}
	}

	return user, nil
}

// 根据Token更新用户信息
func (s *sAuth) UpdateUserByToken(ctx context.Context, token string, user *model.User) error {

	if ttl, err := redis.TTL(ctx, fmt.Sprintf(consts.USER_SESSION, token)); err != nil {
		logger.Error(ctx, err)
	} else {
		if err = redis.SetEX(ctx, fmt.Sprintf(consts.USER_SESSION, token), gjson.MustEncodeString(user), ttl); err != nil {
			logger.Errorf(ctx, "UpdateUserByToken key: %s, error: %v", fmt.Sprintf(consts.USER_SESSION, token), err)
			return err
		}

		if err = s.tokenCache.Set(ctx, fmt.Sprintf(consts.USER_SESSION, token), user, time.Duration(ttl)*time.Second); err != nil {
			logger.Errorf(ctx, "UpdateUserByToken key: %s, error: %v", fmt.Sprintf(consts.USER_SESSION, token), err)
			return err
		}
	}

	return nil
}

// 生成代理商Token
func (s *sAuth) GenResellerToken(ctx context.Context, reseller *model.Reseller, isSaveSession bool) (token string, err error) {

	token = util.NewKey(consts.RESELLER_TOKEN_PREFIX, 32, gconv.String(reseller.UserId))

	if isSaveSession {

		if err = redis.SetEX(ctx, fmt.Sprintf(consts.RESELLER_SESSION, token), gjson.MustEncodeString(reseller), int64(config.Cfg.ResellerLoginRegister.SessionExpire)); err != nil {
			logger.Errorf(ctx, "GenResellerToken key: %s, error: %v", fmt.Sprintf(consts.RESELLER_SESSION, token), err)
			return
		}

		if err = s.tokenCache.Set(ctx, fmt.Sprintf(consts.RESELLER_SESSION, token), reseller, time.Duration(config.Cfg.ResellerLoginRegister.SessionExpire)*time.Second); err != nil {
			logger.Errorf(ctx, "GenResellerToken key: %s, error: %v", fmt.Sprintf(consts.RESELLER_SESSION, token), err)
			return
		}
	}

	return token, err
}

// 根据Token获取代理商信息
func (s *sAuth) GetResellerByToken(ctx context.Context, token string) (*model.Reseller, error) {

	if tokenCache := s.tokenCache.GetVal(ctx, fmt.Sprintf(consts.RESELLER_SESSION, token)); tokenCache != nil {
		return tokenCache.(*model.Reseller), nil
	}

	reply, err := redis.Get(ctx, fmt.Sprintf(consts.RESELLER_SESSION, token))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if reply == nil || reply.IsNil() {
		return nil, errors.New("session is nil")
	}

	reseller := new(model.Reseller)
	if err = reply.Struct(&reseller); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if ttl, err := redis.TTL(ctx, fmt.Sprintf(consts.RESELLER_SESSION, token)); err != nil {
		logger.Error(ctx, err)
	} else {
		if err = s.tokenCache.Set(ctx, fmt.Sprintf(consts.RESELLER_SESSION, token), reseller, time.Second*time.Duration(ttl)); err != nil {
			logger.Errorf(ctx, "GetResellerByToken key: %s, error: %v", fmt.Sprintf(consts.RESELLER_SESSION, token), err)
		}
	}

	return reseller, nil
}

// 根据Token更新代理商信息
func (s *sAuth) UpdateResellerByToken(ctx context.Context, token string, reseller *model.Reseller) error {

	if ttl, err := redis.TTL(ctx, fmt.Sprintf(consts.RESELLER_SESSION, token)); err != nil {
		logger.Error(ctx, err)
	} else {
		if err = redis.SetEX(ctx, fmt.Sprintf(consts.RESELLER_SESSION, token), gjson.MustEncodeString(reseller), ttl); err != nil {
			logger.Errorf(ctx, "UpdateResellerByToken key: %s, error: %v", fmt.Sprintf(consts.RESELLER_SESSION, token), err)
			return err
		}

		if err = s.tokenCache.Set(ctx, fmt.Sprintf(consts.RESELLER_SESSION, token), reseller, time.Duration(ttl)*time.Second); err != nil {
			logger.Errorf(ctx, "UpdateResellerByToken key: %s, error: %v", fmt.Sprintf(consts.RESELLER_SESSION, token), err)
			return err
		}
	}

	return nil
}

// 生成管理员Token
func (s *sAuth) GenAdminToken(ctx context.Context, admin *model.SysAdmin, isSaveSession bool) (token string, err error) {

	token = util.NewKey(consts.ADMIN_TOKEN_PREFIX, 32, admin.Id)

	if isSaveSession {

		if err = redis.SetEX(ctx, fmt.Sprintf(consts.ADMIN_SESSION, token), gjson.MustEncodeString(admin), int64(config.Cfg.AdminLogin.SessionExpire)); err != nil {
			logger.Errorf(ctx, "GenAdminToken key: %s, error: %v", fmt.Sprintf(consts.ADMIN_SESSION, token), err)
			return
		}

		if err = s.tokenCache.Set(ctx, fmt.Sprintf(consts.ADMIN_SESSION, token), admin, time.Duration(config.Cfg.AdminLogin.SessionExpire)*time.Second); err != nil {
			logger.Errorf(ctx, "GenAdminToken key: %s, error: %v", fmt.Sprintf(consts.ADMIN_SESSION, token), err)
			return
		}
	}

	return token, err
}

// 根据Token获取管理员信息
func (s *sAuth) GetAdminByToken(ctx context.Context, token string) (*model.SysAdmin, error) {

	if tokenCache := s.tokenCache.GetVal(ctx, fmt.Sprintf(consts.ADMIN_SESSION, token)); tokenCache != nil {
		return tokenCache.(*model.SysAdmin), nil
	}

	reply, err := redis.Get(ctx, fmt.Sprintf(consts.ADMIN_SESSION, token))
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if reply == nil || reply.IsNil() {
		return nil, errors.New("session is nil")
	}

	admin := new(model.SysAdmin)
	if err = reply.Struct(&admin); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if ttl, err := redis.TTL(ctx, fmt.Sprintf(consts.ADMIN_SESSION, token)); err != nil {
		logger.Error(ctx, err)
	} else {
		if err = s.tokenCache.Set(ctx, fmt.Sprintf(consts.ADMIN_SESSION, token), admin, time.Duration(ttl)*time.Second); err != nil {
			logger.Errorf(ctx, "GetAdminByToken key: %s, error: %v", fmt.Sprintf(consts.ADMIN_SESSION, token), err)
		}
	}

	return admin, nil
}

// 根据Token更新管理员信息
func (s *sAuth) UpdateAdminByToken(ctx context.Context, token string, admin *model.SysAdmin) error {

	if ttl, err := redis.TTL(ctx, fmt.Sprintf(consts.ADMIN_SESSION, token)); err != nil {
		logger.Error(ctx, err)
	} else {
		if err = redis.SetEX(ctx, fmt.Sprintf(consts.ADMIN_SESSION, token), gjson.MustEncodeString(admin), ttl); err != nil {
			logger.Errorf(ctx, "UpdateAdminByToken key: %s, error: %v", fmt.Sprintf(consts.ADMIN_SESSION, token), err)
			return err
		}

		if err = s.tokenCache.Set(ctx, fmt.Sprintf(consts.ADMIN_SESSION, token), admin, time.Duration(ttl)*time.Second); err != nil {
			logger.Errorf(ctx, "UpdateAdminByToken key: %s, error: %v", fmt.Sprintf(consts.ADMIN_SESSION, token), err)
			return err
		}
	}

	return nil
}
