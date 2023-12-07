package logic

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/cache"
	"github.com/iimeta/fastapi-admin/utility/email"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"time"
)

type sEmail struct {
	email *cache.EmailStorage
}

func init() {
	service.RegisterEmail(New())
}

func New() service.IEmail {
	return &sEmail{
		email: cache.NewEmailStorage(redis.Client),
	}
}

// 验证邮件验证码是否正确
func (s *sEmail) Verify(ctx context.Context, channel string, email string, code string) bool {
	return s.email.Verify(ctx, channel, email, code)
}

// 删除邮件验证码记录
func (s *sEmail) Delete(ctx context.Context, channel string, email string) {
	_ = s.email.Del(ctx, channel, email)
}

// 发送邮件
func (s *sEmail) Send(ctx context.Context, channel string, emailAddr string) (string, error) {

	code := util.GenValidateCode(6)

	// 添加发送记录
	if err := s.email.Set(ctx, channel, emailAddr, code, 15*time.Minute); err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	logger.Debugf(ctx, "正在发送邮件验证码, 操作: %s, 收件人: %s, 验证码: %s", consts.CHANNEL_MAP[channel], emailAddr, code)

	data := make(map[string]string)
	data["service_name"] = consts.CHANNEL_MAP[channel]
	data["code"] = code

	template, err := util.RenderTemplate(data)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	_ = email.SendMail(&email.Option{
		To:      []string{emailAddr},
		Subject: consts.CHANNEL_MAP[channel],
		Body:    template,
	})

	return code, nil
}
