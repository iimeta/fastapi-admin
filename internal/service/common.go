// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"

	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

type (
	ICommon interface {
		// 发送邮件验证码
		EmailCode(ctx context.Context, params model.SendEmailReq) (err error)
		// 发送短信验证码
		SmsCode(ctx context.Context, params model.SendSmsReq) error
		// 缓存验证码, rid为发信站点所属代理商ID
		SetCode(ctx context.Context, channel string, account string, code string, exp time.Duration, rid ...int) error
		// 获取验证码
		GetCode(ctx context.Context, channel string, account string) (string, error)
		// 删除验证码
		DelCode(ctx context.Context, channel string, account string) error
		// 校验验证码
		VerifyCode(ctx context.Context, channel string, account string, code string) (pass bool)
		// 校验验证码及其发信站点
		VerifyCodeRid(ctx context.Context, channel string, account string, code string, rid int) bool
	}
)

var (
	localCommon ICommon
)

func Common() ICommon {
	if localCommon == nil {
		panic("implement not found for interface ICommon, forgot register?")
	}
	return localCommon
}

func RegisterCommon(i ICommon) {
	localCommon = i
}
