package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/model"
)

// 发送短信验证码接口请求参数
type SendSmsReq struct {
	g.Meta `path:"/sms-code" tags:"common" method:"post" summary:"发送短信验证码接口"`
	model.SendSmsReq
}

// 发送短信验证码接口响应参数
type SendSmsRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SendSmsRes
}

// 发送邮件验证码接口请求参数
type SendEmailReq struct {
	g.Meta `path:"/email-code" tags:"common" method:"post" summary:"发送邮件验证码接口"`
	model.SendEmailReq
}

// 发送邮件验证码接口响应参数
type SendEmailRes struct {
	g.Meta `mime:"application/json" example:"json"`
	*model.SendEmailRes
}
