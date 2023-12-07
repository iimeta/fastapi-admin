// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ISms interface {
		// 验证短信验证码是否正确
		Verify(ctx context.Context, channel string, mobile string, code string) bool
		// 删除短信验证码记录
		Delete(ctx context.Context, channel string, mobile string)
		// 发送短信
		Send(ctx context.Context, channel string, mobile string) (string, error)
	}
)

var (
	localSms ISms
)

func Sms() ISms {
	if localSms == nil {
		panic("implement not found for interface ISms, forgot register?")
	}
	return localSms
}

func RegisterSms(i ISms) {
	localSms = i
}
