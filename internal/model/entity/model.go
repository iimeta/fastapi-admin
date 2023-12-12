package entity

type Model struct {
	Id        string   `bson:"_id,omitempty"`        // ID
	Corp      string   `bson:"corp,omitempty"`       // 公司[OpenAI;Baidu;Xfyun;Aliyun;Midjourney]
	Name      string   `bson:"name,omitempty"`       // 模型名称
	Model     string   `bson:"model,omitempty"`      // 模型
	Type      int      `bson:"type,omitempty"`       // 模型类型[1:文生文; 2:文生图; 3:图生文; 4:图生图; 5:文生语音; 6:语音生文]
	BaseUrl   string   `bson:"base_url,omitempty"`   // 默认官方模型地址
	Path      string   `bson:"path,omitempty"`       // 默认官方模型地址路径
	Proxy     string   `bson:"proxy,omitempty"`      // 代理
	Keys      []string `bson:"keys,omitempty"`       // 密钥
	Remark    string   `bson:"remark,omitempty"`     // 备注
	Status    int      `bson:"status,omitempty"`     // 状态[1:正常; 2:禁用; -1:删除]
	Creator   string   `bson:"creator,omitempty"`    // 创建人
	Updater   string   `bson:"updater,omitempty"`    // 更新人
	CreatedAt int64    `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt int64    `bson:"updated_at,omitempty"` // 更新时间
}
