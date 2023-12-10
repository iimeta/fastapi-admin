package model

// 登录用户信息接口响应参数
type UserInfoRes struct {
	Id       int    `json:"id"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
}

// 用户配置信息响应参数
type UserSettingRes struct {
	User    *User        `json:"user,omitempty"`
	Setting *SettingInfo `json:"setting,omitempty"`
}

// 用户信息更新接口请求参数
type UserDetailUpdateReq struct {
	Avatar   string `json:"avatar,omitempty"`
	Nickname string `json:"nickname,omitempty" v:"required|max-length:30"`
	Gender   int    `json:"gender,omitempty" v:"in:0,1,2"`
	Motto    string `json:"motto,omitempty" v:"max-length:1024"`
	Birthday string `json:"birthday,omitempty" v:"length:0,10"`
}

// 用户密码更新接口请求参数
type UserPasswordUpdateReq struct {
	OldPassword string `json:"old_password,omitempty" v:"required"`
	NewPassword string `json:"new_password,omitempty" v:"required|min-length:6"`
}

// 用户手机号更新接口请求参数
type UserMobileUpdateReq struct {
	Mobile   string `json:"mobile,omitempty" v:"required"`
	Password string `json:"password,omitempty" v:"required"`
	Code     string `json:"code,omitempty" v:"required|length:0,6"`
}

// 用户邮箱更新接口请求参数
type UserEmailUpdateReq struct {
	Email    string `json:"email,omitempty" v:"required"`
	Password string `json:"password,omitempty" v:"required"`
	Code     string `json:"code,omitempty" v:"required|length:0,6"`
}

type User struct {
	Id        string `json:"_id,omitempty"`        // ID
	UserId    int    `json:"user_id,omitempty"`    // 用户ID
	Nickname  string `json:"nickname,omitempty"`   // 用户昵称
	Avatar    string `json:"avatar,omitempty"`     // 用户头像地址
	Gender    int    `json:"gender,omitempty"`     // 用户性别  0:未知  1:男   2:女
	Mobile    string `json:"mobile,omitempty"`     // 手机号
	Email     string `json:"email,omitempty"`      // 用户邮箱
	CreatedAt int64  `json:"created_at,omitempty"` // 注册时间
	UpdatedAt int64  `json:"updated_at,omitempty"` // 更新时间
}

type SettingInfo struct {
	ThemeMode           string `json:"theme_mode,omitempty"`
	ThemeBagImg         string `json:"theme_bag_img,omitempty"`
	ThemeColor          string `json:"theme_color,omitempty"`
	NotifyCueTone       string `json:"notify_cue_tone,omitempty"`
	KeyboardEventNotify string `json:"keyboard_event_notify,omitempty"`
}
