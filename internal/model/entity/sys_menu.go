package entity

type SysMenu struct {
	Id        string `bson:"_id,omitempty"`        // ID
	Pid       string `bson:"pid,omitempty"`        // 父ID
	Name      string `bson:"name,omitempty"`       // 名称
	Perm      string `bson:"perm,omitempty"`       // 权限
	Type      int    `bson:"type,omitempty"`       // 类型
	Route     string `bson:"route,omitempty"`      // 路由/API
	Sort      int    `bson:"sort,omitempty"`       // 排序
	Level     int    `bson:"level,omitempty"`      // 层级
	Hidden    bool   `bson:"hidden,omitempty"`     // 隐藏
	Remark    string `bson:"remark,omitempty"`     // 备注
	Status    int    `bson:"status,omitempty"`     // 状态[1:正常, 2:禁用, -1:删除]
	Creator   string `bson:"creator,omitempty"`    // 创建人
	Updater   string `bson:"updater,omitempty"`    // 更新人
	CreatedAt int64  `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt int64  `bson:"updated_at,omitempty"` // 更新时间
}
