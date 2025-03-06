package consts

const (
	ACTION_LOGIN                 = "login"
	ACTION_REGISTER              = "register"
	ACTION_FORGET_ACCOUNT        = "forget_account"
	ACTION_CHANGE_MOBILE         = "change_phone"
	ACTION_CHANGE_EMAIL          = "change_email"
	ACTION_WARNING_NOTICE        = "warning_notice"
	ACTION_EXHAUSTION_NOTICE     = "exhaustion_notice"
	ACTION_EXPIRE_WARNING_NOTICE = "expire_warning_notice"
	ACTION_EXPIRE_NOTICE         = "expire_notice"
)

var ACTION_MAP = map[string]string{
	ACTION_LOGIN:                 "登录",
	ACTION_REGISTER:              "注册",
	ACTION_FORGET_ACCOUNT:        "找回密码",
	ACTION_CHANGE_EMAIL:          "换绑邮箱",
	ACTION_CHANGE_MOBILE:         "换绑手机号",
	ACTION_WARNING_NOTICE:        "额度不足提醒",
	ACTION_EXHAUSTION_NOTICE:     "额度耗尽通知",
	ACTION_EXPIRE_WARNING_NOTICE: "额度过期提醒",
	ACTION_EXPIRE_NOTICE:         "额度过期通知",
}
