package common

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
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
	"time"
)

type sCommon struct{}

func init() {
	service.RegisterCommon(New())
}

func New() service.ICommon {
	return &sCommon{}
}

// 发送短信验证码
func (s *sCommon) SmsCode(ctx context.Context, params model.SendSmsReq) (*model.SendSmsRes, error) {

	if !config.Cfg.App.Debug {
		defer func() {
			val, _ := redis.Incr(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Mobile))
			if val == 1 {
				_, _ = redis.Expire(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Mobile), 30*60) // 锁定30分钟
			}
		}()

		val, err := redis.GetInt(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Mobile))
		if err == nil && val >= 5 {
			return nil, errors.New("发送验证码过于频繁, 请稍后再试")
		}
	}

	switch params.Channel {
	// 需要判断账号是否存在
	case consts.CHANNEL_LOGIN, consts.CHANNEL_FORGET_ACCOUNT:
		if !dao.User.IsAccountExist(ctx, params.Mobile) {
			return nil, errors.New("账号不存在或密码错误")
		}

	// 需要判断账号是否存在
	case consts.CHANNEL_REGISTER, consts.CHANNEL_CHANGE_MOBILE:
		if dao.User.IsAccountExist(ctx, params.Mobile) {
			return nil, errors.New("手机号已被他人使用")
		}

	default:
		return nil, errors.New("发送异常")
	}

	code := grand.Digits(6)

	// 添加发送记录
	if err := s.SetCode(ctx, params.Channel, params.Mobile, code, 15*time.Minute); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	// 发送短信验证码
	// TODO ... 请求第三方短信接口
	logger.Debugf(ctx, "正在发送短信验证码: %s", code)

	if config.Cfg.App.Debug {
		return &model.SendSmsRes{
			IsDebug: true,
			SmsCode: code,
		}, nil
	}

	return nil, nil
}

// 发送邮件验证码
func (s *sCommon) EmailCode(ctx context.Context, params model.SendEmailReq) (*model.SendEmailRes, error) {

	if !config.Cfg.App.Debug {
		defer func() {
			val, _ := redis.Incr(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Email))
			if val == 1 {
				_, _ = redis.Expire(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Email), 30*60) // 锁定30分钟
			}
		}()

		val, err := redis.GetInt(ctx, fmt.Sprintf(consts.LOCK_CODE, params.Email))
		if err == nil && val >= 5 {
			return nil, errors.New("发送验证码过于频繁, 请稍后再试")
		}
	}

	switch params.Channel {
	// 需要判断账号是否存在
	case consts.CHANNEL_LOGIN:
		if !dao.User.IsAccountExist(ctx, params.Email) {
			return nil, errors.New("账号不存在或密码错误")
		}

	// 需要判断账号是否存在
	case consts.CHANNEL_FORGET_ACCOUNT:
		if !dao.User.IsAccountExist(ctx, params.Email) {
			return nil, errors.New("账号不存在")
		}

	// 需要判断账号是否存在
	case consts.CHANNEL_REGISTER, consts.CHANNEL_CHANGE_EMAIL:
		if dao.User.IsAccountExist(ctx, params.Email) {
			return nil, errors.New("邮箱已被他人使用")
		}

	default:
		return nil, errors.New("发送异常")
	}

	code := grand.Digits(6)

	// 添加发送记录
	if err := s.SetCode(ctx, params.Channel, params.Email, code, 15*time.Minute); err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	logger.Debugf(ctx, "正在发送邮件验证码, 操作: %s, 收件人: %s, 验证码: %s", consts.CHANNEL_MAP[params.Channel], params.Email, code)

	data := make(map[string]string)
	data["service_name"] = consts.CHANNEL_MAP[params.Channel]
	data["code"] = code

	template, err := util.RenderTemplate(data)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	// 发送邮件验证码
	_ = email.SendMail(&email.Option{
		To:      []string{params.Email},
		Subject: consts.CHANNEL_MAP[params.Channel],
		Body:    template,
	})

	if config.Cfg.App.Debug {
		return &model.SendEmailRes{
			IsDebug: true,
			Code:    code,
		}, nil
	}

	return nil, nil
}

func (s *sCommon) SetCode(ctx context.Context, channel string, email string, code string, exp time.Duration) error {

	pipe := redis.Pipeline(ctx)
	pipe.Del(ctx, s.failName(channel, email))
	pipe.Set(ctx, s.name(channel, email), code, exp)
	_, err := redis.Pipelined(ctx, pipe)

	return err
}

func (s *sCommon) GetCode(ctx context.Context, channel, account string) (string, error) {
	return redis.GetStr(ctx, s.name(channel, account))
}

func (s *sCommon) DelCode(ctx context.Context, channel, account string) error {
	_, err := redis.Del(ctx, s.name(channel, account))
	return err
}

func (s *sCommon) VerifyCode(ctx context.Context, channel, account, code string) (pass bool) {

	defer func() {
		if !pass {
			// 3分钟内同一个邮件验证码错误次数超过5次, 删除验证码
			num, _ := redis.Incr(ctx, s.failName(channel, account))
			if num >= 5 {
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
