package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
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
func (s *sCommon) EmailCode(ctx context.Context, params model.SendEmailReq) error {

	defer func() {
		if val, _ := redis.Incr(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Email)); val == 1 {
			_, _ = redis.Expire(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Email), 30*60) // 锁定30分钟
		}
	}()

	if val, err := redis.GetInt(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Email)); err == nil && val >= 5 {
		return errors.New("发送验证码过于频繁, 请稍后再试")
	}

	siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, params.Domain)

	switch params.Channel {
	// 需要判断账号是否存在
	case consts.CHANNEL_LOGIN:
		//if !dao.User.IsAccountExist(ctx, params.Email) {
		//	return errors.New("账号不存在或密码错误")
		//}

	// 需要判断账号是否存在
	case consts.CHANNEL_FORGET_ACCOUNT:
		if !dao.User.IsAccountExist(ctx, params.Email) {
			return errors.New("账号不存在")
		}

	// 需要判断账号是否存在
	case consts.CHANNEL_REGISTER, consts.CHANNEL_CHANGE_EMAIL:

		if siteConfig != nil && len(siteConfig.SupportEmailSuffix) > 0 {

			isSupport := false
			for _, emailSuffix := range siteConfig.SupportEmailSuffix {
				if isSupport = gstr.HasSuffix(params.Email, emailSuffix); isSupport {
					break
				}
			}

			if !isSupport {
				return errors.New(fmt.Sprintf("邮箱仅支持 %s 后缀", siteConfig.SupportEmailSuffix))
			}
		}

		if dao.User.IsAccountExist(ctx, params.Email) {
			return errors.New("邮箱已被他人使用")
		}

	default:
		return errors.New("发送异常")
	}

	if params.Channel == consts.CHANNEL_REGISTER && siteConfig != nil && siteConfig.RegisterTips != "" {
		return errors.New(siteConfig.RegisterTips)
	}

	code := grand.Digits(6)

	if err := s.SetCode(ctx, params.Channel, params.Email, code, 15*time.Minute); err != nil {
		logger.Error(ctx, err)
		return err
	}

	logger.Infof(ctx, "正在发送邮件验证码, 操作: %s, 收件人: %s, 验证码: %s", consts.CHANNEL_MAP[params.Channel], params.Email, code)

	data := make(map[string]string)
	data["service_name"] = consts.CHANNEL_MAP[params.Channel]
	data["code"] = code

	dialer := email.NewDefaultDialer()

	if siteConfig != nil {

		if siteConfig.Host == "" {
			return errors.New("发信邮箱未配置, 请联系管理员")
		}

		dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password)
		data["copyright"] = siteConfig.Copyright
		data["jump_url"] = siteConfig.JumpUrl
	}

	template, err := util.RenderTemplate(data)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	message := email.NewMessage([]string{params.Email}, consts.CHANNEL_MAP[params.Channel], template)

	// 发送邮件验证码
	_ = email.SendMail(message, dialer)

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

	switch params.Channel {
	// 需要判断账号是否存在
	case consts.CHANNEL_LOGIN, consts.CHANNEL_FORGET_ACCOUNT:
		if !dao.User.IsAccountExist(ctx, params.Phone) {
			return errors.New("账号不存在或密码错误")
		}

	// 需要判断账号是否存在
	case consts.CHANNEL_REGISTER, consts.CHANNEL_CHANGE_MOBILE:
		if dao.User.IsAccountExist(ctx, params.Phone) {
			return errors.New("手机号已被他人使用")
		}

	default:
		return errors.New("发送异常")
	}

	code := grand.Digits(6)

	// 添加发送记录
	if err := s.SetCode(ctx, params.Channel, params.Phone, code, 15*time.Minute); err != nil {
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

// 额度过期时间转换
func ConvQuotaExpiresAt(quotaExpiresAt string) int64 {
	if quotaExpiresAt == "" {
		return 0
	}
	return gtime.NewFromStrLayout(quotaExpiresAt, time.DateTime).TimestampMilli() + 999
}

// 解析密钥
func (s *sCommon) ParseSecretKey(ctx context.Context, secretKey string) (int, int, error) {

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
