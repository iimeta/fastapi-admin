package model

// 发送短信验证码接口请求参数
type SendSmsReq struct {
	Phone   string `json:"phone,omitempty" v:"required|length:0,11"`                                                 // 手机号
	Action  string `json:"action,omitempty" v:"required|in:login,register,forget_account,change_email,change_phone"` // 动作
	Channel string `json:"channel,omitempty" v:"required|in:user,reseller,admin"`                                    // 渠道
	Domain  string `json:"domain,omitempty"`                                                                         // 域名
}
