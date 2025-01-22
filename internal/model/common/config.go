package common

import "time"

type Core struct {
	SecretKeyPrefix string `bson:"secret_key_prefix" json:"secret_key_prefix"`
	ErrorPrefix     string `bson:"error_prefix"      json:"error_prefix"`
	ChannelPrefix   string `bson:"channel_prefix"    json:"channel_prefix"`
}

type Http struct {
	Open     bool          `bson:"open"      json:"open"` // 开关
	Timeout  time.Duration `bson:"timeout"   json:"timeout"`
	ProxyUrl string        `bson:"proxy_url" json:"proxy_url"`
}

type Email struct {
	Open     bool   `bson:"open"      json:"open"`      // 开关
	Host     string `bson:"host"      json:"host"`      // smtp.xxx.com
	Port     int    `bson:"port"      json:"port"`      // 端口号
	UserName string `bson:"user_name" json:"user_name"` // 登录账号
	Password string `bson:"password"  json:"password"`  // 登录密码
	FromName string `bson:"from_name" json:"from_name"` // 发送人名称
}

type Statistics struct {
	Open        bool          `bson:"open"         json:"open"` // 开关
	Cron        string        `bson:"cron"         json:"cron"`
	Days        int           `bson:"days"         json:"days"`
	Limit       int64         `bson:"limit"        json:"limit"`
	LockMinutes time.Duration `bson:"lock_minutes" json:"lock_minutes"`
}

type Api struct {
	Retry                   int   `bson:"retry"                       json:"retry"`
	ModelKeyErrDisable      int64 `bson:"model_key_err_disable"       json:"model_key_err_disable"`
	ModelAgentErrDisable    int64 `bson:"model_agent_err_disable"     json:"model_agent_err_disable"`
	ModelAgentKeyErrDisable int64 `bson:"model_agent_key_err_disable" json:"model_agent_key_err_disable"`
}

type Midjourney struct {
	Open            bool   `bson:"open"              json:"open"` // 开关
	CdnUrl          string `bson:"cdn_url"           json:"cdn_url"`
	ApiBaseUrl      string `bson:"api_base_url"      json:"api_base_url"`
	ApiSecret       string `bson:"api_secret"        json:"api_secret"`
	ApiSecretHeader string `bson:"api_secret_header" json:"api_secret_header"`
	CdnOriginalUrl  string `bson:"cdn_original_url"  json:"cdn_original_url"`
}

type Gcp struct {
	GetTokenUrl string `bson:"get_token_url" json:"get_token_url"`
}

type Log struct {
	Open    bool     `bson:"open"    json:"open"`    // 开关
	Records []string `bson:"records" json:"records"` // 日志记录
}

type Error struct {
	Open         bool     `bson:"open"          json:"open"` // 开关
	ShieldUser   []string `bson:"shield_user"   json:"shield_user"`
	AutoDisabled []string `bson:"auto_disabled" json:"auto_disabled"`
	NotRetry     []string `bson:"not_retry"     json:"not_retry"`
	NotShield    []string `bson:"not_shield"    json:"not_shield"`
}

type Debug struct {
	Open bool `bson:"open" json:"open"` // 开关
}
