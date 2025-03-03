package consts

const (
	ACTION_LOGIN             = "login"
	ACTION_REGISTER          = "register"
	ACTION_FORGET_ACCOUNT    = "forget_account"
	ACTION_CHANGE_MOBILE     = "change_phone"
	ACTION_CHANGE_EMAIL      = "change_email"
	ACTION_WARNING_THRESHOLD = "warning_threshold"
	ACTION_EXHAUSTION_NOTICE = "exhaustion_notice"
)

var ACTION_MAP = map[string]string{
	ACTION_LOGIN:             "登录",
	ACTION_REGISTER:          "注册",
	ACTION_FORGET_ACCOUNT:    "找回密码",
	ACTION_CHANGE_EMAIL:      "换绑邮箱",
	ACTION_CHANGE_MOBILE:     "换绑手机号",
	ACTION_WARNING_THRESHOLD: "额度不足提醒",
	ACTION_EXHAUSTION_NOTICE: "额度耗尽提醒",
}
