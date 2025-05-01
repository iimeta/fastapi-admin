package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/email"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"strings"
	"time"
)

type sCommon struct{}

func init() {
	service.RegisterCommon(New())
}

func New() service.ICommon {
	return &sCommon{}
}

// 发送邮件验证码
func (s *sCommon) EmailCode(ctx context.Context, params model.SendEmailReq) (err error) {

	defer func() {
		if err == nil {
			if val, _ := redis.Incr(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Email)); val == 1 {
				_, _ = redis.Expire(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Email), 30*60) // 锁定30分钟
			}
		}
	}()

	if val, err := redis.GetInt(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Email)); err == nil && val >= 5 {
		return errors.New("发送验证码过于频繁, 请稍后再试")
	}

	siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, params.Domain)
	if siteConfig != nil {
		if siteConfig.Host == "" && (!config.Cfg.Email.Open || config.Cfg.Email.Host == "") {
			return errors.New("发信邮箱未配置, 请联系管理员")
		}
	} else if !config.Cfg.Email.Open || config.Cfg.Email.Host == "" {
		return errors.New("发信邮箱未配置, 请联系管理员")
	}

	switch params.Action {
	case consts.ACTION_LOGIN:
		if params.Channel == consts.USER_CHANNEL && !config.Cfg.UserLoginRegister.EmailLogin {
			return errors.New("未开启邮箱登录, 请联系管理员")
		}

		if params.Channel == consts.RESELLER_CHANNEL && !config.Cfg.ResellerLoginRegister.EmailLogin {
			return errors.New("未开启邮箱登录, 请联系管理员")
		}

		if params.Channel == consts.ADMIN_CHANNEL && !config.Cfg.AdminLogin.EmailLogin {
			return errors.New("未开启邮箱登录, 请联系作者")
		}
	case consts.ACTION_FORGET_ACCOUNT:
		if params.Channel == consts.USER_CHANNEL && !config.Cfg.UserLoginRegister.EmailRetrieve {
			return errors.New("未开启找回密码, 请联系管理员")
		}

		if params.Channel == consts.RESELLER_CHANNEL && !config.Cfg.ResellerLoginRegister.EmailRetrieve {
			return errors.New("未开启找回密码, 请联系管理员")
		}

		if params.Channel == consts.ADMIN_CHANNEL && !config.Cfg.AdminLogin.EmailRetrieve {
			return errors.New("未开启找回密码, 请联系作者")
		}
	case consts.ACTION_REGISTER:

		if params.Channel == consts.USER_CHANNEL && !config.Cfg.UserLoginRegister.EmailRegister {
			return errors.New("未开启用户注册, 请联系管理员")
		}

		if params.Channel == consts.RESELLER_CHANNEL && !config.Cfg.ResellerLoginRegister.EmailRegister {
			return errors.New("未开启代理商注册, 请联系管理员")
		}

		if siteConfig != nil {

			if siteConfig.RegisterTips != "" {
				return errors.New(siteConfig.RegisterTips)
			}

			if len(siteConfig.SupportEmailSuffix) > 0 {

				isSupport := false
				for _, emailSuffix := range siteConfig.SupportEmailSuffix {
					if isSupport = gstr.HasSuffix(params.Email, emailSuffix); isSupport {
						break
					}
				}

				if !isSupport {
					return errors.Newf("邮箱仅支持 %s 后缀", siteConfig.SupportEmailSuffix)
				}
			}
		}

		if params.Channel == consts.USER_CHANNEL && dao.User.IsAccountExist(ctx, params.Email) {
			return errors.New("邮箱已被他人使用")
		}

		if params.Channel == consts.RESELLER_CHANNEL && dao.Reseller.IsAccountExist(ctx, params.Email) {
			return errors.New("邮箱已被他人使用")
		}

	case consts.ACTION_CHANGE_EMAIL:

		if siteConfig != nil && len(siteConfig.SupportEmailSuffix) > 0 {

			isSupport := false
			for _, emailSuffix := range siteConfig.SupportEmailSuffix {
				if isSupport = gstr.HasSuffix(params.Email, emailSuffix); isSupport {
					break
				}
			}

			if !isSupport {
				return errors.Newf("邮箱仅支持 %s 后缀", siteConfig.SupportEmailSuffix)
			}
		}

		if params.Channel == consts.USER_CHANNEL && dao.User.IsAccountExist(ctx, params.Email) {
			return errors.New("邮箱已被他人使用")
		}

		if params.Channel == consts.RESELLER_CHANNEL && dao.Reseller.IsAccountExist(ctx, params.Email) {
			return errors.New("邮箱已被他人使用")
		}
	}

	code := grand.Digits(6)

	if err = s.SetCode(ctx, params.Action, params.Email, code, 15*time.Minute); err != nil {
		logger.Error(ctx, err)
		return err
	}

	logger.Infof(ctx, "正在发送邮件验证码, 操作: %s, 收件人: %s, 验证码: %s", consts.ACTION_MAP[params.Action], params.Email, code)

	data := make(map[string]string)
	data["service_name"] = consts.ACTION_MAP[params.Action]
	data["code"] = code

	dialer := email.NewDefaultDialer()

	if siteConfig != nil {

		data["copyright"] = siteConfig.Copyright
		data["jump_url"] = siteConfig.JumpUrl

		if siteConfig.Host != "" {
			dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password)
		}
	}

	template, err := util.RenderTemplate(data, "verify_code")
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	message := email.NewMessage([]string{params.Email}, consts.ACTION_MAP[params.Action], template)

	// 发送邮件验证码
	if err = email.SendMail(message, dialer); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 发送短信验证码
func (s *sCommon) SmsCode(ctx context.Context, params model.SendSmsReq) error {

	defer func() {
		if val, _ := redis.Incr(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Phone)); val == 1 {
			_, _ = redis.Expire(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Phone), 30*60) // 锁定30分钟
		}
	}()

	if val, err := redis.GetInt(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Phone)); err == nil && val >= 5 {
		return errors.New("发送验证码过于频繁, 请稍后再试")
	}

	switch params.Action {
	// 需要判断账号是否存在
	case consts.ACTION_LOGIN, consts.ACTION_FORGET_ACCOUNT:
		if !dao.User.IsAccountExist(ctx, params.Phone) {
			return errors.New("账号不存在或密码错误")
		}

	// 需要判断账号是否存在
	case consts.ACTION_REGISTER, consts.ACTION_CHANGE_MOBILE:
		if dao.User.IsAccountExist(ctx, params.Phone) {
			return errors.New("手机号已被他人使用")
		}

	default:
		return errors.New("发送异常")
	}

	code := grand.Digits(6)

	// 添加发送记录
	if err := s.SetCode(ctx, params.Action, params.Phone, code, 15*time.Minute); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 发送短信验证码
	// TODO ... 请求第三方短信接口
	logger.Debugf(ctx, "正在发送短信验证码: %s", code)

	return nil
}

// 缓存验证码
func (s *sCommon) SetCode(ctx context.Context, channel string, account string, code string, exp time.Duration) error {

	pipe := redis.Pipeline(ctx)
	pipe.Del(ctx, s.failName(channel, account))
	pipe.Set(ctx, s.name(channel, account), code, exp)
	_, err := redis.Pipelined(ctx, pipe)

	return err
}

// 获取验证码
func (s *sCommon) GetCode(ctx context.Context, channel, account string) (string, error) {
	return redis.GetStr(ctx, s.name(channel, account))
}

// 删除验证码
func (s *sCommon) DelCode(ctx context.Context, channel, account string) error {
	_, err := redis.Del(ctx, s.name(channel, account))
	return err
}

// 校验验证码
func (s *sCommon) VerifyCode(ctx context.Context, channel, account, code string) (pass bool) {

	defer func() {
		if !pass {
			// 3分钟内同一个邮件验证码错误次数超过5次, 删除验证码
			if num, _ := redis.Incr(ctx, s.failName(channel, account)); num >= 5 {
				pipe := redis.Pipeline(ctx)
				pipe.Del(ctx, s.name(channel, account))
				pipe.Del(ctx, s.failName(channel, account))
				_, _ = redis.Pipelined(ctx, pipe)
			} else if num == 1 {
				_, _ = redis.Expire(ctx, s.failName(channel, account), 180)
			}
		}
	}()

	value, err := s.GetCode(ctx, channel, account)
	if err != nil || len(value) == 0 {
		return false
	}

	if value == code {
		return true
	}

	return false
}

func (s *sCommon) name(channel, account string) string {
	return fmt.Sprintf("code:%s:%s", channel, gmd5.MustEncryptString(account))
}

func (s *sCommon) failName(channel, account string) string {
	return fmt.Sprintf("code:fail:%s:%s", channel, gmd5.MustEncryptString(account))
}

// 解析密钥
func ParseSecretKey(ctx context.Context, secretKey string) (int, int, error) {

	if !gstr.HasPrefix(secretKey, config.Cfg.Core.SecretKeyPrefix) {
		return 0, 0, errors.ERR_INVALID_API_KEY
	}

	secretKey = strings.TrimPrefix(secretKey, config.Cfg.Core.SecretKeyPrefix)

	userId, err := gregex.ReplaceString("[a-zA-Z-]*", "", secretKey[:len(secretKey)/2])
	if err != nil {
		logger.Error(ctx, err)
		return 0, 0, err
	}

	appId, err := gregex.ReplaceString("[a-zA-Z-]*", "", secretKey[len(secretKey)/2:])
	if err != nil {
		logger.Error(ctx, err)
		return 0, 0, err
	}

	return gconv.Int(userId), gconv.Int(appId), nil
}
