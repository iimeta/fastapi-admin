// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/internal/model"
)

type (
	ICommon interface {
		// 发送短信验证码
		SmsCode(ctx context.Context, params model.SendSmsReq) (*model.SendSmsRes, error)
		// 发送邮件验证码
		EmailCode(ctx context.Context, params model.SendEmailReq) (*model.SendEmailRes, error)
		// 公共设置
		Setting(ctx context.Context) error
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
