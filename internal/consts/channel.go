package consts

const (
	CHANNEL_LOGIN          = "login"
	CHANNEL_REGISTER       = "register"
	CHANNEL_FORGET_ACCOUNT = "forget_account"
	CHANNEL_CHANGE_MOBILE  = "change_mobile"
	CHANNEL_CHANGE_EMAIL   = "change_email"
)

var CHANNEL_MAP = map[string]string{
	CHANNEL_LOGIN:          "登录",
	CHANNEL_REGISTER:       "注册",
	CHANNEL_FORGET_ACCOUNT: "找回密码",
	CHANNEL_CHANGE_EMAIL:   "换绑邮箱",
	CHANNEL_CHANGE_MOBILE:  "换绑手机号",
}
