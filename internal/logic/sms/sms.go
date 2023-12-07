package sms

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/cache"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"time"
)

type sSms struct {
	sms *cache.SmsStorage
}

func init() {
	service.RegisterSms(New())
}

func New() service.ISms {
	return &sSms{
		sms: cache.NewSmsStorage(redis.Client),
	}
}

// 验证短信验证码是否正确
func (s *sSms) Verify(ctx context.Context, channel string, mobile string, code string) bool {
	return s.sms.Verify(ctx, channel, mobile, code)
}

// 删除短信验证码记录
func (s *sSms) Delete(ctx context.Context, channel string, mobile string) {
	_ = s.sms.Del(ctx, channel, mobile)
}

// 发送短信
func (s *sSms) Send(ctx context.Context, channel string, mobile string) (string, error) {

	code := util.GenValidateCode(6)

	// 添加发送记录
	if err := s.sms.Set(ctx, channel, mobile, code, 15*time.Minute); err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	// TODO ... 请求第三方短信接口
	logger.Debugf(ctx, "正在发送短信验证码: %s", code)

	return code, nil
}
