package model

// 发送邮件验证码接口请求参数
type SendEmailReq struct {
	Email   string `json:"email,omitempty" v:"required"`   // 邮箱
	Action  string `json:"action,omitempty" v:"required"`  // 动作
	Channel string `json:"channel,omitempty" v:"required"` // 渠道
	Domain  string `json:"domain,omitempty"`               // 域名
}
