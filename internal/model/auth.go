package model

// 登录接口请求参数
type LoginReq struct {
	Account  string `json:"account,omitempty" v:"required"`                            // 账号
	Password string `json:"password,omitempty" v:"required"`                           // 密码
	Terminal string `json:"terminal,omitempty" v:"required|in:web,h5,ios,windows,mac"` // 终端
	Channel  string `json:"channel,omitempty" v:"required|in:user,admin"`              // 渠道
}

// 登录接口响应参数
type LoginRes struct {
	Type      string `json:"type,omitempty"`       // Token 类型
	Token     string `json:"token,omitempty"`      // token
	ExpiresIn int    `json:"expires_in,omitempty"` // 过期时间
}

// 注册接口请求参数
type RegisterReq struct {
	Account  string `json:"account,omitempty" v:"required"`                            // 账号
	Password string `json:"password,omitempty" v:"required|min-length:6"`              // 密码
	Nickname string `json:"nickname,omitempty" v:"required|length:2,30"`               // 昵称
	Terminal string `json:"terminal,omitempty" v:"required|in:web,h5,ios,windows,mac"` // 终端
	Code     string `json:"code,omitempty" v:"required"`                               // 验证码
}

// Token 刷新接口响应参数
type RefreshRes struct {
	Type        string `json:"type,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
}

// 找回密码接口请求参数
type ForgetReq struct {
	Account  string `json:"account,omitempty" v:"required"`               // 账号
	Password string `json:"password,omitempty" v:"required|min-length:6"` // 密码
	Code     string `json:"code,omitempty" v:"required"`                  // 验证码
}
