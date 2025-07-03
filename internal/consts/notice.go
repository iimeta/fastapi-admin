package consts

const (
	SCENE_CODE                 = "code"                 // 验证码
	SCENE_LOGIN                = "login"                // 登录通知
	SCENE_REGISTER             = "register"             // 注册通知
	SCENE_FORGET_PASSWORD      = "forget_password"      // 找回密码
	SCENE_CHANGE_PASSWORD      = "change_password"      // 修改密码
	SCENE_CHANGE_EMAIL         = "change_email"         // 修改邮箱
	SCENE_QUOTA_WARNING        = "quota_warning"        // 额度不足提醒
	SCENE_QUOTA_EXHAUSTION     = "quota_exhaustion"     // 额度耗尽通知
	SCENE_QUOTA_EXPIRE_WARNING = "quota_expire_warning" // 额度过期提醒
	SCENE_QUOTA_EXPIRE         = "quota_expire"         // 额度过期通知
	SCENE_NOTICE               = "notice"               // 通知公告
)

var NOTICE_MAP = map[string]string{
	SCENE_QUOTA_WARNING:        "warning_notice",
	SCENE_QUOTA_EXHAUSTION:     "exhaustion_notice",
	SCENE_QUOTA_EXPIRE_WARNING: "expire_warning_notice",
	SCENE_QUOTA_EXPIRE:         "expire_notice",
}
