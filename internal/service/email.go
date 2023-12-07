// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IEmail interface {
		// 验证邮件验证码是否正确
		Verify(ctx context.Context, channel string, email string, code string) bool
		// 删除邮件验证码记录
		Delete(ctx context.Context, channel string, email string)
		// 发送邮件
		Send(ctx context.Context, channel string, emailAddr string) (string, error)
	}
)

var (
	localEmail IEmail
)

func Email() IEmail {
	if localEmail == nil {
		panic("implement not found for interface IEmail, forgot register?")
	}
	return localEmail
}

func RegisterEmail(i IEmail) {
	localEmail = i
}
