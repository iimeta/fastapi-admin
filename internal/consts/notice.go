package consts

const (
	SCENE_CODE                 = "code"                 // 安全验证
	SCENE_REGISTER             = "register"             // 注册
	SCENE_LOGIN                = "login"                // 登录
	SCENE_FORGET_PASSWORD      = "forget_password"      // 找回密码
	SCENE_CHANGE_PASSWORD      = "change_password"      // 修改密码
	SCENE_CHANGE_EMAIL         = "change_email"         // 修改邮箱
	SCENE_CHANGE_PHONE         = "change_phone"         // 修改手机号
	SCENE_QUOTA_RECHARGE       = "quota_recharge"       // 充值额度
	SCENE_QUOTA_WARNING        = "quota_warning"        // 额度不足提醒
	SCENE_QUOTA_EXHAUSTION     = "quota_exhaustion"     // 额度耗尽通知
	SCENE_QUOTA_EXPIRE_WARNING = "quota_expire_warning" // 额度过期提醒
	SCENE_QUOTA_EXPIRE         = "quota_expire"         // 额度过期通知
	SCENE_NOTICE               = "notice"               // 消息通知
	SCENE_NOTICE_REGISTER      = "notice_register"      // 注册通知
	SCENE_NOTICE_LOGIN         = "notice_login"         // 登录通知
)

const (
	NOTICE_CHANNEL_WEB   = "web"   // 站内信
	NOTICE_CHANNEL_EMAIL = "email" // 邮件
)

var QUOTA_NOTICE = map[string]string{
	SCENE_QUOTA_WARNING:        "warning_notice",
	SCENE_QUOTA_EXHAUSTION:     "exhaustion_notice",
	SCENE_QUOTA_EXPIRE_WARNING: "expire_warning_notice",
	SCENE_QUOTA_EXPIRE:         "expire_notice",
}

var SCENE = map[string]string{
	SCENE_CODE:                 "安全验证",
	SCENE_REGISTER:             "注册",
	SCENE_LOGIN:                "登录",
	SCENE_FORGET_PASSWORD:      "找回密码",
	SCENE_CHANGE_PASSWORD:      "修改密码",
	SCENE_CHANGE_EMAIL:         "修改邮箱",
	SCENE_CHANGE_PHONE:         "修改手机号",
	SCENE_QUOTA_RECHARGE:       "充值额度",
	SCENE_QUOTA_WARNING:        "额度不足提醒",
	SCENE_QUOTA_EXHAUSTION:     "额度耗尽通知",
	SCENE_QUOTA_EXPIRE_WARNING: "额度过期提醒",
	SCENE_QUOTA_EXPIRE:         "额度过期通知",
	SCENE_NOTICE:               "消息通知",
	SCENE_NOTICE_REGISTER:      "注册通知",
	SCENE_NOTICE_LOGIN:         "登录通知",
}
