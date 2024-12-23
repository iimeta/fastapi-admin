package model

// 发送邮件验证码接口请求参数
type SendEmailReq struct {
	Email   string `json:"email,omitempty" v:"required"`
	Channel string `json:"channel,omitempty" v:"required|in:login,register,forget_account,change_email,change_phone"`
	Domain  string `json:"domain,omitempty"` // 域名
}

// 发送邮件验证码接口响应参数
type SendEmailRes struct {
	IsDebug bool   `json:"is_debug"`
	Code    string `json:"code"`
}
