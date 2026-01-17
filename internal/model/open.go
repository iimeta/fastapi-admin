package model

// 系统配置接口请求参数
type SysConfigReq struct {
	Domain string `json:"domain,omitempty"` // 域名
	Path   string `json:"path,omitempty"`   // 路径
}

// 系统配置接口响应参数
type SysConfigRes struct {
	*SysConfig
}
