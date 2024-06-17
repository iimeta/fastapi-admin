package model

// 发送短信验证码接口请求参数
type SendSmsReq struct {
	Phone   string `json:"phone,omitempty" v:"required|length:0,11"`
	Channel string `json:"channel,omitempty" v:"required|in:login,register,forget_account,change_email,change_phone"`
}

// 发送短信验证码接口响应参数
type SendSmsRes struct {
	IsDebug bool   `json:"is_debug"`
	SmsCode string `json:"code"`
}
